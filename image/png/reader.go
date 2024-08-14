// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// pngパッケージは、PNG画像のデコーダとエンコーダを実装します。
//
// PNGの仕様は https://www.w3.org/TR/PNG/ にあります。
package png

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/io"
)

// FormatErrorは、入力が有効なPNGではないことを報告します。
type FormatError string

func (e FormatError) Error() string

// UnsupportedErrorは、入力が有効だが未実装のPNG機能を使用していることを報告します。
type UnsupportedError string

func (e UnsupportedError) Error() string

// Decodeは、rからPNG画像を読み取り、それを [image.Image] として返します。
// 返されるImageの型は、PNGの内容に依存します。
func Decode(r io.Reader) (image.Image, error)

// DecodeConfigは、画像全体をデコードすることなく、PNG画像のカラーモデルと寸法を返します。
func DecodeConfig(r io.Reader) (image.Config, error)
