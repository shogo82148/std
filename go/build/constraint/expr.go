// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package constraint implements parsing and evaluation of build constraint lines.
// See https://golang.org/cmd/go/#hdr-Build_constraints for documentation about build constraints themselves.
//
// This package parses both the original “// +build” syntax and the “//go:build” syntax that was added in Go 1.17.
// See https://golang.org/design/draft-gobuild for details about the “//go:build” syntax.
package constraint

// An Expr is a build tag constraint expression.
// The underlying concrete type is *AndExpr, *OrExpr, *NotExpr, or *TagExpr.
type Expr interface {
	String() string

	Eval(ok func(tag string) bool) bool

	isExpr()
}

// A TagExpr is an Expr for the single tag Tag.
type TagExpr struct {
	Tag string
}

func (x *TagExpr) Eval(ok func(tag string) bool) bool

func (x *TagExpr) String() string

// A NotExpr represents the expression !X (the negation of X).
type NotExpr struct {
	X Expr
}

func (x *NotExpr) Eval(ok func(tag string) bool) bool

func (x *NotExpr) String() string

// An AndExpr represents the expression X && Y.
type AndExpr struct {
	X, Y Expr
}

func (x *AndExpr) Eval(ok func(tag string) bool) bool

func (x *AndExpr) String() string

// An OrExpr represents the expression X || Y.
type OrExpr struct {
	X, Y Expr
}

func (x *OrExpr) Eval(ok func(tag string) bool) bool

func (x *OrExpr) String() string

// A SyntaxError reports a syntax error in a parsed build expression.
type SyntaxError struct {
	Offset int
	Err    string
}

func (e *SyntaxError) Error() string

// Parse parses a single build constraint line of the form “//go:build ...” or “// +build ...”
// and returns the corresponding boolean expression.
func Parse(line string) (Expr, error)

// IsGoBuild reports whether the line of text is a “//go:build” constraint.
// It only checks the prefix of the text, not that the expression itself parses.
func IsGoBuild(line string) bool

// An exprParser holds state for parsing a build expression.

// IsPlusBuild reports whether the line of text is a “// +build” constraint.
// It only checks the prefix of the text, not that the expression itself parses.
func IsPlusBuild(line string) bool

// PlusBuildLines returns a sequence of “// +build” lines that evaluate to the build expression x.
// If the expression is too complex to convert directly to “// +build” lines, PlusBuildLines returns an error.
func PlusBuildLines(x Expr) ([]string, error)
