// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"github.com/shogo82148/std/image/color"
)

// YCbCrSubsampleRatio is the chroma subsample ratio used in a YCbCr image.
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

// YCbCr is an in-memory image of Y'CbCr colors. There is one Y sample per
// pixel, but each Cb and Cr sample can span one or more pixels.
// YStride is the Y slice index delta between vertically adjacent pixels.
// CStride is the Cb and Cr slice index delta between vertically adjacent pixels
// that map to separate chroma samples.
// It is not an absolute requirement, but YStride and len(Y) are typically
// multiples of 8, and:
//
//	For 4:4:4, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/1.
//	For 4:2:2, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/2.
//	For 4:2:0, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/4.
//	For 4:4:0, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/2.
//	For 4:1:1, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/4.
//	For 4:1:0, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/8.
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

func (p *YCbCr) YCbCrAt(x, y int) color.YCbCr

// YOffset returns the index of the first element of Y that corresponds to
// the pixel at (x, y).
func (p *YCbCr) YOffset(x, y int) int

// COffset returns the index of the first element of Cb or Cr that corresponds
// to the pixel at (x, y).
func (p *YCbCr) COffset(x, y int) int

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *YCbCr) SubImage(r Rectangle) Image

func (p *YCbCr) Opaque() bool

// NewYCbCr returns a new YCbCr image with the given bounds and subsample
// ratio.
func NewYCbCr(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *YCbCr

// NYCbCrA is an in-memory image of non-alpha-premultiplied Y'CbCr-with-alpha
// colors. A and AStride are analogous to the Y and YStride fields of the
// embedded YCbCr.
type NYCbCrA struct {
	YCbCr
	A       []uint8
	AStride int
}

func (p *NYCbCrA) ColorModel() color.Model

func (p *NYCbCrA) At(x, y int) color.Color

func (p *NYCbCrA) NYCbCrAAt(x, y int) color.NYCbCrA

// AOffset returns the index of the first element of A that corresponds to the
// pixel at (x, y).
func (p *NYCbCrA) AOffset(x, y int) int

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *NYCbCrA) SubImage(r Rectangle) Image

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *NYCbCrA) Opaque() bool

// NewNYCbCrA returns a new NYCbCrA image with the given bounds and subsample
// ratio.
func NewNYCbCrA(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *NYCbCrA
