// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/simd/archsimd/_gen/simdgen/types"
	"github.com/shogo82148/std/simd/archsimd/_gen/unify"
)

type Operation struct {
	rawOperation

	// Go is the Go method name of this operation.
	//
	// It is derived from the raw Go method name by adding optional suffixes.
	// Currently, "Masked" is the only suffix.
	Go string

	// Documentation is the doc string for this API.
	//
	// It is computed from the raw documentation:
	//
	// - "NAME" is replaced by the Go method name.
	//
	// - For masked operation, a sentence about masking is added.
	Documentation string

	// In is the sequence of parameters to the Go method.
	//
	// For masked operations, this will have the mask operand appended.
	In []types.Operand
}

func (o *Operation) IsMasked() bool

func (o *Operation) SkipMaskedMethod() bool

func (o *Operation) DecodeUnified(v *unify.Value) error

func (o *Operation) VectorWidth() int
