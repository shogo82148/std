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

<<<<<<< HEAD
// Posはファイルセット内のソース位置のコンパクトなエンコーディングです。
// より便利ながらもはるかに大きい表現のために、Positionに変換することができます。
=======
// Pos is a compact encoding of a source position within a file set.
// It can be converted into a [Position] for a more convenient, but much
// larger, representation.
>>>>>>> upstream/master
//
// 与えられたファイルに対するPosの値は、[base、base+size]の範囲内の数値です。
// baseとsizeは、ファイルがファイルセットに追加される際に指定されます。
// Posの値と対応するファイルベースの差は、その位置（Posの値によって表される）からのバイトオフセットに対応します。
// したがって、ファイルベースオフセットは、ファイル内の最初のバイトを表すPosの値です。
//
<<<<<<< HEAD
// 特定のソースオフセット（バイト単位で測定される）のPos値を作成するには、
// まずFileSet.AddFileを使用して対応するファイルを現在のファイルセットに追加し、
// そのファイルのためにFile.Pos(offset)を呼び出します。
// 特定のファイルセットfsetに対するPos値pを持つ場合、対応するPosition値はfset.Position(p)を呼び出すことで得られます。
=======
// To create the Pos value for a specific source offset (measured in bytes),
// first add the respective file to the current file set using [FileSet.AddFile]
// and then call [File.Pos](offset) for that file. Given a Pos value p
// for a specific file set fset, the corresponding [Position] value is
// obtained by calling fset.Position(p).
>>>>>>> upstream/master
//
// Pos値は通常の比較演算子を使って直接比較することができます：
// 2つのPos値pとqが同じファイルにある場合、pとqを比較することは、対応するソースファイルのオフセットを比較することと等価です。
// pとqが異なるファイルにある場合、qによって指定されるファイルがpによって指定されるファイルよりも前に対応するファイルセットに追加された場合、p < qはtrueです。
type Pos int

<<<<<<< HEAD
// Posのゼロ値はNoPosであり、ファイルおよび行情報は関連付けられていません。
// また、NoPos.IsValid()はfalseです。NoPosは常に他のPos値よりも小さくなります。
// NoPosに対応するPosition値はPositionのゼロ値です。
=======
// The zero value for [Pos] is NoPos; there is no file and line information
// associated with it, and NoPos.IsValid() is false. NoPos is always
// smaller than any other [Pos] value. The corresponding [Position] value
// for NoPos is the zero value for [Position].
>>>>>>> upstream/master
const NoPos Pos = 0

// IsValid は位置が有効かどうかを報告します。
func (p Pos) IsValid() bool

<<<<<<< HEAD
// Fileは、FileSetに属するファイルのハンドルです。
// Fileには名前、サイズ、行オフセット表があります。
=======
// A File is a handle for a file belonging to a [FileSet].
// A File has a name, size, and line offset table.
>>>>>>> upstream/master
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

<<<<<<< HEAD
// MergeLineは、次の行と行を結合します。これは、行の末尾の改行文字をスペースで置き換えることに似ています（残りのオフセットは変更されません）。行番号を取得するには、Position.Lineなどを参照してください。無効な行番号が指定された場合、MergeLineはパニックを起こします。
func (f *File) MergeLine(line int)

// LinesはSetLinesで指定された形式の効果的な行オフセットの表を返します。
// 呼び出し元は結果を変更してはいけません。
=======
// MergeLine merges a line with the following line. It is akin to replacing
// the newline character at the end of the line with a space (to not change the
// remaining offsets). To obtain the line number, consult e.g. [Position.Line].
// MergeLine will panic if given an invalid line number.
func (f *File) MergeLine(line int)

// Lines returns the effective line offset table of the form described by [File.SetLines].
// Callers must not mutate the result.
>>>>>>> upstream/master
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

