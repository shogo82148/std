// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package io は基本的なI/Oプリミティブに対するインターフェースを提供します。
// その主な役割は、パッケージosなどの既存の実装を共有の公開インターフェースにラップすることで、
// 機能を抽象化し、関連するプリミティブを提供することです。
//
// これらのインターフェースとプリミティブは、低レベルの操作をさまざまな実装でラップするため、
// 明示的に指示がない限り、クライアントは並行実行には安全ではないと想定すべきではありません。
package io

import (
	"github.com/shogo82148/std/errors"
)

// 値の発見を求める。
const (
	SeekStart   = 0
	SeekCurrent = 1
	SeekEnd     = 2
)

// ErrShortWriteは、要求されたバイト数よりも少ないバイト数の書き込みを受け付けたが、
// 明示的なエラーを返すことができなかったことを意味します。
var ErrShortWrite = errors.New("short write")

// ErrShortBuffer は、読み込みに提供されたバッファよりも長いバッファが必要だったことを意味します。
var ErrShortBuffer = errors.New("short buffer")

// EOFは、入力がもう利用できない場合にReadが返すエラーです。
// （EOF自体ではなく、EOFをラップしたエラーを返すのではなく、
// EOFをテストするために呼び出し元で==を使用するためです。）
// 関数は、入力の優雅な終了を示すために、EOFのみを返すべきです。
// もしEOFが構造化されたデータストリームで予期しない場所で発生した場合、
// 適切なエラーはErrUnexpectedEOFまたはその他の詳細を示すエラーです。
var EOF = errors.New("EOF")

// ErrUnexpectedEOFは、固定サイズのブロックまたはデータ構造の読み取り途中にEOFが出現したことを意味します。
var ErrUnexpectedEOF = errors.New("unexpected EOF")

// ErrNoProgressは、一部のReaderのクライアントがデータやエラーを返さずに、
// 複数回のRead呼び出しが失敗した場合に返されます。
// 通常は、破損したReaderの実装を示しています。
var ErrNoProgress = errors.New("multiple Read calls return no data or error")

// Readerは基本的なReadメソッドをラップするインターフェースです。
// Readは最大でlen(p)バイトをpに読み込みます。読み込んだバイト数（0 <= n <= len(p)）と遭遇したエラーを返します。Readがn < len(p)を返しても、呼び出し中にp全体を作業領域として使用する場合があります。データが利用可能であるがlen(p)バイトでない場合、Readは通常、追加のデータを待つ代わりに利用可能なデータを返します。
// Readがn > 0バイトの読み込みに成功した後にエラーやファイルの終わりの状態に遭遇すると、読み込んだバイト数を返します。同じ呼び出しから（nilではない）エラーを返すことも、次の呼び出しからエラー（およびn == 0）を返すこともあります。入力ストリームの終わりで非ゼロのバイト数を返すReaderのインスタンスは、err == EOFまたはerr == nilのいずれかを返す場合があります。次のReadは0, EOFを返すべきです。
// 呼び出し元は、n > 0バイトが返される前にエラーerrを処理する必要があります。これにより、いくつかのバイトを読んだ後に発生するI/Oエラーや、許可されるEOFの両方の動作が正しく処理されます。
// len(p) == 0の場合、Readは常にn == 0を返すべきです。EOFなどのエラー条件が既知の場合、非nilのエラーを返す場合があります。
// Readの実装は、len(p) == 0の場合を除き、0バイトカウントとnilエラーを返すことは避けるべきです。呼び出し元は、0とnilの返り値は何も発生しなかったことを示しており、特にEOFを示しているわけではないと扱うべきです。
// 実装はpを保持してはいけません。
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writerは基本的なWriteメソッドを包むインターフェースです。
//
// Writeはpから基礎となるデータストリームにlen(p)バイトを書き込みます。
// それはpから書き込まれたバイト数（0 <= n <= len(p)）と、
// 書き込みが早期に停止させたエラーを返します。
// Writeは、n < len(p)の場合には非nilのエラーを返さなければなりません。
// Writeは、スライスデータを一時的にでも変更してはいけません。
//
// 実装はpを保持してはいけません。
type Writer interface {
	Write(p []byte) (n int, err error)
}

// Closerは基本のCloseメソッドをラップするインターフェースです。
//
// 最初の呼び出し後のCloseの振る舞いは未定義です。
// 特定の実装は独自の振る舞いを文書化することがあります。
type Closer interface {
	Close() error
}

