// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !386 && !amd64
// +build !386,!amd64

// SHA256 block step.
// In its own file so that a faster assembly or C version
// can be substituted easily.

package sha256
