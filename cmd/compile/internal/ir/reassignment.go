// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

// A ReassignOracle efficiently answers queries about whether local
// variables are reassigned. This helper works by looking for function
// params and short variable declarations (e.g.
// https://go.dev/ref/spec#Short_variable_declarations) that are
// neither address taken nor subsequently re-assigned. It is intended
// to operate much like "ir.StaticValue" and "ir.Reassigned", but in a
// way that does just a single walk of the containing function (as
// opposed to a new walk on every call).
type ReassignOracle struct {
	fn *Func
	// maps candidate name to its defining assignment (or
	// for params, defining func).
	singleDef map[*Name]Node

	// funcAssigns tracks all known simple assignments (OAS) to
	// func-typed PAUTO variables. Only func-typed variables are
	// tracked because this data is used exclusively for callee
	// resolution in escape analysis. Deletion means the candidate was
	// invalidated (e.g., addr-taken, non-simple assignment form, or too
	// many assignments). Assignments inside nested closures are accepted
	// because the only alternative value is nil, which panics on call.
	funcAssigns map[*Name][]*AssignStmt
}

// Init initializes the oracle based on the IR in function fn, laying
// the groundwork for future calls to the StaticValue and Reassigned
// methods. If the fn's IR is subsequently modified, Init must be
// called again.
func (ro *ReassignOracle) Init(fn *Func)

// StaticValue method has the same semantics as the ir package function
// of the same name; see comments on [StaticValue].
func (ro *ReassignOracle) StaticValue(n Node) Node

// Reassigned method has the same semantics as the ir package function
// of the same name; see comments on [Reassigned] for more info.
func (ro *ReassignOracle) Reassigned(n *Name) bool

// FuncAssignments returns all known simple assignments to a func-typed
// variable. For variables defined with := and a non-zero value, the
// defining assignment is included. Returns nil if the variable is not
// func-typed, was invalidated (addr-taken, non-simple assignment,
// too many assignments), or has no tracked assignments. Assignments
// inside nested closures are accepted because the only alternative
// value is nil, which panics on call.
func (ro *ReassignOracle) FuncAssignments(name *Name) []*AssignStmt
