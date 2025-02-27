(*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *)
open Hh_prelude
module SN = Naming_special_names

let on_expr_ (env, expr_, err_acc) =
  let err =
    match expr_ with
    | Aast.(Cast ((_, Hprim (Tint | Tbool | Tfloat | Tstring)), _)) -> err_acc
    | Aast.(Cast ((_, Happly ((_, tycon_nm), _)), _))
      when String.(
             equal tycon_nm SN.Collections.cDict
             || equal tycon_nm SN.Collections.cVec) ->
      err_acc
    | Aast.(Cast ((_, Aast.Hvec_or_dict (_, _)), _)) -> err_acc
    | Aast.(Cast ((_, Aast.Hany), _)) ->
      (* We end up with a `Hany` when we have an arity error for dict/vec
         - we don't error on this case to preserve behaviour
      *)
      err_acc
    | Aast.(Cast ((pos, _), _)) ->
      (Naming_phase_error.naming @@ Naming_error.Object_cast pos) :: err_acc
    | _ -> err_acc
  in
  Ok (env, expr_, err)

let pass =
  Naming_phase_pass.(
    bottom_up Ast_transform.{ identity with on_expr_ = Some on_expr_ })
