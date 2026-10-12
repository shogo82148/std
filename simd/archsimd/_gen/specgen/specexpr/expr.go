// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpr

type Expr interface {
	String() string
	eval(b *Bindings) (any, error)
	preorder(yield func(Expr) bool) bool
}

// Literal is an [Expr] that evaluates to a literal value.
//
// For [Int] and [SymbolicWidth], you probably just want to use those types
// directly. They're literal values, so you could wrap them in a Literal, but
// they are valid expressions on their own.
type Literal struct {
	Val any
}

func (e *Literal) String() string

// Variable is an [Expr] that evaluates to the value of the named variable.
type Variable string

func (e Variable) String() string

// Func is a function that can be used in an expression. Use the [Func.Apply]
// method to create an [Expr].
type Func struct {
	Name string
	Func func([]any) (any, error)
}

func MakeFunc1[T any](name string, fn func(T) (any, error)) func(e Expr) *Apply

func MakeFunc2[T, U any](name string, fn func(T, U) (any, error)) func(e1, e2 Expr) *Apply

func MakeFunc(name string, fn func([]any) (any, error)) *Func

// MakeField returns a Func that projects field fieldName from a value of type
// T. The returned functions are memoized, so only one *Func is created per type
// and field and this is efficient to call repeatedly.
func MakeField[T any](fieldName string) *Func

func (f *Func) Apply(args ...Expr) *Apply

// Apply is an [Expr] that applies a [Func] to a sequence of arguments.
type Apply struct {
	Func *Func
	Args []Expr
}

func (e *Apply) String() string

// BinExpr is a binary [Expr].
type BinExpr struct {
	Op   BinOp
	X, Y Expr
}

func (e *BinExpr) String() string

type BinOp byte

const (
	_       BinOp = iota
	OpTimes
	OpDiv

	// Comparison operators
	OpEqual
	OpNotEqual
	OpGreaterThan
	OpLessThan
	OpGreaterOrEqual
	OpLessOrEqual
)
