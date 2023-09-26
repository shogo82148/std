// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || dragonfly || freebsd || (js && wasm) || linux || netbsd || openbsd || solaris
// +build aix dragonfly freebsd js,wasm linux netbsd openbsd solaris

package x509

// Possible directories with certificate files; stop after successfully
// reading at least one file from a directory.
