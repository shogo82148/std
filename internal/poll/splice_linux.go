// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

// Splice transfers at most remain bytes of data from src to dst, using the
// splice system call to minimize copies of data from and to userspace.
//
// Splice gets a pipe buffer from the pool or creates a new one if needed, to serve as a buffer for the data transfer.
// src and dst must both be stream-oriented sockets.
//
// If err != nil, sc is the system call which caused the error.
func Splice(dst, src *FD, remain int64) (written int64, handled bool, sc string, err error)
