// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/fmt"
)

// BuiltinPkg is a fake package that declares the universe block.
var BuiltinPkg *Pkg

// LocalPkg is the package being compiled.
var LocalPkg *Pkg

// UnsafePkg is package unsafe.
var UnsafePkg *Pkg

// BlankSym is the blank (_) symbol.
var BlankSym *Sym

// numImport tracks how often a package with a given name is imported.
// It is used to provide a better error message (by using the package
// path to disambiguate) if a package that appears multiple times with
// the same name appears in an error message.
var NumImport = make(map[string]int)

// Format implements formatting for a Sym.
// The valid formats are:
//
//	%v	Go syntax: Name for symbols in the local package, PkgName.Name for imported symbols.
//	%+v	Debug syntax: always include PkgName. prefix even for local names.
//	%S	Short syntax: Name only, no matter what.
func (s *Sym) Format(f fmt.State, verb rune)

func (s *Sym) String() string

var BasicTypeNames = []string{
	TINT:        "int",
	TUINT:       "uint",
	TINT8:       "int8",
	TUINT8:      "uint8",
	TINT16:      "int16",
	TUINT16:     "uint16",
	TINT32:      "int32",
	TUINT32:     "uint32",
	TINT64:      "int64",
	TUINT64:     "uint64",
	TUINTPTR:    "uintptr",
	TFLOAT32:    "float32",
	TFLOAT64:    "float64",
	TCOMPLEX64:  "complex64",
	TCOMPLEX128: "complex128",
	TBOOL:       "bool",
	TANY:        "any",
	TSTRING:     "string",
	TNIL:        "nil",
	TIDEAL:      "untyped number",
	TBLANK:      "blank",
}

// Format implements formatting for a Type.
// The valid formats are:
//
//	%v	Go syntax
//	%+v	Debug syntax: Go syntax with a KIND- prefix for all but builtins.
//	%L	Go syntax for underlying type if t is named
//	%S	short Go syntax: drop leading "func" in function type
//	%-S	special case for method receiver symbol
func (t *Type) Format(s fmt.State, verb rune)

// String returns the Go syntax for the type t.
func (t *Type) String() string

// LinkString returns a string description of t, suitable for use in
// link symbols.
//
// The description corresponds to type identity. That is, for any pair
// of types t1 and t2, Identical(t1, t2) == (t1.LinkString() ==
// t2.LinkString()) is true. Thus it's safe to use as a map key to
// implement a type-identity-keyed map.
func (t *Type) LinkString() string

// NameString generates a user-readable, mostly unique string
// description of t. NameString always returns the same description
// for identical types, even across compilation units.
//
// NameString qualifies identifiers by package name, so it has
// collisions when different packages share the same names and
// identifiers. It also does not distinguish function-scope defined
// types from package-scoped defined types or from each other.
func (t *Type) NameString() string

// SplitVargenSuffix returns name split into a base string and a ·N
// suffix, if any.
func SplitVargenSuffix(name string) (base, suffix string)

// TypeHash computes a hash value for type t to use in type switch statements.
func TypeHash(t *Type) uint32
