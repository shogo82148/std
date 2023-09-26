// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gif

import (
	"io"
)

// header, palette and trailer are parts of a valid 2x1 GIF image.

// lzw.NewReader wants an io.ByteReader, this ensures we're compatible.
var _ io.ByteReader = (*blockReader)(nil)

// testGIF is a simple GIF that we can modify to test different scenarios.
