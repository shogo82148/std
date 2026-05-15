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

// EncodeはイメージmをPNG形式でwに書き込みます。任意のイメージをエンコードできますが、
// [image.NRGBA] でないイメージは非可逆的にエンコードされる可能性があります。
//
// wに書き込まれた正確なバイト数はGo 1の互換性保証の対象外です。
// テストを含む呼び出し元は、正確に書き込まれたバイト数に依存してはいけません。
func Encode(w io.Writer, m image.Image) error

// EncodeはイメージmをPNG形式でwに書き込みます。
//
// wに書き込まれた正確なバイト数はGo 1の互換性保証の対象外です。
// テストを含む呼び出し元は、正確に書き込まれたバイト数に依存してはいけません。
func (enc *Encoder) Encode(w io.Writer, m image.Image) error
