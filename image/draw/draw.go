// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// drawパッケージは、画像合成関数を提供します。
//
// このパッケージの紹介については、「Go image/drawパッケージ」を参照してください：
// https://golang.org/doc/articles/image_draw.html
package draw

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/image/color"
)

// Imageは、単一のピクセルを変更するSetメソッドを持つimage.Imageです。
type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}

// RGBA64Imageは、単一のピクセルを変更するSetRGBA64メソッドで、[Image] と [image.RGBA64Image] の
// インターフェースを拡張します。SetRGBA64はSetを呼び出すのと同等ですが、具体的な色の
// タイプを [color.Color] インターフェースタイプに変換する際の割り当てを避けることができます。
type RGBA64Image interface {
	image.RGBA64Image
	Set(x, y int, c color.Color)
	SetRGBA64(x, y int, c color.RGBA64)
}

// Quantizerは、画像のパレットを生成します。
type Quantizer interface {
	Quantize(p color.Palette, m image.Image) color.Palette
}

// Opは、ポーター-ダフ合成演算子です。
type Op int

const (
	// Overは ``(src in mask) over dst'' を指定します。
	Over Op = iota
	// Srcは ``src in mask'' を指定します。
	Src
)

// Drawは、この [Op] とともにDraw関数を呼び出すことで [Drawer] インターフェースを実装します。
func (op Op) Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)

// Drawerは [Draw] メソッドを含みます。
type Drawer interface {
	Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
}

// FloydSteinbergは、Floyd-Steinberg誤差拡散を持つ [Src] [Op] の [Drawer] です。
var FloydSteinberg Drawer = floydSteinberg{}

// Drawは、nilマスクで [DrawMask] を呼び出します。
func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)

// DrawMaskは、dstのr.Minをsrcのspとmaskのmpに揃え、その後dstの矩形rを
// ポーター-ダフ合成の結果で置き換えます。nilのマスクは不透明として扱われます。
func DrawMask(dst Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op Op)
