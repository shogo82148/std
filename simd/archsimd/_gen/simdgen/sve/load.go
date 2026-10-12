// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sve

import (
	"github.com/shogo82148/std/simd/archsimd/_gen/unify"
)

// Load parses the ARM64 ISA XML files at path and returns the SVE / SVE2
// instruction definitions as simdgen unify values.
func Load(path string) ([]*unify.Value, error)
