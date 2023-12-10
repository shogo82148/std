// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージgifは、GIF画像のデコーダとエンコーダを実装します。
//
// GIFの仕様は https://www.w3.org/Graphics/GIF/spec-gif89a.txt にあります。
package gif

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/io"
)

// 廃棄方法。
const (
	DisposalNone       = 0x01
	DisposalBackground = 0x02
	DisposalPrevious   = 0x03
)

// Decodeは、rからGIF画像を読み取り、最初の埋め込み画像を [image.Image] として返します。
func Decode(r io.Reader) (image.Image, error)

// GIFは、GIFファイルに保存されている可能性のある複数の画像を表します。
type GIF struct {
	Image []*image.Paletted
	Delay []int
	// LoopCountは、表示中にアニメーションが再開される回数を制御します。
	// LoopCountが0の場合、無限にループします。
	// LoopCountが-1の場合、各フレームを一度だけ表示します。
	// それ以外の場合、アニメーションはLoopCount+1回ループします。
	LoopCount int
	// Disposalは、フレームごとの連続した廃棄方法です。後方互換性のために、
	// nil DisposalはEncodeAllに渡すことが有効であり、それぞれのフレームの廃棄方法が
	// 0（指定なしの廃棄）であることを意味します。
	Disposal []byte
	// Configは、グローバルカラーテーブル（パレット）、幅、高さです。nilまたは
	// 空のcolor.Palette Config.ColorModelは、各フレームが独自の
	// カラーテーブルを持ち、グローバルカラーテーブルがないことを意味します。各フレームの範囲は、
	// 二つの点 (0, 0) と (Config.Width, Config.Height) で定義される
	// 矩形内になければなりません。
	//
	// 後方互換性のため、ゼロ値のConfigはEncodeAllに渡すことが有効であり、
	// 全体のGIFの幅と高さが最初のフレームの範囲のRectangle.Max点と等しいことを意味します。
	Config image.Config
	// BackgroundIndexは、DisposalBackground廃棄方法で使用するための、
	// グローバルカラーテーブル内の背景インデックスです。
	BackgroundIndex byte
}

// DecodeAllは、rからGIF画像を読み取り、連続するフレームとタイミング情報を返します。
func DecodeAll(r io.Reader) (*GIF, error)

// DecodeConfigは、画像全体をデコードすることなく、GIF画像のグローバルカラーモデルと
// 寸法を返します。
func DecodeConfig(r io.Reader) (image.Config, error)
