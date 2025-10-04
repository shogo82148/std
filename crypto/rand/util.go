// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

<<<<<<< HEAD
// Prime は与えられたビット長の数が高確率で素数である場合に返します。
// Prime は [rand.Read] によって返されるエラーまたは bits < 2 の場合にはエラーを返します。
func Prime(rand io.Reader, bits int) (*big.Int, error)

// Intは[0、max)の一様乱数を返します。max <= 0の場合、パニックを発生させます。
=======
// Prime returns a number of the given bit length that is prime with high probability.
// Prime will return error for any error returned by rand.Read or if bits < 2.
func Prime(rand io.Reader, bits int) (*big.Int, error)

// Int returns a uniform random value in [0, max). It panics if max <= 0, and
// returns an error if rand.Read returns one.
>>>>>>> upstream/release-branch.go1.25
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
