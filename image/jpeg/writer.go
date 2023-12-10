// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jpeg

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/io"
)

// DefaultQualityは、デフォルトの品質エンコーディングパラメータです。
const DefaultQuality = 75

// Optionsは、エンコーディングパラメータです。
// Qualityは1から100までの範囲で、高いほど良いです。
type Options struct {
	Quality int
}

// Encodeは、与えられたオプションでJPEG 4:2:0ベースラインフォーマットでImage mをwに書き込みます。
// nilの *[Options] が渡された場合、デフォルトのパラメータが使用されます。
func Encode(w io.Writer, m image.Image, o *Options) error
