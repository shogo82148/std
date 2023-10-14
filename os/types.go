// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/io/fs"
)

// Getpagesizeは、基礎となるシステムのメモリページサイズを返します。
func Getpagesize() int

// Fileはオープンされたファイルディスクリプタを表します。
type File struct {
	*file
}

// FileInfoはファイルを記述し、StatおよびLstatによって返されます。
type FileInfo = fs.FileInfo

// FileModeはファイルのモードと許可ビットを表します。
// ビットはすべてのシステムで同じ定義を持っているため、
// ファイルの情報をシステム間で移動する際に移植性があります。
// すべてのビットがすべてのシステムで適用されるわけではありません。
// 必須のビットはModeDirであり、ディレクトリに対して適用されます。
type FileMode = fs.FileMode

// 定義されたファイルモードのビットは、FileModeの最上位ビットです。
// 最下位の9ビットは、標準のUnixのrwxrwxrwxパーミッションです。
// これらのビットの値は、パブリックAPIの一部と見なされ、
// ワイヤープロトコルやディスクの表現で使用される場合があります。
// これらの値は変更しないでくださいが、新しいビットが追加されるかもしれません。
const (

	// 単一の文字は、Stringメソッドの書式設定で使用される略語です。
	ModeDir        = fs.ModeDir
	ModeAppend     = fs.ModeAppend
	ModeExclusive  = fs.ModeExclusive
	ModeTemporary  = fs.ModeTemporary
	ModeSymlink    = fs.ModeSymlink
	ModeDevice     = fs.ModeDevice
	ModeNamedPipe  = fs.ModeNamedPipe
	ModeSocket     = fs.ModeSocket
	ModeSetuid     = fs.ModeSetuid
	ModeSetgid     = fs.ModeSetgid
	ModeCharDevice = fs.ModeCharDevice
	ModeSticky     = fs.ModeSticky
	ModeIrregular  = fs.ModeIrregular

	// タイプビット用のマスク。通常のファイルでは、何も設定されません。
	ModeType = fs.ModeType

	ModePerm = fs.ModePerm
)

// SameFileはfi1とfi2が同じファイルを表しているかどうかを報告します。
// 例えば、Unixでは、2つの基礎となる構造体のデバイスとinodeフィールドが同一であることを意味します。他のシステムでは、決定はパス名に基づく場合もあります。
// SameFileは、このパッケージのStatによって返された結果にのみ適用されます。
// それ以外の場合はfalseを返します。
func SameFile(fi1, fi2 FileInfo) bool
