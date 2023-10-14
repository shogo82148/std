// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package generate implements the “go generate” command.
package generate

import (
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/cmd/go/internal/base"
)

var CmdGenerate = &base.Command{
	Run:       runGenerate,
	UsageLine: "go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]",
	Short:     "generate Go files by processing source",
	Long: `
Generate runs commands described by directives within existing
files. Those commands can run any process but the intent is to
create or update Go source files.

Go generate is never run automatically by go build, go test,
and so on. It must be run explicitly.

Go generate scans the file for directives, which are lines of
the form,

	//go:generate command argument...

(note: no leading spaces and no space in "//go") where command
is the generator to be run, corresponding to an executable file
that can be run locally. It must either be in the shell path
(gofmt), a fully qualified path (/usr/you/bin/mytool), or a
command alias, described below.

Note that go generate does not parse the file, so lines that look
like directives in comments or multiline strings will be treated
as directives.

The arguments to the directive are space-separated tokens or
double-quoted strings passed to the generator as individual
arguments when it is run.

Quoted strings use Go syntax and are evaluated before execution; a
quoted string appears as a single argument to the generator.

To convey to humans and machine tools that code is generated,
generated source should have a line that matches the following
regular expression (in Go syntax):

	^// Code generated .* DO NOT EDIT\.$

This line must appear before the first non-comment, non-blank
text in the file.

Go generate sets several variables when it runs the generator:

	$GOARCH
		The execution architecture (arm, amd64, etc.)
	$GOOS
		The execution operating system (linux, windows, etc.)
	$GOFILE
		The base name of the file.
	$GOLINE
		The line number of the directive in the source file.
	$GOPACKAGE
		The name of the package of the file containing the directive.
	$GOROOT
		The GOROOT directory for the 'go' command that invoked the
		generator, containing the Go toolchain and standard library.
	$DOLLAR
		A dollar sign.
	$PATH
		The $PATH of the parent process, with $GOROOT/bin
		placed at the beginning. This causes generators
		that execute 'go' commands to use the same 'go'
		as the parent 'go generate' command.

Other than variable substitution and quoted-string evaluation, no
special processing such as "globbing" is performed on the command
line.

As a last step before running the command, any invocations of any
environment variables with alphanumeric names, such as $GOFILE or
$HOME, are expanded throughout the command line. The syntax for
variable expansion is $NAME on all operating systems. Due to the
order of evaluation, variables are expanded even inside quoted
strings. If the variable NAME is not set, $NAME expands to the
empty string.

A directive of the form,

	//go:generate -command xxx args...

specifies, for the remainder of this source file only, that the
string xxx represents the command identified by the arguments. This
can be used to create aliases or to handle multiword generators.
For example,

	//go:generate -command foo go tool foo

specifies that the command "foo" represents the generator
"go tool foo".

Generate processes packages in the order given on the command line,
one at a time. If the command line lists .go files from a single directory,
they are treated as a single package. Within a package, generate processes the
source files in a package in file name order, one at a time. Within
a source file, generate runs generators in the order they appear
in the file, one at a time. The go generate tool also sets the build
tag "generate" so that files may be examined by go generate but ignored
during build.

For packages with invalid code, generate processes only source files with a
valid package clause.

If any generator returns an error exit status, "go generate" skips
all further processing for that package.

The generator is run in the package's source directory.

Go generate accepts two specific flags:

	-run=""
		if non-empty, specifies a regular expression to select
		directives whose full original source text (excluding
		any trailing spaces and final newline) matches the
		expression.

	-skip=""
		if non-empty, specifies a regular expression to suppress
		directives whose full original source text (excluding
		any trailing spaces and final newline) matches the
		expression. If a directive matches both the -run and
		the -skip arguments, it is skipped.

It also accepts the standard build flags including -v, -n, and -x.
The -v flag prints the names of packages and files as they are
processed.
The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.
	`,
}

// A Generator represents the state of a single Go source file
// being scanned for generator commands.
type Generator struct {
	r        io.Reader
	path     string
	dir      string
	file     string
	pkg      string
	commands map[string][]string
	lineNum  int
	env      []string
}
