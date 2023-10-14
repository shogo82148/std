// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cov

import (
	"github.com/shogo82148/std/cmd/internal/bio"
	"github.com/shogo82148/std/os"
)

type MReader struct {
	f        *os.File
	rdr      *bio.Reader
	fileView []byte
	off      int64
}

func NewMreader(f *os.File) (*MReader, error)

func (r *MReader) Read(p []byte) (int, error)

func (r *MReader) ReadByte() (byte, error)

func (r *MReader) Seek(offset int64, whence int) (int64, error)
