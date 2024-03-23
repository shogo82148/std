// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

/*
パッケージmultipartは、RFC 2046で定義されているMIMEマルチパートの解析を実装します。

この実装は、HTTP（RFC 2388）と一般的なブラウザが生成するマルチパートボディに対して十分です。

# 制限

悪意のある入力に対する保護として、このパッケージは処理するMIMEデータのサイズに制限を設けています。

[Reader.NextPart] と [Reader.NextRawPart] は、パート内のヘッダーの数を10000に制限し、[Reader.ReadForm] はすべての
FileHeaders内のヘッダーの合計数を10000に制限します。
これらの制限は、GODEBUG=multipartmaxheaders=<values>の設定で調整できます。

さらに、Reader.ReadFormはフォーム内のパートの数を1000に制限します。
この制限は、GODEBUG=multipartmaxparts=<value>の設定で調整できます。
*/
package multipart

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/textproto"
)

// Partは、マルチパートボディの単一の部分を表します。
type Part struct {
	// ボディのヘッダー（存在する場合）は、Goのhttp.Requestヘッダーと同様に
	// キーが正規化されています。例えば、"foo-bar"は"Foo-Bar"に変更されます。
	Header textproto.MIMEHeader

	mr *Reader

	disposition       string
	dispositionParams map[string]string

	// r is either a reader directly reading from mr, or it's a
	// wrapper around such a reader, decoding the
	// Content-Transfer-Encoding
	r io.Reader

	n       int
	total   int64
	err     error
	readErr error
}

// FormNameは、pのContent-Dispositionが"type"の"form-data"である場合、
// nameパラメータを返します。それ以外の場合は空文字列を返します。
func (p *Part) FormName() string

// FileNameは、[Part] のContent-Dispositionヘッダーのfilenameパラメータを返します。
// 空でない場合、filenameはfilepath.Base（プラットフォーム依存）を通過してから返されます。
func (p *Part) FileName() string

// NewReaderは、指定されたMIME境界を使用してrから読み取る新しいマルチパート [Reader] を作成します。
//
// 境界は通常、メッセージの"Content-Type"ヘッダーの"boundary"パラメータから取得します。
// そのようなヘッダーを解析するには、[mime.ParseMediaType] を使用します。
func NewReader(r io.Reader, boundary string) *Reader

// Readは、ヘッダーの後と次のパート（存在する場合）が始まる前の、パートのボディを読み取ります。
func (p *Part) Read(d []byte) (n int, err error)

func (p *Part) Close() error

// Readerは、MIMEマルチパートボディ内のパートを反復処理するためのものです。
// Readerの基礎となるパーサーは、必要に応じて入力を消費します。シークはサポートされていません。
type Reader struct {
	bufReader *bufio.Reader
	tempDir   string

	currentPart *Part
	partsRead   int

	nl               []byte
	nlDashBoundary   []byte
	dashBoundaryDash []byte
	dashBoundary     []byte
}

// NextPartは、マルチパートの次のパートまたはエラーを返します。
// パートがこれ以上ない場合、エラー [io.EOF] が返されます。
//
// 特別なケースとして、"Content-Transfer-Encoding"ヘッダーの値が
// "quoted-printable"である場合、そのヘッダーは代わりに隠され、
// ボディはRead呼び出し中に透明にデコードされます。
func (r *Reader) NextPart() (*Part, error)

// NextRawPartは、マルチパートの次のパートまたはエラーを返します。
// パートがこれ以上ない場合、エラー [io.EOF] が返されます。
//
// [Reader.NextPart] とは異なり、"Content-Transfer-Encoding: quoted-printable"に対する特別な処理はありません。
func (r *Reader) NextRawPart() (*Part, error)
