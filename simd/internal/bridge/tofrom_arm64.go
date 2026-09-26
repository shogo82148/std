// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && arm64

package bridge

func (x Float32x4nclm) ToArch() any

func (x Float64x2nclm) ToArch() any

func (x Int16x8nclm) ToArch() any

func (x Int32x4nclm) ToArch() any

func (x Int64x2nclm) ToArch() any

func (x Int8x16nclm) ToArch() any

func (x Mask16x8nclm) ToArch() any

func (x Mask32x4nclm) ToArch() any

func (x Mask64x2nclm) ToArch() any

func (x Mask8x16nclm) ToArch() any

func (x Uint16x8nclm) ToArch() any

func (x Uint32x4nclm) ToArch() any

func (x Uint64x2nclm) ToArch() any

func (x Uint8x16nclm) ToArch() any
