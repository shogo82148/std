// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// FMAは、一度の丸めのみを使用して計算されたx * y + zを返します。
// （つまり、FMAはx、y、およびzの融合乗算-加算結果を返します。）
func FMA(x, y, z float64) float64
