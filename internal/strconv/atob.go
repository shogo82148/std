// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

// ParseBoolは文字列で表されるブール値を返します。
// 1、t、T、TRUE、true、True、0、f、F、FALSE、false、Falseを受け入れます。
// その他の値はエラーを返します。
func ParseBool(str string) (bool, error)

// FormatBoolはbの値に応じて"true"または"false"を返します。
func FormatBool(b bool) string

// AppendBoolはbの値に応じて、"true"または"false"をdstに追加し、拡張されたバッファを返します。
func AppendBool(dst []byte, b bool) []byte
