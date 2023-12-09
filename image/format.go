// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// ErrFormatは、デコードが未知のフォーマットに遭遇したことを示します。
var ErrFormat = errors.New("image: unknown format")

<<<<<<< HEAD
// RegisterFormat registers an image format for use by [Decode].
// Name is the name of the format, like "jpeg" or "png".
// Magic is the magic prefix that identifies the format's encoding. The magic
// string can contain "?" wildcards that each match any one byte.
// [Decode] is the function that decodes the encoded image.
// [DecodeConfig] is the function that decodes just its configuration.
=======
// RegisterFormatは、Decodeによって使用される画像フォーマットを登録します。
// Nameはフォーマットの名前で、"jpeg"や"png"のようなものです。
// Magicは、フォーマットのエンコーディングを識別するマジックプレフィックスです。マジック
// 文字列は、それぞれ任意の1バイトにマッチする"?"ワイルドカードを含むことができます。
// Decodeは、エンコードされた画像をデコードする関数です。
// DecodeConfigは、その設定だけをデコードする関数です。
>>>>>>> release-branch.go1.21
func RegisterFormat(name, magic string, decode func(io.Reader) (Image, error), decodeConfig func(io.Reader) (Config, error))

// Decodeは、登録されたフォーマットでエンコードされた画像をデコードします。
// 返される文字列は、フォーマット登録時に使用されたフォーマット名です。
// フォーマットの登録は、通常、コーデック固有のパッケージのinit関数によって行われます。
func Decode(r io.Reader) (Image, string, error)

// DecodeConfigは、登録されたフォーマットでエンコードされた画像のカラーモデルと寸法をデコードします。
// 返される文字列は、フォーマット登録時に使用されたフォーマット名です。
// フォーマットの登録は、通常、コーデック固有のパッケージのinit関数によって行われます。
func DecodeConfig(r io.Reader) (Config, string, error)
