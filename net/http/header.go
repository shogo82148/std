// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

// Headerは、HTTPヘッダー内のキーと値のペアを表します。
<<<<<<< HEAD
//
=======

>>>>>>> release-branch.go1.22
// キーは、[CanonicalHeaderKey] が返すように正規化された形式である必要があります。
type Header map[string][]string

// Addは、ヘッダーにキーと値のペアを追加します。
// キーに関連付けられた既存の値に追加します。
// キーは大文字小文字を区別せず、[CanonicalHeaderKey] によって正規化されます。
func (h Header) Add(key, value string)

// Setは、キーに関連付けられたヘッダーエントリを単一の要素値に設定します。
// キーに関連付けられた既存の値を置き換えます。
// キーは大文字小文字を区別せず、[textproto.CanonicalMIMEHeaderKey] によって正規化されます。
// 非正規のキーを使用するには、直接マップに割り当ててください。
func (h Header) Set(key, value string)

// Getは、指定されたキーに関連付けられた最初の値を取得します。
// キーに関連付けられた値がない場合、Getは""を返します。
// 大文字小文字を区別せず、[textproto.CanonicalMIMEHeaderKey] が提供されたキーを正規化することに注意してください。
// Getは、すべてのキーが正規形式で格納されていると想定しています。非正規のキーを使用するには、直接マップにアクセスしてください。
func (h Header) Get(key string) string

// Valuesは、指定されたキーに関連付けられたすべての値を返します。
// 大文字小文字を区別せず、[textproto.CanonicalMIMEHeaderKey] が提供されたキーを正規化することに注意してください。
// 非正規のキーを使用するには、直接マップにアクセスしてください。
// 返されるスライスはコピーではありません。
func (h Header) Values(key string) []string

// Delは、キーに関連付けられた値を削除します。
// キーは大文字小文字を区別せず、[CanonicalHeaderKey] によって正規化されます。
func (h Header) Del(key string)

// Writeは、ワイヤーフォーマットでヘッダーを書き込みます。
func (h Header) Write(w io.Writer) error

// Cloneは、hのコピーを返します。hがnilの場合はnilを返します。
func (h Header) Clone() Header

// ParseTimeは、HTTP/1.1で許可されている3つのフォーマットのうちの1つ、[TimeFormat]、[time.RFC850]、および [time.ANSIC] をそれぞれ試して、時間ヘッダー（Date：ヘッダーなど）を解析します。
func ParseTime(text string) (t time.Time, err error)

// WriteSubsetはワイヤーフォーマットでヘッダーを書き込みます。
// excludeがnilでない場合、exclude[key] == trueのキーは書き込まれません。
// excludeマップをチェックする前にキーは正規化されません。
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error

// CanonicalHeaderKeyは、ヘッダーキーsの正規形式を返します。
// 正規化により、最初の文字とハイフンに続く任意の文字が大文字に変換されます。
// それ以外の文字は小文字に変換されます。
// たとえば、「accept-encoding」の正規キーは「Accept-Encoding」です。
// sにスペースまたは無効なヘッダーフィールドバイトが含まれている場合、変更せずに返されます。
func CanonicalHeaderKey(s string) string
