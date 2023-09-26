// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	. "runtime"
)

// minRate is an arbitrary minimum for allocRate, scanRate, and growthRate.
// These values just cannot be zero.

// float64Stream is a function that generates an infinite stream of
// float64 values when called repeatedly.
