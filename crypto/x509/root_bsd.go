// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || freebsd || netbsd || openbsd
// +build dragonfly freebsd netbsd openbsd

package x509

// Possible certificate files; stop after finding one.

// Possible directories with certificate files; stop after successfully
// reading at least one file from a directory.
