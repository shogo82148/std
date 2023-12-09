// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package image implements a basic 2-D image library.
//
// The fundamental interface is called [Image]. An [Image] contains colors, which
// are described in the image/color package.
//
// Values of the [Image] interface are created either by calling functions such
// as [NewRGBA] and [NewPaletted], or by calling [Decode] on an [io.Reader] containing
// image data in a format such as GIF, JPEG or PNG. Decoding any particular
// image format requires the prior registration of a decoder function.
// Registration is typically automatic as a side effect of initializing that
// format's package so that, to decode a PNG image, it suffices to have
//
//	import _ "image/png"
//
// in a program's main package. The _ means to import a package purely for its
// initialization side effects.
//
// See "The Go image package" for more details:
// https://golang.org/doc/articles/image_package.html
//
// # Security Considerations
//
// The image package can be used to parse arbitrarily large images, which can
// cause resource exhaustion on machines which do not have enough memory to
// store them. When operating on arbitrary images, [DecodeConfig] should be called
// before [Decode], so that the program can decide whether the image, as defined
// in the returned header, can be safely decoded with the available resources. A
// call to [Decode] which produces an extremely large image, as defined in the
// header returned by [DecodeConfig], is not considered a security issue,
// regardless of whether the image is itself malformed or not. A call to
// [DecodeConfig] which returns a header which does not match the image returned
// by [Decode] may be considered a security issue, and should be reported per the
// [Go Security Policy](https://go.dev/security/policy).
package image

import (
	"github.com/shogo82148/std/image/color"
)

// Configは、画像のカラーモデルと寸法を保持します。
type Config struct {
	ColorModel    color.Model
	Width, Height int
}

// Imageは、カラーモデルから取得した [color.Color] 値の有限の長方形グリッドです。
type Image interface {
	ColorModel() color.Model

	Bounds() Rectangle

	At(x, y int) color.Color
}

// RGBA64Imageは、そのピクセルを直接color.RGBA64に変換できる [Image] です。
type RGBA64Image interface {
	RGBA64At(x, y int) color.RGBA64
	Image
}

// PalettedImageは、色が限定的なパレットから来る可能性がある画像です。
// もしmがPalettedImageで、m.ColorModel()が [color.Palette] pを返すなら、
// m.At(x, y)はp[m.ColorIndexAt(x, y)]と等価であるべきです。もしmの
// カラーモデルがcolor.Paletteでないなら、ColorIndexAtの振る舞いは
// 定義されていません。
type PalettedImage interface {
	ColorIndexAt(x, y int) uint8
	Image
}

// RGBAは、Atメソッドが [color.RGBA] 値を返すインメモリイメージです。
type RGBA struct {
	// Pixは、画像のピクセルをR, G, B, Aの順序で保持します。ピクセルは
	// (x, y)はPix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4]で始まります。
	Pix []uint8
	// Strideは、垂直方向の隣接ピクセル間のPixストライド（バイト単位）です。
	Stride int
	// Rectは、画像の境界です。
	Rect Rectangle
}

func (p *RGBA) ColorModel() color.Model

func (p *RGBA) Bounds() Rectangle

func (p *RGBA) At(x, y int) color.Color

func (p *RGBA) RGBA64At(x, y int) color.RGBA64

func (p *RGBA) RGBAAt(x, y int) color.RGBA

// PixOffsetは、(x, y)のピクセルに対応するPixの最初の要素のインデックスを返します。
func (p *RGBA) PixOffset(x, y int) int

func (p *RGBA) Set(x, y int, c color.Color)

func (p *RGBA) SetRGBA64(x, y int, c color.RGBA64)

func (p *RGBA) SetRGBA(x, y int, c color.RGBA)

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *RGBA) SubImage(r Rectangle) Image

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (p *RGBA) Opaque() bool

// NewRGBAは、指定された境界を持つ新しい [RGBA] イメージを返します。
func NewRGBA(r Rectangle) *RGBA

