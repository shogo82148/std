// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
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
	In []Operand
}

func (o *Operation) IsMasked() bool

func (o *Operation) SkipMaskedMethod() bool

func (o *Operation) DecodeUnified(v *unify.Value) error

func (o *Operation) VectorWidth() int

type Operand struct {
	Class string

	Go     *string
	AsmPos int

	Base     *string
	ElemBits *int
	Bits     *int

	Const *string
	// Optional immediate arg offsets. If this field is non-nil,
	// This operand will be an immediate operand:
	// The compiler will right-shift the user-passed value by ImmOffset and set it as the AuxInt
	// field of the operation.
	ImmOffset *string
	Name      *string
	Lanes     *int
	// TreatLikeAScalarOfSize means only the lower $TreatLikeAScalarOfSize bits of the vector
	// is used, so at the API level we can make it just a scalar value of this size; Then we
	// can overwrite it to a vector of the right size during intrinsics stage.
	TreatLikeAScalarOfSize *int
	// If non-nil, it means the [Class] field is overwritten here, right now this is used to
	// overwrite the results of AVX2 compares to masks.
	OverwriteClass *string
	// If non-nil, it means the [Base] field is overwritten here. This field exist solely
	// because Intel's XED data is inconsistent. e.g. VANDNP[SD] marks its operand int.
	OverwriteBase *string
	// If non-nil, it means the [ElementBits] field is overwritten. This field exist solely
	// because Intel's XED data is inconsistent. e.g. AVX512 VPMADDUBSW marks its operand
	// elemBits 16, which should be 8.
	OverwriteElementBits *int
	// FixedReg is the name of the fixed registers
	FixedReg *string
}
