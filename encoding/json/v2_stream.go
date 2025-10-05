// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/encoding/json/jsontext"
)

// Decoderは入力ストリームからJSON値を読み込み、デコードします。
type Decoder struct {
	dec  *jsontext.Decoder
	opts jsonv2.Options
	err  error
}

// NewDecoderは、rから読み込む新しいデコーダを返します。
//
// デコーダは独自にバッファリングを行い、要求されたJSON値以外のデータもrから読み込む場合があります。
func NewDecoder(r io.Reader) *Decoder

// UseNumberは、Decoderが数値をfloat64ではなく [Number] としてinterface値にアンマーシャルするようにします。
func (dec *Decoder) UseNumber()

// DisallowUnknownFieldsは、デコード先が構造体の場合に、入力に無視されないエクスポートされたフィールドと一致しないオブジェクトキーが含まれているとエラーを返すようにDecoderに指示します。
func (dec *Decoder) DisallowUnknownFields()

// Decodeは、次のJSONエンコードされた値を入力から読み込み、vが指す値に格納します。
//
// JSONからGo値への変換の詳細については [Unmarshal] のドキュメントを参照してください。
func (dec *Decoder) Decode(v any) error

// Bufferedは、Decoderのバッファに残っているデータのリーダーを返します。
// このリーダーは次に [Decoder.Decode] が呼ばれるまで有効です。
func (dec *Decoder) Buffered() io.Reader

// Encoderは、出力ストリームにJSON値を書き込みます。
type Encoder struct {
	w    io.Writer
	opts jsonv2.Options
	err  error

	buf       bytes.Buffer
	indentBuf bytes.Buffer

	indentPrefix string
	indentValue  string
}

// NewEncoderは、wに書き込む新しいエンコーダを返します。
func NewEncoder(w io.Writer) *Encoder

// Encodeは、vのJSONエンコーディングをストリームに書き込み、
// 続けて改行文字を書き込みます。
//
// Go値からJSONへの変換の詳細については[Marshal]のドキュメントを参照してください。
func (enc *Encoder) Encode(v any) error

// SetIndentは、エンコーダに対して以降にエンコードされる値を、パッケージレベル関数Indent(dst, src, prefix, indent)でインデントされたかのように整形するよう指示します。
// SetIndent("", "")を呼び出すとインデントが無効になります。
func (enc *Encoder) SetIndent(prefix, indent string)

// SetEscapeHTMLは、JSONの引用符付き文字列内で問題となるHTML文字をエスケープするかどうかを指定します。
// デフォルトの動作では、&, <, > をそれぞれ \u0026, \u003c, \u003e にエスケープし、
// HTMLにJSONを埋め込む際に発生しうる安全上の問題を回避します。
//
// HTML以外の用途でエスケープが可読性を損なう場合は、SetEscapeHTML(false)でこの動作を無効化できます。
func (enc *Encoder) SetEscapeHTML(on bool)

// RawMessageは生のエンコード済みJSON値です。
// [Marshaler] と [Unmarshaler] を実装しており、
// JSONデコードを遅延させたり、事前にJSONエンコードを計算するために利用できます。
type RawMessage = jsontext.Value

// Tokenは、次のいずれかの型の値を保持します:
//
//   - [Delim]（4つのJSON区切り記号 [ ] { }）
//   - bool（JSONの真偽値）
//   - float64（JSONの数値）
//   - [Number]（JSONの数値）
//   - string（JSONの文字列リテラル）
//   - nil（JSONのnull）
type Token any

// DelimはJSON配列またはオブジェクトの区切り記号で、[ ] { } のいずれかです。
type Delim rune

func (d Delim) String() string

// Tokenは入力ストリームから次のJSONトークンを返します。
// 入力ストリームの終端では、Tokenはnilまたは[io.EOF]を返します。
//
// Tokenは返す区切り記号[ ] { }が正しくネストされ対応していることを保証します。
// 入力中に予期しない区切り記号が現れた場合、エラーを返します。
//
// 入力ストリームは基本的なJSON値（bool, string, number, null）と、配列やオブジェクトの開始・終了を示す [Delim] 型の区切り記号[ ] { }で構成されます。
// カンマやコロンは省略されます。
func (dec *Decoder) Token() (Token, error)

// Moreは、現在パース中の配列またはオブジェクトに次の要素が存在するかどうかを報告します。
func (dec *Decoder) More() bool

// InputOffsetは、現在のデコーダ位置の入力ストリームのバイトオフセットを返します。
// オフセットは直前に返されたトークンの終了位置と、次のトークンの開始位置を示します。
func (dec *Decoder) InputOffset() int64
