// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// bzip2パッケージは、bzip2の解凍を実装します。
package bzip2

import "github.com/shogo82148/std/io"

// StructuralErrorは、bzip2データが構文的に無効であることが判明した場合に返されます。
type StructuralError string

func (s StructuralError) Error() string

// NewReaderは、rからbzip2データを解凍する[io.Reader]を返します。
// rが[io.ByteReader]も実装していない場合、
// 解凍器はrから必要以上のデータを読み取る可能性があります。
func NewReader(r io.Reader) io.Reader
