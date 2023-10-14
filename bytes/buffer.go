// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

<<<<<<< HEAD
// Bufferはバイトの可変サイズのバッファであり、ReadとWriteのメソッドを持っています。
// Bufferのゼロ値は、使い-readyになった空のバッファです。
=======
// A Buffer is a variable-sized buffer of bytes with [Buffer.Read] and [Buffer.Write] methods.
// The zero value for Buffer is an empty buffer ready to use.
>>>>>>> upstream/master
type Buffer struct {
	buf      []byte
	off      int
	lastRead readOp
}

// ErrTooLargeは、バッファにデータを格納するためのメモリを割り当てることができない場合にpanicに渡されます。
var ErrTooLarge = errors.New("bytes.Buffer: too large")

<<<<<<< HEAD
// Bytesは、バッファの未読部分を保持する長さb.Len()のスライスを返します。
// このスライスは、次のバッファの変更までしか有効ではありません（つまり、
// Read、Write、Reset、またはTruncateなどのメソッドの次の呼び出しまでのみ有効です）。
// スライスは、次のバッファの変更まで少なくともバッファの内容をエイリアスしているため、
// スライスへの直接の変更は将来の読み取り結果に影響を与えます。
func (b *Buffer) Bytes() []byte

// AvailableBufferはb.Available()の容量を持つ空のバッファを返します。
// このバッファは追加され、直後のWrite呼び出しに渡すことが想定されています。
// バッファはbに対する次の書き込み操作までの間のみ有効です。
func (b *Buffer) AvailableBuffer() []byte

// Stringはバッファの未読部分の内容を文字列として返します。もしBufferがnilポインタであれば、"<nil>"を返します。
=======
// Bytes returns a slice of length b.Len() holding the unread portion of the buffer.
// The slice is valid for use only until the next buffer modification (that is,
// only until the next call to a method like [Buffer.Read], [Buffer.Write], [Buffer.Reset], or [Buffer.Truncate]).
// The slice aliases the buffer content at least until the next buffer modification,
// so immediate changes to the slice will affect the result of future reads.
func (b *Buffer) Bytes() []byte

// AvailableBuffer returns an empty buffer with b.Available() capacity.
// This buffer is intended to be appended to and
// passed to an immediately succeeding [Buffer.Write] call.
// The buffer is only valid until the next write operation on b.
func (b *Buffer) AvailableBuffer() []byte

// String returns the contents of the unread portion of the buffer
// as a string. If the [Buffer] is a nil pointer, it returns "<nil>".
>>>>>>> upstream/master
//
// より効率的に文字列を構築するには、strings.Builder型を参照してください。
func (b *Buffer) String() string

// Lenはバッファの未読部分のバイト数を返します。
// b.Len() == len(b.Bytes())。
func (b *Buffer) Len() int

// Capはバッファの基礎となるバイトスライスの容量、つまりバッファのデータのために割り当てられた総スペースを返します。
func (b *Buffer) Cap() int

// Availableはバッファ内で未使用のバイト数を返します。
func (b *Buffer) Available() int

// Truncateはバッファから最初のnバイト以外の未読データを削除し、同じ割り当てられたストレージを使用し続けます。
// nが負数またはバッファの長さよりも大きい場合、パニックが発生します。
func (b *Buffer) Truncate(n int)

<<<<<<< HEAD
// Resetはバッファを空にリセットしますが、将来の書き込みのために基礎となるストレージは保持されます。
// ResetはTruncate(0)と同じです。
func (b *Buffer) Reset()

// Growは必要に応じてバッファの容量を増やし、残りのnバイトの空間を保証します。Grow(n)の後、バッファには少なくともnバイトを別の割り当てなしで書き込むことができます。
// nが負数の場合、Growはパニックを引き起こします。
// バッファを拡大できない場合、ErrTooLargeでパニックを引き起こします。
func (b *Buffer) Grow(n int)

// Write はバッファーに p の内容を追加し、必要に応じてバッファーを拡張します。戻り値 n は p の長さであり、err は常に nil です。バッファーが大きすぎる場合、Write は ErrTooLarge とともにパニックを発生させます。
func (b *Buffer) Write(p []byte) (n int, err error)

// WriteStringは、必要に応じてバッファを拡張し、sの内容をバッファに追加します。戻り値nはsの長さであり、errは常にnilです。バッファが大きすぎる場合、WriteStringはErrTooLargeとともにパニックを発生させます。
func (b *Buffer) WriteString(s string) (n int, err error)

