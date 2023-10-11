// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// signが0以上の場合、正の無限大を返し、signが0より小さい場合は負の無限大を返します。
func Inf(sign int) float64

// NaNはIEEE 754の「非数値」を返します。
func NaN() float64

// IsNaNは、fがIEEE 754の"非数"値であるかどうかを報告します。
func IsNaN(f float64) (is bool)

// IsInfは、fが無限大であるかどうかをsignに基づいて報告します。
// sign > 0の場合、fが正の無限大であるかどうかを報告します。
// sign < 0の場合、fが負の無限大であるかどうかを報告します。
// sign == 0の場合、fがどちらかの無限大であるかどうかを報告します。
func IsInf(f float64, sign int) bool
