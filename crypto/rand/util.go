// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

<<<<<<< HEAD
// Primeは与えられたビット長で高い確率で素数である数を返します。
// Primeはrand.Readが返すエラーやbits < 2の場合にエラーを返します。
func Prime(rand io.Reader, bits int) (*big.Int, error)
=======
// Prime returns a number of the given bit length that is prime with high probability.
// Prime will return error for any error returned by rand.Read or if bits < 2.
//
// Since Go 1.26, a secure source of random bytes is always used, and the Reader is
// ignored unless GODEBUG=cryptocustomrand=1 is set. This setting will be removed
// in a future Go release. Instead, use [testing/cryptotest.SetGlobalRandom].
func Prime(r io.Reader, bits int) (*big.Int, error)
>>>>>>> upstream/release-branch.go1.26

// Intは[0, max)の範囲の一様乱数値を返します。max <= 0の場合はパニックし、
// rand.Readがエラーを返した場合はエラーを返します。
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
