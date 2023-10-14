// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package buffer provides a pool-allocated byte buffer.
package buffer

// buffer adapted from go/src/fmt/print.go
type Buffer []byte

func New() *Buffer

func (b *Buffer) Free()

func (b *Buffer) Reset()

func (b *Buffer) Write(p []byte) (int, error)

func (b *Buffer) WriteString(s string) (int, error)

func (b *Buffer) WriteByte(c byte) error

func (b *Buffer) WritePosInt(i int)

// WritePosIntWidth writes non-negative integer i to the buffer, padded on the left
// by zeroes to the given width. Use a width of 0 to omit padding.
func (b *Buffer) WritePosIntWidth(i, width int)

func (b *Buffer) String() string
