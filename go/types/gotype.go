// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// Build this command explicitly: go build gotype.go

/*
The gotype command, like the front-end of a Go compiler, parses and
type-checks a single Go package. Errors are reported if the analysis
fails; otherwise gotype is quiet (unless -v is set).

Without a list of paths, gotype reads from standard input, which
must provide a single Go source file defining a complete package.

With a single directory argument, gotype checks the Go files in
that directory, comprising a single package. Use -t to include the
(in-package) _test.go files. Use -x to type check only external
test files.

Otherwise, each path must be the filename of a Go file belonging
to the same package.

Imports are processed by importing directly from the source of
imported packages (default), or by importing from compiled and
installed packages (by setting -c to the respective compiler).

The -c flag must be set to a compiler ("gc", "gccgo") when type-
checking packages containing imports with relative import paths
(import "./mypkg") because the source importer cannot know which
files to include for such packages.

Usage:

	gotype [flags] [path...]

The flags are:

	-t
		include local test files in a directory (ignored if -x is provided)
	-x
		consider only external test files in a directory
	-e
		report all errors (not just the first 10)
	-v
		verbose mode
	-c
		compiler used for installed packages (gc, gccgo, or source); default: source

Flags controlling additional output:

	-ast
		print AST (forces -seq)
	-trace
		print parse trace (forces -seq)
	-comments
		parse comments (ignored unless -ast or -trace is provided)
	-panic
		panic on first error

Examples:

To check the files a.go, b.go, and c.go:

	gotype a.go b.go c.go

To check an entire package including (in-package) tests in the directory dir and print the processed files:

	gotype -t -v dir

To check the external test package (if any) in the current directory, based on installed packages compiled with
cmd/compile:

	gotype -c=gc -x .

To verify the output of a pipe:

	echo "package foo" | gotype
*/
package main
