// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !android

// Test that pthread_cancel works as expected
// (NPTL uses SIGRTMIN to implement thread cancellation)
// See https://golang.org/issue/6997
package cgotest
