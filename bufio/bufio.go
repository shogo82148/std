// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bufio はバッファードI/Oを実装しています。io.Readerまたはio.Writerオブジェクトをラップして、
// バッファリングやテキストI/Oのための支援を提供する別のオブジェクト（ReaderまたはWriter）を作成します。
package bufio

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

// Readerはio.Readerオブジェクトに対してバッファリングを実装します。
type Reader struct {
	buf          []byte
	rd           io.Reader
	r, w         int
	err          error
	lastByte     int
	lastRuneSize int
}

// NewReaderSizeは、バッファの最低限のサイズが指定された新しい [Reader] を返します。
// もし引数のio.Readerが既に十分なサイズの [Reader] であれば、それは基本となる [Reader] を返します。
func NewReaderSize(rd io.Reader, size int) *Reader

// NewReaderはデフォルトサイズのバッファを持つ新しい [Reader] を返します。
func NewReader(rd io.Reader) *Reader

// Sizeはバイト単位での基礎バッファのサイズを返します。
func (b *Reader) Size() int

// Resetはバッファに保持されたデータを破棄し、すべての状態をリセットし、
// バッファリーダーをrから読み取るように切り替えます。
// [Reader] のゼロ値に対してResetを呼び出すと、内部バッファがデフォルトのサイズに初期化されます。
// b.Reset(b)（つまり、[Reader] を自身にリセットする）は何もしません。
func (b *Reader) Reset(r io.Reader)

// Peek returns the next n bytes without advancing the reader. The bytes stop
// being valid at the next read call. If Peek returns fewer than n bytes, it
// also returns an error explaining why the read is short. The error is
// [ErrBufferFull] if n is larger than b's buffer size.
//
// Calling Peek prevents a [Reader.UnreadByte] or [Reader.UnreadRune] call from succeeding
// until the next read operation.
func (b *Reader) Peek(n int) ([]byte, error)

// Discard は次の n バイトをスキップし、スキップしたバイト数を返します。
//
// もし Discard が n バイト未満をスキップした場合、エラーも返します。
// もし 0 <= n <= b.Buffered() のとき、Discard は io.Reader の下の方から読み取らずに必ず成功することが保証されています。
func (b *Reader) Discard(n int) (discarded int, err error)

// Readはデータをpに読み込みます。
// pに読み込まれたバイト数を返します。
// バイトは基礎となる [Reader] のReadから最大1つ取り出されますので、nはlen(p)より少ない場合があります。
// len(p)バイトを正確に読み取るには、io.ReadFull(b, p)を使用してください。
// 基礎となる [Reader] がio.EOFで非ゼロの数を返す可能性がある場合、このReadメソッドも同様です。詳細は [io.Reader] ドキュメントを参照してください。
func (b *Reader) Read(p []byte) (n int, err error)

// ReadByteは1バイトを読み取って返します。
// バイトが利用できない場合、エラーを返します。
func (b *Reader) ReadByte() (byte, error)

// UnreadByteは最後のバイトを未読状態に戻します。直前に読み込まれたバイトのみが未読状態に戻すことができます。
//
// UnreadByteは、[Reader] に対して最後に呼び出されたメソッドが読み込み操作ではない場合にエラーを返します。特に、 [Reader.Peek] 、 [Reader.Discard] 、および [Reader.WriteTo] は読み込み操作とはみなされません。
func (b *Reader) UnreadByte() error

// ReadRuneは、1つのUTF-8エンコードされたユニコード文字を読み込み、
// そのルーンとバイトサイズを返します。エンコードされたルーンが無効な場合は、1バイトを消費し、
// サイズが1のunicode.ReplacementChar（U+FFFD）を返します。
func (b *Reader) ReadRune() (r rune, size int, err error)

// UnreadRuneは最後のルーンを戻します。もし、 [Reader] に最も最近呼び出されたメソッドが [Reader.ReadRune] でない場合、 [Reader.UnreadRune] はエラーを返します。（この点で、 [Reader.UnreadByte] よりも厳格です。[Reader.UnreadByte] はどの読み取り操作からも最後のバイトを戻します。）
func (b *Reader) UnreadRune() error

// Bufferedは現在のバッファから読み取ることができるバイト数を返します。
func (b *Reader) Buffered() int

<<<<<<< HEAD
// ReadSliceは入力内の最初のデリミタの出現まで読み取り、バッファ内のバイトを指すスライスを返します。
// バイトは次の読み取り時には無効になります。
// ReadSliceがデリミタを見つける前にエラーに遭遇した場合、バッファ内のすべてのデータとエラー自体（通常はio.EOF）を返します。
// バッファがデリミタなしで満杯になると、ReadSliceはErrBufferFullエラーで失敗します。
// ReadSliceから返されるデータは次のI/O操作によって上書きされるため、ほとんどのクライアントは
// ReadBytesまたはReadStringを代わりに使用すべきです。
// ReadSliceは、lineの終了がデリミタでない場合にのみerr！= nilを返します。
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

