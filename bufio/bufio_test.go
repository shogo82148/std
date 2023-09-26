// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio_test

import (
	. "bufio"
)

// Reads from a reader and rot13s the result.

// A StringReader delivers its data one string segment at a time via Read.
type StringReader struct {
	data []string
	step int
}

// TestReader wraps a []byte and returns reads of a specific length.
