// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.

use std::fmt::Display;

use proc_macro2::TokenStream;
use proc_macro2::TokenTree;

fn mismatch(ta: Option<TokenTree>, tb: Option<TokenTree>, ax: &TokenStream, bx: &TokenStream) -> ! {
    panic!(
        "Mismatch!\nLeft:  {}\nRight: {}\nwhen comparing\nLeft:  {}\nRight: {}\n",
        ta.map_or("None".into(), |t| t.to_string()),
        tb.map_or("None".into(), |t| t.to_string()),
        ax,
        bx
    );
}

pub fn assert_pat_eq<E: Display>(a: Result<TokenStream, E>, b: TokenStream) {
    let a = match a {
        Err(err) => {
            panic!("Unexpected error '{}'", err);
        }
        Ok(ok) => ok,
    };

    fn inner_cmp(a: TokenStream, b: TokenStream, ax: &TokenStream, bx: &TokenStream) {
        let mut ia = a.into_iter();
        let mut ib = b.into_iter();

        loop {
            let t_a = ia.next();
            let t_b = ib.next();
            if t_a.is_none() && t_b.is_none() {
                break;
            }
            if t_a.is_none() || t_b.is_none() {
                mismatch(t_a, t_b, ax, bx);
            }
            let t_a = t_a.unwrap();
            let t_b = t_b.unwrap();

            match (&t_a, &t_b) {
                (TokenTree::Ident(a), TokenTree::Ident(b)) if *a == b.to_string() => {}
                (TokenTree::Literal(a), TokenTree::Literal(b))
                    if a.to_string() == b.to_string() => {}
                (TokenTree::Punct(a), TokenTree::Punct(b)) if a.to_string() == b.to_string() => {}
                (TokenTree::Group(ga), TokenTree::Group(gb))
                    if ga.delimiter() == gb.delimiter() =>
                {
                    inner_cmp(ga.stream(), gb.stream(), ax, bx);
                }
                (TokenTree::Group(_), _)
                | (TokenTree::Ident(_), _)
                | (TokenTree::Punct(_), _)
                | (TokenTree::Literal(_), _) => mismatch(Some(t_a), Some(t_b), ax, bx),
            }
        }
    }

    inner_cmp(a.clone(), b.clone(), &a, &b);
}

pub fn assert_error<E: Display>(a: Result<TokenStream, E>, b: &str) {
    match a {
        Ok(a) => panic!("Expected error but got:\n{}", a),
        Err(e) => {
            let a = format!("{}", e);
            assert_eq!(a, b, "Incorrect error")
        }
    }
}
