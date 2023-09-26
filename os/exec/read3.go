// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// This is a test program that verifies that it can read from
// descriptor 3 and that no other descriptors are open.
// This is not done via TestHelperProcess and GO_WANT_HELPER_PROCESS
// because we want to ensure that this program does not use cgo,
// because C libraries can open file descriptors behind our backs
// and confuse the test. See issue 25628.
package main
