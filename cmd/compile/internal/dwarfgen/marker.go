// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dwarfgen

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/internal/src"
)

// A ScopeMarker tracks scope nesting and boundaries for later use
// during DWARF generation.
type ScopeMarker struct {
	parents []ir.ScopeID
	marks   []ir.Mark
}

// Push records a transition to a new child scope of the current scope.
func (m *ScopeMarker) Push(pos src.XPos)

// Pop records a transition back to the current scope's parent.
func (m *ScopeMarker) Pop(pos src.XPos)

// Unpush removes the current scope, which must be empty.
func (m *ScopeMarker) Unpush()

// WriteTo writes the recorded scope marks to the given function,
// and resets the marker for reuse.
func (m *ScopeMarker) WriteTo(fn *ir.Func)
