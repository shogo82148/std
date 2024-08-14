// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// csvパッケージは、カンマ区切り値(CSV)ファイルの読み書きを行います。
// 多くの種類のCSVファイルがありますが、このパッケージはRFC 4180で説明されている形式をサポートしています。
// ただし、[Writer] はデフォルトで改行文字としてCRLFではなくLFを使用します。
//
// CSVファイルには、レコードごとに1つ以上のフィールドを含むゼロ以上のレコードが含まれています。
// 各レコードは改行文字で区切られます。最後のレコードはオプションで改行文字に続くことができます。
//
//	field1,field2,field3
//
// 空白はフィールドの一部と見なされます。
//
// 改行文字の前のキャリッジリターンは、静かに削除されます。
//
// 空行は無視されます。空白文字のみで構成される行（末尾の改行文字を除く）は、空行と見なされません。
//
// クォート文字 "で始まり、終わるフィールドは、クォートフィールドと呼ばれます。
// 開始と終了の引用符はフィールドの一部ではありません。
//
// ソース：
//
//	normal string,"quoted-field"
//
// は、次のフィールドを生成します。
//
//	{`normal string`, `quoted-field`}
//
// クォートフィールド内の引用符の後に2番目の引用符が続く場合、
// 1つの引用符として扱われます。
//
//	"the ""word"" is true","a ""quoted-field"""
//
// の結果は次のとおりです。
//
//	{`the "word" is true`, `a "quoted-field"`}
//
// 改行とカンマは、クォートフィールド内に含めることができます。
//
//	"Multi-line
//	field","comma is ,"
//
// の結果は次のとおりです。
//
//	{`Multi-line
//	field`, `comma is ,`}
package csv

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// ParseErrorは、解析エラーの場合に返されます。
// 行番号と行番号は1から始まります。
type ParseError struct {
	StartLine int
	Line      int
	Column    int
	Err       error
}

func (e *ParseError) Error() string

func (e *ParseError) Unwrap() error

// [ParseError.Err] で返される可能性のあるエラーです。
var (
	ErrBareQuote  = errors.New("bare \" in non-quoted-field")
	ErrQuote      = errors.New("extraneous or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")

	// Deprecated: ErrTrailingComma はもう使用されません。
	ErrTrailingComma = errors.New("extra delimiter at end of line")
)

// Readerは、CSVエンコードされたファイルからレコードを読み取ります。
//
// [NewReader] によって返された場合、ReaderはRFC 4180に準拠した入力を想定しています。
// 最初の [Reader.Read] または [Reader.ReadAll] 呼び出しの前に、エクスポートされたフィールドを変更して詳細をカスタマイズできます。
//
// Readerは、入力のすべての\r\nシーケンスをプレーンな\nに変換するため、
// 複数行のフィールド値を含む場合でも、返されるデータが入力ファイルが使用する行末の規約に依存しないようにします。
type Reader struct {
	// Commaはフィールドの区切り文字です。
	// NewReaderによってカンマ（','）に設定されます。
	// Commaは有効なルーンである必要があり、\r、\n、
	// またはUnicode置換文字（0xFFFD）であってはなりません。
	Comma rune

	// Commentが0でない場合、Comment文字はコメント文字です。
	// 先行する空白がないComment文字で始まる行は無視されます。
	// 先行する空白がある場合、TrimLeadingSpaceがtrueであっても、Comment文字はフィールドの一部になります。
	// Commentは有効なルーンである必要があり、\r、\n、
	// またはUnicode置換文字（0xFFFD）であってはなりません。
	// また、Commaと等しくてはなりません。
	Comment rune

	// FieldsPerRecordは、レコードごとに期待されるフィールド数です。
	// FieldsPerRecordが正の場合、Readは各レコードが指定されたフィールド数を持つことを要求します。
	// FieldsPerRecordが0の場合、Readは最初のレコードのフィールド数に設定し、
	// 以降のレコードは同じフィールド数を持つ必要があります。
	// FieldsPerRecordが負の場合、チェックは行われず、レコードは可変長のフィールド数を持つ場合があります。
	FieldsPerRecord int

	// LazyQuotesがtrueの場合、引用符は引用符で囲まれていないフィールドに表示される場合があります。
	LazyQuotes bool

	// TrimLeadingSpaceがtrueの場合、フィールドの先頭の空白は無視されます。
	// これは、フィールド区切り文字であるCommaが空白である場合でも行われます。
	TrimLeadingSpace bool

	// ReuseRecordは、パフォーマンスのために、Readの呼び出しが前回の呼び出しの返されたスライスのバッキング配列を共有するスライスを返すかどうかを制御します。
	// デフォルトでは、Readの各呼び出しは、呼び出し元が所有する新しく割り当てられたメモリを返します。
	ReuseRecord bool

	// Deprecated: TrailingComma はもう使用されません。
	TrailingComma bool

	r *bufio.Reader

	// numLine is the current line being read in the CSV file.
	numLine int

	// offset is the input stream byte offset of the current reader position.
	offset int64

	// rawBuffer is a line buffer only used by the readLine method.
	rawBuffer []byte

	// recordBuffer holds the unescaped fields, one after another.
	// The fields can be accessed by using the indexes in fieldIndexes.
	// E.g., For the row `a,"b","c""d",e`, recordBuffer will contain `abc"de`
	// and fieldIndexes will contain the indexes [1, 2, 5, 6].
	recordBuffer []byte

	// fieldIndexes is an index of fields inside recordBuffer.
	// The i'th field ends at offset fieldIndexes[i] in recordBuffer.
	fieldIndexes []int

	// fieldPositions is an index of field positions for the
	// last record returned by Read.
	fieldPositions []position

	// lastRecord is a record cache and only used when ReuseRecord == true.
	lastRecord []string
}

// NewReaderは、rから読み取る新しいReaderを返します。
func NewReader(r io.Reader) *Reader

// Readはrから1つのレコード（フィールドのスライス）を読み込みます。
// レコードに予期しない数のフィールドが含まれている場合、
// Readはエラー [ErrFieldCount] とともにレコードを返します。
// パースできないフィールドが含まれている場合、
// Readは部分的なレコードとパースエラーを返します。
// 部分的なレコードには、エラーが発生する前に読み取られたすべてのフィールドが含まれます。
// 読み取るデータがない場合、Readはnil、io.EOFを返します。
// [Reader.ReuseRecord] がtrueの場合、返されるスライスは複数のRead呼び出し間で共有できます。
func (r *Reader) Read() (record []string, err error)

// FieldPosは、Readで最後に返されたスライス内の指定されたインデックスのフィールドの開始に対応する行と列を返します。
// 行と列の番号付けは1から始まります。列はルーンではなくバイトで数えられます。
//
// インデックスが範囲外で呼び出された場合、panicします。
func (r *Reader) FieldPos(field int) (line, column int)

// InputOffsetは、現在のリーダーの位置の入力ストリームバイトオフセットを返します。
// オフセットは、最後に読み取られた行の終わりと次の行の始まりの場所を示します。
func (r *Reader) InputOffset() int64

// ReadAllは、rから残りのすべてのレコードを読み込みます。
// 各レコードはフィールドのスライスです。
// 成功した呼び出しはerr == nilを返します。err == [io.EOF] ではありません。
// ReadAllはEOFまで読み込むように定義されているため、エラーとして扱いません。
func (r *Reader) ReadAll() (records [][]string, err error)
