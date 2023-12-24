// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// gosymパッケージは、gcコンパイラによって生成されたGoバイナリに埋め込まれた
// Goのシンボルと行番号のテーブルへのアクセスを実装します。
package gosym

// Symは、単一のシンボルテーブルエントリを表します。
type Sym struct {
	Value  uint64
	Type   byte
	Name   string
	GoType uint64
	// このシンボルが関数シンボルである場合、対応するFunc
	Func *Func

	goVersion version
}

// Staticは、このシンボルが静的（ファイルの外部からは見えない）であるかどうかを報告します。
func (s *Sym) Static() bool

// PackageNameは、シンボル名のパッケージ部分を返します。
// パッケージ部分がない場合は空の文字列を返します。
func (s *Sym) PackageName() string

// ReceiverNameは、このシンボルのレシーバタイプ名を返します。
// レシーバ名がない場合は空の文字列を返します。レシーバ名は、
// s.Nameがパッケージ名で完全に指定されている場合にのみ検出されます。
func (s *Sym) ReceiverName() string

// BaseNameは、パッケージ名やレシーバ名を除いたシンボル名を返します。
func (s *Sym) BaseName() string

// Funcは、単一の関数に関する情報を収集します。
type Func struct {
	Entry uint64
	*Sym
	End       uint64
	Params    []*Sym
	Locals    []*Sym
	FrameSize int
	LineTable *LineTable
	Obj       *Obj
}

// Objは、シンボルテーブル内の一連の関数を表します。
//
// バイナリを別々のObjに分割する具体的な方法は、シンボルテーブル形式の内部詳細です。
//
// Goの初期のバージョンでは、各ソースファイルが異なるObjになりました。
//
// Go 1とGo 1.1では、各パッケージはすべてのGoソースに対して1つのObjを生成し、
// Cソースファイルごとに1つのObjを生成しました。
//
// Go 1.2では、プログラム全体に対して単一のObjが存在します。
type Obj struct {
	// Funcsは、Obj内の関数のリストです。
	Funcs []Func

	// Go 1.1以前では、PathsはObjを生成したソースファイル名に対応するシンボルのリストです。
	// Go 1.2では、Pathsはnilです。
	// ソースファイルのリストを取得するには、Table.Filesのキーを使用します。
	Paths []Sym
}

// TableはGoのシンボルテーブルを表します。プログラムからデコードされたすべての
// シンボルを保存し、シンボル、名前、アドレス間の変換を行うメソッドを提供します。
type Table struct {
	Syms  []Sym
	Funcs []Func
	Files map[string]*Obj
	Objs  []Obj

	go12line *LineTable
}

// NewTableはGoのシンボルテーブル（ELFの".gosymtab"セクション）をデコードし、
// メモリ内表現を返します。
// Go 1.3以降、Goのシンボルテーブルにはシンボルデータが含まれなくなりました。
func NewTable(symtab []byte, pcln *LineTable) (*Table, error)

// PCToFuncは、プログラムカウンタpcを含む関数を返します。
// そのような関数がない場合はnilを返します。
func (t *Table) PCToFunc(pc uint64) *Func

// PCToLineは、プログラムカウンタに対する行番号情報を検索します。
// 情報がない場合は、fn == nilを返します。
func (t *Table) PCToLine(pc uint64) (file string, line int, fn *Func)

<<<<<<< HEAD
// LineToPC looks up the first program counter on the given line in
// the named file. It returns [UnknownFileError] or [UnknownLineError] if
// there is an error looking up this line.
=======
// LineToPCは、指定されたファイルの指定された行で最初のプログラムカウンタを検索します。
// この行を検索中にエラーが発生した場合、UnknownPathErrorまたはUnknownLineErrorを返します。
>>>>>>> release-branch.go1.21
func (t *Table) LineToPC(file string, line int) (pc uint64, fn *Func, err error)

// LookupSymは、指定された名前を持つテキスト、データ、またはbssシンボルを返します。
// そのようなシンボルが見つからない場合はnilを返します。
func (t *Table) LookupSym(name string) *Sym

// LookupFuncは、指定された名前を持つテキスト、データ、またはbssシンボルを返します。
// そのようなシンボルが見つからない場合はnilを返します。
func (t *Table) LookupFunc(name string) *Func

// SymByAddrは、指定されたアドレスで開始するテキスト、データ、またはbssシンボルを返します。
func (t *Table) SymByAddr(addr uint64) *Sym

// UnknownFileErrorは、シンボルテーブル内で特定のファイルを見つけることができなかったことを表すエラーです。
type UnknownFileError string

func (e UnknownFileError) Error() string

// UnknownLineErrorは、行をプログラムカウンタにマッピングできなかったことを表すエラーです。
// これは、行がファイルの範囲を超えているか、指定された行にコードがないためです。
type UnknownLineError struct {
	File string
	Line int
}

func (e *UnknownLineError) Error() string

// DecodingError represents an error during the decoding of
// the symbol table.
type DecodingError struct {
	off int
	msg string
	val any
}

func (e *DecodingError) Error() string
