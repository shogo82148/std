// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// colorパッケージは基本的な色ライブラリを実装します。
package color

// Colorは、アルファ乗算済みの16ビットチャンネルRGBAに自身を変換できます。
// この変換は損失を伴う可能性があります。
type Color interface {
	RGBA() (r, g, b, a uint32)
}

// RGBAは、伝統的な32ビットのアルファ乗算済みカラーを表し、赤、緑、青、アルファそれぞれが8ビットです。
//
// アルファ乗算済みの色成分Cは、アルファ(A)によってスケーリングされているため、
// 0 <= C <= Aの有効な値を持ちます。
type RGBA struct {
	R, G, B, A uint8
}

func (c RGBA) RGBA() (r, g, b, a uint32)

// RGBA64は、64ビットのアルファ乗算済みカラーを表し、赤、緑、青、アルファそれぞれが16ビットです。
//
// アルファ乗算済みの色成分Cは、アルファ(A)によってスケーリングされているため、
// 0 <= C <= Aの有効な値を持ちます。
type RGBA64 struct {
	R, G, B, A uint16
}

func (c RGBA64) RGBA() (r, g, b, a uint32)

// NRGBAは、アルファ非乗算の32ビットカラーを表します。
type NRGBA struct {
	R, G, B, A uint8
}

func (c NRGBA) RGBA() (r, g, b, a uint32)

// NRGBA64は、アルファ非乗算の64ビットカラーを表し、
// 赤、緑、青、アルファそれぞれが16ビットです。
type NRGBA64 struct {
	R, G, B, A uint16
}

func (c NRGBA64) RGBA() (r, g, b, a uint32)

// Alphaは、8ビットのアルファカラーを表します。
type Alpha struct {
	A uint8
}

func (c Alpha) RGBA() (r, g, b, a uint32)

// Alpha16は、16ビットのアルファカラーを表します。
type Alpha16 struct {
	A uint16
}

func (c Alpha16) RGBA() (r, g, b, a uint32)

// Grayは、8ビットのグレースケールカラーを表します。
type Gray struct {
	Y uint8
}

func (c Gray) RGBA() (r, g, b, a uint32)

// Gray16は、16ビットのグレースケールカラーを表します。
type Gray16 struct {
	Y uint16
}

func (c Gray16) RGBA() (r, g, b, a uint32)

// Modelは、任意の [Color] を自身のカラーモデルのものに変換できます。この変換は損失を伴う可能性があります。
type Model interface {
	Convert(c Color) Color
}

// ModelFuncは、変換を実装するためにfを呼び出す [Model] を返します。
func ModelFunc(f func(Color) Color) Model

// 標準のカラータイプのモデル。
var (
	RGBAModel    Model = ModelFunc(rgbaModel)
	RGBA64Model  Model = ModelFunc(rgba64Model)
	NRGBAModel   Model = ModelFunc(nrgbaModel)
	NRGBA64Model Model = ModelFunc(nrgba64Model)
	AlphaModel   Model = ModelFunc(alphaModel)
	Alpha16Model Model = ModelFunc(alpha16Model)
	GrayModel    Model = ModelFunc(grayModel)
	Gray16Model  Model = ModelFunc(gray16Model)
)

// Paletteは色のパレットです。
type Palette []Color

// Convertは、ユークリッドのR,G,B空間でcに最も近いパレット色を返します。
func (p Palette) Convert(c Color) Color

// Indexは、ユークリッドのR,G,B,A空間でcに最も近いパレット色のインデックスを返します。
func (p Palette) Index(c Color) int

// 標準的な色。
var (
	Black       = Gray16{0}
	White       = Gray16{0xffff}
	Transparent = Alpha16{0}
	Opaque      = Alpha16{0xffff}
)
