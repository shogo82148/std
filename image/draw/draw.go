// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージdrawは、画像合成関数を提供します。
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

<<<<<<< HEAD
// RGBA64Image extends both the [Image] and [image.RGBA64Image] interfaces with a
// SetRGBA64 method to change a single pixel. SetRGBA64 is equivalent to
// calling Set, but it can avoid allocations from converting concrete color
// types to the [color.Color] interface type.
=======
// RGBA64Imageは、単一のピクセルを変更するSetRGBA64メソッドで、Imageとimage.RGBA64Imageの
// インターフェースを拡張します。SetRGBA64はSetを呼び出すのと同等ですが、具体的な色の
// タイプをcolor.Colorインターフェースタイプに変換する際の割り当てを避けることができます。
>>>>>>> release-branch.go1.21
type RGBA64Image interface {
	image.RGBA64Image
	Set(x, y int, c color.Color)
	SetRGBA64(x, y int, c color.RGBA64)
}

// Quantizerは、画像のパレットを生成します。
type Quantizer interface {
<<<<<<< HEAD
=======
	// Quantizeは、最大 cap(p) - len(p) の色をpに追加し、mをパレット化された画像に
	// 変換するための適切なパレットを返します。
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// Draw implements the [Drawer] interface by calling the Draw function with this
// [Op].
func (op Op) Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)

// Drawer contains the [Draw] method.
type Drawer interface {
	Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
}

// FloydSteinberg is a [Drawer] that is the [Src] [Op] with Floyd-Steinberg error
// diffusion.
var FloydSteinberg Drawer = floydSteinberg{}

// Draw calls [DrawMask] with a nil mask.
=======
// Drawは、このOpとともにDraw関数を呼び出すことでDrawerインターフェースを実装します。
func (op Op) Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)

// DrawerはDrawメソッドを含みます。
type Drawer interface {
	// Drawは、dstのr.Minをsrcのspに揃え、その後dstの矩形rをsrcをdstに描画した結果で置き換えます。
	Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
}

// FloydSteinbergは、Floyd-Steinberg誤差拡散を持つSrc OpのDrawerです。
var FloydSteinberg Drawer = floydSteinberg{}

// Drawは、nilマスクでDrawMaskを呼び出します。
>>>>>>> release-branch.go1.21
func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)

// DrawMaskは、dstのr.Minをsrcのspとmaskのmpに揃え、その後dstの矩形rを
// ポーター-ダフ合成の結果で置き換えます。nilのマスクは不透明として扱われます。
func DrawMask(dst Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op Op)
