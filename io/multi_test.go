// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io_test

import (
	. "io"
)

// readerFunc is an io.Reader implemented by the underlying func.

// byteAndEOFReader is a Reader which reads one byte (the underlying
// byte) and io.EOF at once in its Read call.
