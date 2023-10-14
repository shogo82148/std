// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This test uses the Pdeathsig field of syscall.SysProcAttr, so it only works
// on platforms that support that.

//go:build linux || (freebsd && amd64)

// sanitizers_test checks the use of Go with sanitizers like msan, asan, etc.
// See https://github.com/google/sanitizers.
package sanitizers_test
