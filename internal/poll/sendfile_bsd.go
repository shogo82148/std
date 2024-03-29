// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd

package poll

// SendFile wraps the sendfile system call.
func SendFile(dstFD *FD, src int, pos, remain int64) (written int64, err error, handled bool)
