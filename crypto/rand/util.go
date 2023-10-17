// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// Prime は与えられたビット長の数が高確率で素数である場合に返します。
// Prime は [rand.Read] によって返されるエラーまたは bits < 2 の場合にはエラーを返します。
func Prime(rand io.Reader, bits int) (*big.Int, error)

// Intは[0、max)の一様乱数を返します。max <= 0の場合、パニックを発生させます。
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