// Seekerは基本のSeekメソッドをラップするインターフェースです。
// Seekは、オフセットを次のReadまたはWriteのために設定します。
// whenceに従って解釈されます。
// SeekStartはファイルの先頭を基準とします。
// SeekCurrentは現在のオフセットを基準とします。
// SeekEndは末尾を基準とします。
// (例えば、offset = -2はファイルの最後から1つ前のバイトを指定します)。
// Seekは、ファイルの先頭を基準とした新たなオフセットまたはエラーを返します。
// ファイルの先頭より前のオフセットにシークすることはエラーです。
// 任意の正のオフセットにシークすることは許可されるかもしれませんが、
// 新しいオフセットが基になるオブジェクトのサイズを超える場合、
// その後のI/O操作の振る舞いは実装に依存します。
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

// ReadWriterは基本的なReadとWriteメソッドをグループ化するインターフェースです。
type ReadWriter interface {
	Reader
	Writer
}

// ReadCloserは基本的なReadとCloseメソッドをまとめるインターフェースです。
type ReadCloser interface {
	Reader
	Closer
}

// WriteCloserは基本的なWriteとCloseメソッドをグループ化するインターフェースです。
type WriteCloser interface {
	Writer
	Closer
}

// ReadWriteCloserは、基本的なRead、Write、Closeメソッドをグループ化するインタフェースです。
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// ReadSeekerは、基本的なReadとSeekメソッドをグループ化するインターフェースです。
type ReadSeeker interface {
	Reader
	Seeker
}

// ReadSeekCloserは、基本的なRead、Seek、Closeメソッドをグループ化するインターフェースです。
type ReadSeekCloser interface {
	Reader
	Seeker
	Closer
}

// WriteSeekerは基本的なWriteとSeekメソッドをグループ化するインターフェースです。
type WriteSeeker interface {
	Writer
	Seeker
}

// ReadWriteSeekerは基本的なRead、Write、Seekメソッドをグループ化するインターフェースです。
type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}

// ReaderFromはReadFromメソッドをラップするインターフェースです。
//
// ReadFromはエラーが発生するか、EOFに到達するまでrからデータを読み取ります。
// 返り値のnは読み取られたバイト数です。
// 読み取り中にEOF以外のエラーも返されます。
//
// Copy関数はReaderFromが利用可能な場合に使用します。
type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}

// WriterToはWriteToメソッドをラップするインタフェースです。
//
// WriteToはデータを書き込み、書き込むデータがなくなるか
// エラーが発生するまでwに書き込みます。戻り値nは書き込まれた
// バイト数です。書き込み中にエラーが発生した場合はそれも返されます。
//
// Copy関数は利用可能であればWriterToを使用します。
type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}

// ReaderAtは基本的なReadAtメソッドをラップするインターフェースです。
//
// ReadAtは、基になる入力ソースのオフセットoffから始まるpにlen(p)バイトを読み込みます。
// 読み込まれたバイト数(0 <= n <= len(p))と、エラーが発生した場合のエラーを返します。
//
// ReadAtがn < len(p)を返す場合、なぜ追加のバイトが返されなかったかを説明する非nilのエラーが返されます。
// この点で、ReadAtはReadよりも厳格です。
//
// ReadAtがn < len(p)を返す場合でも、呼び出し中にpをスクラッチ領域として使用することがあります。
// データが一部利用可能であるがlen(p)バイトではない場合、ReadAtはすべてのデータが利用可能になるかエラーが発生するまでブロックされます。
// この点で、ReadAtはReadと異なります。
//
// ReadAtが返すn = len(p)バイトが入力ソースの末尾にある場合、ReadAtはerr == EOFまたはerr == nilのどちらかを返す可能性があります。
//
// ReadAtがシークオフセットを持つ入力ソースから読み込んでいる場合、
// ReadAtは基になるシークオフセットに影響を与えないし、影響を受けません。
//
// ReadAtのクライアントは、同じ入力ソースに対して並行してReadAt呼び出しを実行できます。
//
// 実装はpを保持してはいけません。
type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}

// WriterAtは基本的なWriteAtメソッドをラップするインターフェースです。
//
// WriteAtはpからlen(p)バイトを下層のデータストリームのoffセットに書き込みます。
// 書き込まれたバイト数n(0 <= n <= len(p))と書き込みが早期に終了した原因となった
// エラーが返されます。もしn < len(p)なら、WriteAtは非nilのエラーを返さなければなりません。
//
// もしWriteAtがseekオフセットを持つ宛先に書き込みを行っている場合、
// WriteAtは下層のseekオフセットに影響を与えず、また影響を受けてはいけません。
//
// WriteAtのクライアントは、範囲が重ならない場合には同じ宛先で並列して
// WriteAt呼び出しを実行することができます。
//
// 実装はpを保持してはいけません。
type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}

