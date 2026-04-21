// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// Primeは、指定されたビット長で高い確率で素数である数値を返します。
// Primeは、rand.Readが返すエラーまたはbits < 2の場合にエラーを返します。
//
// Go 1.26以降、安全なランダムバイトソースが常に使用され、GODEBUG=cryptocustomrand=1が
// 設定されない限りReaderは無視されます。この設定は将来のGoリリースで削除されます。
// 代わりに [testing/cryptotest.SetGlobalRandom] を使用してください。
func Prime(r io.Reader, bits int) (*big.Int, error)

// Intは[0, max)の範囲の一様乱数値を返します。max <= 0の場合はパニックし、
// rand.Readがエラーを返した場合はエラーを返します。
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
