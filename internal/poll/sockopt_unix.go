// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package poll

// SetsockoptByte wraps the setsockopt network call with a byte argument.
func (fd *FD) SetsockoptByte(level, name int, arg byte) error
