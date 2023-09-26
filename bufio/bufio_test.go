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

// A writeCountingDiscard is like ioutil.Discard and counts the number of times
// Write is called on it.

// An onlyReader only implements io.Reader, no matter what other methods the underlying implementation may have.

// An onlyWriter only implements io.Writer, no matter what other methods the underlying implementation may have.

// A scriptedReader is an io.Reader that executes its steps sequentially.

// eofReader returns the number of bytes read and io.EOF for the read that consumes the last of the content.
