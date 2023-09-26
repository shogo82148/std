// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// mkpost processes the output of cgo -godefs to
// modify the generated types. It is used to clean up
// the syscall API in an architecture specific manner.
//
// mkpost is run after cgo -godefs by mkall.sh.
package main