// RGBA64は、Atメソッドが [color.RGBA64] 値を返すインメモリイメージです。
type RGBA64 struct {
	// Pixは、画像のピクセルをR, G, B, Aの順序で、ビッグエンディアン形式で保持します。ピクセルは
	// (x, y)はPix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*8]で始まります。
	Pix []uint8
	// Strideは、垂直方向の隣接ピクセル間のPixストライド（バイト単位）です。
	Stride int
	// Rectは、画像の境界です。
	Rect Rectangle
}

func (p *RGBA64) ColorModel() color.Model

func (p *RGBA64) Bounds() Rectangle

func (p *RGBA64) At(x, y int) color.Color

func (p *RGBA64) RGBA64At(x, y int) color.RGBA64

// PixOffsetは、(x, y)のピクセルに対応するPixの最初の要素のインデックスを返します。
func (p *RGBA64) PixOffset(x, y int) int

func (p *RGBA64) Set(x, y int, c color.Color)

func (p *RGBA64) SetRGBA64(x, y int, c color.RGBA64)

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *RGBA64) SubImage(r Rectangle) Image

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (p *RGBA64) Opaque() bool

// NewRGBA64は、指定された境界を持つ新しい [RGBA64] イメージを返します。
func NewRGBA64(r Rectangle) *RGBA64

// NRGBAは、Atメソッドが [color.NRGBA] 値を返すインメモリイメージです。
type NRGBA struct {
	// Pixは、画像のピクセルをR, G, B, Aの順序で保持します。ピクセルは
	// (x, y)はPix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4]で始まります。
	Pix []uint8
	// Strideは、垂直方向の隣接ピクセル間のPixストライド（バイト単位）です。
	Stride int
	// Rectは、画像の境界です。
	Rect Rectangle
}

func (p *NRGBA) ColorModel() color.Model

func (p *NRGBA) Bounds() Rectangle

func (p *NRGBA) At(x, y int) color.Color

func (p *NRGBA) RGBA64At(x, y int) color.RGBA64

func (p *NRGBA) NRGBAAt(x, y int) color.NRGBA

// PixOffsetは、(x, y)のピクセルに対応するPixの最初の要素のインデックスを返します。
func (p *NRGBA) PixOffset(x, y int) int

func (p *NRGBA) Set(x, y int, c color.Color)

func (p *NRGBA) SetRGBA64(x, y int, c color.RGBA64)

func (p *NRGBA) SetNRGBA(x, y int, c color.NRGBA)

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *NRGBA) SubImage(r Rectangle) Image

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (p *NRGBA) Opaque() bool

// NewNRGBAは、指定された境界を持つ新しい [NRGBA] イメージを返します。
func NewNRGBA(r Rectangle) *NRGBA

// NRGBA64は、Atメソッドが [color.NRGBA64] 値を返すインメモリイメージです。
type NRGBA64 struct {
	// Pixは、画像のピクセルをR, G, B, Aの順序で、ビッグエンディアン形式で保持します。ピクセルは
	// (x, y)はPix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*8]で始まります。
	Pix []uint8
	// Strideは、垂直方向の隣接ピクセル間のPixストライド（バイト単位）です。
	Stride int
	// Rectは、画像の境界です。
	Rect Rectangle
}

func (p *NRGBA64) ColorModel() color.Model

func (p *NRGBA64) Bounds() Rectangle

func (p *NRGBA64) At(x, y int) color.Color

func (p *NRGBA64) RGBA64At(x, y int) color.RGBA64

func (p *NRGBA64) NRGBA64At(x, y int) color.NRGBA64

// PixOffsetは、(x, y)のピクセルに対応するPixの最初の要素のインデックスを返します。
func (p *NRGBA64) PixOffset(x, y int) int

func (p *NRGBA64) Set(x, y int, c color.Color)

func (p *NRGBA64) SetRGBA64(x, y int, c color.RGBA64)

func (p *NRGBA64) SetNRGBA64(x, y int, c color.NRGBA64)

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *NRGBA64) SubImage(r Rectangle) Image

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (p *NRGBA64) Opaque() bool

// NewNRGBA64は、指定された境界を持つ新しい [NRGBA64] イメージを返します。
func NewNRGBA64(r Rectangle) *NRGBA64

