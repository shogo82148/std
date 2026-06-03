// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && (amd64 || wasm || arm64)

package bridge

func (x Float32s) ToArch() any

func (x Float64s) ToArch() any

func (x Int16s) ToArch() any

func (x Int32s) ToArch() any

func (x Int64s) ToArch() any

func (x Int8s) ToArch() any

func (x Mask16s) ToArch() any

func (x Mask32s) ToArch() any

func (x Mask64s) ToArch() any

func (x Mask8s) ToArch() any

func (x Uint16s) ToArch() any

func (x Uint32s) ToArch() any

func (x Uint64s) ToArch() any

func (x Uint8s) ToArch() any
