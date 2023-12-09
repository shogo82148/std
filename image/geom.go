// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"github.com/shogo82148/std/image/color"
)

// Pointは、X、Y座標のペアです。軸は右と下に増加します。
type Point struct {
	X, Y int
}

// Stringは、pの文字列表現を"(3,4)"のように返します。
func (p Point) String() string

// Addはベクトルp+qを返します。
func (p Point) Add(q Point) Point

// Subはベクトルp-qを返します。
func (p Point) Sub(q Point) Point

// Mulはベクトルp*kを返します。
func (p Point) Mul(k int) Point

// Divはベクトルp/kを返します。
func (p Point) Div(k int) Point

// Inは、pがr内にあるかどうかを報告します。
func (p Point) In(r Rectangle) bool

// Modは、p.X-q.Xがrの幅の倍数で、p.Y-q.Yがrの高さの倍数となるような、r内の点qを返します。
func (p Point) Mod(r Rectangle) Point

// Eqは、pとqが等しいかどうかを報告します。
func (p Point) Eq(q Point) bool

// ZPはゼロ [Point] です。
//
// Deprecated: 代わりにリテラルの [image.Point] を使用してください。
var ZP Point

// Ptは [Point]{X, Y}の省略形です。
func Pt(X, Y int) Point

// Rectangleは、Min.X <= X < Max.X、Min.Y <= Y < Max.Yの点を含みます。
// Min.X <= Max.XおよびYについても同様に成り立つ場合、それは整形されています。
// 点は常に整形されています。矩形のメソッドは、整形された入力に対して常に整形された出力を返します。
//
// Rectangleは、その境界が矩形自体である [Image] でもあります。Atは、
// 矩形内の点に対してcolor.Opaqueを、それ以外の場合はcolor.Transparentを返します。
type Rectangle struct {
	Min, Max Point
}

// Stringは、rの文字列表現を"(3,4)-(6,5)"のように返します。
func (r Rectangle) String() string

// Dxは、rの幅を返します。
func (r Rectangle) Dx() int

// Dyは、rの高さを返します。
func (r Rectangle) Dy() int

// Sizeは、rの幅と高さを返します。
func (r Rectangle) Size() Point

// Addは、pによって移動された矩形rを返します。
func (r Rectangle) Add(p Point) Rectangle

// Subは、-pによって移動された矩形rを返します。
func (r Rectangle) Sub(p Point) Rectangle

// Insetは、n（負の場合もあり）によって内側に移動された矩形rを返します。
// rの寸法のいずれかが2*n未満の場合、rの中心近くの空の矩形が返されます。
func (r Rectangle) Inset(n int) Rectangle

// Intersectは、rとsの両方に含まれる最大の矩形を返します。もし
// 二つの矩形が重ならない場合、ゼロ矩形が返されます。
func (r Rectangle) Intersect(s Rectangle) Rectangle

// Unionは、rとsの両方を含む最小の矩形を返します。
func (r Rectangle) Union(s Rectangle) Rectangle

// Emptyは、矩形が点を含まないかどうかを報告します。
func (r Rectangle) Empty() bool

// Eqは、rとsが同じ点の集合を含むかどうかを報告します。すべての空の
// 矩形は等しいとみなされます。
func (r Rectangle) Eq(s Rectangle) bool

// Overlapsは、rとsが非空の交差点を持つかどうかを報告します。
func (r Rectangle) Overlaps(s Rectangle) bool

// Inは、rのすべての点がs内にあるかどうかを報告します。
func (r Rectangle) In(s Rectangle) bool

// Canonは、rの正規化されたバージョンを返します。返される矩形は、必要に応じて最小座標と最大座標が交換され、
// 正しく形成されています。
func (r Rectangle) Canon() Rectangle

// Atは、[Image] インターフェースを実装します。
func (r Rectangle) At(x, y int) color.Color

// RGBA64Atは、[RGBA64Image] インターフェースを実装します。
func (r Rectangle) RGBA64At(x, y int) color.RGBA64

// Boundsは、[Image] インターフェースを実装します。
func (r Rectangle) Bounds() Rectangle

// ColorModelは、[Image] インターフェースを実装します。
func (r Rectangle) ColorModel() color.Model

// ZRはゼロ [Rectangle] です。
//
// Deprecated: 代わりにリテラルの [image.Rectangle] を使用してください。
var ZR Rectangle

// Rectは [Rectangle]{Pt(x0, y0), Pt(x1, y1)}の省略形です。返される
// 矩形は、必要に応じて最小座標と最大座標が交換され、正しく形成されています。
func Rect(x0, y0, x1, y1 int) Rectangle
