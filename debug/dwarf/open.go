// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
パッケージdwarfは、実行可能ファイルからロードされたDWARFデバッグ情報へのアクセスを提供します。
これは、DWARF 2.0標準で定義されています。
http://dwarfstd.org/doc/dwarf-2.0.0.pdf

# セキュリティ

このパッケージは、敵対的な入力に対して強化されていないため、https://go.dev/security/policy の範囲外です。
特に、オブジェクトファイルを解析する際には基本的な検証しか行われません。
そのため、信頼できない入力を解析する場合は注意が必要です。
不正なファイルを解析すると、大量のリソースを消費したり、パニックを引き起こす可能性があるためです。
*/
package dwarf

import (
	"github.com/shogo82148/std/encoding/binary"
)

// Dataは、実行可能ファイル（例えば、ELFまたはMach-O実行可能ファイル）からロードされた
// DWARFデバッグ情報を表します。
type Data struct {
	// raw data
	abbrev   []byte
	aranges  []byte
	frame    []byte
	info     []byte
	line     []byte
	pubnames []byte
	ranges   []byte
	str      []byte

	// New sections added in DWARF 5.
	addr       []byte
	lineStr    []byte
	strOffsets []byte
	rngLists   []byte

	// parsed data
	abbrevCache map[uint64]abbrevTable
	bigEndian   bool
	order       binary.ByteOrder
	typeCache   map[Offset]Type
	typeSigs    map[uint64]*typeUnit
	unit        []unit
}

<<<<<<< HEAD
// New returns a new [Data] object initialized from the given parameters.
// Rather than calling this function directly, clients should typically use
// the DWARF method of the File type of the appropriate package [debug/elf],
// [debug/macho], or [debug/pe].
=======
// Newは、指定されたパラメータから初期化された新しいDataオブジェクトを返します。
// この関数を直接呼び出す代わりに、クライアントは通常、適切なパッケージdebug/elf、debug/macho、またはdebug/peのFile型のDWARFメソッドを使用する必要があります。
>>>>>>> release-branch.go1.21
//
// []byte引数は、オブジェクトファイルの対応するデバッグセクションからのデータです。
// たとえば、ELFオブジェクトの場合、abbrevは".debug_abbrev"セクションの内容です。
func New(abbrev, aranges, frame, info, line, pubnames, ranges, str []byte) (*Data, error)

// AddTypesは、DWARFデータに1つの.debug_typesセクションを追加します。
// DWARFバージョン4のデバッグ情報を持つ典型的なオブジェクトには、複数の.debug_typesセクションがあります。
// 名前はエラー報告のみに使用され、1つの.debug_typesセクションを別のセクションと区別するために使用されます。
func (d *Data) AddTypes(name string, types []byte) error

// AddSectionは、名前で指定された別のDWARFセクションを追加します。
// 名前は、".debug_addr"、".debug_str_offsets"などのDWARFセクション名である必要があります。
// このアプローチは、DWARF 5以降で追加された新しいDWARFセクションに使用されます。
func (d *Data) AddSection(name string, contents []byte) error
