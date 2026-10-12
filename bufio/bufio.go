// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// bufioパッケージはバッファードI/Oを実装しています。io.Readerまたはio.Writerオブジェクトをラップして、
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
// 新しいReaderは [NewReader] または [NewReaderSize] を呼び出すことで作成されます。
// あるいは、[Reader.Reset] を呼び出した後にReaderのゼロ値を使用することもできます。
type Reader struct {
	buf          []byte
	rd           io.Reader
	r, w         int
	err          error
	lastByte     int
	lastRuneSize int
}

// NewReaderSizeは、バッファーの最低限のサイズが指定された新しい [Reader] を返します。
// もし引数のio.Readerが既に十分なサイズの [Reader] であれば、それは基礎となる [Reader] を返します。
func NewReaderSize(rd io.Reader, size int) *Reader

// NewReaderはデフォルトサイズのバッファーを持つ新しい [Reader] を返します。
func NewReader(rd io.Reader) *Reader

// Sizeはバイト単位での基礎バッファーのサイズを返します。
func (b *Reader) Size() int

// Resetはバッファーに保持されたデータを破棄し、すべての状態をリセットし、
// バッファーリーダーをrから読み取るように切り替えます。
// [Reader] のゼロ値に対してResetを呼び出すと、内部バッファーがデフォルトのサイズに初期化されます。
// b.Reset(b)（つまり、[Reader] を自身にリセットする）は何もしません。
func (b *Reader) Reset(r io.Reader)

// Peekはリーダーを進めずに次のnバイトを返します。返されたバイトは
// 次の読み取り呼び出しで無効になります。必要に応じて、Peekはnバイトを
// 利用可能にするためにバッファーにさらにバイトを読み込みます。Peekが
// nバイト未満を返す場合、読み取りが不足した理由を説明するエラーも返します。
// nがbのバッファーサイズより大きい場合、エラーは [ErrBufferFull] です。
//
// Peekを呼び出すと、次の読み取り操作まで [Reader.UnreadByte] や [Reader.UnreadRune] の
// 呼び出しは成功しなくなります。
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

// Bufferedは現在のバッファーから読み取ることができるバイト数を返します。
func (b *Reader) Buffered() int

// ReadSliceは入力内の最初のデリミターの出現まで読み取り、バッファー内のバイトを指すスライスを返します。
// バイトは次の読み取り時には無効になります。
// ReadSliceがデリミターを見つける前にエラーに遭遇した場合、バッファー内のすべてのデータとエラー自体（通常はio.EOF）を返します。
// バッファーがデリミターなしで満杯になると、ReadSliceは [ErrBufferFull] エラーで失敗します。
// ReadSliceから返されるデータは次のI/O操作によって上書きされるため、ほとんどのクライアントは
// [Reader.ReadBytes] またはReadStringを代わりに使用すべきです。
// ReadSliceは、lineの終了がデリミターでない場合にのみerr！= nilを返します。
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

// ReadLineは低レベルの行読み取りプリミティブです。ほとんどの呼び出し元は、 [Reader.ReadBytes]('\n') または [Reader.ReadString]('\n') を使用するか、 [Scanner] を使用する必要があります。
//
// ReadLineは、改行文字を含まない1行だけを返そうとします。もし行がバッファーに対して長すぎる場合、isPrefixが設定され、行の先頭が返されます。それ以降の行は、将来の呼び出しで返されます。最後のフラグメントを返す際には、isPrefixはfalseになります。返されるバッファーは、次のReadLine呼び出しまでの間のみ有効です。ReadLineは、nilではない行を返すか、エラーを返すか、どちらかを返しますが、両方を返すことはありません。
//
// ReadLineから返されるテキストには、行末の("\r\n"または"\n")は含まれません。入力が最後の行末で終わっている場合、特定の表示やエラーは与えられません。ReadLineの後に [Reader.UnreadByte] を呼び出すと、常に最後に読み取られたバイト（おそらく行末に属する文字）がアンリードされます。ただし、そのバイトがReadLineによって返された行の一部でない場合でもです。
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)

