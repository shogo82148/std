// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo && !windows

// Issue 18146: pthread_create failure during syscall.Exec.

package cgotest