// ReadLineは低レベルの行読み取りプリミティブです。ほとんどの呼び出し元は、ReadBytes（'\n'）またはReadString（'\n'）を使用するか、Scannerを使用する必要があります。
// ReadLineは、改行文字を含まない1行だけを返そうとします。もし行がバッファーに対して長すぎる場合、isPrefixが設定され、行の先頭が返されます。それ以降の行は、将来の呼び出しで返されます。最後のフラグメントを返す際には、isPrefixはfalseになります。返されるバッファーは、次のReadLine呼び出しまでの間のみ有効です。ReadLineは、nilではない行を返すか、エラーを返すか、どちらかを返しますが、両方を返すことはありません。
// ReadLineから返されるテキストには、行末の("\r\n"または"\n")は含まれません。入力が最後の行末で終わっている場合、特定の表示やエラーは与えられません。ReadLineの後にUnreadByteを呼び出すと、常に最後に読み取られたバイト（おそらく行末に属する文字）がアンリードされます。ただし、そのバイトがReadLineによって返された行の一部でない場合でもです。
=======
// ReadSlice reads until the first occurrence of delim in the input,
// returning a slice pointing at the bytes in the buffer.
// The bytes stop being valid at the next read.
// If ReadSlice encounters an error before finding a delimiter,
// it returns all the data in the buffer and the error itself (often io.EOF).
// ReadSlice fails with error [ErrBufferFull] if the buffer fills without a delim.
// Because the data returned from ReadSlice will be overwritten
// by the next I/O operation, most clients should use
// [Reader.ReadBytes] or ReadString instead.
// ReadSlice returns err != nil if and only if line does not end in delim.
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

// ReadLine is a low-level line-reading primitive. Most callers should use
// [Reader.ReadBytes]('\n') or [Reader.ReadString]('\n') instead or use a [Scanner].
//
// ReadLine tries to return a single line, not including the end-of-line bytes.
// If the line was too long for the buffer then isPrefix is set and the
// beginning of the line is returned. The rest of the line will be returned
// from future calls. isPrefix will be false when returning the last fragment
// of the line. The returned buffer is only valid until the next call to
// ReadLine. ReadLine either returns a non-nil line or it returns an error,
// never both.
//
// The text returned from ReadLine does not include the line end ("\r\n" or "\n").
// No indication or error is given if the input ends without a final line end.
// Calling [Reader.UnreadByte] after ReadLine will always unread the last byte read
// (possibly a character belonging to the line end) even if that byte is not
// part of the line returned by ReadLine.
>>>>>>> upstream/master
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)

// ReadBytesは入力内のデリミタの最初の出現まで読み取り、
// データとデリミタを含むスライスを返します。
// ReadBytesがデリミタを見つける前にエラーが発生した場合、
// エラーが発生する前に読み取られたデータとエラー自体（通常はio.EOF）を返します。
// Returned dataがデリミタで終わっていない場合、ReadBytesはerr != nilを返します。
// 簡単な使用のためには、Scannerがより便利です。
func (b *Reader) ReadBytes(delim byte) ([]byte, error)

// ReadStringは、入力内で最初のデリミタが現れるまで読み込み、デリミタを含むデータの文字列を返します。
// ReadStringがデリミタを見つける前にエラーに遭遇した場合、エラー自体（通常はio.EOF）とエラーが発生する前に読み取ったデータを返します。
// ReadStringは、返されたデータの最後がデリミタで終わっていない場合、err != nilを返します。
// 単純な使用法の場合は、Scannerがより便利です。
func (b *Reader) ReadString(delim byte) (string, error)

<<<<<<< HEAD
// WriteToはio.WriterToを実装します。
// これは基礎となるReaderのReadメソッドを複数回呼び出すことがあります。
// 基礎となるreaderがWriteToメソッドをサポートしている場合、
// これはバッファリングせずに基礎となるWriteToを呼び出します。
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)

// Writerはio.Writerオブジェクトに対してバッファリングを行います。
// Writerに書き込む際にエラーが発生した場合、以降のデータの受け入れや、Flushメソッドの呼び出しはエラーを返します。
// 全てのデータが書き込まれた後、クライアントはFlushメソッドを呼び出して、全てのデータが基になるio.Writerに転送されることを保証する必要があります。
=======
// WriteTo implements io.WriterTo.
// This may make multiple calls to the [Reader.Read] method of the underlying [Reader].
// If the underlying reader supports the [Reader.WriteTo] method,
// this calls the underlying [Reader.WriteTo] without buffering.
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)

// Writer implements buffering for an [io.Writer] object.
// If an error occurs writing to a [Writer], no more data will be
// accepted and all subsequent writes, and [Writer.Flush], will return the error.
// After all data has been written, the client should call the
// [Writer.Flush] method to guarantee all data has been forwarded to
// the underlying [io.Writer].
>>>>>>> upstream/master
type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}