<<<<<<< HEAD
// LineStartは指定された行の開始位置のPos値を返します。
// AddLineColumnInfoを使用して設定された代替の位置は無視されます。
// LineStartは、1ベースの行番号が無効な場合にパニックを引き起こします。
func (f *File) LineStart(line int) Pos

// AddLineInfoは、Column = 1引数を持つAddLineColumnInfoと同様です。
// Go 1.11より前のコードの後方互換性のためにここにあります。
=======
// LineStart returns the [Pos] value of the start of the specified line.
// It ignores any alternative positions set using [File.AddLineColumnInfo].
// LineStart panics if the 1-based line number is invalid.
func (f *File) LineStart(line int) Pos

// AddLineInfo is like [File.AddLineColumnInfo] with a column = 1 argument.
// It is here for backward-compatibility for code prior to Go 1.11.
>>>>>>> upstream/master
func (f *File) AddLineInfo(offset int, filename string, line int)

// AddLineColumnInfoは、与えられたファイルオフセットに対して代替のファイル、行、および列番号の情報を追加します。オフセットは、以前に追加された代替の行情報のオフセットよりも大きく、ファイルサイズよりも小さい必要があります。それ以外の場合、情報は無視されます。
// AddLineColumnInfoは通常、//line filename:line:columnなどの行ディレクティブの代替位置情報を登録するために使用されます。
func (f *File) AddLineColumnInfo(offset int, filename string, line, column int)

// Posは与えられたファイルオフセットのPos値を返します。
// オフセットはf.Size()以下でなければなりません。
// f.Pos(f.Offset(p)) == p。
func (f *File) Pos(offset int) Pos

<<<<<<< HEAD
// Offsetは与えられたファイル位置pのオフセットを返します。
// pはそのファイル内で有効なPosの値でなければなりません。
// f.Offset(f.Pos(offset)) == offset。
func (f *File) Offset(p Pos) int

// Lineは与えられたファイル位置pの行番号を返します。
// pはそのファイル内のPos値またはNoPosでなければなりません。
=======
// Offset returns the offset for the given file position p;
// p must be a valid [Pos] value in that file.
// f.Offset(f.Pos(offset)) == offset.
func (f *File) Offset(p Pos) int

// Line returns the line number for the given file position p;
// p must be a [Pos] value in that file or [NoPos].
>>>>>>> upstream/master
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
<<<<<<< HEAD
// ファイルセット内の各ファイルのバイトオフセットは、異なる（整数）間隔、すなわち間隔[base、base+size]にマッピングされます。 baseはファイルの最初のバイトを表し、sizeは対応するファイルサイズです。 Pos値はそのような間隔内の値です。 Pos値が属する間隔を決定することで、ファイル、そのファイルのベース、そしてPos値が表しているバイトオフセット（位置）を計算することができます。
//
// 新しいファイルを追加する際には、ファイルベースが必要です。それはすでにファイルセット内の任意のファイルの間隔の終わりを過ぎた整数値である必要があります。便宜上、FileSet.Baseはそのような値を提供します。それは単純に最後に追加されたファイルのPos間隔の終わりの位置に+1した値です。後で間隔を拡張する必要がない場合は、FileSet.BaseをFileSet.AddFileの引数として使用する必要があります。
//
// FileSetが不要な場合、FileはFileSetから削除することができます。これにより、長時間実行されるアプリケーションでメモリ使用量を削減することができます。
=======
// The byte offsets for each file in a file set are mapped into
// distinct (integer) intervals, one interval [base, base+size]
// per file. [FileSet.Base] represents the first byte in the file, and size
// is the corresponding file size. A [Pos] value is a value in such
// an interval. By determining the interval a [Pos] value belongs
// to, the file, its file base, and thus the byte offset (position)
// the [Pos] value is representing can be computed.
//
// When adding a new file, a file base must be provided. That can
// be any integer value that is past the end of any interval of any
// file already in the file set. For convenience, [FileSet.Base] provides
// such a value, which is simply the end of the Pos interval of the most
// recently added file, plus one. Unless there is a need to extend an
// interval later, using the [FileSet.Base] should be used as argument
// for [FileSet.AddFile].
//
// A [File] may be removed from a FileSet when it is no longer needed.
// This may reduce memory usage in a long-running application.
>>>>>>> upstream/master
type FileSet struct {
	mutex sync.RWMutex
	base  int
	files []*File
	last  atomic.Pointer[File]
}

