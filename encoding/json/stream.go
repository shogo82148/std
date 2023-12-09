// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"github.com/shogo82148/std/io"
)

// Decoderは、入力ストリームからJSON値を読み取り、デコードします。
type Decoder struct {
	r       io.Reader
	buf     []byte
	d       decodeState
	scanp   int
	scanned int64
	scan    scanner
	err     error

	tokenState int
	tokenStack []int
}

// NewDecoderは、rから読み取る新しいデコーダを返します。
//
// デコーダは自身のバッファリングを導入し、
// 要求されたJSON値を超えてrからデータを読み取る可能性があります。
func NewDecoder(r io.Reader) *Decoder

<<<<<<< HEAD
// UseNumber causes the Decoder to unmarshal a number into an interface{} as a
// [Number] instead of as a float64.
=======
// UseNumberは、Decoderに数値をfloat64ではなく、
// Numberとしてinterface{}にアンマーシャルするよう指示します。
>>>>>>> release-branch.go1.21
func (dec *Decoder) UseNumber()

// DisallowUnknownFieldsは、デコーダに、宛先が構造体であり、入力に宛先の
// いずれの非無視、エクスポートされたフィールドとも一致しないオブジェクトキーが含まれている場合に
// エラーを返すよう指示します。
func (dec *Decoder) DisallowUnknownFields()

// Decodeは、入力から次のJSONエンコードされた値を読み取り、
// それをvが指す値に格納します。
//
<<<<<<< HEAD
// See the documentation for [Unmarshal] for details about
// the conversion of JSON into a Go value.
func (dec *Decoder) Decode(v any) error

// Buffered returns a reader of the data remaining in the Decoder's
// buffer. The reader is valid until the next call to [Decoder.Decode].
=======
// JSONをGoの値に変換する詳細については、Unmarshalのドキュメンテーションを参照してください。
func (dec *Decoder) Decode(v any) error

// Bufferedは、Decoderのバッファに残っているデータのリーダーを返します。
// リーダーは次のDecode呼び出しまで有効です。
>>>>>>> release-branch.go1.21
func (dec *Decoder) Buffered() io.Reader

// Encoderは、JSON値を出力ストリームに書き込みます。
type Encoder struct {
	w          io.Writer
	err        error
	escapeHTML bool

	indentBuf    []byte
	indentPrefix string
	indentValue  string
}

// NewEncoderは、wに書き込む新しいエンコーダを返します。
func NewEncoder(w io.Writer) *Encoder

// Encodeは、vのJSONエンコーディングをストリームに書き込み、
// その後に改行文字を追加します。
//
<<<<<<< HEAD
// See the documentation for [Marshal] for details about the
// conversion of Go values to JSON.
=======
// Goの値をJSONに変換する詳細については、Marshalのドキュメンテーションを参照してください。
>>>>>>> release-branch.go1.21
func (enc *Encoder) Encode(v any) error

// SetIndentは、エンコーダに対して、次にエンコードされる各値を、パッケージレベルの関数Indent(dst, src, prefix, indent)で
// インデントされているかのようにフォーマットするよう指示します。
// SetIndent("", "")を呼び出すと、インデントが無効になります。
func (enc *Encoder) SetIndent(prefix, indent string)

// SetEscapeHTMLは、問題のあるHTML文字がJSONの引用符で囲まれた文字列内でエスケープされるべきかどうかを指定します。
// デフォルトの動作は、&, <, >を\u0026, \u003c, \u003eにエスケープして、
// JSONをHTMLに埋め込む際に生じる可能性のある特定の安全性問題を回避します。
//
// エスケープが出力の可読性を妨げる非HTML設定では、SetEscapeHTML(false)でこの動作を無効にします。
func (enc *Encoder) SetEscapeHTML(on bool)

<<<<<<< HEAD
// RawMessage is a raw encoded JSON value.
// It implements [Marshaler] and [Unmarshaler] and can
// be used to delay JSON decoding or precompute a JSON encoding.
=======
// RawMessageは、生のエンコードされたJSON値です。
// これはMarshalerとUnmarshalerを実装しており、
// JSONのデコードを遅延させるか、JSONのエンコードを事前に計算するために使用できます。
>>>>>>> release-branch.go1.21
type RawMessage []byte

// MarshalJSONは、mのJSONエンコーディングとしてmを返します。
func (m RawMessage) MarshalJSON() ([]byte, error)

// UnmarshalJSONは、*mをdataのコピーに設定します。
func (m *RawMessage) UnmarshalJSON(data []byte) error

var _ Marshaler = (*RawMessage)(nil)
var _ Unmarshaler = (*RawMessage)(nil)

// Tokenは、以下の型のいずれかの値を保持します:
//
<<<<<<< HEAD
//   - [Delim], for the four JSON delimiters [ ] { }
//   - bool, for JSON booleans
//   - float64, for JSON numbers
//   - [Number], for JSON numbers
//   - string, for JSON string literals
//   - nil, for JSON null
=======
//	Delim、JSONの4つの区切り文字 [ ] { } のため
//	bool、JSONのブール値のため
//	float64、JSONの数値のため
//	Number、JSONの数値のため
//	string、JSONの文字列リテラルのため
//	nil、JSONのnullのため
>>>>>>> release-branch.go1.21
type Token any

// Delimは、JSON配列またはオブジェクトの区切り文字であり、[ ] { }のいずれかです。
type Delim rune

func (d Delim) String() string

<<<<<<< HEAD
// Token returns the next JSON token in the input stream.
// At the end of the input stream, Token returns nil, [io.EOF].
=======
// Tokenは、入力ストリームの次のJSONトークンを返します。
// 入力ストリームの終わりでは、Tokenはnil, io.EOFを返します。
>>>>>>> release-branch.go1.21
//
// Tokenは、返す区切り文字[ ] { }が適切にネストされ、
// マッチしていることを保証します：もしTokenが入力で予期しない
// 区切り文字に遭遇した場合、エラーを返します。
//
<<<<<<< HEAD
// The input stream consists of basic JSON values—bool, string,
// number, and null—along with delimiters [ ] { } of type [Delim]
// to mark the start and end of arrays and objects.
// Commas and colons are elided.
=======
// 入力ストリームは、基本的なJSON値—bool, string,
// number, null—と、配列とオブジェクトの開始と終了を
// マークするための区切り文字[ ] { }のタイプDelimで構成されています。
// コンマとコロンは省略されます。
>>>>>>> release-branch.go1.21
func (dec *Decoder) Token() (Token, error)

// Moreは、解析中の現在の配列またはオブジェクトに別の要素があるかどうかを報告します。
func (dec *Decoder) More() bool

// InputOffsetは、現在のデコーダ位置の入力ストリームバイトオフセットを返します。
// オフセットは、最近返されたトークンの終わりと次のトークンの始まりの位置を示します。
func (dec *Decoder) InputOffset() int64
