// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package image implements a basic 2-D image library.
//
// The fundamental interface is called Image. An Image contains colors, which
// are described in the image/color package.
//
// Values of the Image interface are created either by calling functions such
// as NewRGBA and NewPaletted, or by calling Decode on an io.Reader containing
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
// http://golang.org/doc/articles/image_package.html
package image

import (
	"github.com/shogo82148/std/image/color"
)

// Config holds an image's color model and dimensions.
type Config struct {
	ColorModel    color.Model
	Width, Height int
}

// Image is a finite rectangular grid of color.Color values taken from a color
// model.
type Image interface {
	ColorModel() color.Model

	Bounds() Rectangle

	At(x, y int) color.Color
}

// PalettedImage is an image whose colors may come from a limited palette.
// If m is a PalettedImage and m.ColorModel() returns a PalettedColorModel p,
// then m.At(x, y) should be equivalent to p[m.ColorIndexAt(x, y)]. If m's
// color model is not a PalettedColorModel, then ColorIndexAt's behavior is
// undefined.
type PalettedImage interface {
	ColorIndexAt(x, y int) uint8
	Image
}

// RGBA is an in-memory image whose At method returns color.RGBA values.
type RGBA struct {
	Pix []uint8

	Stride int

	Rect Rectangle
}

func (p *RGBA) ColorModel() color.Model

func (p *RGBA) Bounds() Rectangle

func (p *RGBA) At(x, y int) color.Color

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA) PixOffset(x, y int) int

func (p *RGBA) Set(x, y int, c color.Color)

func (p *RGBA) SetRGBA(x, y int, c color.RGBA)

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA) SubImage(r Rectangle) Image

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA) Opaque() bool

// NewRGBA returns a new RGBA with the given bounds.
func NewRGBA(r Rectangle) *RGBA

// RGBA64 is an in-memory image whose At method returns color.RGBA64 values.
type RGBA64 struct {
	Pix []uint8

	Stride int

	Rect Rectangle
}

func (p *RGBA64) ColorModel() color.Model

func (p *RGBA64) Bounds() Rectangle

func (p *RGBA64) At(x, y int) color.Color

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA64) PixOffset(x, y int) int

func (p *RGBA64) Set(x, y int, c color.Color)

func (p *RGBA64) SetRGBA64(x, y int, c color.RGBA64)

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA64) SubImage(r Rectangle) Image

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA64) Opaque() bool

// NewRGBA64 returns a new RGBA64 with the given bounds.
func NewRGBA64(r Rectangle) *RGBA64

// NRGBA is an in-memory image whose At method returns color.NRGBA values.
type NRGBA struct {
	Pix []uint8

	Stride int

	Rect Rectangle
}

func (p *NRGBA) ColorModel() color.Model

func (p *NRGBA) Bounds() Rectangle

func (p *NRGBA) At(x, y int) color.Color

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *NRGBA) PixOffset(x, y int) int

func (p *NRGBA) Set(x, y int, c color.Color)

func (p *NRGBA) SetNRGBA(x, y int, c color.NRGBA)

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *NRGBA) SubImage(r Rectangle) Image

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *NRGBA) Opaque() bool

// NewNRGBA returns a new NRGBA with the given bounds.
func NewNRGBA(r Rectangle) *NRGBA

// NRGBA64 is an in-memory image whose At method returns color.NRGBA64 values.
type NRGBA64 struct {
	Pix []uint8

	Stride int

	Rect Rectangle
}

func (p *NRGBA64) ColorModel() color.Model

func (p *NRGBA64) Bounds() Rectangle

func (p *NRGBA64) At(x, y int) color.Color

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *NRGBA64) PixOffset(x, y int) int

func (p *NRGBA64) Set(x, y int, c color.Color)

func (p *NRGBA64) SetNRGBA64(x, y int, c color.NRGBA64)

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *NRGBA64) SubImage(r Rectangle) Image

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *NRGBA64) Opaque() bool

// NewNRGBA64 returns a new NRGBA64 with the given bounds.
func NewNRGBA64(r Rectangle) *NRGBA64

// Alpha is an in-memory image whose At method returns color.Alpha values.
type Alpha struct {
	Pix []uint8

	Stride int

	Rect Rectangle
}

func (p *Alpha) ColorModel() color.Model

func (p *Alpha) Bounds() Rectangle

func (p *Alpha) At(x, y int) color.Color

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Alpha) PixOffset(x, y int) int

func (p *Alpha) Set(x, y int, c color.Color)

func (p *Alpha) SetAlpha(x, y int, c color.Alpha)

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Alpha) SubImage(r Rectangle) Image

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Alpha) Opaque() bool

// NewAlpha returns a new Alpha with the given bounds.
func NewAlpha(r Rectangle) *Alpha

// Alpha16 is an in-memory image whose At method returns color.Alpha64 values.
type Alpha16 struct {
	Pix []uint8

	Stride int

	Rect Rectangle
}

func (p *Alpha16) ColorModel() color.Model

func (p *Alpha16) Bounds() Rectangle

func (p *Alpha16) At(x, y int) color.Color

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Alpha16) PixOffset(x, y int) int

func (p *Alpha16) Set(x, y int, c color.Color)

func (p *Alpha16) SetAlpha16(x, y int, c color.Alpha16)

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Alpha16) SubImage(r Rectangle) Image

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Alpha16) Opaque() bool

// NewAlpha16 returns a new Alpha16 with the given bounds.
func NewAlpha16(r Rectangle) *Alpha16

// Gray is an in-memory image whose At method returns color.Gray values.
type Gray struct {
	Pix []uint8

	Stride int

	Rect Rectangle
}

func (p *Gray) ColorModel() color.Model

func (p *Gray) Bounds() Rectangle

func (p *Gray) At(x, y int) color.Color

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Gray) PixOffset(x, y int) int

func (p *Gray) Set(x, y int, c color.Color)

func (p *Gray) SetGray(x, y int, c color.Gray)

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Gray) SubImage(r Rectangle) Image

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Gray) Opaque() bool

// NewGray returns a new Gray with the given bounds.
func NewGray(r Rectangle) *Gray

// Gray16 is an in-memory image whose At method returns color.Gray16 values.
type Gray16 struct {
	Pix []uint8

	Stride int

	Rect Rectangle
}

func (p *Gray16) ColorModel() color.Model

func (p *Gray16) Bounds() Rectangle

func (p *Gray16) At(x, y int) color.Color

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Gray16) PixOffset(x, y int) int

func (p *Gray16) Set(x, y int, c color.Color)

func (p *Gray16) SetGray16(x, y int, c color.Gray16)

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Gray16) SubImage(r Rectangle) Image

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Gray16) Opaque() bool

// NewGray16 returns a new Gray16 with the given bounds.
func NewGray16(r Rectangle) *Gray16

// Paletted is an in-memory image of uint8 indices into a given palette.
type Paletted struct {
	Pix []uint8

	Stride int

	Rect Rectangle

	Palette color.Palette
}

func (p *Paletted) ColorModel() color.Model

func (p *Paletted) Bounds() Rectangle

func (p *Paletted) At(x, y int) color.Color

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Paletted) PixOffset(x, y int) int

func (p *Paletted) Set(x, y int, c color.Color)

func (p *Paletted) ColorIndexAt(x, y int) uint8

func (p *Paletted) SetColorIndex(x, y int, index uint8)

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Paletted) SubImage(r Rectangle) Image

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Paletted) Opaque() bool

// NewPaletted returns a new Paletted with the given width, height and palette.
func NewPaletted(r Rectangle, p color.Palette) *Paletted
