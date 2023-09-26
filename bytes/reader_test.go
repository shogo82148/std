// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes_test

import (
	. "bytes"
	"io"
)

var UnreadRuneErrorTests = []struct {
	name string
	f    func(*Reader)
}{
	{"Read", func(r *Reader) { r.Read([]byte{0}) }},
	{"ReadByte", func(r *Reader) { r.ReadByte() }},
	{"UnreadRune", func(r *Reader) { r.UnreadRune() }},
	{"Seek", func(r *Reader) { r.Seek(0, io.SeekCurrent) }},
	{"WriteTo", func(r *Reader) { r.WriteTo(&Buffer{}) }},
}
