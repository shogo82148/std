// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

// A RawConn is a raw network connection.
type RawConn interface {
	Control(f func(fd uintptr)) error

	Read(f func(fd uintptr) (done bool)) error

	Write(f func(fd uintptr) (done bool)) error
}

// Conn is implemented by some types in the net package to provide
// access to the underlying file descriptor or handle.
type Conn interface {
	SyscallConn() (RawConn, error)
}
