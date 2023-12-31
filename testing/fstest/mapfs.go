// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fstest

import (
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/time"
)

// MapFSは、テストで使用するためのシンプルなメモリ内ファイルシステムであり、
// パス名（Openへの引数）からそれらが表すファイルやディレクトリの情報へのマップとして表されます。
//
// マップには、マップに含まれるファイルの親ディレクトリを含める必要はありません。
// 必要に応じてそれらは合成されます。
// しかし、[MapFile.Mode] の [fs.ModeDir] ビットを設定することで、ディレクトリをまだ含めることができます。
// これは、ディレクトリの [fs.FileInfo] に対する詳細な制御が必要であるか、
// 空のディレクトリを作成するために必要かもしれません。
//
// ファイルシステムの操作は、マップから直接読み取るため、
// 必要に応じてマップを編集することでファイルシステムを変更できます。
// その意味するところは、ファイルシステムの操作はマップの変更と同時に実行してはならず、
// それはレース条件を引き起こす可能性があるということです。
// 別の意味するところは、ディレクトリのオープンや読み取りには、
// マップ全体を反復処理する必要があるため、MapFSは通常、
// 数百エントリ以上またはディレクトリの読み取りを使用しないで使用する必要があります。
type MapFS map[string]*MapFile

// MapFileは、[MapFS] 内の単一のファイルを説明します。
type MapFile struct {
	Data    []byte
	Mode    fs.FileMode
	ModTime time.Time
	Sys     any
}

var _ fs.FS = MapFS(nil)
var _ fs.File = (*openMapFile)(nil)

// Openは、指定された名前のファイルを開きます。
func (fsys MapFS) Open(name string) (fs.File, error)

func (fsys MapFS) ReadFile(name string) ([]byte, error)

func (fsys MapFS) Stat(name string) (fs.FileInfo, error)

func (fsys MapFS) ReadDir(name string) ([]fs.DirEntry, error)

func (fsys MapFS) Glob(pattern string) ([]string, error)

func (fsys MapFS) Sub(dir string) (fs.FS, error)
