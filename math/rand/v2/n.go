//go:build goexperiment.genericmethods

// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package rand

// Nは半開区間 [0,n) の疑似乱数を返します。
// 型パラメータIntには任意の整数型を指定できます。
// n <= 0 の場合、パニックします。
func (r *Rand) N[Int intType](n Int) Int
