// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Asinは、xの弧度法におけるアークサインを返します。
//
// 特別なケースは次のとおりです：
//
//	Asin(±0) = ±0
//	Asin(x) = NaN if x < -1 or x > 1
func Asin(x float64) float64

// Acosは、xのarccosine（ラジアン単位）を返します。
//
// 特殊な場合は以下です：
//
//	Acos(x) = NaN if x < -1 or x > 1
func Acos(x float64) float64