// MinReadはBuffer.ReadFromによってRead呼び出しに渡される最小のスライスサイズです。
// Bufferは、rの内容を保持するために必要なものを超えて少なくともMinReadバイトを持っている限り、ReadFromは基礎となるバッファを拡大しません。
const MinRead = 512

// ReadFromは、rからEOFまでデータを読み取り、バッファに追加していきます。必要に応じてバッファのサイズが拡大されます。返り値nは読み取られたバイト数です。読み取り中にio.EOF以外のエラーが発生した場合、それも返されます。バッファがあまりに大きくなると、ReadFromはErrTooLargeでパニックを引き起こします。
=======
// Reset resets the buffer to be empty,
// but it retains the underlying storage for use by future writes.
// Reset is the same as [Buffer.Truncate](0).
func (b *Buffer) Reset()

// Grow grows the buffer's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to the
// buffer without another allocation.
// If n is negative, Grow will panic.
// If the buffer can't grow it will panic with [ErrTooLarge].
func (b *Buffer) Grow(n int)

// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil. If the
// buffer becomes too large, Write will panic with [ErrTooLarge].
func (b *Buffer) Write(p []byte) (n int, err error)

// WriteString appends the contents of s to the buffer, growing the buffer as
// needed. The return value n is the length of s; err is always nil. If the
// buffer becomes too large, WriteString will panic with [ErrTooLarge].
func (b *Buffer) WriteString(s string) (n int, err error)

// MinRead is the minimum slice size passed to a Read call by
// [Buffer.ReadFrom]. As long as the [Buffer] has at least MinRead bytes beyond
// what is required to hold the contents of r, ReadFrom will not grow the
// underlying buffer.
const MinRead = 512

// ReadFrom reads data from r until EOF and appends it to the buffer, growing
// the buffer as needed. The return value n is the number of bytes read. Any
// error except io.EOF encountered during the read is also returned. If the
// buffer becomes too large, ReadFrom will panic with [ErrTooLarge].
>>>>>>> upstream/master
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)

// WriteTo はバッファが空になるかエラーが発生するまで、データを w に書き込みます。
// 戻り値の n は書き込まれたバイト数です。この値は常に int に収まりますが、io.WriterTo インターフェースに合わせて int64 型です。書き込み中に発生したエラーも返されます。
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)

<<<<<<< HEAD
// WriteByteはバイトcをバッファに追加し、必要に応じてバッファを拡張します。
// 返されるエラーは常にnilですが、bufio.WriterのWriteByteに合わせるために含まれています。
// バッファが大きすぎる場合、WriteByteはErrTooLargeでパニックします。
func (b *Buffer) WriteByte(c byte) error

// WriteRuneはUnicodeコードポイントrのUTF-8エンコーディングをバッファに追加し、その長さと常にnilであるエラーを返します。エラーは常にnilですが、bufio.WriterのWriteRuneとのマッチングのために含まれます。必要に応じてバッファは拡張されます。もしバッファがあまりにも大きくなった場合、WriteRuneはErrTooLargeでパニックを起こします。
=======
// WriteByte appends the byte c to the buffer, growing the buffer as needed.
// The returned error is always nil, but is included to match [bufio.Writer]'s
// WriteByte. If the buffer becomes too large, WriteByte will panic with
// [ErrTooLarge].
func (b *Buffer) WriteByte(c byte) error

// WriteRune appends the UTF-8 encoding of Unicode code point r to the
// buffer, returning its length and an error, which is always nil but is
// included to match [bufio.Writer]'s WriteRune. The buffer is grown as needed;
// if it becomes too large, WriteRune will panic with [ErrTooLarge].
>>>>>>> upstream/master
func (b *Buffer) WriteRune(r rune) (n int, err error)

// Readは、バッファから次のlen(p)バイトを読み取るか、バッファが空になるまで読み取ります。返り値nは読み取られたバイト数です。バッファに返すデータがない場合、errはio.EOFです（len(p)がゼロの場合を除く）；それ以外の場合、nilです。
func (b *Buffer) Read(p []byte) (n int, err error)

<<<<<<< HEAD
// Nextは、バッファから次のnバイトを含むスライスを返し、
// バイトがReadによって返された場合と同様にバッファを進めます。
// バッファにnバイト未満のバイトがある場合、Nextはバッファ全体を返します。
// スライスは、次の読み取りまたは書き込みメソッドの呼び出しまでの間のみ有効です。
=======
// Next returns a slice containing the next n bytes from the buffer,
// advancing the buffer as if the bytes had been returned by [Buffer.Read].
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
// The slice is only valid until the next call to a read or write method.
>>>>>>> upstream/master
func (b *Buffer) Next(n int) []byte

