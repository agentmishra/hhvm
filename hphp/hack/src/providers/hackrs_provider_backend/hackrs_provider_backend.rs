// Copyright (c) Meta Platforms, Inc. and affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.

pub mod naming_table;

#[cfg(test)]
mod test_naming_table;

use std::collections::BTreeMap;
use std::path::PathBuf;
use std::sync::Arc;

use anyhow::Result;
use datastore::ChangesStore;
use datastore::Store;
use decl_parser::DeclParser;
use file_provider::DiskProvider;
use file_provider::FileProvider;
use folded_decl_provider::FoldedDeclProvider;
use folded_decl_provider::LazyFoldedDeclProvider;
use naming_provider::NamingProvider;
use naming_table::NamingTable;
use ocamlrep::ptr::UnsafeOcamlPtr;
use ocamlrep::FromOcamlRep;
use ocamlrep::ToOcamlRep;
use oxidized::global_options::GlobalOptions;
use oxidized_by_ref::direct_decl_parser::ParsedFileWithHashes;
use pos::RelativePath;
use pos::RelativePathCtx;
use pos::TypeName;
use shallow_decl_provider::LazyShallowDeclProvider;
use shallow_decl_provider::ShallowDeclProvider;
use shallow_decl_provider::ShallowDeclStore;
use shm_store::OcamlShmStore;
use ty::decl;
use ty::decl::folded::FoldedClass;
use ty::decl::shallow::NamedDecl;
use ty::reason::BReason as BR;

pub struct HhServerProviderBackend {
    path_ctx: Arc<RelativePathCtx>,
    opts: GlobalOptions,
    decl_parser: DeclParser<BR>,
    file_store: Arc<ChangesStore<RelativePath, FileType>>,
    file_provider: Arc<FileProviderWithChanges>,
    naming_table: Arc<NamingTable>,
    /// Collection of Arcs pointing to the backing stores for the
    /// ShallowDeclStore below, allowing us to invoke push/pop_local_changes.
    shallow_decl_changes_store: Arc<ShallowStoreWithChanges>,
    shallow_decl_store: Arc<ShallowDeclStore<BR>>,
    lazy_shallow_decl_provider: Arc<LazyShallowDeclProvider<BR>>,
    folded_classes_shm: Arc<OcamlShmStore<TypeName, Arc<FoldedClass<BR>>>>,
    folded_classes_store: Arc<ChangesStore<TypeName, Arc<FoldedClass<BR>>>>,
    folded_decl_provider: Arc<LazyFoldedDeclProvider<BR>>,
}

#[derive(serde::Serialize, serde::Deserialize)]
pub struct Config {
    pub path_ctx: RelativePathCtx,
    pub opts: GlobalOptions,
    pub db_path: Option<PathBuf>,
}

impl HhServerProviderBackend {
    pub fn new(config: Config) -> Result<Self> {
        let Config {
            path_ctx,
            opts,
            db_path,
        } = config;
        let path_ctx = Arc::new(path_ctx);
        let file_store = Arc::new(ChangesStore::new(Arc::new(OcamlShmStore::new(
            "File",
            shm_store::Evictability::NonEvictable,
            shm_store::Compression::default(),
        ))));
        let file_provider = Arc::new(FileProviderWithChanges {
            delta_and_changes: Arc::clone(&file_store),
            disk: DiskProvider::new(Arc::clone(&path_ctx), None),
        });
        let decl_parser = DeclParser::with_options(Arc::clone(&file_provider) as _, opts.clone());
        let dependency_graph = Arc::new(depgraph_api::NoDepGraph::default());
        let naming_table = Arc::new(NamingTable::new(db_path)?);

        let shallow_decl_changes_store = Arc::new(ShallowStoreWithChanges::new());
        let shallow_decl_store = shallow_decl_changes_store.as_shallow_decl_store();

        let lazy_shallow_decl_provider = Arc::new(LazyShallowDeclProvider::new(
            Arc::clone(&shallow_decl_store),
            Arc::clone(&naming_table) as _,
            decl_parser.clone(),
        ));

        let folded_classes_shm = Arc::new(OcamlShmStore::new(
            "Decl_Class",
            shm_store::Evictability::Evictable,
            shm_store::Compression::default(),
        ));
        let folded_classes_store =
            Arc::new(ChangesStore::new(Arc::clone(&folded_classes_shm) as _));

        let folded_decl_provider = Arc::new(LazyFoldedDeclProvider::new(
            Arc::new(opts.clone()),
            Arc::clone(&folded_classes_store) as _,
            Arc::clone(&lazy_shallow_decl_provider) as _,
            dependency_graph,
        ));

        Ok(Self {
            path_ctx,
            opts,
            file_store,
            file_provider,
            decl_parser,
            folded_decl_provider,
            naming_table,
            shallow_decl_changes_store,
            shallow_decl_store,
            lazy_shallow_decl_provider,
            folded_classes_shm,
            folded_classes_store,
        })
    }

