// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || (solaris && !illumos)

// This code implements the filelock API using POSIX 'fcntl' locks, which attach
// to an (inode, process) pair rather than a file descriptor. To avoid unlocking
// files prematurely when the same file is opened through different descriptors,
// we allow only one read-lock at a time.
//
// Most platforms provide some alternative API, such as an 'flock' system call
// or an F_OFD_SETLK command for 'fcntl', that allows for better concurrency and
// does not require per-inode bookkeeping in the application.

package filelock
