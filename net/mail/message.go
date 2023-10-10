// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
パッケージメールは、メールメッセージの解析を実装しています。

このパッケージのほとんどは、RFC 5322およびRFC 6532で指定された構文に従います。
主な異なる点：
  - 廃止されたアドレス形式は解析されません。これには、経路情報を埋め込んだアドレスも含まれます。
  - スペーシング（CFWS構文要素）の完全範囲はサポートされておらず、アドレスを複数行にわたって分割することもありません。
  - Unicodeの正規化は行われません。
  - 特殊文字（）[]：; @ \は、名前でクォートされずに使用することが許可されています。
  - 先頭のFrom行は、mbox形式（RFC 4155）と同様に許可されています。
*/
package mail

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/mime"
	"github.com/shogo82148/std/time"
)

// メッセージは解析されたメールメッセージを表します。
type Message struct {
	Header Header
	Body   io.Reader
}

// ReadMessageはrからメッセージを読み取ります。
// ヘッダーが解析され、メッセージの本文はmsg.Bodyから読み取れるようになります。
func ReadMessage(r io.Reader) (msg *Message, err error)

// ParseDateはRFC 5322形式の日付文字列を解析します。
func ParseDate(date string) (time.Time, error)

// Headerはメールメッセージヘッダー内のキーと値のペアを表します。
type Header map[string][]string

// Getは指定されたキーに関連付けられた最初の値を取得します。
// 大文字と小文字の区別はなく、CanonicalMIMEHeaderKeyが使用されます。
// キーに関連付けられた値がない場合、Getは "" を返します。
// キーの複数の値にアクセスする場合や、正準形でないキーを使用する場合は、
// マップに直接アクセスしてください。
func (h Header) Get(key string) string

var ErrHeaderNotPresent = errors.New("mail: header not in message")

// Date ヘッダーフィールドをパースします。
func (h Header) Date() (time.Time, error)

// AddressListは指定されたヘッダーフィールドをアドレスのリストとして解析します。
func (h Header) AddressList(key string) ([]*Address, error)

// Addressは1つのメールアドレスを表します。
// "Barry Gibbs <bg@example.com>"のようなアドレスは、
// Address{Name: "Barry Gibbs", Address: "bg@example.com"}として表されます。
type Address struct {
	Name    string
	Address string
}

// ParseAddressは単一のRFC 5322アドレスを解析します。例：「Barry Gibbs <bg@example.com>」
func ParseAddress(address string) (*Address, error)

// ParseAddressListは与えられた文字列をアドレスのリストとして解析します。
func ParseAddressList(list string) ([]*Address, error)

// AddressParserはRFC 5322形式のアドレスパーサーです。
type AddressParser struct {
	// WordDecoderはRFC 2047でエンコードされた単語のデコーダをオプションで指定します。
	WordDecoder *mime.WordDecoder
}

// Parseは" Gogh Fir <gf@example.com>"または"foo@example.com"という形式の単一のRFC 5322アドレスを解析します。
func (p *AddressParser) Parse(address string) (*Address, error)

// ParseListは与えられた文字列をコンマで区切られたアドレスのリストとして解析します。
// 形式は「Gogh Fir <gf@example.com>」または「foo@example.com」です。
func (p *AddressParser) ParseList(list string) ([]*Address, error)

// Stringはアドレスを有効なRFC 5322形式のアドレスとしてフォーマットします。
// アドレスの名前に非ASCIIの文字が含まれている場合、名前はRFC 2047に従って表示されます。
func (a *Address) String() string
