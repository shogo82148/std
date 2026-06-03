// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// NOTE: This currently only supports the advsimd (NEON) instruction class.

package arm64

import (
	"github.com/shogo82148/std/_gen/unify"
)

// ParseInstructions loads and parses ARM64 instruction definitions from XML files at given path.
func ParseInstructions(path string) ([]*Instruction, error)

// Load loads ARM64 instruction definitions from XML files at given path and returns them as unify values.
func Load(path string) ([]*unify.Value, error)
