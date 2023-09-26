// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io_test

import (
	"bytes"
	. "io"
)

// An version of bytes.Buffer without ReadFrom and WriteTo
type Buffer struct {
	bytes.Buffer
	ReaderFrom
	WriterTo
}

// A version of bytes.Buffer that returns n > 0, EOF on Read
// when the input is exhausted.
