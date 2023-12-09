// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"github.com/shogo82148/std/image/color"
)

// YCbCrSubsampleRatioは、YCbCr画像で使用されるクロマサブサンプル比率です。
type YCbCrSubsampleRatio int

const (
	YCbCrSubsampleRatio444 YCbCrSubsampleRatio = iota
	YCbCrSubsampleRatio422
	YCbCrSubsampleRatio420
	YCbCrSubsampleRatio440
	YCbCrSubsampleRatio411
	YCbCrSubsampleRatio410
)

func (s YCbCrSubsampleRatio) String() string

// YCbCrは、Y'CbCr色のインメモリイメージです。ピクセルごとに1つのYサンプルがありますが、
// 各CbおよびCrサンプルは1つ以上のピクセルに跨ることができます。
// YStrideは、垂直方向の隣接ピクセル間のYスライスインデックスデルタです。
// CStrideは、別々のクロマサンプルにマップされる垂直方向の隣接ピクセル間のCbおよびCrスライスインデックスデルタです。
// 絶対的な要件ではありませんが、通常、YStrideとlen(Y)は8の倍数です、そして：
//
//	4:4:4の場合、CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/1。
//	4:2:2の場合、CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/2。
//	4:2:0の場合、CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/4。
//	4:4:0の場合、CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/2。
//	4:1:1の場合、CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/4。
//	4:1:0の場合、CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/8。
type YCbCr struct {
	Y, Cb, Cr      []uint8
	YStride        int
	CStride        int
	SubsampleRatio YCbCrSubsampleRatio
	Rect           Rectangle
}

func (p *YCbCr) ColorModel() color.Model

func (p *YCbCr) Bounds() Rectangle

func (p *YCbCr) At(x, y int) color.Color

func (p *YCbCr) RGBA64At(x, y int) color.RGBA64

func (p *YCbCr) YCbCrAt(x, y int) color.YCbCr

// YOffsetは、(x, y)のピクセルに対応するYの最初の要素のインデックスを返します。
func (p *YCbCr) YOffset(x, y int) int

// COffsetは、(x, y)のピクセルに対応するCbまたはCrの最初の要素のインデックスを返します。
func (p *YCbCr) COffset(x, y int) int

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *YCbCr) SubImage(r Rectangle) Image

func (p *YCbCr) Opaque() bool

// NewYCbCrは、指定された境界とサブサンプル比率を持つ新しいYCbCrイメージを返します。
func NewYCbCr(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *YCbCr

// NYCbCrAは、非アルファ乗算のY'CbCr-with-alpha色のインメモリイメージです。
// AとAStrideは、埋め込まれたYCbCrのYとYStrideフィールドに対応します。
type NYCbCrA struct {
	YCbCr
	A       []uint8
	AStride int
}

func (p *NYCbCrA) ColorModel() color.Model

func (p *NYCbCrA) At(x, y int) color.Color

func (p *NYCbCrA) RGBA64At(x, y int) color.RGBA64

func (p *NYCbCrA) NYCbCrAAt(x, y int) color.NYCbCrA

// AOffsetは、(x, y)のピクセルに対応するAの最初の要素のインデックスを返します。
func (p *NYCbCrA) AOffset(x, y int) int

// SubImageは、rを通じて見える画像pの一部を表す画像を返します。
// 返される値は、元の画像とピクセルを共有します。
func (p *NYCbCrA) SubImage(r Rectangle) Image

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (p *NYCbCrA) Opaque() bool

// NewNYCbCrAは、指定された境界とサブサンプル比率を持つ新しい [NYCbCrA] イメージを返します。
func NewNYCbCrA(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *NYCbCrA
