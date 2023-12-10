// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// W.Hormann, G.Derflinger:
// "Rejection-Inversion to Generate Variates
// from Monotone Discrete Distributions"
// http://eeyore.wu-wien.ac.at/papers/96-04-04.wh-der.ps.gz

package rand

// Zipfは、Zipf分布に従う変量を生成します。
type Zipf struct {
	r            *Rand
	imax         float64
	v            float64
	q            float64
	s            float64
	oneminusQ    float64
	oneminusQinv float64
	hxm          float64
	hx0minusHxm  float64
}

// NewZipfは、[Zipf] 変量ジェネレータを返します。
// このジェネレータは、P(k)が(v + k) ** (-s)に比例するような値k ∈ [0, imax]を生成します。
// 要件: s > 1 かつ v >= 1。
func NewZipf(r *Rand, s float64, v float64, imax uint64) *Zipf

// Uint64は、Zipfオブジェクトで記述された [Zipf] 分布から抽出された値を返します。
func (z *Zipf) Uint64() uint64
