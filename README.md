# gofmt.v2
strict version of gofmt improved and idempotent

Desc
====

Gofmt formats Go programs. This is improved version.

Features
====

- remove empty lines between import specs in block (-import\_whitespaces enabled by default)
- join all import blocks (-join\_imports)
- fix lines
- show ast (-ast)

Code style changes to original `gofmt`
====

`func` always separated by an empty line if its not sequential block of one-liners:

```
func a() {}
func b() {}
                <-- added
func c() {
    //
}
                <-- added
func d() {}

```

empty blocks are not allowed on multiple lines:

```
func a() {

} -> func a() {}

func b() {
} -> func b() {}

for {
} -> for {}

switch {
} -> switch {}
```

no empty lines are allowed before closing right bracket and after opening left bracket:

```
func a() int {
                <-- would be deleted
    if {
                <-- would be deleted
        fmt.Print(1)
                <-- would be deleted
    }

    return 1
                <-- would be deleted
}
```

by default imports in a single block are concatenated.

Usage:
====

```

	gofmt [flags] [path ...]

The flags are:
	-d
		Do not print reformatted sources to standard output.
		If a file's formatting is different than gofmt's, print diffs
		to standard output.
	-e
		Print all (including spurious) errors.
	-l
		Do not print reformatted sources to standard output.
		If a file's formatting is different from gofmt's, print its name
		to standard output.
	-r rule
		Apply the rewrite rule to the source before reformatting.
	-s
		Try to simplify code (after applying the rewrite rule, if any).
	-w
		Do not print reformatted sources to standard output.
		If a file's formatting is different from gofmt's, overwrite it
		with gofmt's version. If an error occurred during overwriting,
		the original file is restored from an automatic backup.
	-import_whitespaces
		Remove empty lines in import blocks (default true).
	-join_imports
		Join imports blocks (default false).
	-ast
		Print AST.
``````

TODO
====

- rewrite expressions in which `{rvalue ? lvalue} -> {lvalue ? rvalue}`
