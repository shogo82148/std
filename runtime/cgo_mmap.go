// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Support for memory sanitizer.  See runtime/cgo/mmap.go.

//go:build linux && amd64
// +build linux,amd64

package runtime

// _cgo_mmap is filled in by runtime/cgo when it is linked into the
// program, so it is only non-nil when using cgo.
//go:linkname _cgo_mmap _cgo_mmap
