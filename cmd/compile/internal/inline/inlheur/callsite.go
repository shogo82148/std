// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inlheur

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// CallSite records useful information about a potentially inlinable
// (direct) function call. "Callee" is the target of the call, "Call"
// is the ir node corresponding to the call itself, "Assign" is
// the top-level assignment statement containing the call (if the call
// appears in the form of a top-level statement, e.g. "x := foo()"),
// "Flags" contains properties of the call that might be useful for
// making inlining decisions, "Score" is the final score assigned to
// the site, and "ID" is a numeric ID for the site within its
// containing function.
type CallSite struct {
	Callee    *ir.Func
	Call      *ir.CallExpr
	parent    *CallSite
	Assign    ir.Node
	Flags     CSPropBits
	Score     int
	ScoreMask scoreAdjustTyp
	ID        uint
	aux       uint8
}

// CallSiteTab is a table of call sites, keyed by call expr.
// Ideally it would be nice to key the table by src.XPos, but
// this results in collisions for calls on very long lines (the
// front end saturates column numbers at 255). We also wind up
// with many calls that share the same auto-generated pos.
type CallSiteTab map[*ir.CallExpr]*CallSite

type CSPropBits uint32

const (
	CallSiteInLoop CSPropBits = 1 << iota
	CallSiteOnPanicPath
	CallSiteInInitFunc
)

func EncodeCallSiteKey(cs *CallSite) string
