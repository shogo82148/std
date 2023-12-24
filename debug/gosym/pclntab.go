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
<<<<<<< HEAD
// In Go 1.1 and earlier, each function (represented by a [Func]) had its own LineTable,
// and the line number corresponded to a numbering of all source lines in the
// program, across all files. That absolute line number would then have to be
// converted separately to a file name and line number within the file.
=======
// Go 1.1以前では、各関数（Funcによって表される）は独自のLineTableを持ち、
// 行番号はプログラム内のすべてのソース行を通じての番号付けに対応していました。
// その絶対行番号は、別途ファイル名とファイル内の行番号に変換する必要がありました。
>>>>>>> release-branch.go1.21
//
// Go 1.2では、データの形式が変更され、プログラム全体で単一のLineTableが存在し、
// すべてのFuncが共有し、絶対行番号はなく、特定のファイル内の行番号のみが存在します。
//
<<<<<<< HEAD
// For the most part, LineTable's methods should be treated as an internal
// detail of the package; callers should use the methods on [Table] instead.
=======
// 大部分において、LineTableのメソッドはパッケージの内部詳細として扱うべきであり、
// 呼び出し元は代わりにTableのメソッドを使用するべきです。
>>>>>>> release-branch.go1.21
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

// NewLineTableは、エンコードされたデータに対応する新しいPC/行テーブルを返します。
// Textは、対応するテキストセグメントの開始アドレスでなければなりません。
func NewLineTable(data []byte, text uint64) *LineTable
