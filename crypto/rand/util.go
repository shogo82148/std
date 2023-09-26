// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// Prime returns a number, p, of the given size, such that p is prime
// with high probability.
func Prime(rand io.Reader, bits int) (p *big.Int, err error)

// Int returns a uniform random value in [0, max).
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
