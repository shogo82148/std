// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// ReadFileFSは、 [ReadFile] の最適化された実装を提供するファイルシステムによって実装されるインターフェースです。
type ReadFileFS interface {
	FS

	ReadFile(name string) ([]byte, error)
}

// ReadFileはファイルシステムfsysから指定された名前のファイルを読み取り、その内容を返します。
// 成功した呼び出しはnilエラーを返し、[io.EOF] は返しません。
// （ReadFileはファイル全体を読み取るため、最後のReadで予期されるEOFは
// 報告すべきエラーとして扱われません。）
//
// fsysが [ReadFileFS] を実装している場合、ReadFileはfsys.ReadFileを呼び出します。
// そうでない場合、ReadFileはfsys.Openを呼び出し、返された [File] の
// ReadとCloseを使用します。
func ReadFile(fsys FS, name string) ([]byte, error)
