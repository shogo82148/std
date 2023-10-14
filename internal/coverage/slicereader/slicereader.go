// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slicereader

type Reader struct {
	b        []byte
	readonly bool
	off      int64
}

func NewReader(b []byte, readonly bool) *Reader

func (r *Reader) Read(b []byte) (int, error)

func (r *Reader) Seek(offset int64, whence int) (ret int64, err error)

func (r *Reader) Offset() int64

func (r *Reader) ReadUint8() uint8

func (r *Reader) ReadUint32() uint32

func (r *Reader) ReadUint64() uint64

func (r *Reader) ReadULEB128() (value uint64)

func (r *Reader) ReadString(len int64) string
