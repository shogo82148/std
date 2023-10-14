// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !aix && !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !(solaris && go1.20) && !windows

package ld

// Mmap allocates an in-heap output buffer with the given size. It copies
// any old data (if any) to the new buffer.
func (out *OutBuf) Mmap(filesize uint64) error
