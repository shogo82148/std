// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io_test

import (
	. "io"
)

// writerFunc is an Writer implemented by the underlying func.

// readerFunc is an Reader implemented by the underlying func.

// byteAndEOFReader is a Reader which reads one byte (the underlying
// byte) and EOF at once in its Read call.
