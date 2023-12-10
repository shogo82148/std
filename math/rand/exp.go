// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

// ExpFloat64は、レートパラメータ（ラムダ）が1で、平均が1/ラムダ（1）である指数分布に従う、
// 範囲（0, +[math.MaxFloat64]]の指数分布のfloat64を返します。
// 異なるレートパラメータの分布を生成するには、呼び出し元は出力を調整できます：
//
//	sample = ExpFloat64() / desiredRateParameter
func (r *Rand) ExpFloat64() float64
