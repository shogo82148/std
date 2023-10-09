// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

import (
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
)

// Positionはファイル、行、列の位置を含む任意のソース位置を表します。
// Positionは行番号が> 0の場合に有効です。
type Position struct {
	Filename string
	Offset   int
	Line     int
	Column   int
}

// IsValid は位置が有効かどうかを報告します。
func (pos *Position) IsValid() bool

// Stringはいくつかの形式で文字列を返します：
//
//	file:line:column    ファイル名を含む有効な位置
//	file:line           ファイル名を含む有効な位置だが列がない（column == 0）
//	line:column         ファイル名を含まない有効な位置
//	line                ファイル名もなく列もない有効な位置（column == 0）
//	file                ファイル名を含む無効な位置
//	-                   ファイル名もない無効な位置
func (pos Position) String() string

// Posはファイルセット内のソース位置のコンパクトなエンコーディングです。
// より便利ながらもはるかに大きい表現のために、Positionに変換することができます。
//
// 与えられたファイルに対するPosの値は、[base、base+size]の範囲内の数値です。
// baseとsizeは、ファイルがファイルセットに追加される際に指定されます。
// Posの値と対応するファイルベースの差は、その位置（Posの値によって表される）からのバイトオフセットに対応します。
// したがって、ファイルベースオフセットは、ファイル内の最初のバイトを表すPosの値です。
//
// 特定のソースオフセット（バイト単位で測定される）のPos値を作成するには、
// まずFileSet.AddFileを使用して対応するファイルを現在のファイルセットに追加し、
// そのファイルのためにFile.Pos(offset)を呼び出します。
// 特定のファイルセットfsetに対するPos値pを持つ場合、対応するPosition値はfset.Position(p)を呼び出すことで得られます。
//
// Pos値は通常の比較演算子を使って直接比較することができます：
// 2つのPos値pとqが同じファイルにある場合、pとqを比較することは、対応するソースファイルのオフセットを比較することと等価です。
// pとqが異なるファイルにある場合、qによって指定されるファイルがpによって指定されるファイルよりも前に対応するファイルセットに追加された場合、p < qはtrueです。
type Pos int

// Posのゼロ値はNoPosであり、ファイルおよび行情報は関連付けられていません。
// また、NoPos.IsValid()はfalseです。NoPosは常に他のPos値よりも小さくなります。
// NoPosに対応するPosition値はPositionのゼロ値です。
const NoPos Pos = 0

// IsValid は位置が有効かどうかを報告します。
func (p Pos) IsValid() bool

// Fileは、FileSetに属するファイルのハンドルです。
// Fileには名前、サイズ、行オフセット表があります。
type File struct {
	name string
	base int
	size int

	// lines と infos はミューテックスによって保護されています。
	mutex sync.Mutex
	lines []int
	infos []lineInfo
}

// NameはAddFileで登録されたファイルfのファイル名を返します。
func (f *File) Name() string

// Baseは、AddFileで登録されたファイルfの基本オフセットを返します。
func (f *File) Base() int

// SizeはAddFileで登録されたファイルfのサイズを返します。
func (f *File) Size() int

// LineCount returns the number of lines in file f.
func (f *File) LineCount() int

// AddLineは新しい行の行オフセットを追加します。
// 行オフセットは前の行のオフセットよりも大きく、ファイルのサイズよりも小さい必要があります。そうでない場合、行オフセットは無視されます。
func (f *File) AddLine(offset int)

// MergeLineは、次の行と行を結合します。これは、行の末尾の改行文字をスペースで置き換えることに似ています（残りのオフセットは変更されません）。行番号を取得するには、Position.Lineなどを参照してください。無効な行番号が指定された場合、MergeLineはパニックを起こします。
func (f *File) MergeLine(line int)

// LinesはSetLinesで指定された形式の効果的な行オフセットの表を返します。
// 呼び出し元は結果を変更してはいけません。
func (f *File) Lines() []int

// SetLinesはファイルの行オフセットを設定し、成功したかどうかを報告します。
// 行オフセットとは、各行の最初の文字のオフセットです。
// たとえば、"ab\nc\n"という内容の場合、行オフセットは{0、3}です。
// 空のファイルは空の行オフセットテーブルを持ちます。
// 各行のオフセットは、前の行のオフセットよりも大きく、ファイルサイズよりも小さくなければなりません。
// それ以外の場合、SetLinesは失敗し、falseを返します。
// SetLinesが返された後は、与えられたスライスを変更しないでください。
func (f *File) SetLines(lines []int) bool

// SetLinesForContentは与えられたファイルの内容に対して行のオフセットを設定します。
// 位置を変更する//lineコメントは無視されます。
func (f *File) SetLinesForContent(content []byte)

// LineStartは指定された行の開始位置のPos値を返します。
// AddLineColumnInfoを使用して設定された代替の位置は無視されます。
// LineStartは、1ベースの行番号が無効な場合にパニックを引き起こします。
func (f *File) LineStart(line int) Pos

// AddLineInfoは、Column = 1引数を持つAddLineColumnInfoと同様です。
// Go 1.11より前のコードの後方互換性のためにここにあります。
func (f *File) AddLineInfo(offset int, filename string, line int)

