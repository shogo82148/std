// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inlheur

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/pgo"
)

// DumpInlCallSiteScores is invoked by the inliner if the debug flag
// "-d=dumpinlcallsitescores" is set; it dumps out a human-readable
// summary of all (potentially) inlinable callsites in the package,
// along with info on call site scoring and the adjustments made to a
// given score. Here profile is the PGO profile in use (may be
// nil), budgetCallback is a callback that can be invoked to find out
// the original pre-adjustment hairyness limit for the function, and
// inlineHotMaxBudget is the constant of the same name used in the
// inliner. Sample output lines:
//
// Score  Adjustment  Status  Callee  CallerPos ScoreFlags
// 115    40          DEMOTED cmd/compile/internal/abi.(*ABIParamAssignment).Offset     expand_calls.go:1679:14|6       panicPathAdj
// 76     -5n           PROMOTED runtime.persistentalloc   mcheckmark.go:48:45|3   inLoopAdj
// 201    0           --- PGO  unicode.DecodeRuneInString        utf8.go:312:30|1
// 7      -5          --- PGO  internal/abi.Name.DataChecked     type.go:625:22|0        inLoopAdj
//
// In the dump above, "Score" is the final score calculated for the
// callsite, "Adjustment" is the amount added to or subtracted from
// the original hairyness estimate to form the score. "Status" shows
// whether anything changed with the site -- did the adjustment bump
// it down just below the threshold ("PROMOTED") or instead bump it
// above the threshold ("DEMOTED"); this will be blank ("---") if no
// threshold was crossed as a result of the heuristics. Note that
// "Status" also shows whether PGO was involved. "Callee" is the name
// of the function called, "CallerPos" is the position of the
// callsite, and "ScoreFlags" is a digest of the specific properties
// we used to make adjustments to callsite score via heuristics.
func DumpInlCallSiteScores(profile *pgo.Profile, budgetCallback func(fn *ir.Func, profile *pgo.Profile) (int32, bool))
