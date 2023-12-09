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

// Uniformは、一様な色の無限大の [Image] です。
// これは [color.Color]、[color.Model]、および [Image] インターフェースを実装します。
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

// NewUniformは、指定された色の新しい [Uniform] イメージを返します。
func NewUniform(c color.Color) *Uniform
