// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pe

const COFFSymbolSize = 18

// COFFSymbolは、単一のCOFFシンボルテーブルレコードを表します。
type COFFSymbol struct {
	Name               [8]uint8
	Value              uint32
	SectionNumber      int16
	Type               uint16
	StorageClass       uint8
	NumberOfAuxSymbols uint8
}

// FullNameは、シンボルsymの実際の名前を見つけます。通常、名前は
// sym.Nameに格納されますが、それが8文字より長い場合、代わりに
// COFF文字列テーブルstに格納されます。
func (sym *COFFSymbol) FullName(st StringTable) (string, error)

// Symbolは、NameフィールドがGoの文字列に置き換えられ、
// NumberOfAuxSymbolsが存在しないCOFFSymbolと似ています。
type Symbol struct {
	Name          string
	Value         uint32
	SectionNumber int16
	Type          uint16
	StorageClass  uint8
}

// COFFSymbolAuxFormat5は、セクション定義シンボルに付随するauxシンボルの予想される形式を説明します。
// PEフォーマットは、関数定義のためのフォーマット1、.beおよび.efシンボルのためのフォーマット2など、
// いくつかの異なるauxシンボルフォーマットを定義します。フォーマット5は、セクション定義に関連する追加情報を保持し、
// 再配置の数+行番号、およびCOMDAT情報を含みます。ここで何が起こっているのかについての詳細は、
// https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#auxiliary-format-5-section-definitions を参照してください。
type COFFSymbolAuxFormat5 struct {
	Size           uint32
	NumRelocs      uint16
	NumLineNumbers uint16
	Checksum       uint32
	SecNum         uint16
	Selection      uint8
	_              [3]uint8
}

// これらの定数は、AuxFormat5の 'Selection'
// フィールドの可能な値を構成します。
const (
	IMAGE_COMDAT_SELECT_NODUPLICATES = 1
	IMAGE_COMDAT_SELECT_ANY          = 2
	IMAGE_COMDAT_SELECT_SAME_SIZE    = 3
	IMAGE_COMDAT_SELECT_EXACT_MATCH  = 4
	IMAGE_COMDAT_SELECT_ASSOCIATIVE  = 5
	IMAGE_COMDAT_SELECT_LARGEST      = 6
)

å
func (f *File) COFFSymbolReadSectionDefAux(idx int) (*COFFSymbolAuxFormat5, error)
