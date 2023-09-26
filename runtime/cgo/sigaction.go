// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (linux && amd64) || (freebsd && amd64) || (linux && arm64)
// +build linux,amd64 freebsd,amd64 linux,arm64

package cgo

// Import "unsafe" because we use go:linkname.
import _ "github.com/shogo82148/std/unsafe"

//go:cgo_import_static x_cgo_sigaction
//go:linkname x_cgo_sigaction x_cgo_sigaction
//go:linkname _cgo_sigaction _cgo_sigaction
