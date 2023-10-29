// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dwarf

import (
	"github.com/shogo82148/std/errors"
)

// LineReaderは、単一のコンパイルユニットのDWARF「line」セクションからLineEntry構造体のシーケンスを読み取ります。
// LineEntryは、PCの増加順に発生し、各LineEntryは、そのLineEntryのPCから次のLineEntryのPCの直前までの命令のメタデータを提供します。
// 最後のエントリには、EndSequenceフィールドが設定されます。
type LineReader struct {
	buf buf

	// Original .debug_line section data. Used by Seek.
	section []byte

	str     []byte
	lineStr []byte

	// Header information
	version              uint16
	addrsize             int
	segmentSelectorSize  int
	minInstructionLength int
	maxOpsPerInstruction int
	defaultIsStmt        bool
	lineBase             int
	lineRange            int
	opcodeBase           int
	opcodeLengths        []int
	directories          []string
	fileEntries          []*LineFile

	programOffset Offset
	endOffset     Offset

	initialFileEntries int

	// Current line number program state machine registers
	state     LineEntry
	fileIndex int
}

// LineEntryは、DWARF行テーブル内の行を表します。
type LineEntry struct {
	// Addressは、コンパイラによって生成されたマシン命令のプログラムカウンター値です。
	// このLineEntryは、Addressから次のLineEntryのAddressの直前までの各命令に適用されます。
	Address uint64

	// OpIndexは、VLIW命令内の操作のインデックスです。
	// 最初の操作のインデックスは0です。非VLIWアーキテクチャの場合、常に0になります。
	// AddressとOpIndexは、命令ストリーム内の任意の個々の操作を参照できる操作ポインターを形成します。
	OpIndex int

	// Fileは、これらの命令に対応するソースファイルです。
	File *LineFile

	// Lineは、これらの命令に対応するソースコードの行番号です。
	// 行番号は1から始まります。これらの命令がどのソース行にも関連付けられていない場合は、0になる場合があります。
	Line int

	// Columnは、これらの命令のソース行内の列番号です。
	// 列番号は1から始まります。行の「左端」を示すために0になる場合があります。
	Column int

	// IsStmtは、Addressが推奨されるブレークポイントの場所であることを示します。
	// 例えば、行の始まり、文の始まり、または文の明確な部分などです。
	IsStmt bool

	// BasicBlockは、Addressが基本ブロックの開始であることを示します。
	BasicBlock bool

	// PrologueEndは、Addressが、含まれる関数へのエントリにブレークポイントを設定するために、
	// 実行を一時停止する必要があるPCの1つ（可能性がある）であることを示します。
	//
	// DWARF 3で追加されました。
	PrologueEnd bool

	// EpilogueBeginは、Addressが、この関数からの終了時にブレークポイントを設定するために、
	// 実行を一時停止する必要があるPCの1つ（可能性がある）であることを示します。
	//
	// DWARF 3で追加されました。
	EpilogueBegin bool

	// ISAは、これらの命令の命令セットアーキテクチャを表します。
	// 可能なISA値は、適用可能なABI仕様によって定義される必要があります。
	//
	// DWARF 3で追加されました。
	ISA int

	// Discriminatorは、これらの命令が属するブロックを示す任意の整数です。
	// これにより、同じソースファイル、行、列を持つ複数のブロックを区別できます。
	// 特定のソース位置に1つのブロックしか存在しない場合、0にする必要があります。
	//
	// DWARF 3で追加されました。
	Discriminator int

	// EndSequenceは、Addressがターゲットマシン命令のシーケンスの終わりの直後の最初のバイトであることを示します。
	// 設定されている場合、このフィールドとAddressフィールドのみが有意です。
	// 行番号テーブルには、複数の可能性のある不連続な命令シーケンスの情報が含まれる場合があります。
	// 行テーブルの最後のエントリには、常にEndSequenceが設定されている必要があります。
	EndSequence bool
}

// LineFileは、DWARF行テーブルエントリによって参照されるソースファイルです。
type LineFile struct {
	Name   string
	Mtime  uint64
	Length int
}

// LineReaderは、TagCompileUnitを持つEntry cuの行テーブルのための新しいリーダーを返します。
//
// このコンパイルユニットに行テーブルがない場合、nil、nilを返します。
func (d *Data) LineReader(cu *Entry) (*LineReader, error)

// Nextは、この行テーブルの次の行を*entryに設定し、次の行に移動します。
// もうエントリがなく、行テーブルが適切に終了している場合、io.EOFを返します。
//
// 行は常にentry.Addressの増加順に並んでいますが、entry.Lineは前後に移動する場合があります。
func (r *LineReader) Next(entry *LineEntry) error

// LineReaderPosは、行テーブル内の位置を表します。
type LineReaderPos struct {
	// off is the current offset in the DWARF line section.
	off Offset
	// numFileEntries is the length of fileEntries.
	numFileEntries int
	// state and fileIndex are the statement machine state at
	// offset off.
	state     LineEntry
	fileIndex int
}

// Tellは、行テーブル内の現在の位置を返します。
func (r *LineReader) Tell() LineReaderPos

// Seekは、Tellによって返された位置に行テーブルリーダーを復元します。
//
// 引数posは、この行テーブルのTell呼び出しによって返されたものである必要があります。
func (r *LineReader) Seek(pos LineReaderPos)

// Resetは、行テーブルリーダーを行テーブルの先頭に再配置します。
func (r *LineReader) Reset()

// Filesは、現在の行テーブルの位置に基づいて、このコンパイルユニットのファイル名テーブルを返します。
// ファイル名テーブルは、AttrDeclFileなどのこのコンパイルユニットの属性から参照される場合があります。
//
// Entry 0は常にnilです。なぜなら、ファイルインデックス0は「ファイルなし」を表すからです。
//
// コンパイルユニットのファイル名テーブルは固定されていません。Filesは、
// 行テーブルの現在の位置に基づいてファイルテーブルを返します。
// これにより、行テーブルの以前の位置のファイルテーブルよりもエントリが多く含まれる場合がありますが、
// 既存のエントリは変更されません。
func (r *LineReader) Files() []*LineFile

// ErrUnknownPCは、LineReader.ScanPCが行テーブルのエントリによってカバーされていないPCを検出した場合に返されるエラーです。
var ErrUnknownPC = errors.New("ErrUnknownPC")

// SeekPCは、pcを含むLineEntryを*entryに設定し、
// 行テーブルの次のエントリに位置を設定します。
// 必要に応じて、pcを検索するために後方にシークします。
//
// pcがこの行テーブルのエントリによってカバーされていない場合、
// SeekPCはErrUnknownPCを返します。この場合、*entryと最終的なシーク位置は未指定です。
//
// DWARF行テーブルは、順次前方スキャンのみを許可します。
// したがって、最悪の場合、これには行テーブルのサイズに比例する時間がかかります。
// 呼び出し側が繰り返し高速なPC検索を行いたい場合は、
// 適切な行テーブルのインデックスを構築する必要があります。
func (r *LineReader) SeekPC(pc uint64, entry *LineEntry) error
