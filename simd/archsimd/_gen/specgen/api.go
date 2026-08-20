// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specgen

import (
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
}

type Arg struct {
	Name string
	Type specexpr.Type
}

func (f *Func) Signature() string

func (f *Func) Decl() string
