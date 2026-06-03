//go:build goexperiment.genericmethods

// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package rand

// N returns a pseudo-random number in the half-open interval [0,n).
// The type parameter Int can be any integer type.
// It panics if n <= 0.
func (r *Rand) N[Int intType](n Int) Int
