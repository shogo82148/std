// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ioutilパッケージは、いくつかのI/Oユーティリティ関数を実装しています。
//
// Deprecated: Go 1.16以降、同じ機能はパッケージ [io] またはパッケージ [os] で提供されるようになり、
// これらの実装が新しいコードで優先されるべきです。
// 詳細については、特定の関数のドキュメントを参照してください。
package ioutil

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
)

// ReadAllは、rからエラーまたはEOFが発生するまで読み取り、読み取ったデータを返します。
// 成功した呼び出しはerr == nilを返します。err == EOFではありません。
// ReadAllは、EOFをエラーとして報告する必要はありません。
// なぜなら、ReadAllはsrcからEOFまで読み取るように定義されているためです。
//
// Deprecated: Go 1.16以降、この関数は単に [io.ReadAll] を呼び出すだけです。
//
//go:fix inline
>>>>>>> upstream/release-branch.go1.25
func ReadAll(r io.Reader) ([]byte, error)

// ReadFileは、filenameで指定されたファイルを読み取り、その内容を返します。
// 成功した呼び出しはerr == nilを返します。err == EOFではありません。
// ReadFileは、ファイル全体を読み取るため、ReadからのEOFをエラーとして報告する必要はありません。
//
// Deprecated: Go 1.16以降、この関数は単に [os.ReadFile] を呼び出すだけです。
//
//go:fix inline
func ReadFile(filename string) ([]byte, error)

// WriteFileは、filenameで指定されたファイルにデータを書き込みます。
// ファイルが存在しない場合、WriteFileは、パーミッションperm（umaskの前）で作成します。
// それ以外の場合、WriteFileはパーミッションを変更せずに書き込むために切り捨てます。
//
// Deprecated: Go 1.16以降、この関数は単に [os.WriteFile] を呼び出すだけです。
//
//go:fix inline
func WriteFile(filename string, data []byte, perm fs.FileMode) error

// ReadDirは、dirnameで指定されたディレクトリを読み取り、
// ファイル名でソートされたディレクトリの内容の [fs.FileInfo] リストを返します。
// ディレクトリの読み取り中にエラーが発生した場合、
// ReadDirはエラーとともにディレクトリエントリを返しません。
//
// Deprecated: Go 1.16以降、 [os.ReadDir] がより効率的で正確な選択肢となります。
// [os.ReadDir] は [fs.FileInfo] のリストではなく[fs.DirEntry]のリストを返し、
// ディレクトリの読み取り中にエラーが発生した場合でも部分的な結果を返します。
//
// [fs.FileInfo] のリストを引き続き取得する必要がある場合は、次のようにします。
//
//	entries, err := os.ReadDir(dirname)
//	if err != nil { ... }
//	infos := make([]fs.FileInfo, 0, len(entries))
//	for _, entry := range entries {
//		info, err := entry.Info()
//		if err != nil { ... }
//		infos = append(infos, info)
//	}
func ReadDir(dirname string) ([]fs.FileInfo, error)

// NopCloserは、提供されたReader rをラップするCloseメソッドのないReadCloserを返します。
//
// Deprecated: Go 1.16以降、この関数は単に [io.NopCloser] を呼び出すだけです。
//
//go:fix inline
func NopCloser(r io.Reader) io.ReadCloser

// Discardは、何もしないですべての書き込み呼び出しが成功するio.Writerです。
//
// Deprecated: Go 1.16以降、この値は単に [io.Discard] です。
var Discard io.Writer = io.Discard
