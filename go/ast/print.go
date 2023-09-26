// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains printing support for ASTs.

package ast

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/reflect"
)

// A FieldFilter may be provided to Fprint to control the output.
type FieldFilter func(name string, value reflect.Value) bool

// NotNilFilter returns true for field values that are not nil;
// it returns false otherwise.
func NotNilFilter(_ string, v reflect.Value) bool

// Fprint prints the (sub-)tree starting at AST node x to w.
// If fset != nil, position information is interpreted relative
// to that file set. Otherwise positions are printed as integer
// values (file set specific offsets).
//
// A non-nil FieldFilter f may be provided to control the output:
// struct fields for which f(fieldname, fieldvalue) is true are
// are printed; all others are filtered from the output. Unexported
// struct fields are never printed.
func Fprint(w io.Writer, fset *token.FileSet, x interface{}, f FieldFilter) (err error)

// Print prints x to standard output, skipping nil fields.
// Print(fset, x) is the same as Fprint(os.Stdout, fset, x, NotNilFilter).
func Print(fset *token.FileSet, x interface{}) error

// localError wraps locally caught errors so we can distinguish
// them from genuine panics which we don't want to return as errors.
