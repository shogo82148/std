// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

type RF_State int

// These constants are shared between the compiler, which uses them for state functions
// and panic indicators, and the runtime, which turns them into more meaningful strings
// For best code generation, RF_DONE and RF_READY should be 0 and 1.
const (
	RF_DONE          = RF_State(iota)
	RF_READY
	RF_PANIC
	RF_EXHAUSTED
	RF_MISSING_PANIC = 4
)
