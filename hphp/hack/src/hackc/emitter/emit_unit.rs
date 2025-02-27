// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.
mod emit_adata;
mod emit_attribute;
mod emit_body;
mod emit_class;
mod emit_constant;
mod emit_expression;
mod emit_fatal;
mod emit_file_attributes;
mod emit_function;
mod emit_memoize_function;
mod emit_memoize_helpers;
mod emit_memoize_method;
mod emit_method;
mod emit_module;
mod emit_native_opcode;
mod emit_param;
mod emit_property;
mod emit_statement;
mod emit_type_constant;
mod emit_typedef;
mod emit_xhp;
mod generator;
mod reified_generics_helpers;
mod try_finally_rewriter;
mod xhp_attribute;

use emit_class::emit_classes_from_program;
use emit_constant::emit_constants_from_program;
use emit_file_attributes::emit_file_attributes_from_program;
use emit_function::emit_functions_from_program;
use emit_module::emit_module_use_from_program;
use emit_module::emit_modules_from_program;
use emit_typedef::emit_typedefs_from_program;
use env::emitter::Emitter;
use env::Env;
use error::Error;
use error::ErrorKind;
use error::Result;
use ffi::Maybe::*;
use ffi::Slice;
use ffi::Str;
use hhbc::Fatal;
use hhbc::FatalOp;
use hhbc::Unit;
use ocamlrep::rc::RcOc;
use oxidized::ast;
use oxidized::namespace_env;
use oxidized::pos::Pos;

// PUBLIC INTERFACE (ENTRY POINTS)

/// This is the entry point from hh_single_compile & fuzzer
pub fn emit_fatal_unit<'arena>(
    alloc: &'arena bumpalo::Bump,
    op: FatalOp,
    pos: Pos,
    msg: impl AsRef<str> + 'arena,
) -> Result<Unit<'arena>> {
    Ok(Unit {
        fatal: Just(Fatal {
            op,
            loc: pos.into(),
            message: Str::new_str(alloc, msg.as_ref()),
        }),
        ..Unit::default()
    })
}

/// This is the entry point from hh_single_compile
pub fn emit_unit<'a, 'arena, 'decl>(
    emitter: &mut Emitter<'arena, 'decl>,
    namespace: RcOc<namespace_env::Env>,
    tast: &'a ast::Program,
) -> Result<Unit<'arena>> {
    let result = emit_unit_(emitter, namespace, tast);
    match result {
        Err(e) => match e.into_kind() {
            ErrorKind::IncludeTimeFatalException(op, pos, msg) => {
                emit_fatal_unit(emitter.alloc, op, pos, msg)
            }
            ErrorKind::Unrecoverable(x) => Err(Error::unrecoverable(x)),
        },
        _ => result,
    }
}

fn emit_unit_<'a, 'arena, 'decl>(
    emitter: &mut Emitter<'arena, 'decl>,
    namespace: RcOc<namespace_env::Env>,
    prog: &'a ast::Program,
) -> Result<Unit<'arena>> {
    let prog = prog.as_slice();
    let mut functions = emit_functions_from_program(emitter, prog)?;
    let classes = emit_classes_from_program(emitter.alloc, emitter, prog)?;
    let modules = emit_modules_from_program(emitter.alloc, emitter, prog)?;
    let typedefs = emit_typedefs_from_program(emitter, prog)?;
    let (constants, mut const_inits) = {
        let mut env = Env::default(emitter.alloc, namespace);
        emit_constants_from_program(emitter, &mut env, prog)?
    };
    functions.append(&mut const_inits);
    let file_attributes = emit_file_attributes_from_program(emitter, prog)?;
    let adata = emitter.adata_state_mut().take_adata();
    let module_use = emit_module_use_from_program(emitter, prog);
    let symbol_refs = emitter.finish_symbol_refs();
    let fatal = Nothing;

    Ok(Unit {
        classes: Slice::fill_iter(emitter.alloc, classes.into_iter()),
        modules: Slice::fill_iter(emitter.alloc, modules.into_iter()),
        functions: Slice::fill_iter(emitter.alloc, functions.into_iter()),
        typedefs: Slice::fill_iter(emitter.alloc, typedefs.into_iter()),
        constants: Slice::fill_iter(emitter.alloc, constants.into_iter()),
        adata: Slice::fill_iter(emitter.alloc, adata.into_iter()),
        file_attributes: Slice::fill_iter(emitter.alloc, file_attributes.into_iter()),
        module_use,
        symbol_refs,
        fatal,
    })
}

#[derive(Clone, Copy, Debug)]
enum TypeRefinementInHint {
    Allowed,
    Disallowed,
}
