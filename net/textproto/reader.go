// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textproto

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// Readerは、テキストプロトコルネットワーク接続からリクエストまたはレスポンスを読み取るための便利なメソッドを実装します。
type Reader struct {
	R   *bufio.Reader
	dot *dotReader
	buf []byte
}

<<<<<<< HEAD
// NewReaderはrから読み取りを行う新しいReaderを返します。
//
// サービス拒否攻撃を避けるために、提供されたbufio.Readerは
// io.LimitReaderまたは同様のReaderから読み取るようになっている必要があります。
=======
// NewReader returns a new [Reader] reading from r.
//
// To avoid denial of service attacks, the provided [bufio.Reader]
// should be reading from an [io.LimitReader] or similar Reader to bound
// the size of responses.
>>>>>>> upstream/release-branch.go1.22
func NewReader(r *bufio.Reader) *Reader

// ReadLineはrから1行だけ読み取り、返された文字列から最後の\nまたは\r\nを省略します。
func (r *Reader) ReadLine() (string, error)

<<<<<<< HEAD
// ReadLineBytesは、文字列の代わりに[]byteを返すReadLineと同様の機能です。
=======
// ReadLineBytes is like [Reader.ReadLine] but returns a []byte instead of a string.
>>>>>>> upstream/release-branch.go1.22
func (r *Reader) ReadLineBytes() ([]byte, error)

// ReadContinuedLineは、rから可能性がある継続行を読み取ります。
// 最後の余分なASCII空白は省略されます。
// 最初の行以降の行は、スペースまたはタブ文字で始まる場合には、
// 継続行と見なされます。返されるデータでは、
// 継続行は前の行とスペース1つのみで区切られます：
// 改行と先頭の空白は削除されます。
//
// 例えば、次の入力を考えてみてください：
//
//	Line 1
//	  continued...
//	Line 2
//
// ReadContinuedLineの最初の呼び出しは「Line 1 continued...」を返し、
// 2番目の呼び出しは「Line 2」を返します。
//
// 空行は継続されません。
func (r *Reader) ReadContinuedLine() (string, error)

<<<<<<< HEAD
// ReadContinuedLineBytesは、ReadContinuedLineと同様ですが、
// 文字列ではなく[]byteを返します。
=======
// ReadContinuedLineBytes is like [Reader.ReadContinuedLine] but
// returns a []byte instead of a string.
>>>>>>> upstream/release-branch.go1.22
func (r *Reader) ReadContinuedLineBytes() ([]byte, error)

// ReadCodeLineは、下記の形式の応答コード行を読み取ります：
//
//	code message
//
// ここで、codeは3桁のステータスコードであり、messageは行の残りの部分に拡張されます。
// このような行の例は次の通りです：
//
//	220 plan9.bell-labs.com ESMTP
//
// もしステータスのプレフィックスがexpectCodeの数字と一致しない場合、ReadCodeLineはerrを&Error{code, message}に設定して返します。
// 例えば、expectCodeが31である場合、ステータスが[310,319]の範囲にない場合はエラーが返されます。
//
// もし応答が複数行の場合、ReadCodeLineはエラーを返します。
//
// expectCodeが0以下の場合、ステータスコードのチェックは無効になります。
func (r *Reader) ReadCodeLine(expectCode int) (code int, message string, err error)

// ReadResponseは以下の形式の複数行のレスポンスを読み込みます：
//
//	code-message 行1
//	code-message 行2
//	...
//	code message 行n
//
// ここで、codeは3桁のステータスコードです。最初の行はcodeとハイフンで始まります。
// レスポンスは、同じcodeの後にスペースが続く行で終了します。
// メッセージ中の各行は改行（\n）で区切られます。
//
// 別の形式のレスポンスの詳細については、RFC 959（https://www.ietf.org/rfc/rfc959.txt）の
// 36ページを参照してください：
//
//	code-message 行1
//	message 行2
//	...
//	code message 行n
//
// ステータスのプレフィックスがexpectCodeの数字と一致しない場合、
// ReadResponseはerrを設定した&Error{code, message}として返されます。
// たとえば、expectCodeが31の場合、ステータスが[310、319]の範囲内でない場合、エラーが返されます。
//
// expectCode <= 0の場合、ステータスコードのチェックは無効になります。
func (r *Reader) ReadResponse(expectCode int) (code int, message string, err error)

