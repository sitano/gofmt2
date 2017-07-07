package main

import (
	"go/ast"
	"go/token"
)

// JoinImports joins all top import decls
func JoinImports(fset *token.FileSet, f *ast.File) {
	var fst *ast.GenDecl

	for i := 0; i < len(f.Decls); i++ {
		d, ok := (f.Decls[i]).(*ast.GenDecl)
		if !ok || d.Tok != token.IMPORT {
			// Not an import declaration, so we're done.
			// Imports are always first.
			break
		}

		// Select first import block
		if fst == nil {
			fst = d
			continue
		}

		// Append docs data to the first import block
		if d.Doc != nil {
			if fst.Doc != nil {
				fst.Doc.List = append(fst.Doc.List, d.Doc.List...)
			} else {
				fst.Doc = d.Doc
			}
		}

		// Append specs
		fst.Specs = append(fst.Specs, d.Specs...)
		f.Decls = append(f.Decls[:i], f.Decls[i+1:]...)

		i--
	}
}
