// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && amd64

package simd_test

import (
	"github.com/shogo82148/std/fmt"
)

func ExampleInt16s_ToArch() {
	xs := make([]int16, 50)
	ys := make([]int16, 50)
	for i := range xs {
		xs[i] = int16(i + 1)
		ys[i] = 2
	}
	fmt.Println(dotProduct(xs, ys))
	// Output: 2550
}
