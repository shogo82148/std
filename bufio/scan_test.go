// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio_test

import (
	. "bufio"
)

// slowReader is a reader that returns only a few bytes at a time, to test the incremental
// reads in Scanner.Scan.

// Test for issue 5268.

// Test that Scan finishes if we have endless empty reads.

// negativeEOFReader returns an invalid -1 at the end, as though it
// were wrapping the read system call.

// largeReader returns an invalid count that is larger than the number
// of bytes requested.
