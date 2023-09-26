// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Support for sanitizers. See runtime/cgo/sigaction.go.

//go:build (linux && amd64) || (freebsd && amd64) || (linux && arm64) || (linux && ppc64le)

package runtime

// _cgo_sigaction is filled in by runtime/cgo when it is linked into the
// program, so it is only non-nil when using cgo.
//
//go:linkname _cgo_sigaction _cgo_sigaction
