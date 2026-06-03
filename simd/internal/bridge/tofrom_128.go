// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && (amd64 || wasm || arm64)

package bridge

func (x Float32x4) ToArch() any

func (x Float64x2) ToArch() any

func (x Int16x8) ToArch() any

func (x Int32x4) ToArch() any

func (x Int64x2) ToArch() any

func (x Int8x16) ToArch() any

func (x Mask16x8) ToArch() any

func (x Mask32x4) ToArch() any

func (x Mask64x2) ToArch() any

func (x Mask8x16) ToArch() any

func (x Uint16x8) ToArch() any

func (x Uint32x4) ToArch() any

func (x Uint64x2) ToArch() any

func (x Uint8x16) ToArch() any