// ByteReaderはReadByteメソッドをラップするインターフェースです。
//
// ReadByteは入力から次のバイトを読み取り、それを返します。
// エラーが発生した場合、入力バイトは消費されず、返されるバイト値は未定義です。
//
// ReadByteはバイト単位の効率的な処理を提供します。
// ByteReaderを実装していないReaderは、bufio.NewReaderを使用してこのメソッドを追加することができます。
type ByteReader interface {
	ReadByte() (byte, error)
}

// ByteScannerは、基本のReadByteメソッドにUnreadByteメソッドを追加するインターフェースです。
//
// UnreadByteは、次のReadByte呼び出しで最後に読み込まれたバイトを返します。
// 最後の操作がReadByteへの成功した呼び出しでない場合、UnreadByteはエラーを返す可能性があります。
// 最後に読み込まれたバイト（または最後に読み込まれていないバイトの前のバイト）を未読状態に戻すか、
// （Seekerインターフェースをサポートする実装の場合）現在のオフセットの1バイト前にシークします。
type ByteScanner interface {
	ByteReader
	UnreadByte() error
}

// ByteWriterはWriteByteメソッドをラップするインターフェースです。
type ByteWriter interface {
	WriteByte(c byte) error
}

// RuneReaderはReadRuneメソッドをラップするインターフェースです。
//
// ReadRuneは単一のエンコードされたUnicode文字を読み取り、
// その文字とバイトサイズを返します。文字が利用できない場合、errが設定されます。
type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}

// RuneScannerは基本のReadRuneメソッドにUnreadRuneメソッドを追加するインターフェースです。
//
// UnreadRuneは次のReadRune呼び出しで最後に読み取られたルーンを返します。
// もし最後の操作が成功したReadRune呼び出しでない場合、UnreadRuneはエラーを返す、最後に読み取られたルーン（または最後に未読となったルーンの前のルーン）を未読扱いにする、
// または（Seekerインターフェースをサポートする実装の場合）現在のオフセットの直前のルーンの先頭にシークする可能性があります。
type RuneScanner interface {
	RuneReader
	UnreadRune() error
}

// StringWriterは、WriteStringメソッドを包んでいるインターフェースです。
type StringWriter interface {
	WriteString(s string) (n int, err error)
}

// WriteStringはstring sの内容をバイトスライスを受け取るwに書き込みます。
// もしwが [StringWriter] を実装している場合は、 [StringWriter.WriteString] が直接呼び出されます。
// そうでない場合は、 [Writer.Write] が一度だけ呼び出されます。
func WriteString(w Writer, s string) (n int, err error)

// ReadAtLeastは、rからbufに少なくともminバイト読み取るまで読み取ります。
// 読み取られたバイト数と、読み取りが少なかった場合のエラーを返します。
// エラーがEOFの場合、読み取られたバイトがない場合のみです。
// minバイト未満の読み取り後にEOFが発生した場合、ReadAtLeastはErrUnexpectedEOFを返します。
// minがbufの長さよりも大きい場合、ReadAtLeastはErrShortBufferを返します。
// 戻り値のn >= min if and only if err == nil。
// rが少なくともminバイトを読み取った後にエラーが発生した場合、エラーは破棄されます。
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)

// ReadFullはrからbufにちょうどlen(buf)バイト読み込みます。
// 読み込まれたバイト数と、読み込まれたバイト数が少なかった場合のエラーが返されます。
// エラーは、バイトが一つも読み込まれなかった場合にのみEOFです。
// 一部のバイトが読み込まれた後にEOFが発生した場合、ReadFullはErrUnexpectedEOFを返します。
// 返り値では、n == len(buf)であるのはerr == nilの場合のみです。
// rが少なくともlen(buf)バイトを読み込んだ後にエラーが発生した場合、そのエラーは無視されます。
func ReadFull(r Reader, buf []byte) (n int, err error)

// CopyNはsrcからdstにnバイト（またはエラーになるまで）をコピーします。
// コピーされたバイト数とコピー中に最初に遭遇したエラーが返されます。
// コピーが完了し、エラーがない場合、written == nが成り立ちます。
//
// もしdstが [ReaderFrom] インターフェースを実装している場合、
// そのインターフェースを使用してコピーが実装されます。
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)