// AddLineColumnInfoは、与えられたファイルオフセットに対して代替のファイル、行、および列番号の情報を追加します。オフセットは、以前に追加された代替の行情報のオフセットよりも大きく、ファイルサイズよりも小さい必要があります。それ以外の場合、情報は無視されます。
// AddLineColumnInfoは通常、//line filename:line:columnなどの行ディレクティブの代替位置情報を登録するために使用されます。
func (f *File) AddLineColumnInfo(offset int, filename string, line, column int)

// Posは与えられたファイルオフセットのPos値を返します。
// オフセットはf.Size()以下でなければなりません。
// f.Pos(f.Offset(p)) == p。
func (f *File) Pos(offset int) Pos

// Offsetは与えられたファイル位置pのオフセットを返します。
// pはそのファイル内で有効なPosの値でなければなりません。
// f.Offset(f.Pos(offset)) == offset。
func (f *File) Offset(p Pos) int

// Lineは与えられたファイル位置pの行番号を返します。
// pはそのファイル内のPos値またはNoPosでなければなりません。
func (f *File) Line(p Pos) int

// PositionForは、指定されたファイルの位置pに対するPositionの値を返します。
// 位置を変更する可能性のある行コメントが設定されている場合、位置は調整されるかもしれません。そうでない場合は、コメントは無視されます。
// pは、fまたはNoPosのPos値である必要があります。
func (f *File) PositionFor(p Pos, adjusted bool) (pos Position)

// Positionは指定されたファイルの位置pに対するPositionの値を返します。
// f.Position(p)を呼び出すことは、f.PositionFor(p, true)を呼び出すことと等価です。
func (f *File) Position(p Pos) (pos Position)

// FileSetはソースファイルの集合を表します。
// ファイルセットのメソッドは同期されており、複数のゴルーチンが同時に呼び出すことができます。
//
// ファイルセット内の各ファイルのバイトオフセットは、異なる（整数）間隔、すなわち間隔[base、base+size]にマッピングされます。 baseはファイルの最初のバイトを表し、sizeは対応するファイルサイズです。 Pos値はそのような間隔内の値です。 Pos値が属する間隔を決定することで、ファイル、そのファイルのベース、そしてPos値が表しているバイトオフセット（位置）を計算することができます。
//
// 新しいファイルを追加する際には、ファイルベースが必要です。それはすでにファイルセット内の任意のファイルの間隔の終わりを過ぎた整数値である必要があります。便宜上、FileSet.Baseはそのような値を提供します。それは単純に最後に追加されたファイルのPos間隔の終わりの位置に+1した値です。後で間隔を拡張する必要がない場合は、FileSet.BaseをFileSet.AddFileの引数として使用する必要があります。
//
// FileSetが不要な場合、FileはFileSetから削除することができます。これにより、長時間実行されるアプリケーションでメモリ使用量を削減することができます。
type FileSet struct {
	mutex sync.RWMutex
	base  int
	files []*File
	last  atomic.Pointer[File]
}

// NewFileSetは新しいファイルセットを作成します。
func NewFileSet() *FileSet

// Baseは、次のファイルを追加する際にAddFileに提供する必要がある最小のベースオフセットを返します。
func (s *FileSet) Base() int

// AddFileは、指定されたファイル名、ベースオフセット、ファイルサイズを持つ新しいファイルをファイルセットsに追加し、ファイルを返します。複数のファイルは同じ名前を持つことができます。ベースオフセットは、FileSetのBase()より小さくはならず、サイズは負であってはいけません。特別なケースとして、負のベースが提供された場合、FileSetのBase()の現在の値が代わりに使用されます。
// ファイルを追加すると、次のファイルのための最小ベース値として、ファイルセットのBase()の値はbase + size + 1に設定されます。与えられたファイルオフセットoffsに対するPos値pの関係は次のとおりです：
// int(p) = base + offs
// ただし、offsは範囲[0、size]にあり、したがってpは範囲[base、base+size]にあります。便宜上、File.Posはファイル固有の位置値をファイルオフセットから作成するために使用できます。
func (s *FileSet) AddFile(filename string, base, size int) *File

// RemoveFileは、FileSetからファイルを削除し、その後のPos間隔のクエリが負の結果を返すようにします。
// これにより、長寿命のFileSetのメモリ使用量が減少し、無制限のファイルストリームに遭遇した場合でも処理が可能になります。
//
// セットに属さないファイルを削除しても効果はありません。
func (s *FileSet) RemoveFile(file *File)

// ファイルセット内のファイルを追加された順にfに呼び出し、fがfalseを返すまで繰り返します。
func (s *FileSet) Iterate(f func(*File) bool)

// File関数は、位置pを含むファイルを返します。
// 該当するファイルが見つからない場合（たとえばp == NoPosの場合）、結果はnilです。
func (s *FileSet) File(p Pos) (f *File)

// PositionForは、ファイルセット内の位置pをPosition値に変換します。
// adjustedが設定されている場合、位置は位置変更を行うコメントによって調整される可能性があります。
// そうでなければ、そのコメントは無視されます。
// pはsまたはNoPosのPos値でなければなりません。
func (s *FileSet) PositionFor(p Pos, adjusted bool) (pos Position)

// Positionはファイルセット内のPos pをPosition値に変換します。
// s.Position(p)を呼び出すことは、s.PositionFor(p, true)を呼び出すことと同じです。
func (s *FileSet) Position(p Pos) (pos Position)
