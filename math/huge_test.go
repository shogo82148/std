// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math_test

import (
	. "math"
)

// Inputs to test trig_reduce

// Results for trigHuge[i] calculated with https://github.com/robpike/ivy
// using 4096 bits of working precision.   Values requiring less than
// 102 decimal digits (1 << 120, 1 << 240, 1 << 480, 1234567891234567 << 180)
// were confirmed via https://keisan.casio.com/
