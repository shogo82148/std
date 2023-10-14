// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build plan9

package fdtest

// Exists returns true if fd is a valid file descriptor.
func Exists(fd uintptr) bool