// ReadByte はバッファから次のバイトを読み込んで返します。
// バイトが利用できない場合は、エラー io.EOF を返します。
func (b *Buffer) ReadByte() (byte, error)

// ReadRuneはバッファから次のUTF-8エンコードされた
// Unicodeコードポイントを読み取り、返します。
// バイトが利用できない場合は、io.EOFというエラーが返されます。
// バイトが不正なUTF-8エンコーディングの場合、1バイトを消費し、U+FFFD、1を返します。
func (b *Buffer) ReadRune() (r rune, size int, err error)

<<<<<<< HEAD
// UnreadRuneはReadRuneによって返された最後のルーンを未読状態にします。
// バッファ上の直近の読み込みや書き込み操作が成功していない場合、UnreadRuneはエラーを返します。 (この点で、UnreadByteよりも厳格です。UnreadByteはすべての読み込み操作から最後のバイトを未読状態にします。)
=======
// UnreadRune unreads the last rune returned by [Buffer.ReadRune].
// If the most recent read or write operation on the buffer was
// not a successful [Buffer.ReadRune], UnreadRune returns an error.  (In this regard
// it is stricter than [Buffer.UnreadByte], which will unread the last byte
// from any read operation.)
>>>>>>> upstream/master
func (b *Buffer) UnreadRune() error

// UnreadByteは、少なくとも1バイトを読み込んだ最後の成功した読み込み操作で返された最後のバイトを戻します。最後の読み込み以降に書き込みが発生した場合、最後の読み込みがエラーを返した場合、または読み込みが0バイトを読み込んだ場合、UnreadByteはエラーを返します。
func (b *Buffer) UnreadByte() error

// ReadBytesは入力の最初のdelimが現れるまで読み取り、
// デリミタを含むデータを含むスライスを返します。
// ReadBytesがデリミタを見つける前にエラーに遭遇した場合、
// エラー自体（しばしばio.EOF）とエラー前に読み取ったデータを返します。
// 返されたデータの末尾がdelimで終わっていない場合、
// ReadBytesはerr != nilを返します。
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)

// ReadStringは入力の最初のデリミタが現れるまで読み取り、
// デリミタを含むデータを含む文字列を返します。
// ReadStringがデリミタを見つける前にエラーに遭遇する場合、
// エラー自体（通常はio.EOF）とエラーが発生する前に読み取ったデータを返します。
// ReadStringは、返されるデータがdelimで終わっていない場合、err！= nilを返します。
func (b *Buffer) ReadString(delim byte) (line string, err error)

<<<<<<< HEAD
// NewBufferは、bufを初期コンテンツとして使用して新しいBufferを作成および初期化します。
// 新しいBufferは、bufを所有し、この呼び出しの後にbufを使用しないようにする必要があります。
// NewBufferは、既存のデータを読むためにBufferを準備するためのものです。書き込み用の内部バッファの初期サイズを設定するためにも使用できます。そのためには、
// bufは希望する容量を持つ必要がありますが、長さはゼロである必要があります。
//
// ほとんどの場合、new(Buffer)（または単にBuffer変数を宣言する）で
// Bufferを初期化するのに十分です。
func NewBuffer(buf []byte) *Buffer

// NewBufferStringは、文字列sを初期内容として使用して新しいBufferを作成し、初期化します。既存の文字列を読むためのバッファを準備するために使用されます。
//
// ほとんどの場合、new(Buffer)（または単にBuffer変数を宣言する）でBufferを初期化するのに十分です。
=======
// NewBuffer creates and initializes a new [Buffer] using buf as its
// initial contents. The new [Buffer] takes ownership of buf, and the
// caller should not use buf after this call. NewBuffer is intended to
// prepare a [Buffer] to read existing data. It can also be used to set
// the initial size of the internal buffer for writing. To do that,
// buf should have the desired capacity but a length of zero.
//
// In most cases, new([Buffer]) (or just declaring a [Buffer] variable) is
// sufficient to initialize a [Buffer].
func NewBuffer(buf []byte) *Buffer

// NewBufferString creates and initializes a new [Buffer] using string s as its
// initial contents. It is intended to prepare a buffer to read an existing
// string.
//
// In most cases, new([Buffer]) (or just declaring a [Buffer] variable) is
// sufficient to initialize a [Buffer].
>>>>>>> upstream/master
func NewBufferString(s string) *Buffer
