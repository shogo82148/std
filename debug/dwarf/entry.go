// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DWARF debug information entry parser.
// An entry is a sequence of data items of a given format.
// The first word in the entry is an index into what DWARF
// calls the ``abbreviation table.''  An abbreviation is really
// just a type descriptor: it's an array of attribute tag/value format pairs.

package dwarf

import (
	"github.com/shogo82148/std/encoding/binary"
)

// エントリは、属性/値のペアのシーケンスです。
type Entry struct {
	Offset   Offset
	Tag      Tag
	Children bool
	Field    []Field
}

// Fieldは、Entry内の単一の属性/値ペアです。
//
// 値は、DWARFによって定義されたいくつかの「属性クラス」のうちの1つです。
// 各クラスに対応するGoの型は次のとおりです。
//
//	DWARFクラス       Goの型        クラス
//	-----------       -------        -----
//	address           uint64         ClassAddress
//	block             []byte         ClassBlock
//	constant          int64          ClassConstant
//	flag              bool           ClassFlag
//	reference
//	  to info         dwarf.Offset   ClassReference
//	  to type unit    uint64         ClassReferenceSig
//	string            string         ClassString
//	exprloc           []byte         ClassExprLoc
//	lineptr           int64          ClassLinePtr
//	loclistptr        int64          ClassLocListPtr
//	macptr            int64          ClassMacPtr
//	rangelistptr      int64          ClassRangeListPtr
//
// 識別できないまたはベンダー定義の属性の場合、ClassはClassUnknownになります。
type Field struct {
	Attr  Attr
	Val   any
	Class Class
}

// Classは、属性値のDWARF 4クラスです。
//
// 一般的に、特定の属性の値は、DWARFによって定義されたいくつかの可能なクラスのうちの1つを取ることができます。
// それぞれのクラスは、属性のわずかに異なる解釈につながります。
//
// DWARFバージョン4は、以前のDWARFバージョンよりも属性値クラスを細かく区別します。
// リーダーは、以前のDWARFバージョンの粗いクラスを、適切なDWARF 4クラスに明確に区別します。
// たとえば、DWARF 2は、定数だけでなくすべてのタイプのセクションオフセットに対して「constant」を使用しますが、
// リーダーは、セクションオフセットを参照するDWARF 2ファイルの属性を、Class*Ptrクラスの1つに正規化します。
// これらのクラスは、DWARF 3でのみ定義されていたにもかかわらずです。
type Class int

const (
	// ClassUnknownは、未知のDWARFクラスの値を表します。
	ClassUnknown Class = iota

	// ClassAddressは、ターゲットマシン上のアドレスであるuint64型の値を表します。
	ClassAddress

	// ClassBlockは、属性に依存する[]byte型の値を表します。
	ClassBlock

	// ClassConstantは、定数であるint64型の値を表します。
	// この定数の解釈は、属性に依存します。
	ClassConstant

	// ClassExprLocは、エンコードされたDWARF式またはロケーション記述を含む[]byte型の値を表します。
	ClassExprLoc

	// ClassFlagは、bool型の値を表します。
	ClassFlag

	// ClassLinePtrは、int64オフセットである値を表します。
	// このオフセットは、"line"セクション内の位置を指します。
	ClassLinePtr

	// ClassLocListPtrは、int64オフセットである値を表します。
	// このオフセットは、"loclist"セクション内の位置を指します。
	ClassLocListPtr

	// ClassMacPtrは、int64オフセットである値を表します。
	// このオフセットは、"mac"セクション内の位置を指します。
	ClassMacPtr

	// ClassRangeListPtrは、int64オフセットである値を表します。
	// このオフセットは、"rangelist"セクション内の位置を指します。
	ClassRangeListPtr

	// ClassReferenceは、infoセクション内のEntryのオフセットを表す値を表します。
	// (Reader.Seekで使用するため)。
	// DWARF仕様は、ClassReferenceとClassReferenceSigをクラス"reference"に結合します。
	ClassReference

	// ClassReferenceSigは、型Entryを参照するuint64型のシグネチャを表す値を表します。
	ClassReferenceSig

	// ClassStringは、文字列を表す値を表します。
	// コンパイルユニットがAttrUseUTF8フラグ（強く推奨）を指定している場合、
	// 文字列値はUTF-8でエンコードされます。それ以外の場合、エンコーディングは未指定です。
	ClassString

	// ClassReferenceAltは、代替オブジェクトファイルのDWARF "info"セクション内のオフセットを表すint64型の値を表します。
	ClassReferenceAlt

	// ClassStringAltは、代替オブジェクトファイルのDWARF文字列セクション内のオフセットを表すint64型の値を表します。
	ClassStringAlt

	// ClassAddrPtrは、"addr"セクション内のint64オフセットである値を表します。
	ClassAddrPtr

	// ClassLocListは、"loclists"セクション内のint64オフセットである値を表します。
	ClassLocList

	// ClassRngListは、"rnglists"セクションのベースからのuint64オフセットを表す値を表します。
	ClassRngList

	// ClassRngListsPtrは、"rnglists"セクション内のint64オフセットである値を表します。
	// これらは、ClassRngList値のベースとして使用されます。
	ClassRngListsPtr

	// ClassStrOffsetsPtrは、"str_offsets"セクション内のint64オフセットである値を表します。
	ClassStrOffsetsPtr
)

