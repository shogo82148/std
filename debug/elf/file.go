// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
elfパッケージは、ELFオブジェクトファイルへのアクセスを実装します。

# セキュリティ

このパッケージは敵対的な入力に対して強化されるように設計されておらず、
https://go.dev/security/policy の範囲外です。特に、オブジェクトファイルを解析する際には基本的な
検証のみが行われます。そのため、信頼できない入力を解析する際には注意が必要です。なぜなら、
形式が不正なファイルを解析すると、大量のリソースを消費したり、パニックを引き起こす可能性があるからです。
*/
package elf

import (
	"github.com/shogo82148/std/debug/dwarf"
	"github.com/shogo82148/std/encoding/binary"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// FileHeaderはELFファイルヘッダーを表します。
type FileHeader struct {
	Class      Class
	Data       Data
	Version    Version
	OSABI      OSABI
	ABIVersion uint8
	ByteOrder  binary.ByteOrder
	Type       Type
	Machine    Machine
	Entry      uint64
}

// Fileは開いているELFファイルを表します。
type File struct {
	FileHeader
	Sections  []*Section
	Progs     []*Prog
	closer    io.Closer
	gnuNeed   []verneed
	gnuVersym []byte
}

// SectionHeaderは単一のELFセクションヘッダーを表します。
type SectionHeader struct {
	Name      string
	Type      SectionType
	Flags     SectionFlag
	Addr      uint64
	Offset    uint64
	Size      uint64
	Link      uint32
	Info      uint32
	Addralign uint64
	Entsize   uint64

	// FileSizeは、ファイル内のこのセクションのサイズをバイト単位で表します。
	// セクションが圧縮されている場合、FileSizeは圧縮データのサイズであり、
	// Size（上記）は非圧縮データのサイズです。
	FileSize uint64
}

// Sectionは、ELFファイル内の単一のセクションを表します。
type Section struct {
	SectionHeader

	// ReadAtメソッドのためにReaderAtを埋め込みます。
	// ReadとSeekを避けるために、SectionReaderを直接埋め込まないでください。
	// クライアントがReadとSeekを使用したい場合は、
	// 他のクライアントとのシークオフセットの競合を避けるために
	// Open()を使用する必要があります。
	//
	// セクションがランダムアクセス形式で簡単に利用できない場合、
	// ReaderAtはnilになる可能性があります。例えば、圧縮されたセクションは
	// ReaderAtがnilになるかもしれません。
	io.ReaderAt
	sr *io.SectionReader

	compressionType   CompressionType
	compressionOffset int64
}

// DataはELFセクションの内容を読み取り、返します。
// セクションがELFファイル内で圧縮されて保存されていても、
// Dataは非圧縮データを返します。
//
// [SHT_NOBITS] セクションの場合、Dataは常に非nilのエラーを返します。
func (s *Section) Data() ([]byte, error)

// Openは、ELFセクションを読み取る新しいReadSeekerを返します。
// セクションがELFファイル内で圧縮されて保存されていても、
// ReadSeekerは非圧縮データを読み取ります。
//
// [SHT_NOBITS] セクションの場合、開いたリーダーへのすべての呼び出しは
// 非nilのエラーを返します。
func (s *Section) Open() io.ReadSeeker

// ProgHeaderは、単一のELFプログラムヘッダーを表します。
type ProgHeader struct {
	Type   ProgType
	Flags  ProgFlag
	Off    uint64
	Vaddr  uint64
	Paddr  uint64
	Filesz uint64
	Memsz  uint64
	Align  uint64
}

// Progは、ELFバイナリ内の単一のELFプログラムヘッダーを表します。
type Prog struct {
	ProgHeader

	// ReadAtメソッドのためにReaderAtを埋め込みます。
	// ReadとSeekを避けるために、SectionReaderを直接埋め込まないでください。
	// クライアントがReadとSeekを使用したい場合は、
	// 他のクライアントとのシークオフセットの競合を避けるために
	// Open()を使用する必要があります。
	io.ReaderAt
	sr *io.SectionReader
}

// Openは、ELFプログラム本体を読み取る新しいReadSeekerを返します。
func (p *Prog) Open() io.ReadSeeker

// Symbolは、ELFシンボルテーブルセクションのエントリを表します。
type Symbol struct {
	Name        string
	Info, Other byte
	Section     SectionIndex
	Value, Size uint64

	// VersionとLibraryは、動的シンボルテーブルにのみ存在します。
	Version string
	Library string
}

type FormatError struct {
	off int64
	msg string
	val any
}

func (e *FormatError) Error() string

// Openは [os.Open] を使用して指定された名前のファイルを開き、ELFバイナリとしての使用を準備します。
func Open(name string) (*File, error)

// Closeは [File] を閉じます。
// [File] が [Open] ではなく [NewFile] を直接使用して作成された場合、
// Closeは何も影響を与えません。
func (f *File) Close() error

// SectionByTypeは、指定されたタイプを持つf内の最初のセクションを返します。
// そのようなセクションがない場合はnilを返します。
func (f *File) SectionByType(typ SectionType) *Section

// NewFileは、基礎となるリーダー内のELFバイナリにアクセスするための新しい [File] を作成します。
// ELFバイナリは、ReaderAtの位置0で開始することが期待されます。
func NewFile(r io.ReaderAt) (*File, error)

// ErrNoSymbolsは、[File.Symbols] と [File.DynamicSymbols] によって返されます。
// ファイルにそのようなセクションがない場合に返されます。
var ErrNoSymbols = errors.New("no symbol section")

// Sectionは、指定された名前を持つセクションを返します。
// そのようなセクションがない場合はnilを返します。
func (f *File) Section(name string) *Section

func (f *File) DWARF() (*dwarf.Data, error)

// Symbolsは、fのシンボルテーブルを返します。シンボルは、f内に出現する順序でリストされます。
//
// Go 1.0との互換性のため、Symbolsはインデックス0のnullシンボルを省略します。
// シンボルをsymtabとして取得した後、外部から供給されたインデックスxは
// symtab[x]ではなく、symtab[x-1]に対応します。
func (f *File) Symbols() ([]Symbol, error)

// DynamicSymbolsは、fの動的シンボルテーブルを返します。シンボルは、f内に出現する順序でリストされます。
//
// fがシンボルバージョンテーブルを持っている場合、返される [File.Symbols] は
// 初期化されたVersionとLibraryフィールドを持ちます。
//
// [File.Symbols] との互換性のため、[File.DynamicSymbols] はインデックス0のnullシンボルを省略します。
// シンボルをsymtabとして取得した後、外部から供給されたインデックスxは
// symtab[x]ではなく、symtab[x-1]に対応します。
func (f *File) DynamicSymbols() ([]Symbol, error)

type ImportedSymbol struct {
	Name    string
	Version string
	Library string
}

// ImportedSymbolsは、動的ロード時に他のライブラリによって満たされることが期待される
// バイナリfによって参照されるすべてのシンボルの名前を返します。
// 弱いシンボルは返しません。
func (f *File) ImportedSymbols() ([]ImportedSymbol, error)

// ImportedLibrariesは、動的リンク時にバイナリとリンクされることが期待される
// バイナリfによって参照されるすべてのライブラリの名前を返します。
func (f *File) ImportedLibraries() ([]string, error)

// DynStringは、ファイルの動的セクションで指定されたタグにリストされている文字列を返します。
//
// タグは、文字列値を取るものでなければなりません：[DT_NEEDED]、[DT_SONAME]、[DT_RPATH]、または
// [DT_RUNPATH]。
func (f *File) DynString(tag DynTag) ([]string, error)

// DynValueは、ファイルの動的セクションで指定されたタグにリストされている値を返します。
func (f *File) DynValue(tag DynTag) ([]uint64, error)
