package main

import (
	"go/ast"
	"go/token"
)

// FixLines fixes empty and missed lines in the source file
func FixLines(fset *token.FileSet, f *ast.File) {
	for _, d := range f.Decls {
		d, ok := d.(*ast.FuncDecl)
		if !ok || d.Body == nil {
			continue
		}

		if fset.Position(d.Body.Lbrace).Line >= fset.Position(d.Body.Rbrace).Line {
			continue
		}

		if len(d.Body.List) > 0 {
			// stmt\n\n}
			s := d.Body.List[len(d.Body.List)-1]
			send := fset.Position(s.End()).Line
			rl := fset.Position(d.Body.Rbrace).Line
			if send+1 < rl {
				// [send+1 .. rl)
				for k := rl - 1; k >= send+1; k-- {
					if !isLineComment(fset, f, k) {
						fset.File(d.Body.Pos()).MergeLine(k)
					} else {
						break
					}
				}
			}
			// {\n\nstmt
			s = d.Body.List[0]
			sstart := fset.Position(s.Pos()).Line
			ll := fset.Position(d.Body.Lbrace).Line
			if ll+1 < sstart {
				// [ll + 1 .. sstart)
				for k := sstart - 1; k >= ll+1; k-- {
					if !isLineComment(fset, f, k) {
						fset.File(d.Body.Pos()).MergeLine(k)
					} else {
						break
					}
				}
			}
		} else {
			ll := fset.Position(d.Body.Lbrace).Line
			rl := fset.Position(d.Body.Rbrace).Line
			// [ll + 1 .. rl)
			for k := rl - 1; k >= ll+1; k-- {
				if !isLineComment(fset, f, k) {
					fset.File(d.Body.Pos()).MergeLine(k)
				} else {
					break
				}
			}
		}
	}
}

func isLineComment(fset *token.FileSet, f *ast.File, l int) bool {
	for _, c := range f.Comments {
		if l < fset.Position(c.Pos()).Line {
			break
		}
		if l >= fset.Position(c.Pos()).Line && l <= fset.Position(c.Pos()).Line {
			return true
		}
	}
	return false
}