func (i Class) GoString() string

// Valは、Entry内の属性Attrに関連付けられた値を返します。
// そのような属性が存在しない場合は、nilを返します。
//
// 一般的なイディオムは、nilの返却値のチェックを、
// 期待される動的型を持つ値のチェックと統合することです。例えば、以下のようになります。
//
//	v, ok := e.Val(AttrSibling).(int64)
func (e *Entry) Val(a Attr) any

// AttrFieldは、Entry内の属性Attrに関連付けられたFieldを返します。
// そのような属性が存在しない場合は、nilを返します。
func (e *Entry) AttrField(a Attr) *Field

// Offsetは、DWARF情報内のEntryの位置を表します。
// (Reader.Seekを参照してください。)
type Offset uint32

// Readerは、DWARF「info」セクションからEntry構造体を読み取ることを可能にします。
// Entry構造体はツリー形式で配置されています。ReaderのNext関数は、
// ツリーの先行順序の走査から連続するエントリを返します。
// エントリに子がある場合、そのChildrenフィールドはtrueになり、子は
// Tag 0のEntryで終了します。
type Reader struct {
	b            buf
	d            *Data
	err          error
	unit         int
	lastUnit     bool
	lastChildren bool
	lastSibling  Offset
	cu           *Entry
}

// Readerは、Dataのための新しいReaderを返します。
// このリーダーは、DWARF「info」セクションのバイトオフセット0に位置しています。
func (d *Data) Reader() *Reader

// AddressSizeは、現在のコンパイルユニットのアドレスのバイト数を返します。
func (r *Reader) AddressSize() int

// ByteOrderは、現在のコンパイルユニットのバイトオーダーを返します。
func (r *Reader) ByteOrder() binary.ByteOrder

// Seekは、エンコードされたエントリストリーム内のオフセットoffにReaderを配置します。
// オフセット0は最初のエントリを示すために使用できます。
func (r *Reader) Seek(off Offset)

// Nextは、エンコードされたエントリストリームから次のエントリを読み取ります。
// セクションの終わりに達した場合、nil、nilを返します。
// 現在のオフセットが無効であるか、オフセットのデータを有効なEntryとしてデコードできない場合、エラーを返します。
func (r *Reader) Next() (*Entry, error)

// SkipChildrenは、Nextによって返された最後のEntryに関連付けられた子エントリをスキップします。
// そのEntryに子がない場合、またはNextが呼び出されていない場合、SkipChildrenは何もしません。
func (r *Reader) SkipChildren()

// SeekPCは、pcを含むコンパイルユニットのEntryを返し、
// そのユニットの子を読み取るためにリーダーを配置します。
// pcがどのユニットにも含まれていない場合、SeekPCはErrUnknownPCを返し、
// リーダーの位置は未定義です。
//
// コンパイルユニットが実行可能ファイルの複数の領域を記述できるため、
// SeekPCは最悪の場合、すべてのコンパイルユニットのすべての範囲を検索する必要があります。
// SeekPCの各呼び出しは、前回の呼び出しのコンパイルユニットから検索を開始するため、
// 一般的には、PCをソートすると、連続する高速なPC検索を行う場合に効果的です。
// 呼び出し元が繰り返し高速なPC検索を行いたい場合は、Rangesメソッドを使用して適切なインデックスを構築する必要があります。
func (r *Reader) SeekPC(pc uint64) (*Entry, error)

// Rangesは、eによってカバーされるPC範囲、つまり[low, high)のペアのスライスを返します。
// TagCompileUnitやTagSubprogramなど、一部のエントリタイプのみがPC範囲を持っています。
// それ以外の場合、エラーを返さずにnilを返します。
func (d *Data) Ranges(e *Entry) ([][2]uint64, error)
