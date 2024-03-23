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

<<<<<<< HEAD
// ErrInvalidMediaParameterは、ParseMediaTypeによって
// メディアタイプの値が見つかったが、オプションのパラメータの解析でエラーがあった場合に返されます
var ErrInvalidMediaParameter = errors.New("mime: invalid media parameter")

// ParseMediaTypeは、RFC 1521に従ってメディアタイプの値と任意の
// パラメータを解析します。メディアタイプは、
// Content-TypeとContent-Dispositionヘッダー（RFC 2183）の値です。
// 成功すると、ParseMediaTypeは空白を削除して小文字に変換したメディアタイプと
// nilでないマップを返します。
// オプションのパラメータの解析でエラーがある場合、
// メディアタイプとエラーErrInvalidMediaParameterが返されます。
// 返されるマップ、paramsは、小文字の属性から
// ケースが保持された属性値へのマップです。
=======
// ErrInvalidMediaParameter is returned by [ParseMediaType] if
// the media type value was found but there was an error parsing
// the optional parameters
var ErrInvalidMediaParameter = errors.New("mime: invalid media parameter")

// ParseMediaType parses a media type value and any optional
// parameters, per RFC 1521.  Media types are the values in
// Content-Type and Content-Disposition headers (RFC 2183).
// On success, ParseMediaType returns the media type converted
// to lowercase and trimmed of white space and a non-nil map.
// If there is an error parsing the optional parameter,
// the media type will be returned along with the error
// [ErrInvalidMediaParameter].
// The returned map, params, maps from the lowercase
// attribute to the attribute value with its case preserved.
>>>>>>> upstream/master
func ParseMediaType(v string) (mediatype string, params map[string]string, err error)
