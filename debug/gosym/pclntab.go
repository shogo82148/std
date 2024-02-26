// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Line tables
 */

package gosym

import (
	"github.com/shogo82148/std/encoding/binary"
	"github.com/shogo82148/std/sync"
)

// LineTableは、プログラムカウンタを行番号にマッピングするデータ構造です。
//
// Go 1.1以前では、各関数（[Func] によって表される）は独自のLineTableを持ち、
// 行番号はプログラム内のすべてのソース行を通じての番号付けに対応していました。
// その絶対行番号は、別途ファイル名とファイル内の行番号に変換する必要がありました。
//
// Go 1.2では、データの形式が変更され、プログラム全体で単一のLineTableが存在し、
// すべてのFuncが共有し、絶対行番号はなく、特定のファイル内の行番号のみが存在します。
//
// 大部分において、LineTableのメソッドはパッケージの内部詳細として扱うべきであり、
// 呼び出し元は代わりに [Table] のメソッドを使用するべきです。
type LineTable struct {
	Data []byte
	PC   uint64
	Line int

	// This mutex is used to keep parsing of pclntab synchronous.
	mu sync.Mutex

	// Contains the version of the pclntab section.
	version version

	// Go 1.2/1.16/1.18 state
	binary      binary.ByteOrder
	quantum     uint32
	ptrsize     uint32
	textStart   uint64
	funcnametab []byte
	cutab       []byte
	funcdata    []byte
	functab     []byte
	nfunctab    uint32
	filetab     []byte
	pctab       []byte
	nfiletab    uint32
	funcNames   map[uint32]string
	strings     map[uint32]string
	// fileMap varies depending on the version of the object file.
	// For ver12, it maps the name to the index in the file table.
	// For ver116, it maps the name to the offset in filetab.
	fileMap map[string]uint32
}

// PCToLineは、指定されたプログラムカウンタに対応する行番号を返します。
//
// Deprecated: 代わりにTableのPCToLineメソッドを使用してください。
func (t *LineTable) PCToLine(pc uint64) int

// LineToPCは、指定された行番号に対応するプログラムカウンタを返します。
// ただし、maxpcより前のプログラムカウンタのみを考慮します。
//
// Deprecated: 代わりにTableのLineToPCメソッドを使用してください。
func (t *LineTable) LineToPC(line int, maxpc uint64) uint64

<<<<<<< HEAD
// NewLineTableは、エンコードされたデータに対応する新しいPC/行テーブルを返します。
// Textは、対応するテキストセグメントの開始アドレスでなければなりません。
=======
// NewLineTable returns a new PC/line table
// corresponding to the encoded data.
// Text must be the start address of the
// corresponding text segment, with the exact
// value stored in the 'runtime.text' symbol.
// This value may differ from the start
// address of the text segment if
// binary was built with cgo enabled.
>>>>>>> upstream/master
func NewLineTable(data []byte, text uint64) *LineTable
