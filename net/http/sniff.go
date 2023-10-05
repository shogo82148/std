// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

// DetectContentTypeは、指定されたデータのContent-Typeを決定するために
// https://mimesniff.spec.whatwg.org/ で説明されているアルゴリズムを実装します。
// データの最初の512バイトまでしか考慮しません。
// DetectContentTypeは常に有効なMIMEタイプを返します。
// より具体的なMIMEタイプを決定できない場合は、"application/octet-stream"を返します。
func DetectContentType(data []byte) string
