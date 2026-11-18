// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specgen

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/go/types"
	"github.com/shogo82148/std/simd/archsimd/_gen/specgen/specexpr"
)

// Func represents a function or method in the SIMD API.
type Func struct {
	Name string

	// Doc is the function documentation, split into classified paragraphs.
	Doc []Paragraph

	// Commutative indicates that this operation produces the same result
	// regardless of the order of its arguments in Recv and In.
	Commutative bool

	// The general category of this function or method (useful for overview
	// documentation).
	Category string

	// Experiment indicates that this operation is experimental and should
	// remain behind a GOEXPERIMENT (from the `//specgen:experiment`
	// directive).
	Experiment bool

	// Recv, if non-zero, is the shape of the receiver. The name of the receiver
	// is always "x".
	Recv Arg

	In  []Arg
	Out []Arg

	// Pos is the source position of this definition.
	Pos token.Position

	// specFunc and instance describe the underlying spec function and its
	// instantiation that led to this API function.
	specFunc      *specFunc
	typeParamVars map[*types.TypeParam]specexpr.Variable
	instance      *specexpr.Bindings
}

type Arg struct {
	Name string
	Type specexpr.Type
}

// TypeDecl represents a vector or mask type in the SIMD API.
//
// TypeDecls come from exported generic type declarations in spec whose
// underlying type is a Vec, such as
//
//	// Vector is a ...
//	//
//	//specgen:name {{.T}}
//	type Vector[E Elt, W Width] Vec[E, W]
//
// Each instantiation of the type parameters produces one TypeDecl.
type TypeDecl struct {
	// Name is the API type name, such as "Int32x4" or "Mask8s". It always
	// equals Shape.String().
	Name string

	// Shape is the vector shape of this type.
	Shape specexpr.Vector

	// Doc is the type documentation, split into classified paragraphs.
	Doc []Paragraph

	// Pos is the source position of the spec type declaration.
	Pos token.Position
}

func (f *Func) Signature() string

func (f *Func) Decl() string

// SpecFunc returns information about the spec package function that generated
// f. name is the name of the function in the spec package, sig is its
// uninstantiated signature type, and typeArgs is a slice of the type arguments
// it was instantiated on to construct f.
func (f *Func) SpecFunc() (name string, sig *types.Signature, typeArgs []types.Type)
