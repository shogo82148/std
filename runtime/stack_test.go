// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test stack split logic by calling functions of every frame size
// from near 0 up to and beyond the default segment size (4k).
// Each of those functions reports its SP + stack limit, and then
// the test (the caller) checks that those make sense.  By not
// doing the actual checking and reporting from the suspect functions,
// we minimize the possibility of crashes during the test itself.
//
// Exhaustive test for http://golang.org/issue/3310.
// The linker used to get a few sizes near the segment size wrong:
//
// --- FAIL: TestStackSplit (0.01 seconds)
// 	stack_test.go:22: after runtime_test.stack3812: sp=0x7f7818d5d078 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3816: sp=0x7f7818d5d078 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3820: sp=0x7f7818d5d070 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3824: sp=0x7f7818d5d070 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3828: sp=0x7f7818d5d068 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3832: sp=0x7f7818d5d068 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3836: sp=0x7f7818d5d060 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3840: sp=0x7f7818d5d060 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3844: sp=0x7f7818d5d058 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3848: sp=0x7f7818d5d058 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3852: sp=0x7f7818d5d050 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3856: sp=0x7f7818d5d050 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3860: sp=0x7f7818d5d048 < limit=0x7f7818d5d080
// 	stack_test.go:22: after runtime_test.stack3864: sp=0x7f7818d5d048 < limit=0x7f7818d5d080
// FAIL

package runtime_test

import (
	. "runtime"
)

// See stack.h.
const (
	StackGuard = 256
	StackLimit = 128
)

var Used byte
