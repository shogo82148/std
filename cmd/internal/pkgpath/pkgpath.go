// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pkgpath determines the package path used by gccgo/GoLLVM symbols.
// This package is not used for the gc compiler.
package pkgpath

// ToSymbolFunc returns a function that may be used to convert a
// package path into a string suitable for use as a symbol.
// cmd is the gccgo/GoLLVM compiler in use, and tmpdir is a temporary
// directory to pass to os.CreateTemp().
// For example, this returns a function that converts "net/http"
// into a string like "net..z2fhttp". The actual string varies for
// different gccgo/GoLLVM versions, which is why this returns a function
// that does the conversion appropriate for the compiler in use.
func ToSymbolFunc(cmd, tmpdir string) (func(string) string, error)
