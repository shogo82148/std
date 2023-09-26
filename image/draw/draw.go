// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package draw provides image composition functions.
//
// See "The Go image/draw package" for an introduction to this package:
// http://golang.org/doc/articles/image_draw.html
package draw

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/image/color"
)

// m is the maximum color value returned by image.Color.RGBA.

// Op is a Porter-Duff compositing operator.
type Op int

const (
	// Over specifies ``(src in mask) over dst''.
	Over Op = iota
	// Src specifies ``src in mask''.
	Src
)

// A draw.Image is an image.Image with a Set method to change a single pixel.
type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}

// Draw calls DrawMask with a nil mask.
func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result of a Porter-Duff composition. A nil mask is treated as opaque.
func DrawMask(dst Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op Op)
