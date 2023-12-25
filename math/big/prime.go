// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

// ProbablyPrimeは、xがおそらく素数であるかどうかを報告します。
// これは、n個の擬似ランダムに選ばれた基底とともにミラー-ラビンテストを適用し、
// ベイリー-PSWテストも行います。
//
// xが素数の場合、ProbablyPrimeはtrueを返します。
// xがランダムに選ばれ、素数でない場合、ProbablyPrimeはおそらくfalseを返します。
// ランダムに選ばれた非素数に対してtrueを返す確率は最大で¼ⁿです。
//
// ProbablyPrimeは、2⁶⁴未満の入力に対しては100%正確です。
// エラー確率の詳細な議論については、Menezesらの「Handbook of Applied Cryptography」、1997年、pp. 145-149、
// およびFIPS 186-4 Appendix Fを参照してください。
//
// ProbablyPrimeは、敵がテストを欺くために作成した可能性のある素数を判断するのに適していません。
//
// Go 1.8以降、ProbablyPrime(0)が許可され、Baillie-PSWテストのみが適用されます。
// Go 1.8以前では、ProbablyPrimeはMiller-Rabinテストのみを適用し、ProbablyPrime(0)はパニックを引き起こしました。
func (x *Int) ProbablyPrime(n int) bool
