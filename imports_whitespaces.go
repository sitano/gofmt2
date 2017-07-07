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

		/*

			Case:

			import (

				x
			)

		*/
		for fset.Position(d.Specs[0].Pos()).Line > fset.Position(d.Pos()).Line+1 {
			fset.File(d.Pos()).MergeLine(fset.Position(d.Pos()).Line + 1)
		}

		/*

			Case:

			import (
				a

				b
			)

		*/
		i := 0
		for j, s := range d.Specs {
			if j > i {
				prev := fset.Position(d.Specs[j-1].Pos()).Line
				cur := fset.Position(s.Pos()).Line
				if cur > 1+prev {
					// j begins a new run. End this one.
					for k := prev + 1; k < cur; k++ {
						fset.File(s.Pos()).MergeLine(prev + 1)
					}
					i = j
				}
			}
		}

		/*

			Case:

			import (
				a

			)

		*/
		if len(d.Specs) > 0 {
			lastSpec := d.Specs[len(d.Specs)-1]
			lastLine := fset.Position(lastSpec.Pos()).Line
			rParenLine := fset.Position(d.Rparen).Line
			for rParenLine > lastLine+1 {
				rParenLine--
				fset.File(d.Rparen).MergeLine(rParenLine)
			}
		}
	}
}
