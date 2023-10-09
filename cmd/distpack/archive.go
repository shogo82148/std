// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/time"
)

// Archiveは、書き込むためのアーカイブを表します：ファイルの集合体です。
// ディレクトリはファイルに含まれており、明示的にリストされていません。
type Archive struct {
	Files []File
}

// Fileはアーカイブに書き込む単一のファイルを表します。
type File struct {
	Name string
	Time time.Time
	Mode fs.FileMode
	Size int64
	Src  string
}

// Infoはファイルに関するFileInfoを返します。tar.FileInfoHeaderやzip.FileInfoHeaderと組み合わせて使用します。
func (f *File) Info() fs.FileInfo

// NewArchiveはディレクトリdirに含まれるすべてのファイルを含む新しいアーカイブを返します。
// アーカイブはAddやFilterなどのメソッドを使って後から修正することができます。
func NewArchive(dir string) (*Archive, error)

// Addは指定された名前と情報を持つファイルをアーカイブに追加します。
// ファイルの内容はオペレーティングシステムのファイルsrcから取得します。
// 1回以上Addを呼び出した後、アーカイブのファイルを再ソートするためにSortを呼び出す必要があります。
func (a *Archive) Add(name, src string, info fs.FileInfo)

// Sort sorts the files in the archive.
// It is only necessary to call Sort after calling Add or RenameGoMod.
// NewArchive returns a sorted archive, and the other methods
// preserve the sorting of the archive.
func (a *Archive) Sort()

// CloneはArchiveのコピーを返します。
// コピーに対して行われるAddやFilterなどのメソッド呼び出しは、元のデータに影響を与えません。
// また、元のデータに対する呼び出しも、コピーには影響しません。
func (a *Archive) Clone() *Archive

// AddPrefixはアーカイブ内のすべてのファイル名に接頭辞を追加します。
func (a *Archive) AddPrefix(prefix string)

// Filterはkeep(name)がfalseを返すアーカイブからファイルを除外します。
func (a *Archive) Filter(keep func(name string) bool)

// SetModeはアーカイブ内のすべてのファイルのモードを変更します
// モードは(name, m)になります。ここで、mはファイルの現在のモードです。
func (a *Archive) SetMode(mode func(name string, m fs.FileMode) fs.FileMode)

// Removeはアーカイブから任意のパターンにマッチするファイルを削除します。
// パターンはpath.Matchの文法を使用しており、**/で始まるか/**で終わることができます。
// これにより、メインのマッチングの前や後に任意のパス要素（パス要素がない場合も含む）がマッチします。
func (a *Archive) Remove(patterns ...string)

// SetTimeはアーカイブ内のすべてのファイルの変更時刻をtに設定します。
func (a *Archive) SetTime(t time.Time)

// RenameGoModはアーカイブ内のgo.modファイルを_go.modに名前変更します。
// モジュール形式では、他のgo.modファイルを含むことができないため、この名前変更が必要です。
func (a *Archive) RenameGoMod()
