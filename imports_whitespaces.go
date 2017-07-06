package main

import (
	"go/ast"
	"go/token"
)

// RemoveImportWhitespaces removes empty lines between imports specs
func RemoveImportWhitespaces(fset *token.FileSet, f *ast.File) {
	for _, d := range f.Decls {
		d, ok := d.(*ast.GenDecl)
		if !ok || d.Tok != token.IMPORT {
			// Not an import declaration, so we're done.
			// Imports are always first.
			break
		}

		if !d.Lparen.IsValid() {
			// Not a block: sorted by default.
			continue
		}

		// Identify and sort runs of specs on successive lines.
		i := 0
		for j, s := range d.Specs {
			if j > i {
				prev := fset.Position(d.Specs[j-1].End()).Line
				cur := fset.Position(s.Pos()).Line
				if cur > 1+prev {
					// j begins a new run. End this one.
					for k := prev + 1; k < cur; k++ {
						fset.File(s.Pos()).MergeLine(k)
					}
					i = j
				}
			}
		}
	}
}
