// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parse input AST and prepare Prog structure.

package main

// ParseGo populates f with information learned from the Go source code
// which was read from the named file. It gathers the C preamble
// attached to the import "C" comment, a list of references to C.xxx,
// a list of exported functions, and the actual AST, to be rewritten and
// printed.
func (f *File) ParseGo(abspath string, src []byte)