// Alphaは、Atメソッドが [color.Alpha] 値を返すインメモリイメージです。
type Alpha struct {
	// Pixは、画像のピクセルをアルファ値として保持します。ピクセルは
	// (x, y)はPix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1]で始まります。
	Pix []uint8
	// Strideは、垂直方向の隣接ピクセル間のPixストライド（バイト単位）です。
	Stride int
	// Rectは、画像の境界です。
	Rect Rectangle
}

func (p *Alpha) ColorModel() color.Model

func (p *Alpha) Bounds() Rectangle

func (p *Alpha) At(x, y int) color.Color

func (p *Alpha) RGBA64At(x, y int) color.RGBA64

func (p *Alpha) AlphaAt(x, y int) color.Alpha

// PixOffsetは、(x, y)のピクセルに対応するPixの最初の要素のインデックスを返します。
func (p *Alpha) PixOffset(x, y int) int

func (p *Alpha) Set(x, y int, c color.Color)

func (p *Alpha) SetRGBA64(x, y int, c color.RGBA64)

func (p *Alpha) SetAlpha(x, y int, c color.Alpha)

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *Alpha) SubImage(r Rectangle) Image

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (p *Alpha) Opaque() bool

// NewAlphaは、指定された境界を持つ新しい [Alpha] イメージを返します。
func NewAlpha(r Rectangle) *Alpha

// Alpha16は、Atメソッドが [color.Alpha16] 値を返すインメモリイメージです。
type Alpha16 struct {
	// Pixは、画像のピクセルをアルファ値として、ビッグエンディアン形式で保持します。ピクセルは
	// (x, y)はPix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2]で始まります。
	Pix []uint8
	// Strideは、垂直方向の隣接ピクセル間のPixストライド（バイト単位）です。
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}

func (p *Alpha16) ColorModel() color.Model

func (p *Alpha16) Bounds() Rectangle

func (p *Alpha16) At(x, y int) color.Color

func (p *Alpha16) RGBA64At(x, y int) color.RGBA64

func (p *Alpha16) Alpha16At(x, y int) color.Alpha16

// PixOffsetは、(x, y)のピクセルに対応するPixの最初の要素のインデックスを返します。
func (p *Alpha16) PixOffset(x, y int) int

func (p *Alpha16) Set(x, y int, c color.Color)

func (p *Alpha16) SetRGBA64(x, y int, c color.RGBA64)

func (p *Alpha16) SetAlpha16(x, y int, c color.Alpha16)

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *Alpha16) SubImage(r Rectangle) Image

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (p *Alpha16) Opaque() bool

// NewAlpha16は、指定された境界を持つ新しい [Alpha16] イメージを返します。
func NewAlpha16(r Rectangle) *Alpha16

// Grayは、Atメソッドが [color.Gray] 値を返すインメモリイメージです。
type Gray struct {
	// Pixは、画像のピクセルをグレー値として保持します。ピクセルは
	// (x, y)はPix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1]で始まります。
	Pix []uint8
	// Strideは、垂直方向の隣接ピクセル間のPixストライド（バイト単位）です。
	Stride int
	// Rectは、画像の境界です。
	Rect Rectangle
}

func (p *Gray) ColorModel() color.Model

func (p *Gray) Bounds() Rectangle

func (p *Gray) At(x, y int) color.Color

func (p *Gray) RGBA64At(x, y int) color.RGBA64

func (p *Gray) GrayAt(x, y int) color.Gray

// PixOffsetは、(x, y)のピクセルに対応するPixの最初の要素のインデックスを返します。
func (p *Gray) PixOffset(x, y int) int

func (p *Gray) Set(x, y int, c color.Color)

func (p *Gray) SetRGBA64(x, y int, c color.RGBA64)

func (p *Gray) SetGray(x, y int, c color.Gray)

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *Gray) SubImage(r Rectangle) Image

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (p *Gray) Opaque() bool

// NewGrayは、指定された境界を持つ新しい [Gray] イメージを返します。
func NewGray(r Rectangle) *Gray

