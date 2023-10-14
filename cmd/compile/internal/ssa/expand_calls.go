// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/compile/internal/abi"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

type Abi1RO uint8

const (
	// Register offsets for fields of built-in aggregate types; the ones not listed are zero.
	RO_complex_imag = 1
	RO_string_len   = 1
	RO_slice_len    = 1
	RO_slice_cap    = 2
	RO_iface_data   = 1
)

// ParamAssignmentForArgName returns the ABIParamAssignment for f's arg with matching name.
func ParamAssignmentForArgName(f *Func, name *ir.Name) *abi.ABIParamAssignment

// ArgOpAndRegisterFor converts an abi register index into an ssa Op and corresponding
// arg register index.
func ArgOpAndRegisterFor(r abi.RegIndex, abiConfig *abi.ABIConfig) (Op, int64)
