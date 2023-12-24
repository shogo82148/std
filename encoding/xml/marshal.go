// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/reflect"
)

const (
	// Headerは、Marshalの出力と一緒に使用するのに適した一般的なXMLヘッダーです。
	// これはこのパッケージの出力に自動的に追加されるものではなく、便宜上提供されています。
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

// Marshalは、vのXMLエンコーディングを返します。
//
// Marshalは、配列またはスライスを処理するために、各要素をマーシャリングします。
// Marshalは、ポインタが指す値をマーシャリングするか、ポインタがnilの場合は何も書き込まないことで、ポインタを処理します。
// Marshalは、インターフェース値が含む値をマーシャリングするか、インターフェース値がnilの場合は何も書き込まないことで、インターフェース値を処理します。
// Marshalは、その他のすべてのデータを処理するために、データを含む1つ以上のXML要素を書き込みます。
//
// XML要素の名前は、以下の優先順位で取得されます：
//   - データが構造体の場合、XMLNameフィールドのタグ
//   - Name型のXMLNameフィールドの値
//   - データを取得するために使用された構造体フィールドのタグ
//   - データを取得するために使用された構造体フィールドの名前
//   - マーシャルされた型の名前
//
// 構造体のXML要素には、構造体のエクスポートされた各フィールドのマーシャルされた要素が含まれますが、以下の例外があります：
//   - 上記で説明したXMLNameフィールドは省略されます。
//   - タグ "-" のフィールドは省略されます。
//   - タグ "name,attr" のフィールドは、XML要素内で指定された名前の属性になります。
//   - タグ ",attr" のフィールドは、XML要素内でフィールド名の属性になります。
//   - タグ ",chardata" のフィールドは、文字データとして書き込まれ、XML要素としては書き込まれません。
//   - タグ ",cdata" のフィールドは、<![CDATA[ ... ]]>タグで囲まれた文字データとして書き込まれ、XML要素としては書き込まれません。
//   - タグ ",innerxml" のフィールドは、通常のマーシャリング手順に従わず、そのまま書き込まれます。
//   - タグ ",comment" のフィールドは、通常のマーシャリング手順に従わず、XMLコメントとして書き込まれます。これには "--" 文字列を含めることはできません。
//   - "omitempty" オプションを含むタグのフィールドは、フィールド値が空の場合に省略されます。空の値は false、0、nil ポインタまたはインターフェース値、長さゼロの配列、スライス、マップ、文字列です。
//   - 匿名の構造体フィールドは、その値のフィールドが外部の構造体の一部であるかのように処理されます。
//   - Marshalerを実装するフィールドは、そのMarshalXMLメソッドを呼び出して書き込まれます。
//   - encoding.TextMarshalerを実装するフィールドは、そのMarshalTextメソッドの結果をテキストとしてエンコードして書き込まれます。
//
// フィールドがタグ "a>b>c" を使用する場合、要素cは親要素aとbの内部にネストされます。
// 同じ親を名指す隣接するフィールドは、1つのXML要素内に囲まれます。
//
// 構造体フィールドのXML名がフィールドタグと構造体のXMLNameフィールドの両方によって定義されている場合、
// 名前は一致しなければなりません。
//
// 例については、MarshalIndentを参照してください。
//
// Marshalは、チャネル、関数、またはマップをマーシャルするように求められた場合、エラーを返します。
func Marshal(v any) ([]byte, error)

// Marshalerは、自身を有効なXML要素にマーシャルできるオブジェクトが実装するインターフェースです。
//
// MarshalXMLは、レシーバをゼロ個以上のXML要素としてエンコードします。
// 通常、配列やスライスは、エントリごとに一つの要素としてエンコードされます。
// startを要素タグとして使用することは必須ではありませんが、そうすることで
// UnmarshalがXML要素を正しい構造体フィールドにマッチさせることができます。
// 一般的な実装戦略の一つは、所望のXMLに対応するレイアウトを持つ別の
// 値を構築し、それをe.EncodeElementを使用してエンコードすることです。
// もう一つの一般的な戦略は、e.EncodeTokenを繰り返し呼び出して、
// XML出力を一つずつトークンとして生成することです。
// エンコードされたトークンのシーケンスは、ゼロ個以上の有効な
// XML要素を構成しなければなりません。
type Marshaler interface {
	MarshalXML(e *Encoder, start StartElement) error
}

// MarshalerAttrは、自身を有効なXML属性にマーシャルできるオブジェクトが実装するインターフェースです。
//
// MarshalXMLAttrは、レシーバのエンコードされた値を持つXML属性を返します。
// 属性名としてnameを使用することは必須ではありませんが、そうすることで
// Unmarshalが属性を正しい構造体フィールドにマッチさせることができます。
// MarshalXMLAttrがゼロ属性Attr{}を返す場合、出力には属性が生成されません。
// MarshalXMLAttrは、フィールドタグに"attr"オプションを持つ構造体フィールドのみで使用されます。
type MarshalerAttr interface {
	MarshalXMLAttr(name Name) (Attr, error)
}

// MarshalIndentはMarshalと同様に動作しますが、各XML要素は新しい
// インデントされた行から始まり、その行はprefixで始まり、ネストの深さに応じて
// indentの一つ以上のコピーに続きます。
func MarshalIndent(v any, prefix, indent string) ([]byte, error)

// Encoderは、XMLデータを出力ストリームに書き込みます。
type Encoder struct {
	p printer
}

// NewEncoderは、wに書き込む新しいエンコーダを返します。
func NewEncoder(w io.Writer) *Encoder

// Indentは、エンコーダを設定して、各要素が新しいインデントされた行から始まるXMLを生成します。
// その行はprefixで始まり、ネストの深さに応じてindentの一つ以上のコピーに続きます。
func (enc *Encoder) Indent(prefix, indent string)

// Encodeは、vのXMLエンコーディングをストリームに書き込みます。
//
// Goの値をXMLに変換する詳細については、Marshalのドキュメンテーションを参照してください。
//
// Encodeは、戻る前にFlushを呼び出します。
func (enc *Encoder) Encode(v any) error

// EncodeElementは、vのXMLエンコーディングをストリームに書き込みます。
// この際、エンコーディングの最も外側のタグとしてstartを使用します。
//
// Goの値をXMLに変換する詳細については、Marshalのドキュメンテーションを参照してください。
//
// EncodeElementは、戻る前にFlushを呼び出します。
func (enc *Encoder) EncodeElement(v any, start StartElement) error

// EncodeTokenは、与えられたXMLトークンをストリームに書き込みます。
// StartElementとEndElementトークンが適切にマッチしていない場合、エラーを返します。
//
// EncodeTokenはFlushを呼び出しません。なぜなら、通常これはEncodeやEncodeElement
// （またはそれらの間に呼び出されるカスタムMarshalerのMarshalXML）のような大きな操作の一部であり、
// それらは終了時にFlushを呼び出します。
// Encoderを作成し、EncodeやEncodeElementを使用せずに直接EncodeTokenを呼び出す呼び出し元は、
// XMLが基礎となるライターに書き込まれることを確認するために、終了時にFlushを呼び出す必要があります。
//
// EncodeTokenは、"xml"をTargetに設定したProcInstを、ストリームの最初のトークンとしてのみ書き込むことを許可します。
func (enc *Encoder) EncodeToken(t Token) error

// Flushは、バッファリングされたXMLを基礎となるライターにフラッシュします。
// いつ必要かについての詳細は、EncodeTokenのドキュメンテーションを参照してください。
func (enc *Encoder) Flush() error

// エンコーダを閉じます。これは、これ以上データが書き込まれないことを示します。
// バッファリングされたXMLを基礎となるライターにフラッシュし、
// 書き込まれたXMLが無効である場合（例えば、閉じられていない要素を含む場合）にエラーを返します。
func (enc *Encoder) Close() error

// UnsupportedTypeErrorは、MarshalがXMLに変換できないタイプに遭遇したときに返されます。
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string
