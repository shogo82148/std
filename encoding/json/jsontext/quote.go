// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

// AppendQuoteは、srcを表す二重引用符付きのJSON文字列リテラルを
// dstに追加し、拡張されたバッファを返します。
// RFC 8785の3.2.2.2節に従い、最小限の文字列表現を使用します。
// 不正なUTF-8バイトはUnicodeの置換文字に置き換えられ、
// 不正なUTF-8が存在した場合はエラーが返されます。
// dstはsrcと重複していてはいけません。
func AppendQuote[Bytes ~[]byte | ~string](dst []byte, src Bytes) ([]byte, error)

// AppendUnquoteは、srcをデコードした二重引用符付きJSON文字列リテラルとして
// dstに追加し、拡張されたバッファを返します。
// 入力srcは前後に空白のないJSON文字列である必要があります。
// 不正なUTF-8バイトはUnicodeの置換文字に置き換えられ、
// 不正なUTF-8が存在した場合は最後にエラーが返されます。
// JSON文字列リテラルの後に余分なバイトがある場合はエラーとなります。
// dstはsrcと重複していてはいけません。
func AppendUnquote[Bytes ~[]byte | ~string](dst []byte, src Bytes) ([]byte, error)
