// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fstest

import (
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/time"
)

// MapFSはテストで使用するためのシンプルなインメモリーファイルシステムです。
// パス名（Openの引数）から、それが表すファイル、ディレクトリー、またはシンボリックリンクの情報へのマップとして表現されます。
//
// マップには、マップに含まれるファイルの親ディレクトリーを含める必要はありません。
// 必要に応じてそれらは合成されます。
// しかし、[MapFile.Mode] の [fs.ModeDir] ビットを設定することで、ディレクトリーをまだ含めることができます。
// これは、ディレクトリーの [fs.FileInfo] に対する詳細な制御が必要であるか、
// 空のディレクトリーを作成するために必要かもしれません。
//
// ファイルシステムの操作は、マップから直接読み取るため、
// 必要に応じてマップを編集することでファイルシステムを変更できます。
// その意味するところは、ファイルシステムの操作はマップの変更と同時に実行してはならず、
// それはレース条件を引き起こす可能性があるということです。
// 別の意味するところは、ディレクトリーのオープンや読み取りには、
// マップ全体を反復処理する必要があるため、MapFSは通常、
// 数百エントリー以上またはディレクトリーの読み取りを使用しないで使用する必要があります。
type MapFS map[string]*MapFile

// MapFileは、[MapFS] 内の単一のファイルを説明します。
type MapFile struct {
	Data    []byte
	Mode    fs.FileMode
	ModTime time.Time
	Sys     any
}

var _ fs.FS = MapFS(nil)
var _ fs.ReadLinkFS = MapFS(nil)
var _ fs.File = (*openMapFile)(nil)

// Openは、シンボリックリンクを辿った後に指定されたファイルを開きます。
func (fsys MapFS) Open(name string) (fs.File, error)

// ReadLinkは、指定されたシンボリックリンクのリンク先を返します。
func (fsys MapFS) ReadLink(name string) (string, error)

// Lstatは、指定されたファイルを説明するFileInfoを返します。
// ファイルがシンボリックリンクの場合、返されるFileInfoはそのシンボリックリンク自体を説明します。
// Lstatはリンクを辿ろうとはしません。
func (fsys MapFS) Lstat(name string) (fs.FileInfo, error)

func (fsys MapFS) ReadFile(name string) ([]byte, error)

func (fsys MapFS) Stat(name string) (fs.FileInfo, error)

func (fsys MapFS) ReadDir(name string) ([]fs.DirEntry, error)

func (fsys MapFS) Glob(pattern string) ([]string, error)

func (fsys MapFS) Sub(dir string) (fs.FS, error)