// CopyはsrcからdstにEOFに到達するか、エラーが発生するまでコピーします。コピーされたバイト数と最初に遭遇したエラーが返されます。
// 成功したコピーではerr == nilを返します。err == EOFではありません。
// CopyはsrcからEOFに到達するまで読み取るように定義されているため、ReadからのEOFは報告されるエラーではありません。
//
// srcが [WriterTo] インターフェースを実装している場合、コピーはsrc.WriteTo(dst)の呼び出しで実装されます。
// それ以外の場合、dstが [ReaderFrom] インターフェースを実装している場合、コピーはdst.ReadFrom(src)の呼び出しで実装されます。
func Copy(dst Writer, src Reader) (written int64, err error)

// CopyBuffer は Copy と同じですが、一時的なバッファを割り当てる代わりに、
// 提供されたバッファを使用してステージングします（必要な場合）。buf が nil の場合、一時的なバッファが割り当てられます。
// もし buf の長さがゼロなら、CopyBuffer はパニックを引き起こします。
//
// もし src が [WriterTo] を実装しているか、dst が [ReaderFrom] を実装している場合、
// コピーを実行するために buf は使用されません。
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)

// LimitReaderは、rから読み取るReaderを返します。
// しかし、nバイト後にEOFで停止します。
// 内部の実装は*LimitedReaderです。
func LimitReader(r Reader, n int64) Reader

// LimitedReaderはRから読み込みますが、返されるデータ量をNバイトに制限します。Readの各呼び出しは、新しい残りの量を反映するためにNを更新します。また、N <= 0または基になるRがEOFを返す場合に、ReadはEOFを返します。
type LimitedReader struct {
	R Reader
	N int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error)

// NewSectionReaderは、rからオフセットoffで読み取りを開始し、nバイト後にEOFで停止するSectionReaderを返します。
func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader

// SectionReaderは、元となる [ReaderAt] の一部に対してRead、Seek、ReadAtを実装します。
type SectionReader struct {
	r     ReaderAt
	base  int64
	off   int64
	limit int64
	n     int64
}

func (s *SectionReader) Read(p []byte) (n int, err error)

func (s *SectionReader) Seek(offset int64, whence int) (int64, error)

func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)

// Size はセクションのサイズをバイト単位で返します。
func (s *SectionReader) Size() int64

<<<<<<< HEAD
// OffsetWriterは、基準オフセットから基準オフセット+オフセットの範囲で下位のライターへの書き込みをマッピングします。
=======
// Outer returns the underlying ReaderAt and offsets for the section.
//
// The returned values are the same that were passed to NewSectionReader
// when the SectionReader was created.
func (s *SectionReader) Outer() (r ReaderAt, off int64, n int64)

// An OffsetWriter maps writes at offset base to offset base+off in the underlying writer.
>>>>>>> upstream/master
type OffsetWriter struct {
	w    WriterAt
	base int64
	off  int64
}

// NewOffsetWriterは、オフセットoffから書き込むOffsetWriterを返します。
func NewOffsetWriter(w WriterAt, off int64) *OffsetWriter

func (o *OffsetWriter) Write(p []byte) (n int, err error)

func (o *OffsetWriter) WriteAt(p []byte, off int64) (n int, err error)

func (o *OffsetWriter) Seek(offset int64, whence int) (int64, error)

// TeeReaderは、rから読み取ったものをwに書き込むReaderを返します。
// これを通じて実行されるrからの全ての読み取りは、
// 対応するwへの書き込みとマッチングされます。
// 内部バッファリングはありません -
// 読み取りが完了する前に書き込みが完了する必要があります。
// 書き込み中にエラーが発生した場合、読み取りエラーとして報告されます。
func TeeReader(r Reader, w Writer) Reader

// Discardは、何もせずにすべての書き込み呼び出しに成功するWriterです。
var Discard Writer = discard{}

// discardは、最適化としてReaderFromを実装しています。そのため、io.DiscardへのCopyに不要な作業を避けることができます。
var _ ReaderFrom = discard{}

// NopCloser は、提供された [Reader] r を包む、Close メソッドの動作がない [ReadCloser] を返します。
// r が [WriterTo] を実装している場合、返された ReadCloser は WriterTo を実装し、
// 呼び出しを r に転送します。
func NopCloser(r Reader) ReadCloser

// ReadAllはrからエラーまたはEOFが発生するまで読み取り、読み取ったデータを返します。
// 成功した呼び出しではerr == nil、err == EOFではありません。ReadAllはsrcからEOFが発生するまで読み取るように定義されているため、ReadからのEOFはエラーとして報告されません。
func ReadAll(r Reader) ([]byte, error)
