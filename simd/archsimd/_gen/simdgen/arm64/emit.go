// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arm64

import (
	"github.com/shogo82148/std/_gen/unify"
)

// Emit generates the unify.Value representation of this operand
func (op *Operand) Emit() *unify.Value

// EmitAll generates instruction definitions for all arrangements of this instruction.
// Returns nil for instructions with UnsupportedArngs.
func (instruction *Instruction) EmitAll() []*unify.Value
