// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// gifパッケージは、GIF画像のデコーダとエンコーダを実装します。
//
// GIF仕様は https://www.w3.org/Graphics/GIF/spec-gif89a.txt にあります。
//
// 信頼できない入力をデコードする場合は、[Decode] または [DecodeAll] を呼び出す前に
// [DecodeConfig] でディメンションを読み取ってください。これらの関数と [image] パッケージ
// ドキュメントの「Security Considerations」セクションを参照してください。
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

// DecodeはrからGIF画像を読み取り、最初に埋め込まれた画像を
// [image.Image] として返します。
//
// 信頼できないソースから画像をデコードする場合、
// 最初に DecodeConfig を呼び出して画像サイズをチェックして
// 下さい。詳算外の大きなメモリ割り当てを会を回避できます。
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

// DecodeAllはrからGIF画像を読み取り、連続したフレームと
// タイミング情報を返します。
//
// [Decode] と同様に、画像記述子の幅と高さからフレームごとに
// パレットバッファを割り当てます。[DecodeAll] はすべてのデコードされたフレームを
// メモリに留めます。信頼できない入力の場合、最初に [DecodeConfig] を呼び出して
// 詰輿画面サイズを検証し、過大なメモリを要求する入力を拒否して下さい。
func DecodeAll(r io.Reader) (*GIF, error)

// DecodeConfigはGIF画像全体をデコードせずに、グローバルカラーモデルと
// ディメンションを返します。
//
// 詰輿画面記述子とグローバルカラーテーブルのみを読み取り、
// フレームのピクセルバッファを割り当てません。幅と高さを検証したい場合、
// [Decode] または [DecodeAll] を呼び出す前に使用して下さい。
func DecodeConfig(r io.Reader) (image.Config, error)
