// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gccgo
// +build gccgo

package main

/*
int supports_sse2() {
#if defined(__i386__) || defined(__x86_64__)
	return __builtin_cpu_supports("sse2");
#else
	return 0;
#endif
}
*/
