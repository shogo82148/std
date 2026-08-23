// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specgen

import (
	"github.com/shogo82148/std/go/types"
	"github.com/shogo82148/std/simd/archsimd/_gen/specgen/specexpr"
)

// Func represents a function or method in the SIMD API.
type Func struct {
	Name string

	// Doc is the function documentation, without any leading comment markers
	Doc string

	// Recv, if non-zero, is the shape of the receiver. The name of the receiver
	// is always "x".
	Recv Arg

	In  []Arg
	Out []Arg

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

func (f *Func) Signature() string

func (f *Func) Decl() string

// SpecFunc returns information about the spec package function that generated
// f. name is the name of the function in the spec package, sig is its
// uninstantiated signature type, and typeArgs is a slice of the type arguments
// it was instantiated on to construct f.
func (f *Func) SpecFunc() (name string, sig *types.Signature, typeArgs []types.Type)
