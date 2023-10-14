// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textproto

// MIMEHeaderは、キーと値のセットへのマッピングを表すMIMEスタイルのヘッダーです。
type MIMEHeader map[string][]string

// Addは、キーと値のペアをヘッダーに追加します。
// キーに関連付けられている既存の値に追記されます。
func (h MIMEHeader) Add(key, value string)

// Setは、キーに関連付けられたヘッダーエントリを単一の要素"value"に設定します。既存の値は置き換えられます。
func (h MIMEHeader) Set(key, value string)

// Getは与えられたキーに関連付けられた最初の値を取得します。
// 大文字と小文字を区別しないです。CanonicalMIMEHeaderKeyが提供されたキーを正規化するために使用されます。
// キーに関連付けられた値がない場合、Getは "" を返します。
// 正規化されていないキーを使用する場合は、直接マップにアクセスしてください。
func (h MIMEHeader) Get(key string) string

// Valuesは与えられたキーに関連付けられたすべての値を返します。
// 大文字小文字を区別しません。CanonicalMIMEHeaderKeyを使って提供されたキーを正規化します。
// 正規化されていないキーを使用する場合は、マップに直接アクセスしてください。
// 返されるスライスはコピーされません。
func (h MIMEHeader) Values(key string) []string

// Delはkeyに関連付けられた値を削除します。
func (h MIMEHeader) Del(key string)
