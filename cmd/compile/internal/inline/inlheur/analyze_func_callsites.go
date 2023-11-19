// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inlheur

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// UpdateCallsiteTable handles updating of callerfn's call site table
// after an inlined has been carried out, e.g. the call at 'n' as been
// turned into the inlined call expression 'ic' within function
// callerfn. The chief thing of interest here is to make sure that any
// call nodes within 'ic' are added to the call site table for
// 'callerfn' and scored appropriately.
func UpdateCallsiteTable(callerfn *ir.Func, n *ir.CallExpr, ic *ir.InlinedCallExpr)
