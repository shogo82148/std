// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// Mkoptions generates the lists of the X86 features that each GOAMD64 level
// requires: goamd64_gen.go, which lists their GODEBUG options, and the
// compiler's cmd/compile/internal/ssacompile/goamd64_gen.go, which lists the
// fields whose loads the compiler can replace with true. Generating both from
// one table keeps them in sync.
package main
