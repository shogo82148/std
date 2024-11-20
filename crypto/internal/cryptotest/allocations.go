// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptotest

import (
	"github.com/shogo82148/std/testing"
)

// SkipTestAllocations skips the test if there are any factors that interfere
// with allocation optimizations.
func SkipTestAllocations(t *testing.T)
