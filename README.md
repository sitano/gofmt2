# gofmt.v2
strict version of gofmt improved and idempotent

Desc
====

Gofmt formats Go programs. This is improved version.

Features
====

- remove empty lines between import specs in block (-import\_whitespaces enabled by default)
- join all import blocks (-join\_imports)
- show ast

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
``````
