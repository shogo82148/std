// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

// RGBToYCbCrは、RGBトリプルをY'CbCrトリプルに変換します。
func RGBToYCbCr(r, g, b uint8) (uint8, uint8, uint8)

// YCbCrToRGBは、Y'CbCrトリプルをRGBトリプルに変換します。
func YCbCrToRGB(y, cb, cr uint8) (uint8, uint8, uint8)

// YCbCrは、完全に不透明な24ビットのY'CbCr色を表し、
// 1つの輝度成分と2つの色差成分のそれぞれに8ビットずつを持っています。
//
// JPEG、VP8、MPEGファミリー、その他のコーデックはこのカラーモデルを使用します。
// これらのコーデックはしばしばYUVとY'CbCrを同義語として使用しますが、
// 厳密には、YUVという用語はアナログビデオ信号にのみ適用され、
// Y'（ルーマ）はガンマ補正を適用した後のY（輝度）です。
//
// RGBとY'CbCr間の変換は損失を伴い、両者間の変換には少し異なる複数の公式があります。
// このパッケージは、https://www.w3.org/Graphics/JPEG/jfif3.pdf のJFIF仕様に従います。
type YCbCr struct {
	Y, Cb, Cr uint8
}

func (c YCbCr) RGBA() (uint32, uint32, uint32, uint32)

// YCbCrModelはY'CbCr色の [Model] です。
var YCbCrModel Model = ModelFunc(yCbCrModel)

// NYCbCrAは、アルファ非乗算のY'CbCr-with-alpha色を表し、
// 1つの輝度成分、2つの色差成分、1つのアルファ成分それぞれに8ビットずつを持っています。
type NYCbCrA struct {
	YCbCr
	A uint8
}

func (c NYCbCrA) RGBA() (uint32, uint32, uint32, uint32)

// NYCbCrAModelは、アルファ非乗算のY'CbCr-with-alpha色の [Model] です。
var NYCbCrAModel Model = ModelFunc(nYCbCrAModel)

// RGBToCMYKは、RGBトリプルをCMYK四重奏に変換します。
func RGBToCMYK(r, g, b uint8) (uint8, uint8, uint8, uint8)

// CMYKToRGBは、[CMYK] 四重奏をRGBトリプルに変換します。
func CMYKToRGB(c, m, y, k uint8) (uint8, uint8, uint8)

// CMYKは、シアン、マゼンタ、イエロー、ブラックの各色に8ビットずつ持つ、完全に不透明なCMYK色を表します。
//
// それは特定のカラープロファイルに関連付けられていません。
type CMYK struct {
	C, M, Y, K uint8
}

func (c CMYK) RGBA() (uint32, uint32, uint32, uint32)

// CMYKModelはCMYK色の [Model] です。
var CMYKModel Model = ModelFunc(cmykModel)