<<<<<<< HEAD
// DotReaderは、rから読み込まれたドットエンコードされたブロックのデコードされたテキストを使用して、
// Readsを満たす新しいReaderを返します。
// 返されたReaderは、次にrのメソッドが呼び出されるまでの間のみ有効です。
=======
// DotReader returns a new [Reader] that satisfies Reads using the
// decoded text of a dot-encoded block read from r.
// The returned Reader is only valid until the next call
// to a method on r.
>>>>>>> upstream/release-branch.go1.22
//
// ドットエンコーディングは、SMTPなどのテキストプロトコルで使用される一般的なフレーミングです。
// データは、各行が"\r\n"で終わるシーケンスです。シーケンス自体は、単独のドット「.」の行で終了します：".\r\n"。
// ドットで始まる行は、シーケンスの終わりのように見えないように追加のドットでエスケープされます。
//
<<<<<<< HEAD
// ReaderのReadメソッドによって返されるデコードされた形式は、"\r\n"の行末をよりシンプルな"\n"に書き換え、
// 先頭のドットエスケープを削除し、シーケンスの終了行を消費（および破棄）した後にエラーio.EOFで停止します。
=======
// The decoded form returned by the Reader's Read method
// rewrites the "\r\n" line endings into the simpler "\n",
// removes leading dot escapes if present, and stops with error [io.EOF]
// after consuming (and discarding) the end-of-sequence line.
>>>>>>> upstream/release-branch.go1.22
func (r *Reader) DotReader() io.Reader

// ReadDotBytesはドットエンコーディングを読み込み、デコードされたデータを返します。
//
<<<<<<< HEAD
// ドットエンコーディングの詳細については、DotReaderメソッドのドキュメントを参照してください。
=======
// See the documentation for the [Reader.DotReader] method for details about dot-encoding.
>>>>>>> upstream/release-branch.go1.22
func (r *Reader) ReadDotBytes() ([]byte, error)

// ReadDotLines関数はドットエンコーディングを読み取り、各行から最後の\r\nまたは\nを省いたデコードされたスライスを返します。
//
<<<<<<< HEAD
// dot-encodingの詳細についてはDotReaderメソッドのドキュメントを参照してください。
func (r *Reader) ReadDotLines() ([]string, error)

// ReadMIMEHeaderはrからMIME形式のヘッダーを読み取ります。
// ヘッダーは、連続したキー：値行のシーケンスで、
// 空行で終わる可能性があります。
// 返されるマップmは、CanonicalMIMEHeaderKey(key)をキーとし、
// 入力で遭遇した順に値のシーケンスをマッピングします。
=======
// See the documentation for the [Reader.DotReader] method for details about dot-encoding.
func (r *Reader) ReadDotLines() ([]string, error)

// ReadMIMEHeader reads a MIME-style header from r.
// The header is a sequence of possibly continued Key: Value lines
// ending in a blank line.
// The returned map m maps [CanonicalMIMEHeaderKey](key) to a
// sequence of values in the same order encountered in the input.
>>>>>>> upstream/release-branch.go1.22
//
// 例えば、以下のような入力を考えてください：
//
//	My-Key: Value 1
//	Long-Key: Even
//	       Longer Value
//	My-Key: Value 2
//
// この入力が与えられた場合、ReadMIMEHeaderは以下のマップを返します：
//
//	map[string][]string{
//		"My-Key": {"Value 1", "Value 2"},
//		"Long-Key": {"Even Longer Value"},
//	}
func (r *Reader) ReadMIMEHeader() (MIMEHeader, error)

// CanonicalMIMEHeaderKeyは、MIMEヘッダーキーsの正準形を返します。正準化は、
// 最初の文字とハイフンの後に続くすべての文字を大文字に変換し、
// 残りの文字は小文字に変換します。たとえば、
// "accept-encoding"の正規キーは"Accept-Encoding"です。
// MIMEヘッダーキーはASCIIのみとします。
// sにスペースや無効なヘッダーフィールドバイトが含まれている場合、
// 変更せずに返されます。
func CanonicalMIMEHeaderKey(s string) string
