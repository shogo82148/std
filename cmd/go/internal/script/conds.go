// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

// DefaultConds returns a set of broadly useful script conditions.
//
// Run the 'help' command within a script engine to view a list of the available
// conditions.
func DefaultConds() map[string]Cond

// Condition returns a Cond with the given summary and evaluation function.
func Condition(summary string, eval func(*State) (bool, error)) Cond

// PrefixCondition returns a Cond with the given summary and evaluation function.
func PrefixCondition(summary string, eval func(*State, string) (bool, error)) Cond

// BoolCondition returns a Cond with the given truth value and summary.
// The Cond rejects the use of condition suffixes.
func BoolCondition(summary string, v bool) Cond

// OnceCondition returns a Cond that calls eval the first time the condition is
// evaluated. Future calls reuse the same result.
//
// The eval function is not passed a *State because the condition is cached
// across all execution states and must not vary by state.
func OnceCondition(summary string, eval func() (bool, error)) Cond

// CachedCondition is like Condition but only calls eval the first time the
// condition is evaluated for a given suffix.
// Future calls with the same suffix reuse the earlier result.
//
// The eval function is not passed a *State because the condition is cached
// across all execution states and must not vary by state.
func CachedCondition(summary string, eval func(string) (bool, error)) Cond
