// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

// AllocsPerRunは、関数fの呼び出し中に行われる平均的な割り当ての数を返します。
// 返り値はfloat64型ですが、常に整数値になります。
//
// 割り当ての数を計算するために、まず関数はウォームアップとして一度実行されます。
// 指定された回数の実行における平均的な割り当ての数が測定され、返されます。
//
<<<<<<< HEAD
// AllocsPerRunは、計測中にGOMAXPROCSを1に設定し、戻る前に元に戻します。
=======
// AllocsPerRun sets [runtime.GOMAXPROCS] to 1 during its measurement and will restore
// it before returning.
>>>>>>> upstream/release-branch.go1.25
func AllocsPerRun(runs int, f func()) (avg float64)
