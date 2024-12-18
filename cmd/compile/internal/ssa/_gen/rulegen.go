// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This program generates Go code that applies rewrite rules to a Value.
// The generated code implements a function of type func (v *Value) bool
// which reports whether if did something.
// Ideas stolen from the Swift Java compiler:
// https://bitsavers.org/pdf/dec/tech_reports/WRL-2000-2.pdf

package main

import (
	"github.com/shogo82148/std/go/ast"
)

type Rule struct {
	Rule string
	Loc  string
}

func (r Rule) String() string

// Node can be a Statement or an ast.Expr.
type Node interface{}

// Statement can be one of our high-level statement struct types, or an
// ast.Stmt under some limited circumstances.
type Statement interface{}

// BodyBase is shared by all of our statement pseudo-node types which can
// contain other statements.
type BodyBase struct {
	List    []Statement
	CanFail bool
}

// These types define some high-level statement struct types, which can be used
// as a Statement. This allows us to keep some node structs simpler, and have
// higher-level nodes such as an entire rule rewrite.
//
// Note that ast.Expr is always used as-is; we don't declare our own expression
// nodes.
type (
	File struct {
		BodyBase
		Arch   arch
		Suffix string
	}
	Func struct {
		BodyBase
		Kind   string
		Suffix string
		ArgLen int32
	}
	Switch struct {
		BodyBase
		Expr ast.Expr
	}
	Case struct {
		BodyBase
		Expr ast.Expr
	}
	RuleRewrite struct {
		BodyBase
		Match, Cond, Result string
		Check               string

		Alloc        int
		Loc          string
		CommuteDepth int
	}
	Declare struct {
		Name  string
		Value ast.Expr
	}
	CondBreak struct {
		Cond              ast.Expr
		InsideCommuteLoop bool
	}
	StartCommuteLoop struct {
		Depth int
		V     string
	}
)
