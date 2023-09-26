// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Map predeclared Go types to Type.

// gccExportHeaderEpilog goes at the end of the generated header file.

// gccgoExportFileProlog is written to the _cgo_export.c file when
// using gccgo.
// We use weak declarations, and test the addresses, so that this code
// works with older versions of gccgo.
