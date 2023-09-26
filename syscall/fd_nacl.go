// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// File descriptor support for Native Client.
// We want to provide access to a broader range of (simulated) files than
// Native Client allows, so we maintain our own file descriptor table exposed
// to higher-level packages.

package syscall

// files is the table indexed by a file descriptor.

// A file is an open file, something with a file descriptor.
// A particular *file may appear in files multiple times, due to use of Dup or Dup2.

// A fileImpl is the implementation of something that can be a file.

func Close(fd int) error

func CloseOnExec(fd int)

func Dup(fd int) (int, error)

func Dup2(fd, newfd int) error

func Fstat(fd int, st *Stat_t) error

func Read(fd int, b []byte) (int, error)

func Write(fd int, b []byte) (int, error)

func Pread(fd int, b []byte, offset int64) (int, error)

func Pwrite(fd int, b []byte, offset int64) (int, error)

func Seek(fd int, offset int64, whence int) (int64, error)

// defaulFileImpl implements fileImpl.
// It can be embedded to complete a partial fileImpl implementation.

// naclFile is the fileImpl implementation for a Native Client file descriptor.

// A pipeFile is an in-memory implementation of a pipe.
// The byteq implementation is in net_nacl.go.

func Pipe(fd []int) error
