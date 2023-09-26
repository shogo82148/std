// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// Repeated malloc test.

package main

func OkAmount(size, n uintptr) bool

func AllocAndFree(size, count int)
