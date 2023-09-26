// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package types declares the data types and implements
// the algorithms for type-checking of Go packages. Use
// Config.Check to invoke the type checker for a package.
// Alternatively, create a new type checker with NewChecker
// and invoke it incrementally by calling Checker.Files.
//
// Type-checking consists of several interdependent phases:
//
// Name resolution maps each identifier (ast.Ident) in the program to the
// language object (Object) it denotes.
// Use Info.{Defs,Uses,Implicits} for the results of name resolution.
//
// Constant folding computes the exact constant value (constant.Value)
// for every expression (ast.Expr) that is a compile-time constant.
// Use Info.Types[expr].Value for the results of constant folding.
//
// Type inference computes the type (Type) of every expression (ast.Expr)
// and checks for compliance with the language specification.
// Use Info.Types[expr].Type for the results of type inference.
//
// For a tutorial, see https://golang.org/s/types-tutorial.
package types

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/constant"
	"github.com/shogo82148/std/go/token"
	. "internal/types/errors"
)

// An Error describes a type-checking error; it implements the error interface.
// A "soft" error is an error that still permits a valid interpretation of a
// package (such as "unused variable"); "hard" errors may lead to unpredictable
// behavior if ignored.
type Error struct {
	Fset *token.FileSet
	Pos  token.Pos
	Msg  string
	Soft bool

	go116code  Code
	go116start token.Pos
	go116end   token.Pos
}

// Error returns an error string formatted as follows:
// filename:line:column: message
func (err Error) Error() string

// An ArgumentError holds an error associated with an argument index.
type ArgumentError struct {
	Index int
	Err   error
}

func (e *ArgumentError) Error() string
func (e *ArgumentError) Unwrap() error

// An Importer resolves import paths to Packages.
//
// CAUTION: This interface does not support the import of locally
// vendored packages. See https://golang.org/s/go15vendor.
// If possible, external implementations should implement ImporterFrom.
type Importer interface {
	Import(path string) (*Package, error)
}

// ImportMode is reserved for future use.
type ImportMode int

// An ImporterFrom resolves import paths to packages; it
// supports vendoring per https://golang.org/s/go15vendor.
// Use go/importer to obtain an ImporterFrom implementation.
type ImporterFrom interface {
	Importer

	ImportFrom(path, dir string, mode ImportMode) (*Package, error)
}

// A Config specifies the configuration for type checking.
// The zero value for Config is a ready-to-use default configuration.
type Config struct {
	Context *Context

	GoVersion string

	IgnoreFuncBodies bool

	FakeImportC bool

	go115UsesCgo bool

	_Trace bool

	Error func(err error)

	Importer Importer

	Sizes Sizes

	DisableUnusedImportCheck bool

	_ErrorURL string
}

// Info holds result type information for a type-checked package.
// Only the information for which a map is provided is collected.
// If the package has type errors, the collected information may
// be incomplete.
type Info struct {
	Types map[ast.Expr]TypeAndValue

	Instances map[*ast.Ident]Instance

	Defs map[*ast.Ident]Object

	Uses map[*ast.Ident]Object

	Implicits map[ast.Node]Object

	Selections map[*ast.SelectorExpr]*Selection

	Scopes map[ast.Node]*Scope

	InitOrder []*Initializer

	_FileVersions map[token.Pos]_Version
}

// TypeOf returns the type of expression e, or nil if not found.
// Precondition: the Types, Uses and Defs maps are populated.
func (info *Info) TypeOf(e ast.Expr) Type

// ObjectOf returns the object denoted by the specified id,
// or nil if not found.
//
// If id is an embedded struct field, ObjectOf returns the field (*Var)
// it defines, not the type (*TypeName) it uses.
//
// Precondition: the Uses and Defs maps are populated.
func (info *Info) ObjectOf(id *ast.Ident) Object

