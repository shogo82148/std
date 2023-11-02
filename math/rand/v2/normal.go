// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

// NormFloat64は、標準正規分布（平均 = 0、標準偏差 = 1）に従う
// 範囲[-math.MaxFloat64, +math.MaxFloat64]の正規分布のfloat64を返します。
// 異なる正規分布を生成するために、呼び出し元は出力を調整できます：
//
//	sample = NormFloat64() * desiredStdDev + desiredMean
func (r *Rand) NormFloat64() float64
