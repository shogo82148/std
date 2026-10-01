// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specdoc

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
)

// Finding is an individual finding in a [Report].
type Finding interface {
	String() string
	Decl() Decl
}

// Report collects all findings from specfill across checked declarations.
// It implements the error interface.
type Report struct {
	TypeMismatches     []TypeMismatch
	NameMismatches     []NameMismatch
	DocOrderViolations []DocOrderViolation
	UnknownDecls       []UnknownDecl
	UnexpectedDocs     []UnexpectedDoc
}

// Error formats the report as a string, implementing the error interface.
func (r *Report) Error() string

// Decl identifies an AST declaration and its location in the source code.
type Decl struct {
	Pos  token.Position
	Recv string
	Name string
}

// Compare orders declarations by source position (filename, line, column,
// offset), breaking ties by receiver and name.
func (d Decl) Compare(other Decl) int

// String formats the declaration as "(Recv) Name" or "Name".
func (d Decl) String() string

// TypeMismatch records a disagreement between spec parameter/result types
// and an AST declaration's parameter/result types.
type TypeMismatch struct {
	D       Decl
	DeclSig string
	SpecSig string
	Details []string
}

func (m TypeMismatch) Decl() Decl

func (m TypeMismatch) String() string

// NameMismatch records a disagreement between spec parameter/result names
// and an AST declaration's parameter/result names.
type NameMismatch struct {
	D       Decl
	DeclSig string
	SpecSig string
	Details []string
}

func (m NameMismatch) Decl() Decl

func (m NameMismatch) String() string

// DocOrderViolation records a comment ordering violation in an exported
// declaration's doc comment, which would result in non-idempotent rewriting.
type DocOrderViolation struct {
	D   Decl
	Err error
}

func (v DocOrderViolation) Decl() Decl

func (v DocOrderViolation) String() string

// UnknownDecl records an exported declaration that has no matching entry in spec.
type UnknownDecl struct {
	D Decl
}

func (u UnknownDecl) Decl() Decl

func (u UnknownDecl) String() string

// UnexpectedDoc records an exported declaration containing an existing
// spec-owned doc comment when AllowDocRewrite is false.
type UnexpectedDoc struct {
	D Decl
}

func (u UnexpectedDoc) Decl() Decl

func (u UnexpectedDoc) String() string

// Merge combines other into r.
func (r *Report) Merge(other *Report)

// Empty reports whether the report contains no items.
func (r Report) Empty() bool

// Findings returns all findings in r as a slice of [Finding].
func (r Report) Findings() []Finding

// Print writes the report to w sorted by position. If there are no findings,
// nothing is written.
func (r Report) Print(w io.Writer)
