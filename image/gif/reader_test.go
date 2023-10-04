// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gif

import (
	"io"
)

// lzw.NewReader wants a io.ByteReader, this ensures we're compatible.
var _ io.ByteReader = (*blockReader)(nil)
