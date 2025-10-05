// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// Primeは与えられたビット長で高い確率で素数である数を返します。
// Primeはrand.Readが返すエラーやbits < 2の場合にエラーを返します。
func Prime(rand io.Reader, bits int) (*big.Int, error)

// Intは[0, max)の範囲の一様乱数値を返します。max <= 0の場合はパニックし、
// rand.Readがエラーを返した場合はエラーを返します。
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
