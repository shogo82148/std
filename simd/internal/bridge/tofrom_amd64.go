// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && amd64

package bridge

func (x Float32x16) ToArch() any

func (x Float32x8) ToArch() any

func (x Float64x4) ToArch() any

func (x Float64x8) ToArch() any

func (x Int16x16) ToArch() any

func (x Int16x32) ToArch() any

func (x Int32x16) ToArch() any

func (x Int32x8) ToArch() any

func (x Int64x4) ToArch() any

func (x Int64x8) ToArch() any

func (x Int8x32) ToArch() any

func (x Int8x64) ToArch() any

func (x Mask16x16) ToArch() any

func (x Mask16x32) ToArch() any

func (x Mask32x16) ToArch() any

func (x Mask32x8) ToArch() any

func (x Mask64x4) ToArch() any

func (x Mask64x8) ToArch() any

func (x Mask8x32) ToArch() any

func (x Mask8x64) ToArch() any

func (x Uint16x16) ToArch() any

func (x Uint16x32) ToArch() any

func (x Uint32x16) ToArch() any

func (x Uint32x8) ToArch() any

func (x Uint64x4) ToArch() any

func (x Uint64x8) ToArch() any

func (x Uint8x32) ToArch() any

func (x Uint8x64) ToArch() any
