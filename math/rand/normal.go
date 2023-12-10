// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

<<<<<<< HEAD
// NormFloat64 returns a normally distributed float64 in
// the range -[math.MaxFloat64] through +[math.MaxFloat64] inclusive,
// with standard normal distribution (mean = 0, stddev = 1).
// To produce a different normal distribution, callers can
// adjust the output using:
=======
// NormFloat64は、標準正規分布（平均 = 0、標準偏差 = 1）に従う、
// 範囲 -math.MaxFloat64から+math.MaxFloat64（両端を含む）の正規分布のfloat64を返します。
// 異なる正規分布を生成するには、呼び出し元は出力を調整できます：
>>>>>>> release-branch.go1.21
//
//	sample = NormFloat64() * desiredStdDev + desiredMean
func (r *Rand) NormFloat64() float64