// NewFileSetは新しいファイルセットを作成します。
func NewFileSet() *FileSet

<<<<<<< HEAD
// Baseは、次のファイルを追加する際にAddFileに提供する必要がある最小のベースオフセットを返します。
func (s *FileSet) Base() int

// AddFileは、指定されたファイル名、ベースオフセット、ファイルサイズを持つ新しいファイルをファイルセットsに追加し、ファイルを返します。複数のファイルは同じ名前を持つことができます。ベースオフセットは、FileSetのBase()より小さくはならず、サイズは負であってはいけません。特別なケースとして、負のベースが提供された場合、FileSetのBase()の現在の値が代わりに使用されます。
// ファイルを追加すると、次のファイルのための最小ベース値として、ファイルセットのBase()の値はbase + size + 1に設定されます。与えられたファイルオフセットoffsに対するPos値pの関係は次のとおりです：
// int(p) = base + offs
// ただし、offsは範囲[0、size]にあり、したがってpは範囲[base、base+size]にあります。便宜上、File.Posはファイル固有の位置値をファイルオフセットから作成するために使用できます。
func (s *FileSet) AddFile(filename string, base, size int) *File

// RemoveFileは、FileSetからファイルを削除し、その後のPos間隔のクエリが負の結果を返すようにします。
// これにより、長寿命のFileSetのメモリ使用量が減少し、無制限のファイルストリームに遭遇した場合でも処理が可能になります。
=======
// Base returns the minimum base offset that must be provided to
// [FileSet.AddFile] when adding the next file.
func (s *FileSet) Base() int

// AddFile adds a new file with a given filename, base offset, and file size
// to the file set s and returns the file. Multiple files may have the same
// name. The base offset must not be smaller than the [FileSet.Base], and
// size must not be negative. As a special case, if a negative base is provided,
// the current value of the [FileSet.Base] is used instead.
//
// Adding the file will set the file set's [FileSet.Base] value to base + size + 1
// as the minimum base value for the next file. The following relationship
// exists between a [Pos] value p for a given file offset offs:
//
//	int(p) = base + offs
//
// with offs in the range [0, size] and thus p in the range [base, base+size].
// For convenience, [File.Pos] may be used to create file-specific position
// values from a file offset.
func (s *FileSet) AddFile(filename string, base, size int) *File

// RemoveFile removes a file from the [FileSet] so that subsequent
// queries for its [Pos] interval yield a negative result.
// This reduces the memory usage of a long-lived [FileSet] that
// encounters an unbounded stream of files.
>>>>>>> upstream/master
//
// セットに属さないファイルを削除しても効果はありません。
func (s *FileSet) RemoveFile(file *File)

// ファイルセット内のファイルを追加された順にfに呼び出し、fがfalseを返すまで繰り返します。
func (s *FileSet) Iterate(f func(*File) bool)

<<<<<<< HEAD
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
=======
// File returns the file that contains the position p.
// If no such file is found (for instance for p == [NoPos]),
// the result is nil.
func (s *FileSet) File(p Pos) (f *File)

// PositionFor converts a [Pos] p in the fileset into a [Position] value.
// If adjusted is set, the position may be adjusted by position-altering
// //line comments; otherwise those comments are ignored.
// p must be a [Pos] value in s or [NoPos].
func (s *FileSet) PositionFor(p Pos, adjusted bool) (pos Position)

// Position converts a [Pos] p in the fileset into a Position value.
// Calling s.Position(p) is equivalent to calling s.PositionFor(p, true).
>>>>>>> upstream/master
func (s *FileSet) Position(p Pos) (pos Position)
