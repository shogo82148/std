// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
machoパッケージは、Mach-Oオブジェクトファイルへのアクセスを実装します。

# セキュリティ

このパッケージは、敵対的な入力に対して強化されるように設計されていません、そして
https://go.dev/security/policy の範囲外です。特に、オブジェクトファイルを解析する際には基本的な
検証のみが行われます。そのため、信頼できない入力を解析する際には注意が必要です、なぜなら、
不正なファイルを解析すると、大量のリソースを消費するか、パニックを引き起こす可能性があります。
*/
package macho

import (
	"github.com/shogo82148/std/debug/dwarf"
	"github.com/shogo82148/std/encoding/binary"
	"github.com/shogo82148/std/io"
)

// Fileは、開かれたMach-Oファイルを表します。
type File struct {
	FileHeader
	ByteOrder binary.ByteOrder
	Loads     []Load
	Sections  []*Section

	Symtab   *Symtab
	Dysymtab *Dysymtab

	closer io.Closer
}

// Loadは、任意のMach-Oロードコマンドを表します。
type Load interface {
	Raw() []byte
}

// LoadBytesは、Mach-Oロードコマンドの解釈されていないバイトを表します。
type LoadBytes []byte

func (b LoadBytes) Raw() []byte

// SegmentHeaderは、Mach-O 32ビットまたは64ビットのロードセグメントコマンドのヘッダーです。
type SegmentHeader struct {
	Cmd     LoadCmd
	Len     uint32
	Name    string
	Addr    uint64
	Memsz   uint64
	Offset  uint64
	Filesz  uint64
	Maxprot uint32
	Prot    uint32
	Nsect   uint32
	Flag    uint32
}

// Segmentは、Mach-O 32ビットまたは64ビットのロードセグメントコマンドを表します。
type Segment struct {
	LoadBytes
	SegmentHeader

	// ReadAtメソッドのためにReaderAtを埋め込みます。
	// ReadとSeekを避けるために、SectionReaderを直接埋め込まないでください。
	// クライアントがReadとSeekを使用したい場合は、
	// 他のクライアントとのシークオフセットの競合を避けるために
	// Open()を使用する必要があります。
	io.ReaderAt
	sr *io.SectionReader
}

// Dataはセグメントの内容を読み取り、返します。
func (s *Segment) Data() ([]byte, error)

// Openは、セグメントを読み取る新しいReadSeekerを返します。
func (s *Segment) Open() io.ReadSeeker

type SectionHeader struct {
	Name   string
	Seg    string
	Addr   uint64
	Size   uint64
	Offset uint32
	Align  uint32
	Reloff uint32
	Nreloc uint32
	Flags  uint32
}

// Relocは、Mach-Oの再配置を表します。
type Reloc struct {
	Addr  uint32
	Value uint32
	// Scattered == false かつ Extern == true の場合、Valueはシンボル番号です。
	// Scattered == false かつ Extern == false の場合、Valueはセクション番号です。
	// Scattered == true の場合、Valueはこの再配置が参照する値です。
	Type      uint8
	Len       uint8
	Pcrel     bool
	Extern    bool
	Scattered bool
}

type Section struct {
	SectionHeader
	Relocs []Reloc

	// ReadAtメソッドのためにReaderAtを埋め込みます。
	// ReadとSeekを避けるために、SectionReaderを直接埋め込まないでください。
	// クライアントがReadとSeekを使用したい場合は、
	// 他のクライアントとのシークオフセットの競合を避けるために
	// Open()を使用する必要があります。
	io.ReaderAt
	sr *io.SectionReader
}

// Dataは、Mach-Oセクションの内容を読み取り、返します。
func (s *Section) Data() ([]byte, error)

// Openは、Mach-Oセクションを読み取る新しいReadSeekerを返します。
func (s *Section) Open() io.ReadSeeker

// Dylibは、Mach-Oの動的ライブラリロードコマンドを表します。
type Dylib struct {
	LoadBytes
	Name           string
	Time           uint32
	CurrentVersion uint32
	CompatVersion  uint32
}

// Symtabは、Mach-Oのシンボルテーブルコマンドを表します。
type Symtab struct {
	LoadBytes
	SymtabCmd
	Syms []Symbol
}

// A Dysymtab represents a Mach-O dynamic symbol table command.
type Dysymtab struct {
	LoadBytes
	DysymtabCmd
	IndirectSyms []uint32
}

// Rpathは、Mach-O rpathコマンドを表します。
type Rpath struct {
	LoadBytes
	Path string
}

// Symbolは、Mach-O 32ビットまたは64ビットのシンボルテーブルエントリです。
type Symbol struct {
	Name  string
	Type  uint8
	Sect  uint8
	Desc  uint16
	Value uint64
}

// FormatErrorは、データがオブジェクトファイルの正しい形式でない場合、
// 一部の操作によって返されます。
type FormatError struct {
	off int64
	msg string
	val any
}

func (e *FormatError) Error() string

<<<<<<< HEAD
// Open opens the named file using [os.Open] and prepares it for use as a Mach-O binary.
func Open(name string) (*File, error)

// Close closes the [File].
// If the [File] was created using [NewFile] directly instead of [Open],
// Close has no effect.
func (f *File) Close() error

// NewFile creates a new [File] for accessing a Mach-O binary in an underlying reader.
// The Mach-O binary is expected to start at position 0 in the ReaderAt.
=======
// Openは、os.Openを使用して指定されたファイルを開き、それをMach-Oバイナリとして使用するための準備をします。
func Open(name string) (*File, error)

// Closeは、Fileを閉じます。
// FileがOpenではなくNewFileを直接使用して作成された場合、
// Closeは何も影響を与えません。
func (f *File) Close() error

// NewFileは、基礎となるリーダーでMach-Oバイナリにアクセスするための新しいFileを作成します。
// Mach-Oバイナリは、ReaderAtの位置0で開始することが期待されています。
>>>>>>> release-branch.go1.21
func NewFile(r io.ReaderAt) (*File, error)

// Segmentは、指定された名前の最初のSegmentを返します。そのようなセグメントが存在しない場合はnilを返します。
func (f *File) Segment(name string) *Segment

// Sectionは、指定された名前の最初のセクションを返します。そのような
// セクションが存在しない場合はnilを返します。
func (f *File) Section(name string) *Section

// DWARFは、Mach-OファイルのDWARFデバッグ情報を返します。
func (f *File) DWARF() (*dwarf.Data, error)

// ImportedSymbolsは、バイナリfが参照しているすべてのシンボルの名前を返します。
// これらは、動的ロード時に他のライブラリによって満たされることが期待されています。
func (f *File) ImportedSymbols() ([]string, error)

// ImportedLibrariesは、バイナリfが参照しているすべてのライブラリのパスを返します。
// これらは、動的リンク時にバイナリとリンクされることが期待されています。
func (f *File) ImportedLibraries() ([]string, error)
