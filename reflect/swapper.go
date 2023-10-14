// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

// Swapperは、与えられたスライス内の要素を交換する関数を返します。
//
// 与えられたインターフェースがスライスでない場合、Swapperはパニックを起こします。
func Swapper(slice any) func(i, j int)
