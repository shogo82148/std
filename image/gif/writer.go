// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gif

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/image/draw"
	"github.com/shogo82148/std/io"
)

// Optionsは、エンコーディングパラメータです。
type Options struct {
	// NumColorsは、画像で使用される色の最大数です。
	// 1から256までの範囲です。
	NumColors int

	// Quantizerは、NumColorsのサイズを持つパレットを生成するために使用されます。
	// Quantizerがnilの場合、代わりにpalette.Plan9が使用されます。
	Quantizer draw.Quantizer

	// Drawerは、ソース画像を所望のパレットに変換するために使用されます。
	// Drawerがnilの場合、代わりにdraw.FloydSteinbergが使用されます。
	Drawer draw.Drawer
}

// EncodeAllは、指定されたループカウントとフレーム間の遅延で、
// GIF形式のwにgの画像を書き込みます。
func EncodeAll(w io.Writer, g *GIF) error

// Encodeは、GIF形式で画像mをwに書き込みます。
func Encode(w io.Writer, m image.Image, o *Options) error
