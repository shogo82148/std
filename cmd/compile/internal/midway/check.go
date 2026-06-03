// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package midway

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
)

// CheckPositions checks that all nodes in the files have known positions.
// This converts lack-of-Pos into an early fatal error instead of a later
// weird downstream error (e.g., in the linker, in debugging information).
func CheckPositions(files []*syntax.File, phase string)
