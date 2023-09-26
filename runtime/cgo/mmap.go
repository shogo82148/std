// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (linux && amd64) || (linux && arm64) || (freebsd && amd64)

package cgo

// Import "unsafe" because we use go:linkname.
import _ "github.com/shogo82148/std/unsafe"

//go:cgo_import_static x_cgo_mmap
//go:linkname x_cgo_mmap x_cgo_mmap
//go:linkname _cgo_mmap _cgo_mmap

//go:cgo_import_static x_cgo_munmap
//go:linkname x_cgo_munmap x_cgo_munmap
//go:linkname _cgo_munmap _cgo_munmap
