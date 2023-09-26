// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd
// +build darwin dragonfly freebsd linux netbsd openbsd

package filepath

// IsAbs returns true if the path is absolute.
func IsAbs(path string) bool

// HasPrefix exists for historical compatibility and should not be used.
func HasPrefix(p, prefix string) bool
