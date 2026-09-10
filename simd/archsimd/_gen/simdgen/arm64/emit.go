// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arm64

import (
	"github.com/shogo82148/std/simd/archsimd/_gen/unify"
)

// EmitAll generates instruction definitions for all arrangements of this instruction.
// Returns nil for instructions with UnsupportedArngs.
func (instruction *Instruction) EmitAll() []*unify.Value
