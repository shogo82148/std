// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// W.Hormann, G.Derflinger:
// "Rejection-Inversion to Generate Variates
// from Monotone Discrete Distributions"
// http://eeyore.wu-wien.ac.at/papers/96-04-04.wh-der.ps.gz

package rand

// A Zipf generates Zipf distributed variates.
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

// NewZipf returns a Zipf generating variates p(k) on [0, imax]
// proportional to (v+k)**(-s) where s>1 and k>=0, and v>=1.
func NewZipf(r *Rand, s float64, v float64, imax uint64) *Zipf

// Uint64 returns a value drawn from the Zipf distributed described
// by the Zipf object.
func (z *Zipf) Uint64() uint64
