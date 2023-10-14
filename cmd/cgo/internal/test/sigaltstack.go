// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !android

// Test that the Go runtime still works if C code changes the signal stack.

package cgotest
