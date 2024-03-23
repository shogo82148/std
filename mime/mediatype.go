// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mime

import (
	"github.com/shogo82148/std/errors"
)

// FormatMediaTypeは、メディアタイプtとパラメータparamをRFC 2045とRFC 2616に準拠した
// メディアタイプとしてシリアライズします。タイプとパラメータ名は小文字で書かれます。
// 引数のいずれかが標準違反を引き起こす場合、FormatMediaTypeは空の文字列を返します。
func FormatMediaType(t string, param map[string]string) string

// ErrInvalidMediaParameterは、[ParseMediaType] によって
// メディアタイプの値が見つかったが、オプションのパラメータの解析でエラーがあった場合に返されます
var ErrInvalidMediaParameter = errors.New("mime: invalid media parameter")

// ParseMediaTypeは、RFC 1521に従ってメディアタイプの値と任意の
// パラメータを解析します。メディアタイプは、
// Content-TypeとContent-Dispositionヘッダー（RFC 2183）の値です。
// 成功すると、ParseMediaTypeは空白を削除して小文字に変換したメディアタイプと
// nilでないマップを返します。
// オプションのパラメータの解析でエラーがある場合、
// メディアタイプとエラー [ErrInvalidMediaParameter] が返されます。
// 返されるマップ、paramsは、小文字の属性から
// ケースが保持された属性値へのマップです。
func ParseMediaType(v string) (mediatype string, params map[string]string, err error)
