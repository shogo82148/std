// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

// Readはファイルセットをデシリアライズするためにデコードを呼び出します。sはnilであってはいけません。
func (s *FileSet) Read(decode func(any) error) error

// Writeはエンコードを呼び出してファイルセットsをシリアライズする。
func (s *FileSet) Write(encode func(any) error) error
