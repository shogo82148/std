// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

// SendFile wraps the sendfile system call.
func SendFile(dstFD *FD, src int, remain int64) (int64, error, bool)
