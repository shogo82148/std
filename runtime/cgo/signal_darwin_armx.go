// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin && (arm || arm64)
// +build darwin
// +build arm arm64

package cgo

//go:cgo_import_static x_cgo_panicmem
//go:linkname x_cgo_panicmem x_cgo_panicmem

// use a pointer to avoid relocation of external symbol in __TEXT
// make linker happy
