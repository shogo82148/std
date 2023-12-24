// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big_test

import (
	"github.com/shogo82148/std/fmt"
)

// この例では、big.Ratを使用して、自然対数の基数である定数eの
// 有理数収束のシーケンスの最初の15項を計算する方法を示します。
func Example_eConvergents() {
	for i := 1; i <= 15; i++ {
		r := recur(0, int64(i))

		// rを分数と浮動小数点数の両方として印刷します。
		// big.Ratはfmt.Formatterを実装しているので、%-13sを使用して
		// 分数の左揃えの文字列表現を取得することができます。
		fmt.Printf("%-13s = %s\n", r, r.FloatString(8))
	}

	// Output:
	// 2/1           = 2.00000000
	// 3/1           = 3.00000000
	// 8/3           = 2.66666667
	// 11/4          = 2.75000000
	// 19/7          = 2.71428571
	// 87/32         = 2.71875000
	// 106/39        = 2.71794872
	// 193/71        = 2.71830986
	// 1264/465      = 2.71827957
	// 1457/536      = 2.71828358
	// 2721/1001     = 2.71828172
	// 23225/8544    = 2.71828184
	// 25946/9545    = 2.71828182
	// 49171/18089   = 2.71828183
	// 517656/190435 = 2.71828183
}
