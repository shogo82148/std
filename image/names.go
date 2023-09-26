// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"github.com/shogo82148/std/image/color"
)

var (
	// Black is an opaque black uniform image.
	Black = NewUniform(color.Black)
	// White is an opaque white uniform image.
	White = NewUniform(color.White)
	// Transparent is a fully transparent uniform image.
	Transparent = NewUniform(color.Transparent)
	// Opaque is a fully opaque uniform image.
	Opaque = NewUniform(color.Opaque)
)

// Uniform is an infinite-sized Image of uniform color.
// It implements the color.Color, color.Model, and Image interfaces.
type Uniform struct {
	C color.Color
}

func (c *Uniform) RGBA() (r, g, b, a uint32)

func (c *Uniform) ColorModel() color.Model

func (c *Uniform) Convert(color.Color) color.Color

func (c *Uniform) Bounds() Rectangle

func (c *Uniform) At(x, y int) color.Color

// Opaque scans the entire image and returns whether or not it is fully opaque.
func (c *Uniform) Opaque() bool

func NewUniform(c color.Color) *Uniform
