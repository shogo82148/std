// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package png

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/io"
)

// Encode writes the Image m to w in PNG format. Any Image may be encoded, but
// images that are not image.NRGBA might be encoded lossily.
func Encode(w io.Writer, m image.Image) error
