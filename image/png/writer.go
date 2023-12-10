// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package png

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/io"
)

// Encoderは、PNG画像のエンコーディングを設定します。
type Encoder struct {
	CompressionLevel CompressionLevel

	// BufferPoolは、画像をエンコードする際に一時的な
	// EncoderBuffersを取得するためのバッファプールをオプションで指定します。
	BufferPool EncoderBufferPool
}

// EncoderBufferPoolは、一時的な [EncoderBuffer] 構造体のインスタンスを取得し、
// 返すためのインターフェースです。これは、複数の画像をエンコードする際にバッファを再利用するために使用できます。
type EncoderBufferPool interface {
	Get() *EncoderBuffer
	Put(*EncoderBuffer)
}

// EncoderBufferは、PNG画像のエンコーディングに使用されるバッファを保持します。
type EncoderBuffer encoder

// CompressionLevelは、圧縮レベルを示します。
type CompressionLevel int

const (
	DefaultCompression CompressionLevel = 0
	NoCompression      CompressionLevel = -1
	BestSpeed          CompressionLevel = -2
	BestCompression    CompressionLevel = -3
)

// Encodeは、画像mをPNG形式でwに書き込みます。任意の画像をエンコードできますが、
// [image.NRGBA] でない画像は、損失を伴ってエンコードされる可能性があります。
func Encode(w io.Writer, m image.Image) error

// Encodeは、画像mをPNG形式でwに書き込みます。
func (enc *Encoder) Encode(w io.Writer, m image.Image) error