<<<<<<< HEAD
// NewWriterSizeは、バッファのサイズが指定された最小値を持つ新しいWriterを返します。
// 引数のio.Writerがすでに十分な大きさを持つWriterである場合、基になるWriterを返します。
func NewWriterSize(w io.Writer, size int) *Writer

// NewWriterは、バッファのデフォルトサイズを持つ新しいWriterを返します。
// 引数のio.Writerが既に十分に大きなバッファサイズを持つWriterである場合、基になるWriterを返します。
=======
// NewWriterSize returns a new [Writer] whose buffer has at least the specified
// size. If the argument io.Writer is already a [Writer] with large enough
// size, it returns the underlying [Writer].
func NewWriterSize(w io.Writer, size int) *Writer

// NewWriter returns a new [Writer] whose buffer has the default size.
// If the argument io.Writer is already a [Writer] with large enough buffer size,
// it returns the underlying [Writer].
>>>>>>> upstream/master
func NewWriter(w io.Writer) *Writer

// Sizeはバイト単位で下層のバッファーのサイズを返します。
func (b *Writer) Size() int

<<<<<<< HEAD
// Resetは、フラッシュされていないバッファデータを破棄し、エラーをクリアし、出力をwにリセットします。
// Writerのゼロ値に対してResetを呼び出すと、内部バッファがデフォルトのサイズに初期化されます。
// w.Reset(w)（つまり、Writerを自身にリセットすること）は何もしません。
func (b *Writer) Reset(w io.Writer)

// Flushはバッファされたデータを基になるio.Writerに書き込みます。
=======
// Reset discards any unflushed buffered data, clears any error, and
// resets b to write its output to w.
// Calling Reset on the zero value of [Writer] initializes the internal buffer
// to the default size.
// Calling w.Reset(w) (that is, resetting a [Writer] to itself) does nothing.
func (b *Writer) Reset(w io.Writer)

// Flush writes any buffered data to the underlying [io.Writer].
>>>>>>> upstream/master
func (b *Writer) Flush() error

// Available はバッファ内で未使用のバイト数を返します。
func (b *Writer) Available() int

<<<<<<< HEAD
// AvailableBufferは、b.Available（）容量の空のバッファを返します。
// このバッファは追加されることを意図しており、
// 直後のWrite呼び出しに渡されます。
// このバッファは、b上の次の書き込み操作までの間のみ有効です。
=======
// AvailableBuffer returns an empty buffer with b.Available() capacity.
// This buffer is intended to be appended to and
// passed to an immediately succeeding [Writer.Write] call.
// The buffer is only valid until the next write operation on b.
>>>>>>> upstream/master
func (b *Writer) AvailableBuffer() []byte

// Bufferedは現在のバッファに書き込まれたバイト数を返します。
func (b *Writer) Buffered() int

// Write は p の内容をバッファに書き込みます。
// 書き込まれたバイト数を返します。
// nn < len(p) の場合、短い書き込みの理由を説明するエラーも返ります。
func (b *Writer) Write(p []byte) (nn int, err error)

// WriteByteは1バイトを書き込みます。
func (b *Writer) WriteByte(c byte) error

// WriteRuneは一つのUnicodeコードポイントを書き込み、書き込んだバイト数とエラーを返します。
func (b *Writer) WriteRune(r rune) (size int, err error)

// WriteString関数は文字列を書き込みます。
// 書き込んだバイト数を返します。
// もし書き込んだバイト数がsの長さよりも少ない場合、短い書き込みである理由を説明するエラーも返されます。
func (b *Writer) WriteString(s string) (int, error)

<<<<<<< HEAD
// ReadFrom は io.ReaderFrom インターフェースを実装します。もし基礎となる書き込み先が ReadFrom メソッドをサポートしている場合、これは基礎となる ReadFrom を呼び出します。
// バッファされたデータと基礎となる ReadFrom がある場合、これはバッファを埋めてから ReadFrom を呼び出します。
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)

// ReadWriterはReaderとWriterへのポインタを保存します。
// io.ReadWriterを実装します。
=======
// ReadFrom implements [io.ReaderFrom]. If the underlying writer
// supports the ReadFrom method, this calls the underlying ReadFrom.
// If there is buffered data and an underlying ReadFrom, this fills
// the buffer and writes it before calling ReadFrom.
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)

// ReadWriter stores pointers to a [Reader] and a [Writer].
// It implements [io.ReadWriter].
>>>>>>> upstream/master
type ReadWriter struct {
	*Reader
	*Writer
}

<<<<<<< HEAD
// NewReadWriterはrとwにディスパッチする新しいReadWriterを割り当てます。
=======
// NewReadWriter allocates a new [ReadWriter] that dispatches to r and w.
>>>>>>> upstream/master
func NewReadWriter(r *Reader, w *Writer) *ReadWriter
