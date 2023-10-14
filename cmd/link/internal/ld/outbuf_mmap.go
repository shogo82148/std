// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package ld

// Mmap maps the output file with the given size. It unmaps the old mapping
// if it is already mapped. It also flushes any in-heap data to the new
// mapping.
func (out *OutBuf) Mmap(filesize uint64) (err error)
