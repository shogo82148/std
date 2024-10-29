// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || solaris

package poll

// SendFile wraps the sendfile system call.
//
// It copies data from src (a file descriptor) to dstFD,
// starting at the current position of src.
// It updates the current position of src to after the
// copied data.
//
// If size is zero, it copies the rest of src.
// Otherwise, it copies up to size bytes.
//
// The handled return parameter indicates whether SendFile
// was able to handle some or all of the operation.
// If handled is false, sendfile was unable to perform the copy,
// has not modified the source or destination,
// and the caller should perform the copy using a fallback implementation.
func SendFile(dstFD *FD, src int, size int64) (n int64, err error, handled bool)
