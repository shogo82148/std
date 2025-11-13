// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Annotate Ref in Prog with C types by parsing gcc debug output.
// Conversion of debug output to Go types.

package main

// ProcessCgoDirectives processes the import C preamble:
//  1. discards all #cgo CFLAGS, LDFLAGS, nocallback and noescape directives,
//     so they don't make their way into _cgo_export.h.
//  2. parse the nocallback and noescape directives.
func (f *File) ProcessCgoDirectives()

// Translate rewrites f.AST, the original Go input, to remove
// references to the imported package C, replacing them with
// references to the equivalent Go types, functions, and variables.
// Preconditions: File.loadDebug must be called prior to translate.
func (p *Package) Translate(f *File)

// String returns the current type representation. Format arguments
// are assembled within this method so that any changes in mutable
// values are taken into account.
func (tr *TypeRepr) String() string

// Empty reports whether the result of String would be "".
func (tr *TypeRepr) Empty() bool

// Set modifies the type representation.
// If fargs are provided, repr is used as a format for fmt.Sprintf.
// Otherwise, repr is used unprocessed as the type representation.
func (tr *TypeRepr) Set(repr string, fargs ...any)
