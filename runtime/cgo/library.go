// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.cgobindings

package cgo

import (
	"github.com/shogo82148/std/unsafe"
)

// Library is a C shared library loaded at run time.
//
// Library is meant to be embedded in a struct type declared with a
// //cgo:dynload pragma, whose methods, bound with //cgo:func and //cgo:var
// pragmas, call functions and access variables in the library. See the
// cgo command documentation for details.
//
// The zero value of Library is not usable; use [Open] to load a library.
// For now, a library is never unloaded.
type Library struct {
	handle uintptr

	// syms holds the addresses of the symbols that cmd/cgo registered for
	// the type embedding the Library, indexed by a per-type symbol index
	// assigned by cmd/cgo. Open resolves all of them before it returns, and
	// they are never written again.
	syms []unsafe.Pointer
}

// Open loads the C shared library at path and returns a new *T for it.
//
// T must be a struct type that embeds Library, normally one declared with
// a //cgo:dynload pragma. Open resolves all of the symbols bound to T's
// methods, and fails if any of them is missing from the library. Each call
// to Open returns a new instance, with its own symbol table, even if the
// same library is opened more than once. On Unix, the library is loaded
// with RTLD_NOW | RTLD_LOCAL, so its symbols are bound eagerly and are not
// added to the global symbol namespace.
// For now, libraries are never unloaded.
func Open[T any, PT interface {
	*T
	library() *Library
}](path string) (PT, error)
