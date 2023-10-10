// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// ReadFileFSは、ReadFileの最適化された実装を提供するファイルシステムによって実装されるインターフェースです。
type ReadFileFS interface {
	FS

	// ReadFile reads the named file and returns its contents.
	// A successful call returns a nil error, not io.EOF.
	// (Because ReadFile reads the whole file, the expected EOF
	// from the final Read is not treated as an error to be reported.)
	//
	// The caller is permitted to modify the returned byte slice.
	// This method should return a copy of the underlying data.
	ReadFile(name string) ([]byte, error)
}

// ReadFileはファイルシステムfsから指定された名前のファイルを読み込み、その内容を返します。
// 成功した呼び出しはnilのエラーを返しますが、io.EOFではありません。
// (ReadFileはファイル全体を読み込むため、最後のReadでの予想されるEOFはエラーとして報告されません。)
//
// もしfsがReadFileFSを実装している場合、ReadFileはfs.ReadFileを呼び出します。
// そうでなければ、ReadFileはfs.Openを呼び出し、返されたファイルに対してReadとCloseを使用します。
func ReadFile(fsys FS, name string) ([]byte, error)
