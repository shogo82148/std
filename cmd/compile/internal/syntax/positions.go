// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements helper functions for scope position computations.

package syntax

// StartPos returns the start position of n.
func StartPos(n Node) Pos

// EndPos returns the approximate end position of n in the source.
// For some nodes (*Name, *BasicLit) it returns the position immediately
// following the node; for others (*BlockStmt, *SwitchStmt, etc.) it
// returns the position of the closing '}'; and for some (*ParenExpr)
// the returned position is the end position of the last enclosed
// expression.
// Thus, EndPos should not be used for exact demarcation of the
// end of a node in the source; it is mostly useful to determine
// scope ranges where there is some leeway.
func EndPos(n Node) Pos
