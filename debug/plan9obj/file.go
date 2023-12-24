// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
plan9objパッケージは、Plan 9 a.outオブジェクトファイルへのアクセスを実装します。

# セキュリティ

このパッケージは、敵対的な入力に対して強化されるように設計されていませんし、
https://go.dev/security/policy の範囲外です。特に、オブジェクトファイルを解析する際には基本的な
検証のみが行われます。そのため、信頼できない入力を解析する際には注意が必要です。なぜなら、
不正なファイルを解析すると、大量のリソースを消費したり、パニックを引き起こす可能性があるからです。
*/
package plan9obj

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// FileHeaderは、Plan 9 a.outファイルヘッダーを表します。
type FileHeader struct {
	Magic       uint32
	Bss         uint32
	Entry       uint64
	PtrSize     int
	LoadAddress uint64
	HdrSize     uint64
}

// Fileは、開いているPlan 9 a.outファイルを表します。
type File struct {
	FileHeader
	Sections []*Section
	closer   io.Closer
}

// SectionHeaderは、単一のPlan 9 a.outセクションヘッダーを表します。
// この構造体はディスク上には存在せず、オブジェクトファイルを通じた
// ナビゲーションを容易にします。
type SectionHeader struct {
	Name   string
	Size   uint32
	Offset uint32
}

// Sectionは、Plan 9 a.outファイルの単一のセクションを表します。
type Section struct {
	SectionHeader

	// ReadAtメソッドのためにReaderAtを埋め込みます。
	// ReadとSeekを持つことを避けるために、
	// SectionReaderを直接埋め込むことはありません。
	// クライアントがReadとSeekを使用したい場合は、
	// 他のクライアントとのシークオフセットの競合を避けるために
	// Open()を使用する必要があります。
	io.ReaderAt
	sr *io.SectionReader
}

// Dataは、Plan 9 a.outセクションの内容を読み取り、返します。
func (s *Section) Data() ([]byte, error)

// Openは、Plan 9 a.outセクションを読み取る新しいReadSeekerを返します。
func (s *Section) Open() io.ReadSeeker

// Symは、Plan 9 a.outのシンボルテーブルセクションのエントリを表します。
type Sym struct {
	Value uint64
	Type  rune
	Name  string
}

// Openは、os.Openを使用して指定された名前のファイルを開き、
// それをPlan 9 a.outバイナリとして使用するための準備をします。
func Open(name string) (*File, error)

// Closeは、Fileを閉じます。
// FileがOpenではなくNewFileを直接使用して作成された場合、
// Closeは何も影響を及ぼしません。
func (f *File) Close() error

// NewFileは、基礎となるリーダーでPlan 9バイナリにアクセスするための新しいFileを作成します。
// Plan 9バイナリは、ReaderAtの位置0で開始することが期待されます。
func NewFile(r io.ReaderAt) (*File, error)

// ErrNoSymbolsは、File内にそのようなセクションがない場合に、
// File.Symbolsによって返されるエラーです。
var ErrNoSymbols = errors.New("no symbol section")

// Symbolsは、fのシンボルテーブルを返します。
func (f *File) Symbols() ([]Sym, error)

// Sectionは、指定された名前のセクションを返します。
// そのようなセクションが存在しない場合はnilを返します。
func (f *File) Section(name string) *Section
