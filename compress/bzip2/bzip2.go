// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bzip2 implements bzip2 decompression.
package bzip2

import "github.com/shogo82148/std/io"

// StructuralErrorは、bzip2データが構文的に無効であることが判明した場合に返されます。
type StructuralError string

func (s StructuralError) Error() string

<<<<<<< HEAD
// NewReader returns an io.Reader which decompresses bzip2 data from r.
// If r does not also implement [io.ByteReader],
// the decompressor may read more data than necessary from r.
=======
// NewReaderは、rからbzip2データを解凍するio.Readerを返します。
// rがio.ByteReaderインターフェースを実装していない場合、
// 解凍器はrから必要以上のデータを読み取る可能性があります。
>>>>>>> release-branch.go1.21
func NewReader(r io.Reader) io.Reader
