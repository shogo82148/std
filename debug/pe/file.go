// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
パッケージpeは、PE（Microsoft Windows Portable Executable）ファイルへのアクセスを実装します。

# セキュリティ

このパッケージは、敵対的な入力に対して強化されるように設計されていませんし、
https://go.dev/security/policy の範囲外です。特に、オブジェクトファイルを解析する際には基本的な
検証のみが行われます。そのため、信頼できない入力を解析する際には注意が必要です。なぜなら、
形式が不正なファイルを解析すると、大量のリソースを消費したり、パニックを引き起こす可能性があるからです。
*/
package pe

import (
	"github.com/shogo82148/std/debug/dwarf"
	"github.com/shogo82148/std/io"
)

// Fileは、開かれたPEファイルを表します。
type File struct {
	FileHeader
	OptionalHeader any
	Sections       []*Section
	Symbols        []*Symbol
	COFFSymbols    []COFFSymbol
	StringTable    StringTable

	closer io.Closer
}

// Openは、[os.Open] を使用して指定されたファイルを開き、それをPEバイナリとして使用するための準備をします。
func Open(name string) (*File, error)

// Closeは、[File] を閉じます。
// [File] が [Open] ではなく [NewFile] を直接使用して作成された場合、
// Closeは何も影響を与えません。
func (f *File) Close() error

// NewFileは、基礎となるリーダーでPEバイナリにアクセスするための新しいFileを作成します。
func NewFile(r io.ReaderAt) (*File, error)

// Sectionは、指定された名前の最初のセクションを返します。そのような
// セクションが存在しない場合はnilを返します。
func (f *File) Section(name string) *Section

func (f *File) DWARF() (*dwarf.Data, error)

type ImportDirectory struct {
	OriginalFirstThunk uint32
	TimeDateStamp      uint32
	ForwarderChain     uint32
	Name               uint32
	FirstThunk         uint32

	dll string
}

// ImportedSymbolsは、動的ロード時に他のライブラリによって満たされることが期待されている、
// バイナリfが参照しているすべてのシンボルの名前を返します。
// それは弱いシンボルを返しません。
func (f *File) ImportedSymbols() ([]string, error)

// ImportedLibrariesは、動的リンク時にバイナリとリンクされることが期待されている、
// バイナリfが参照しているすべてのライブラリの名前を返します。
func (f *File) ImportedLibraries() ([]string, error)

// FormatErrorは使用されていません。
// この型は互換性のために保持されています。
type FormatError struct {
}

func (e *FormatError) Error() string
