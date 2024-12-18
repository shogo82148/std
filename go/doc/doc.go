// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package doc extracts source code documentation from a Go AST.
package doc

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/doc/comment"
	"github.com/shogo82148/std/go/token"
)

// Package is the documentation for an entire package.
type Package struct {
	Doc        string
	Name       string
	ImportPath string
	Imports    []string
	Filenames  []string
	Notes      map[string][]*Note

	// Deprecated: For backward compatibility Bugs is still populated,
	// but all new code should use Notes instead.
	Bugs []string

	// declarations
	Consts []*Value
	Types  []*Type
	Vars   []*Value
	Funcs  []*Func

	// Examples is a sorted list of examples associated with
	// the package. Examples are extracted from _test.go files
	// provided to NewFromFiles.
	Examples []*Example

	importByName map[string]string
	syms         map[string]bool
}

// Value is the documentation for a (possibly grouped) var or const declaration.
type Value struct {
	Doc   string
	Names []string
	Decl  *ast.GenDecl

	order int
}

// Type is the documentation for a type declaration.
type Type struct {
	Doc  string
	Name string
	Decl *ast.GenDecl

	// associated declarations
	Consts  []*Value
	Vars    []*Value
	Funcs   []*Func
	Methods []*Func

	// Examples is a sorted list of examples associated with
	// this type. Examples are extracted from _test.go files
	// provided to NewFromFiles.
	Examples []*Example
}

// Func is the documentation for a func declaration.
type Func struct {
	Doc  string
	Name string
	Decl *ast.FuncDecl

	// methods
	// (for functions, these fields have the respective zero value)
	Recv  string
	Orig  string
	Level int

	// Examples is a sorted list of examples associated with this
	// function or method. Examples are extracted from _test.go files
	// provided to NewFromFiles.
	Examples []*Example
}

// A Note represents a marked comment starting with "MARKER(uid): note body".
// Any note with a marker of 2 or more upper case [A-Z] letters and a uid of
// at least one character is recognized. The ":" following the uid is optional.
// Notes are collected in the Package.Notes map indexed by the notes marker.
type Note struct {
	Pos, End token.Pos
	UID      string
	Body     string
}

// Mode values control the operation of [New] and [NewFromFiles].
type Mode int

const (
	// AllDecls says to extract documentation for all package-level
	// declarations, not just exported ones.
	AllDecls Mode = 1 << iota

	// AllMethods says to show all embedded methods, not just the ones of
	// invisible (unexported) anonymous fields.
	AllMethods

	// PreserveAST says to leave the AST unmodified. Originally, pieces of
	// the AST such as function bodies were nil-ed out to save memory in
	// godoc, but not all programs want that behavior.
	PreserveAST
)

// New computes the package documentation for the given package AST.
// New takes ownership of the AST pkg and may edit or overwrite it.
// To have the [Examples] fields populated, use [NewFromFiles] and include
// the package's _test.go files.
func New(pkg *ast.Package, importPath string, mode Mode) *Package

// NewFromFiles computes documentation for a package.
//
// The package is specified by a list of *ast.Files and corresponding
// file set, which must not be nil.
// NewFromFiles uses all provided files when computing documentation,
// so it is the caller's responsibility to provide only the files that
// match the desired build context. "go/build".Context.MatchFile can
// be used for determining whether a file matches a build context with
// the desired GOOS and GOARCH values, and other build constraints.
// The import path of the package is specified by importPath.
//
// Examples found in _test.go files are associated with the corresponding
// type, function, method, or the package, based on their name.
// If the example has a suffix in its name, it is set in the
// [Example.Suffix] field. [Examples] with malformed names are skipped.
//
// Optionally, a single extra argument of type [Mode] can be provided to
// control low-level aspects of the documentation extraction behavior.
//
// NewFromFiles takes ownership of the AST files and may edit them,
// unless the PreserveAST Mode bit is on.
func NewFromFiles(fset *token.FileSet, files []*ast.File, importPath string, opts ...any) (*Package, error)

// Parser returns a doc comment parser configured
// for parsing doc comments from package p.
// Each call returns a new parser, so that the caller may
// customize it before use.
func (p *Package) Parser() *comment.Parser

// Printer returns a doc comment printer configured
// for printing doc comments from package p.
// Each call returns a new printer, so that the caller may
// customize it before use.
func (p *Package) Printer() *comment.Printer

// HTML returns formatted HTML for the doc comment text.
//
// To customize details of the HTML, use [Package.Printer]
// to obtain a [comment.Printer], and configure it
// before calling its HTML method.
func (p *Package) HTML(text string) []byte

// Markdown returns formatted Markdown for the doc comment text.
//
// To customize details of the Markdown, use [Package.Printer]
// to obtain a [comment.Printer], and configure it
// before calling its Markdown method.
func (p *Package) Markdown(text string) []byte

// Text returns formatted text for the doc comment text,
// wrapped to 80 Unicode code points and using tabs for
// code block indentation.
//
// To customize details of the formatting, use [Package.Printer]
// to obtain a [comment.Printer], and configure it
// before calling its Text method.
func (p *Package) Text(text string) []byte
