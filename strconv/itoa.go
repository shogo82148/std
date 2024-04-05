// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

// FormatUintは、2 <= base <= 36の場合、iの文字列表現を指定された基数で返します。
// 結果は、10以上の桁の値に対して小文字の 'a' から 'z' を使用します。
func FormatUint(i uint64, base int) string

// FormatIntは、2 <= base <= 36の場合、iの文字列表現を指定された基数で返します。
// 結果は、10以上の桁の値に対して小文字の 'a' から 'z' を使用します。
func FormatInt(i int64, base int) string

// Itoaは、[FormatInt](int64(i), 10)と同等です。
func Itoa(i int) string

// AppendIntは、[FormatInt] によって生成された整数iの文字列形式をdstに追加し、拡張されたバッファを返します。
func AppendInt(dst []byte, i int64, base int) []byte

// AppendUintは、[FormatUint] によって生成された符号なし整数iの文字列形式をdstに追加し、拡張されたバッファを返します。
func AppendUint(dst []byte, i uint64, base int) []byte