// TypeAndValue reports the type and value (for constants)
// of the corresponding expression.
type TypeAndValue struct {
	mode  operandMode
	Type  Type
	Value constant.Value
}

// IsVoid reports whether the corresponding expression
// is a function call without results.
func (tv TypeAndValue) IsVoid() bool

// IsType reports whether the corresponding expression specifies a type.
func (tv TypeAndValue) IsType() bool

// IsBuiltin reports whether the corresponding expression denotes
// a (possibly parenthesized) built-in function.
func (tv TypeAndValue) IsBuiltin() bool

// IsValue reports whether the corresponding expression is a value.
// Builtins are not considered values. Constant values have a non-
// nil Value.
func (tv TypeAndValue) IsValue() bool

// IsNil reports whether the corresponding expression denotes the
// predeclared value nil.
func (tv TypeAndValue) IsNil() bool

// Addressable reports whether the corresponding expression
// is addressable (https://golang.org/ref/spec#Address_operators).
func (tv TypeAndValue) Addressable() bool

// Assignable reports whether the corresponding expression
// is assignable to (provided a value of the right type).
func (tv TypeAndValue) Assignable() bool

// HasOk reports whether the corresponding expression may be
// used on the rhs of a comma-ok assignment.
func (tv TypeAndValue) HasOk() bool

// Instance reports the type arguments and instantiated type for type and
// function instantiations. For type instantiations, Type will be of dynamic
// type *Named. For function instantiations, Type will be of dynamic type
// *Signature.
type Instance struct {
	TypeArgs *TypeList
	Type     Type
}

// An Initializer describes a package-level variable, or a list of variables in case
// of a multi-valued initialization expression, and the corresponding initialization
// expression.
type Initializer struct {
	Lhs []*Var
	Rhs ast.Expr
}

func (init *Initializer) String() string

// A _Version represents a released Go version.

// Check type-checks a package and returns the resulting package object and
// the first error if any. Additionally, if info != nil, Check populates each
// of the non-nil maps in the Info struct.
//
// The package is marked as complete if no errors occurred, otherwise it is
// incomplete. See Config.Error for controlling behavior in the presence of
// errors.
//
// The package is specified by a list of *ast.Files and corresponding
// file set, and the package path the package is identified with.
// The clean path must not be empty or dot (".").
func (conf *Config) Check(path string, fset *token.FileSet, files []*ast.File, info *Info) (*Package, error)

// AssertableTo reports whether a value of type V can be asserted to have type T.
//
// The behavior of AssertableTo is unspecified in three cases:
//   - if T is Typ[Invalid]
//   - if V is a generalized interface; i.e., an interface that may only be used
//     as a type constraint in Go code
//   - if T is an uninstantiated generic type
func AssertableTo(V *Interface, T Type) bool

// AssignableTo reports whether a value of type V is assignable to a variable
// of type T.
//
// The behavior of AssignableTo is unspecified if V or T is Typ[Invalid] or an
// uninstantiated generic type.
func AssignableTo(V, T Type) bool

// ConvertibleTo reports whether a value of type V is convertible to a value of
// type T.
//
// The behavior of ConvertibleTo is unspecified if V or T is Typ[Invalid] or an
// uninstantiated generic type.
func ConvertibleTo(V, T Type) bool

// Implements reports whether type V implements interface T.
//
// The behavior of Implements is unspecified if V is Typ[Invalid] or an uninstantiated
// generic type.
func Implements(V Type, T *Interface) bool

// Satisfies reports whether type V satisfies the constraint T.
//
// The behavior of Satisfies is unspecified if V is Typ[Invalid] or an uninstantiated
// generic type.
func Satisfies(V Type, T *Interface) bool

// Identical reports whether x and y are identical types.
// Receivers of Signature types are ignored.
func Identical(x, y Type) bool

// IdenticalIgnoreTags reports whether x and y are identical types if tags are ignored.
// Receivers of Signature types are ignored.
func IdenticalIgnoreTags(x, y Type) bool
