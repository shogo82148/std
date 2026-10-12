// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && arm64

package bridge

func (xx Uint64x2nclm) CarrylessMultiplyEven(yy Uint64x2nclm) Uint64x2nclm

func (xx Uint64x2nclm) CarrylessMultiplyOdd(yy Uint64x2nclm) Uint64x2nclm