// Gray16は、Atメソッドが [color.Gray16] 値を返すインメモリイメージです。
type Gray16 struct {
	// Pixは、画像のピクセルをグレー値として、ビッグエンディアン形式で保持します。ピクセルは
	// (x, y)はPix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2]で始まります。
	Pix []uint8
	// Strideは、垂直方向の隣接ピクセル間のPixストライド（バイト単位）です。
	Stride int
	// Rectは、画像の境界です。
	Rect Rectangle
}

func (p *Gray16) ColorModel() color.Model

func (p *Gray16) Bounds() Rectangle

func (p *Gray16) At(x, y int) color.Color

func (p *Gray16) RGBA64At(x, y int) color.RGBA64

func (p *Gray16) Gray16At(x, y int) color.Gray16

// PixOffsetは、(x, y)のピクセルに対応するPixの最初の要素のインデックスを返します。
func (p *Gray16) PixOffset(x, y int) int

func (p *Gray16) Set(x, y int, c color.Color)

func (p *Gray16) SetRGBA64(x, y int, c color.RGBA64)

func (p *Gray16) SetGray16(x, y int, c color.Gray16)

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *Gray16) SubImage(r Rectangle) Image

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (p *Gray16) Opaque() bool

// NewGray16は、指定された境界を持つ新しい [Gray16] イメージを返します。
func NewGray16(r Rectangle) *Gray16

// CMYKは、Atメソッドが [color.CMYK] 値を返すインメモリイメージです。
type CMYK struct {
	// Pixは、画像のピクセルをC, M, Y, Kの順序で保持します。ピクセルは
	// (x, y)はPix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4]で始まります。
	Pix []uint8
	// Strideは、垂直方向の隣接ピクセル間のPixストライド（バイト単位）です。
	Stride int
	// Rectは、画像の境界です。
	Rect Rectangle
}

func (p *CMYK) ColorModel() color.Model

func (p *CMYK) Bounds() Rectangle

func (p *CMYK) At(x, y int) color.Color

func (p *CMYK) RGBA64At(x, y int) color.RGBA64

func (p *CMYK) CMYKAt(x, y int) color.CMYK

// PixOffsetは、(x, y)のピクセルに対応するPixの最初の要素のインデックスを返します。
func (p *CMYK) PixOffset(x, y int) int

func (p *CMYK) Set(x, y int, c color.Color)

func (p *CMYK) SetRGBA64(x, y int, c color.RGBA64)

func (p *CMYK) SetCMYK(x, y int, c color.CMYK)

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *CMYK) SubImage(r Rectangle) Image

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (p *CMYK) Opaque() bool

// NewCMYKは、指定された境界を持つ新しいCMYKイメージを返します。
func NewCMYK(r Rectangle) *CMYK

// Palettedは、指定されたパレットへのuint8インデックスのインメモリイメージです。
type Paletted struct {
	// Pixは、画像のピクセルをパレットインデックスとして保持します。ピクセルは
	// (x, y)はPix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1]で始まります。
	Pix []uint8
	// Strideは、垂直方向の隣接ピクセル間のPixストライド（バイト単位）です。
	Stride int
	// Rectは、画像の境界です。
	Rect Rectangle
	// Paletteは、画像のパレットです。
	Palette color.Palette
}

func (p *Paletted) ColorModel() color.Model

func (p *Paletted) Bounds() Rectangle

func (p *Paletted) At(x, y int) color.Color

func (p *Paletted) RGBA64At(x, y int) color.RGBA64

// PixOffsetは、(x, y)のピクセルに対応するPixの最初の要素のインデックスを返します。
func (p *Paletted) PixOffset(x, y int) int

func (p *Paletted) Set(x, y int, c color.Color)

func (p *Paletted) SetRGBA64(x, y int, c color.RGBA64)

func (p *Paletted) ColorIndexAt(x, y int) uint8

func (p *Paletted) SetColorIndex(x, y int, index uint8)

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *Paletted) SubImage(r Rectangle) Image

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (p *Paletted) Opaque() bool

// NewPalettedは、指定された幅、高さ、およびパレットを持つ新しい [Paletted] イメージを返します。
func NewPaletted(r Rectangle, p color.Palette) *Paletted
