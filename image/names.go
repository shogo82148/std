// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"github.com/shogo82148/std/image/color"
)

var (
	// Blackは、不透明な黒の一様な画像です。
	Black = NewUniform(color.Black)
	// Whiteは、不透明な白の一様な画像です。
	White = NewUniform(color.White)
	// Transparentは、完全に透明な一様な画像です。
	Transparent = NewUniform(color.Transparent)
	// Opaqueは、完全に不透明な一様な画像です。
	Opaque = NewUniform(color.Opaque)
)

<<<<<<< HEAD
// Uniform is an infinite-sized [Image] of uniform color.
// It implements the [color.Color], [color.Model], and [Image] interfaces.
=======
// Uniformは、一様な色の無限大のイメージです。
// これはcolor.Color、color.Model、およびImageインターフェースを実装します。
>>>>>>> release-branch.go1.21
type Uniform struct {
	C color.Color
}

func (c *Uniform) RGBA() (r, g, b, a uint32)

func (c *Uniform) ColorModel() color.Model

func (c *Uniform) Convert(color.Color) color.Color

func (c *Uniform) Bounds() Rectangle

func (c *Uniform) At(x, y int) color.Color

func (c *Uniform) RGBA64At(x, y int) color.RGBA64

// Opaqueは、画像全体をスキャンし、それが完全に不透明であるかどうかを報告します。
func (c *Uniform) Opaque() bool

<<<<<<< HEAD
// NewUniform returns a new [Uniform] image of the given color.
=======
// NewUniformは、指定された色の新しいUniformイメージを返します。
>>>>>>> release-branch.go1.21
func NewUniform(c color.Color) *Uniform
