// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// MaxStruct is the maximum number of fields a struct
// can have and still be SSAable.
const MaxStruct = 4

// StructMakeOp returns the opcode to construct a struct with the
// given number of fields.
func StructMakeOp(nf int) Op
