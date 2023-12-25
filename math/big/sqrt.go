// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

// Sqrtは、zをxの四捨五入した平方根に設定し、それを返します。
//
// もしzの精度が0なら、操作前にxの精度に変更されます。丸めはzの精度と
// 丸めモードに従って行われますが、zの精度は計算されません。具体的には、
// z.Acc()の結果は未定義です。
//
// もしz < 0なら、関数はパニックを引き起こします。その場合、zの値は未定義です。
func (z *Float) Sqrt(x *Float) *Float
