// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Buffered reading and decoding of DWARF data streams.

package dwarf

// Data buffer being decoded.

// Data format, other than byte order.  This affects the handling of
// certain field formats.

// Some parts of DWARF have no data format, e.g., abbrevs.

type DecodeError struct {
	Name   string
	Offset Offset
	Err    string
}

func (e DecodeError) Error() string
