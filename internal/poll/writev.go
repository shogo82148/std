// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package poll

// Writev wraps the writev system call.
func (fd *FD) Writev(v *[][]byte) (int64, error)
