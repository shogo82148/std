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

	// sveMergingPrefixed marks the MOVPRFX-prefixed variant of a merging
	// predicated operation, built by [Operation.sveMergingPrefixedOp]. It exists
	// only to give that variant a machine-op name of its own.
	sveMergingPrefixed bool

	// sveMergeSourceIn0 marks a merging predicated operation whose first input
	// is the value the destination starts out holding, and which therefore has
	// to share that input's register. Merging predication leaves the inactive
	// lanes of the destination alone, so that value is an operand of the
	// operation whether the instruction names it (a constructive one does, as
	// ABS <Zd>, <Pg>/M, <Zn>) or a MOVPRFX has to put it there.
	sveMergeSourceIn0 bool
}

func (o *Operation) IsMasked() bool

func (o *Operation) SkipMaskedMethod() bool

func (o *Operation) DecodeUnified(v *unify.Value) error

func (o *Operation) VectorWidth() int
