// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// runtime·duffzero is a Duff's device for zeroing memory.
// The compiler jumps to computed addresses within
// the routine to zero chunks of memory.
// Do not change duffzero without also
// changing clearfat in cmd/?g/ggen.go.

// runtime·duffcopy is a Duff's device for copying memory.
// The compiler jumps to computed addresses within
// the routine to copy chunks of memory.
// Source and destination must not overlap.
// Do not change duffcopy without also
// changing blockcopy in cmd/?g/cgen.go.

// See the zero* and copy* generators below
// for architecture-specific comments.

// mkduff generates duff_*.s.
package main
