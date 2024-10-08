// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// jpegパッケージは、JPEG画像のデコーダとエンコーダを実装します。
//
// JPEGはITU-T T.81で定義されています：https://www.w3.org/Graphics/JPEG/itu-t81.pdf。
package jpeg

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/io"
)

// FormatErrorは、入力が有効なJPEGではないことを報告します。
type FormatError string

func (e FormatError) Error() string

// UnsupportedErrorは、入力が有効だが未実装のJPEG機能を使用していることを報告します。
type UnsupportedError string

func (e UnsupportedError) Error() string

// Deprecated: Readerは [image/jpeg] パッケージによって使用されておらず、
// 他の人によっても使用されるべきではありません。互換性のために保持されています。
type Reader interface {
	io.ByteReader
	io.Reader
}

// Decodeは、rからJPEG画像を読み取り、それを [image.Image] として返します。
func Decode(r io.Reader) (image.Image, error)

// DecodeConfigは、画像全体をデコードすることなく、JPEG画像のカラーモデルと寸法を返します。
func DecodeConfig(r io.Reader) (image.Config, error)