    pub fn config(&self) -> Config {
        Config {
            path_ctx: (*self.path_ctx).clone(),
            db_path: self.naming_table.db_path(),
            opts: self.opts.clone(),
        }
    }

    pub fn opts(&self) -> &GlobalOptions {
        &self.opts
    }

    pub fn naming_table(&self) -> &NamingTable {
        &self.naming_table
    }

    pub fn file_store(&self) -> &dyn Store<RelativePath, FileType> {
        &*self.file_store
    }

    pub fn shallow_decl_provider(&self) -> &dyn ShallowDeclProvider<BR> {
        &*self.lazy_shallow_decl_provider
    }

    /// Decl-parse the given file, dedup duplicate definitions of the same
    /// symbol (within the file, as well as removing losers of naming conflicts
    /// with other files), and add the parsed decls to the shallow decl store.
    pub fn parse_and_cache_decls<'a>(
        &self,
        path: RelativePath,
        text: &'a [u8],
        arena: &'a bumpalo::Bump,
    ) -> Result<ParsedFileWithHashes<'a>> {
        let hashed_file = self.decl_parser.parse_impl(path, text, arena);
        self.lazy_shallow_decl_provider.dedup_and_add_decls(
            path,
            (hashed_file.iter()).map(|(name, decl, _)| NamedDecl::from(&(*name, *decl))),
        )?;
        Ok(hashed_file)
    }

    /// Directly add the given decls to the shallow decl store (without removing
    /// php_stdlib decls, deduping, or removing naming conflict losers).
    pub fn add_decls(
        &self,
        decls: &[&(&str, oxidized_by_ref::shallow_decl_defs::Decl<'_>)],
    ) -> Result<()> {
        self.shallow_decl_store
            .add_decls(decls.iter().copied().map(Into::into))
    }

    pub fn push_local_changes(&self) {
        self.file_store.push_local_changes();
        self.naming_table.push_local_changes();
        self.shallow_decl_changes_store.push_local_changes();
        self.folded_classes_store.push_local_changes();
    }

    pub fn pop_local_changes(&self) {
        self.file_store.pop_local_changes();
        self.naming_table.pop_local_changes();
        self.shallow_decl_changes_store.pop_local_changes();
        self.folded_classes_store.pop_local_changes();
    }

    pub fn oldify_defs(&self, names: &file_info::Names) -> Result<()> {
        self.folded_classes_store
            .remove_batch(&mut names.classes.iter().map(Into::into))?;
        self.shallow_decl_changes_store.oldify_defs(names)
    }

    pub fn remove_old_defs(&self, names: &file_info::Names) -> Result<()> {
        self.folded_classes_store
            .remove_batch(&mut names.classes.iter().map(Into::into))?;
        self.shallow_decl_changes_store.remove_old_defs(names)
    }

    pub fn remove_defs(&self, names: &file_info::Names) -> Result<()> {
        self.folded_classes_store
            .remove_batch(&mut names.classes.iter().map(Into::into))?;
        self.shallow_decl_changes_store.remove_defs(names)
    }

    pub fn get_old_defs(
        &self,
        names: &file_info::Names,
    ) -> Result<(
        BTreeMap<pos::TypeName, Option<Arc<decl::ShallowClass<BR>>>>,
        BTreeMap<pos::FunName, Option<Arc<decl::FunDecl<BR>>>>,
        BTreeMap<pos::TypeName, Option<Arc<decl::TypedefDecl<BR>>>>,
        BTreeMap<pos::ConstName, Option<Arc<decl::ConstDecl<BR>>>>,
        BTreeMap<pos::ModuleName, Option<Arc<decl::ModuleDecl<BR>>>>,
    )> {
        self.shallow_decl_changes_store.get_old_defs(names)
    }
}

impl rust_provider_backend_api::RustProviderBackend<BR> for HhServerProviderBackend {
    fn file_provider(&self) -> &dyn FileProvider {
        &*self.file_provider
    }

    fn naming_provider(&self) -> &dyn NamingProvider {
        &*self.naming_table
    }

    fn folded_decl_provider(&self) -> &dyn FoldedDeclProvider<BR> {
        &*self.folded_decl_provider
    }

    fn as_any(&self) -> &dyn std::any::Any {
        self
    }
}

#[rustfmt::skip]
impl HhServerProviderBackend {
    /// SAFETY: This method (and all other `get_ocaml_` methods) call into the
    /// OCaml runtime and may trigger a GC. Must be invoked from the main thread
    /// with no concurrent interaction with the OCaml runtime. The returned
    /// `UnsafeOcamlPtr` is unrooted and could be invalidated if the GC is
    /// triggered after this method returns.
    pub unsafe fn get_ocaml_shallow_class(&self, name: &[u8]) -> Option<UnsafeOcamlPtr> {
        if self.shallow_decl_changes_store.classes.has_local_changes() { None }
        else { self.shallow_decl_changes_store.classes_shm.get_ocaml_by_byte_string(name) }
    }
    pub unsafe fn get_ocaml_typedef(&self, name: &[u8]) -> Option<UnsafeOcamlPtr> {
        if self.shallow_decl_changes_store.typedefs.has_local_changes() { None }
        else { self.shallow_decl_changes_store.typedefs_shm.get_ocaml_by_byte_string(name) }
    }
    pub unsafe fn get_ocaml_fun(&self, name: &[u8]) -> Option<UnsafeOcamlPtr> {
        if self.shallow_decl_changes_store.funs.has_local_changes() { None }
        else { self.shallow_decl_changes_store.funs_shm.get_ocaml_by_byte_string(name) }
    }
    pub unsafe fn get_ocaml_const(&self, name: &[u8]) -> Option<UnsafeOcamlPtr> {
        if self.shallow_decl_changes_store.consts.has_local_changes() { None }
        else { self.shallow_decl_changes_store.consts_shm.get_ocaml_by_byte_string(name) }
    }
    pub unsafe fn get_ocaml_module(&self, name: &[u8]) -> Option<UnsafeOcamlPtr> {
        if self.shallow_decl_changes_store.modules.has_local_changes() { None }
        else { self.shallow_decl_changes_store.modules_shm.get_ocaml_by_byte_string(name) }
    }

    pub unsafe fn get_ocaml_folded_class(&self, name: &[u8]) -> Option<UnsafeOcamlPtr> {
        if self.folded_classes_store.has_local_changes() { None }
        else { self.folded_classes_shm.get_ocaml_by_byte_string(name) }
    }
}

#[rustfmt::skip]
struct ShallowStoreWithChanges {
    classes:      Arc<ChangesStore <TypeName, Arc<decl::ShallowClass<BR>>>>,
    classes_shm:  Arc<OcamlShmStore<TypeName, Arc<decl::ShallowClass<BR>>>>,
    typedefs:     Arc<ChangesStore <TypeName, Arc<decl::TypedefDecl<BR>>>>,
    typedefs_shm: Arc<OcamlShmStore<TypeName, Arc<decl::TypedefDecl<BR>>>>,
    funs:         Arc<ChangesStore <pos::FunName, Arc<decl::FunDecl<BR>>>>,
    funs_shm:     Arc<OcamlShmStore<pos::FunName, Arc<decl::FunDecl<BR>>>>,
    consts:       Arc<ChangesStore <pos::ConstName, Arc<decl::ConstDecl<BR>>>>,
    consts_shm:   Arc<OcamlShmStore<pos::ConstName, Arc<decl::ConstDecl<BR>>>>,
    modules:      Arc<ChangesStore <pos::ModuleName, Arc<decl::ModuleDecl<BR>>>>,
    modules_shm:  Arc<OcamlShmStore<pos::ModuleName, Arc<decl::ModuleDecl<BR>>>>,
    store_view: Arc<ShallowDeclStore<BR>>,
}

impl ShallowStoreWithChanges {
    #[rustfmt::skip]
    fn new() -> Self {
        use shm_store::{Compression, Evictability::{Evictable, NonEvictable}};
        let classes_shm =  Arc::new(OcamlShmStore::new("Decl_ShallowClass", NonEvictable, Compression::default()));
        let typedefs_shm = Arc::new(OcamlShmStore::new("Decl_Typedef", Evictable, Compression::default()));
        let funs_shm =     Arc::new(OcamlShmStore::new("Decl_Fun", Evictable, Compression::default()));
        let consts_shm =   Arc::new(OcamlShmStore::new("Decl_GConst", Evictable, Compression::default()));
        let modules_shm =  Arc::new(OcamlShmStore::new("Decl_Module", Evictable, Compression::default()));

        let classes =  Arc::new(ChangesStore::new(Arc::clone(&classes_shm) as _));
        let typedefs = Arc::new(ChangesStore::new(Arc::clone(&typedefs_shm) as _));
        let funs =     Arc::new(ChangesStore::new(Arc::clone(&funs_shm) as _));
        let consts =   Arc::new(ChangesStore::new(Arc::clone(&consts_shm) as _));
        let modules =  Arc::new(ChangesStore::new(Arc::clone(&modules_shm) as _));

        let store_view = Arc::new(ShallowDeclStore::with_no_member_stores(
            Arc::clone(&classes) as _,
            Arc::clone(&typedefs) as _,
            Arc::clone(&funs) as _,
            Arc::clone(&consts) as _,
            Arc::clone(&modules) as _,
        ));
        Self {
            classes,
            typedefs,
            funs,
            consts,
            modules,
            classes_shm,
            typedefs_shm,
            funs_shm,
            consts_shm,
            modules_shm,
            store_view,
        }
    }

    fn push_local_changes(&self) {
        self.classes.push_local_changes();
        self.typedefs.push_local_changes();
        self.funs.push_local_changes();
        self.consts.push_local_changes();
        self.modules.push_local_changes();
    }

    fn pop_local_changes(&self) {
        self.classes.pop_local_changes();
        self.typedefs.pop_local_changes();
        self.funs.pop_local_changes();
        self.consts.pop_local_changes();
        self.modules.pop_local_changes();
    }

    fn as_shallow_decl_store(&self) -> Arc<ShallowDeclStore<BR>> {
        Arc::clone(&self.store_view)
    }

    fn to_old_key<K: AsRef<str> + From<String> + Copy>(key: K) -> K {
        format!("old${}", key.as_ref()).into()
    }

    fn oldify<K, V>(
        &self,
        store: &dyn Store<K, V>,
        names: &mut dyn Iterator<Item = K>,
    ) -> Result<()>
    where
        K: AsRef<str> + From<String> + Copy,
    {
        let mut moves = vec![];
        let mut deletes = vec![];
        for name in names {
            let old_name = Self::to_old_key(name);
            if store.contains_key(name)? {
                moves.push((name, old_name));
            } else if store.contains_key(old_name)? {
                deletes.push(old_name);
            }
        }
        store.move_batch(&mut moves.into_iter())?;
        store.remove_batch(&mut deletes.into_iter())?;
        Ok(())
    }

    fn remove_old<K, V>(
        &self,
        store: &dyn Store<K, V>,
        names: &mut dyn Iterator<Item = K>,
    ) -> Result<()>
    where
        K: AsRef<str> + From<String> + Copy,
    {
        let mut deletes = vec![];
        for name in names {
            let old_name = Self::to_old_key(name);
            if store.contains_key(old_name)? {
                deletes.push(old_name);
            }
        }
        store.remove_batch(&mut deletes.into_iter())
    }

    fn get_old<K, V>(
        &self,
        store: &dyn Store<K, V>,
        names: &mut dyn Iterator<Item = K>,
    ) -> Result<BTreeMap<K, Option<V>>>
    where
        K: AsRef<str> + From<String> + Copy + Ord,
    {
        names
            .map(|name| Ok((name, store.get(Self::to_old_key(name))?)))
            .collect()
    }

    #[rustfmt::skip]
    fn oldify_defs(&self, names: &file_info::Names) -> Result<()> {
        self.oldify(&*self.classes,  &mut names.classes.iter().map(Into::into))?;
        self.oldify(&*self.funs,     &mut names.funs.iter().map(Into::into))?;
        self.oldify(&*self.typedefs, &mut names.types.iter().map(Into::into))?;
        self.oldify(&*self.consts,   &mut names.consts.iter().map(Into::into))?;
        self.oldify(&*self.modules,  &mut names.modules.iter().map(Into::into))?;
        Ok(())
    }

    #[rustfmt::skip]
    fn remove_old_defs(&self, names: &file_info::Names) -> Result<()> {
        self.remove_old(&*self.classes,  &mut names.classes.iter().map(Into::into))?;
        self.remove_old(&*self.funs,     &mut names.funs.iter().map(Into::into))?;
        self.remove_old(&*self.typedefs, &mut names.types.iter().map(Into::into))?;
        self.remove_old(&*self.consts,   &mut names.consts.iter().map(Into::into))?;
        self.remove_old(&*self.modules,  &mut names.modules.iter().map(Into::into))?;
        Ok(())
    }

    #[rustfmt::skip]
    fn remove_defs(&self, names: &file_info::Names) -> Result<()> {
        self.classes .remove_batch(&mut names.classes.iter().map(Into::into))?;
        self.funs    .remove_batch(&mut names.funs.iter().map(Into::into))?;
        self.typedefs.remove_batch(&mut names.types.iter().map(Into::into))?;
        self.consts  .remove_batch(&mut names.consts.iter().map(Into::into))?;
        self.modules .remove_batch(&mut names.modules.iter().map(Into::into))?;
        Ok(())
    }

    #[rustfmt::skip]
    fn get_old_defs(
        &self,
        names: &file_info::Names,
    ) -> Result<(
        BTreeMap<pos::TypeName, Option<Arc<decl::ShallowClass<BR>>>>,
        BTreeMap<pos::FunName, Option<Arc<decl::FunDecl<BR>>>>,
        BTreeMap<pos::TypeName, Option<Arc<decl::TypedefDecl<BR>>>>,
        BTreeMap<pos::ConstName, Option<Arc<decl::ConstDecl<BR>>>>,
        BTreeMap<pos::ModuleName, Option<Arc<decl::ModuleDecl<BR>>>>,
    )> {
        Ok((
            self.get_old(&*self.classes,  &mut names.classes.iter().map(Into::into))?,
            self.get_old(&*self.funs,     &mut names.funs.iter().map(Into::into))?,
            self.get_old(&*self.typedefs, &mut names.types.iter().map(Into::into))?,
            self.get_old(&*self.consts,   &mut names.consts.iter().map(Into::into))?,
            self.get_old(&*self.modules,  &mut names.modules.iter().map(Into::into))?,
        ))
    }
}

#[derive(Clone, Debug, ToOcamlRep, FromOcamlRep)]
#[derive(serde::Serialize, serde::Deserialize)]
pub enum FileType {
    Disk(bstr::BString),
    Ide(bstr::BString),
}

#[derive(Debug)]
struct FileProviderWithChanges {
    // We could use DeltaStore here if not for the fact that the OCaml
    // implementation of `File_provider.get` does not fall back to disk when the
    // given path isn't present in sharedmem/local_changes (it only does so for
    // `File_provider.get_contents`).
    delta_and_changes: Arc<ChangesStore<RelativePath, FileType>>,
    disk: DiskProvider,
}

impl FileProvider for FileProviderWithChanges {
    fn get(&self, file: RelativePath) -> Result<bstr::BString> {
        match self.delta_and_changes.get(file)? {
            Some(FileType::Disk(contents)) => Ok(contents),
            Some(FileType::Ide(contents)) => Ok(contents),
            None => match self.disk.read(file) {
                Ok(contents) => Ok(contents),
                Err(e) if e.kind() == std::io::ErrorKind::NotFound => Ok("".into()),
                Err(e) => Err(e.into()),
            },
        }
    }
}