// ReadBytesは入力内のデリミターの最初の出現まで読み取り、
// データとデリミターを含むスライスを返します。
// ReadBytesがデリミターを見つける前にエラーが発生した場合、
// エラーが発生する前に読み取られたデータとエラー自体（通常はio.EOF）を返します。
// Returned dataがデリミターで終わっていない場合、ReadBytesはerr != nilを返します。
// 簡単な使用のためには、Scannerがより便利です。
func (b *Reader) ReadBytes(delim byte) ([]byte, error)

// ReadStringは、入力内で最初のデリミターが現れるまで読み込み、デリミターを含むデータの文字列を返します。
// ReadStringがデリミターを見つける前にエラーに遭遇した場合、エラー自体（通常はio.EOF）とエラーが発生する前に読み取ったデータを返します。
// ReadStringは、返されたデータの最後がデリミターで終わっていない場合、err != nilを返します。
// 単純な使用法の場合は、Scannerがより便利です。
func (b *Reader) ReadString(delim byte) (string, error)

// WriteToはio.WriterToを実装します。
// これは基礎となる [Reader] の [Reader.Read] メソッドを複数回呼び出すことがあります。
// 基礎となるreaderが [Reader.WriteTo] メソッドをサポートしている場合、
// これはバッファリングせずに基礎となる [Reader.WriteTo] を呼び出します。
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)

// Writerは [io.Writer] オブジェクトに対してバッファリングを行います。
// [Writer] に書き込む際にエラーが発生した場合、以降のデータの受け入れや、 [Writer.Flush] メソッドの呼び出しはエラーを返します。
// 全てのデータが書き込まれた後、クライアントは [Writer.Flush] メソッドを呼び出して、全てのデータが基礎となる [io.Writer] に転送されることを保証する必要があります。
type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}

// NewWriterSizeは、バッファーのサイズが指定された最小値を持つ新しい [Writer] を返します。
// 引数のio.Writerがすでに十分な大きさを持つ [Writer] である場合、基礎となる [Writer] を返します。
func NewWriterSize(w io.Writer, size int) *Writer

// NewWriterは、バッファーのデフォルトサイズを持つ新しい [Writer] を返します。
// 引数のio.Writerが既に十分に大きなバッファーサイズを持つ [Writer] である場合、基礎となる [Writer] を返します。
func NewWriter(w io.Writer) *Writer

// Sizeはバイト単位で基礎となるバッファーのサイズを返します。
func (b *Writer) Size() int

// Resetは、フラッシュされていないバッファーデータを破棄し、エラーをクリアし、出力をwにリセットします。
// [Writer] のゼロ値に対してResetを呼び出すと、内部バッファーがデフォルトのサイズに初期化されます。
// w.Reset(w)（つまり、[Writer] を自身にリセットすること）は何もしません。
func (b *Writer) Reset(w io.Writer)

// Flushはバッファーされたデータを基礎となる [io.Writer] に書き込みます。
func (b *Writer) Flush() error

// Available はバッファー内で未使用のバイト数を返します。
func (b *Writer) Available() int

// AvailableBufferは、b.Available() 容量の空のバッファーを返します。
// このバッファーは追加されることを意図しており、
// 直後の [Writer.Write] 呼び出しに渡されます。
// このバッファーは、b上の次の書き込み操作までの間のみ有効です。
func (b *Writer) AvailableBuffer() []byte

// Bufferedは現在のバッファーに書き込まれたバイト数を返します。
func (b *Writer) Buffered() int

// Write は p の内容をバッファーに書き込みます。
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

// ReadFrom は [io.ReaderFrom] インターフェースを実装します。もし基礎となる書き込み先が ReadFrom メソッドをサポートしている場合、これは基礎となる ReadFrom を呼び出します。
// バッファーされたデータと基礎となる ReadFrom がある場合、これはバッファーを埋めてから ReadFrom を呼び出します。
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)

// ReadWriterは [Reader] と [Writer] へのポインターを保存します。
// [io.ReadWriter] を実装します。
type ReadWriter struct {
	*Reader
	*Writer
}

// NewReadWriterはrとwにディスパッチする新しい [ReadWriter] を割り当てます。
func NewReadWriter(r *Reader, w *Writer) *ReadWriter
