// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inlheur

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/pgo"
)

// SetupScoreAdjustments interprets the value of the -d=inlscoreadj
// debugging option, if set. The value of this flag is expected to be
// a series of "/"-separated clauses of the form adj1:value1. Example:
// -d=inlscoreadj=inLoopAdj=0/passConstToIfAdj=-99
func SetupScoreAdjustments()

// LargestNegativeScoreAdjustment tries to estimate the largest possible
// negative score adjustment that could be applied to a call of the
// function with the specified props. Example:
//
//	func foo() {                  func bar(x int, p *int) int {
//	   ...                          if x < 0 { *p = x }
//	}                               return 99
//	                              }
//
// Function 'foo' above on the left has no interesting properties,
// thus as a result the most we'll adjust any call to is the value for
// "call in loop". If the calculated cost of the function is 150, and
// the in-loop adjustment is 5 (for example), then there is not much
// point treating it as inlinable. On the other hand "bar" has a param
// property (parameter "x" feeds unmodified to an "if" statement") and
// a return property (always returns same constant) meaning that a
// given call _could_ be rescored down as much as -35 points-- thus if
// the size of "bar" is 100 (for example) then there is at least a
// chance that scoring will enable inlining.
func LargestNegativeScoreAdjustment(fn *ir.Func, props *FuncProps) int

// LargestPositiveScoreAdjustment tries to estimate the largest possible
// positive score adjustment that could be applied to a given callsite.
// At the moment we don't have very many positive score adjustments, so
// this is just hard-coded, not table-driven.
func LargestPositiveScoreAdjustment(fn *ir.Func) int

// ScoreCalls assigns numeric scores to each of the callsites in
// function 'fn'; the lower the score, the more helpful we think it
// will be to inline.
//
// Unlike a lot of the other inline heuristics machinery, callsite
// scoring can't be done as part of the CanInline call for a function,
// due to fact that we may be working on a non-trivial SCC. So for
// example with this SCC:
//
//	func foo(x int) {           func bar(x int, f func()) {
//	  if x != 0 {                  f()
//	    bar(x, func(){})           foo(x-1)
//	  }                         }
//	}
//
// We don't want to perform scoring for the 'foo' call in "bar" until
// after foo has been analyzed, but it's conceivable that CanInline
// might visit bar before foo for this SCC.
func ScoreCalls(fn *ir.Func)

// ScoreCallsCleanup resets the state of the callsite cache
// once ScoreCalls is done with a function.
func ScoreCallsCleanup()

// GetCallSiteScore returns the previously calculated score for call
// within fn.
func GetCallSiteScore(fn *ir.Func, call *ir.CallExpr) (int, bool)

// BudgetExpansion returns the amount to relax/expand the base
// inlining budget when the new inliner is turned on; the inliner
// will add the returned value to the hairiness budget.
//
// Background: with the new inliner, the score for a given callsite
// can be adjusted down by some amount due to heuristics, however we
// won't know whether this is going to happen until much later after
// the CanInline call. This function returns the amount to relax the
// budget initially (to allow for a large score adjustment); later on
// in RevisitInlinability we'll look at each individual function to
// demote it if needed.
func BudgetExpansion(maxBudget int32) int32

// DumpInlCallSiteScores is invoked by the inliner if the debug flag
// "-d=dumpinlcallsitescores" is set; it dumps out a human-readable
// summary of all (potentially) inlinable callsites in the package,
// along with info on call site scoring and the adjustments made to a
// given score. Here profile is the PGO profile in use (may be
// nil), budgetCallback is a callback that can be invoked to find out
// the original pre-adjustment hairiness limit for the function, and
// inlineHotMaxBudget is the constant of the same name used in the
// inliner. Sample output lines:
//
// Score  Adjustment  Status  Callee  CallerPos ScoreFlags
// 115    40          DEMOTED cmd/compile/internal/abi.(*ABIParamAssignment).Offset     expand_calls.go:1679:14|6       panicPathAdj
// 76     -5n         PROMOTED runtime.persistentalloc   mcheckmark.go:48:45|3   inLoopAdj
// 201    0           --- PGO  unicode.DecodeRuneInString        utf8.go:312:30|1
// 7      -5          --- PGO  internal/abi.Name.DataChecked     type.go:625:22|0        inLoopAdj
//
// In the dump above, "Score" is the final score calculated for the
// callsite, "Adjustment" is the amount added to or subtracted from
// the original hairiness estimate to form the score. "Status" shows
// whether anything changed with the site -- did the adjustment bump
// it down just below the threshold ("PROMOTED") or instead bump it
// above the threshold ("DEMOTED"); this will be blank ("---") if no
// threshold was crossed as a result of the heuristics. Note that
// "Status" also shows whether PGO was involved. "Callee" is the name
// of the function called, "CallerPos" is the position of the
// callsite, and "ScoreFlags" is a digest of the specific properties
// we used to make adjustments to callsite score via heuristics.
func DumpInlCallSiteScores(profile *pgo.Profile, budgetCallback func(fn *ir.Func, profile *pgo.Profile) (int32, bool))
