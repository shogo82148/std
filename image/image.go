// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// imageパッケージは、基本的な2次元画像ライブラリを実装します。
//
// 基本的なインターフェースは[Image]と呼ばれます。[Image]は色を含み、
// 色はimage/colorパッケージで記述されています。
//
// [Image]インターフェースの値は、[NewRGBA]や[NewPaletted]などの関数を呼び出すか、
// GIF、JPEG、PNGなどの形式の画像データを含む[io.Reader]に対して[Decode]を
// 呼び出すことによって作成されます。特定の画像形式をデコードするには、
// 事前にデコーダ関数を登録する必要があります。
// 登録は通常、その形式のパッケージを初期化する際の副作用として自動的に行われるため、
// PNG画像をデコードするには、プログラムのmainパッケージに
//
//	import _ "image/png"
//
// と記述するだけで十分です。_は、初期化の副作用のためだけにパッケージを
// インポートすることを意味します。
//
// 詳細については「The Go image package」を参照してください:
// https://golang.org/doc/articles/image_package.html
//
// # セキュリティに関する考慮事項
//
// imageパッケージは任意の大きさの画像をパースするために使用できますが、
// それらを格納するのに十分なメモリを持たないマシンではリソースの枯渇を引き起こす可能性があります。
// 任意の画像を扱う場合は、[Decode]の前に[DecodeConfig]を呼び出して、
// 返されたヘッダーで定義されている画像を利用可能なリソースで安全にデコードできるかどうかを
// プログラムが判断できるようにするべきです。[DecodeConfig]が返すヘッダーで
// 定義されているとおりの非常に大きな画像を生成する[Decode]の呼び出しは、
// 画像自体が不正な形式であるかどうかにかかわらず、セキュリティ上の問題とはみなされません。
// [Decode]が返す画像と一致しないヘッダーを返す[DecodeConfig]の呼び出しは
// セキュリティ上の問題とみなされる可能性があり、
// [Go Security Policy](https://go.dev/security/policy)に従って報告するべきです。
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
	// Rectは画像の境界です。
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
