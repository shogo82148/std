// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ioutil

import (
	"github.com/shogo82148/std/os"
)

// TempFileは、ディレクトリdirに新しい一時ファイルを作成し、
// ファイルを読み書きするために開き、結果の *[os.File] を返します。
// ファイル名は、patternを取り、ランダムな文字列を末尾に追加して生成されます。
// patternに"*"が含まれている場合、ランダムな文字列が最後の"*"に置き換えられます。
// dirが空の文字列の場合、TempFileは一時ファイルのデフォルトディレクトリを使用します（[os.TempDir] を参照）。
// 同時にTempFileを呼び出す複数のプログラムは、同じファイルを選択しません。
// 呼び出し元は、f.Name()を使用してファイルのパス名を見つけることができます。
// ファイルが不要になったら、呼び出し元の責任でファイルを削除する必要があります。
//
// Deprecated: Go 1.17以降、この関数は単に [os.CreateTemp] を呼び出すだけです。
//
//go:fix inline
func TempFile(dir, pattern string) (f *os.File, err error)

// TempDirは、ディレクトリdirに新しい一時ディレクトリを作成し、
// ディレクトリ名を生成するためにpatternを取り、ランダムな文字列を末尾に追加します。
// patternに"*"が含まれている場合、ランダムな文字列が最後の"*"に置き換えられます。
// TempDirは、新しいディレクトリの名前を返します。
// dirが空の文字列の場合、TempDirは一時ファイルのデフォルトディレクトリを使用します（[os.TempDir] を参照）。
// 同時にTempDirを呼び出す複数のプログラムは、同じディレクトリを選択しません。
// 呼び出し元は、ディレクトリが不要になったら削除する責任があります。
//
// Deprecated: Go 1.17以降、この関数は単に [os.MkdirTemp] を呼び出すだけです。
//
//go:fix inline
func TempDir(dir, pattern string) (name string, err error)
