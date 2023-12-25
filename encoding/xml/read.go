// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

// Unmarshalは、XMLエンコードされたデータを解析し、結果を
// vが指す値に格納します。vは任意の構造体、スライス、または文字列でなければなりません。
// vに収まらない形式の良いデータは破棄されます。
//
// Unmarshalはreflectパッケージを使用するため、エクスポートされた（大文字の）フィールドにのみ割り当てることができます。
// Unmarshalは、XML要素名をタグ値と構造体フィールド名にマッチさせるために、大文字と小文字を区別する比較を使用します。
//
// Unmarshalは、以下のルールを使用してXML要素を構造体にマップします。
// ルールでは、フィールドのタグは、構造体フィールドのタグに関連付けられた
// 'xml'キーの値を指します（上記の例を参照してください）。
//
//   - 構造体がタグが",innerxml"の[]byte型またはstring型のフィールドを持つ場合、
//     Unmarshalはそのフィールドに要素内にネストされた生のXMLを蓄積します。
//     他のルールは依然として適用されます。
//
//   - 構造体がName型のフィールドXMLNameを持つ場合、
//     Unmarshalはそのフィールドに要素名を記録します。
//
//   - XMLNameフィールドが"名前"または"名前空間-URL 名前"の形式の関連タグを持つ場合、
//     XML要素は指定された名前（およびオプションで名前空間）を持たなければならず、
//     そうでない場合、Unmarshalはエラーを返します。
//
//   - XML要素が、",attr"を含む関連タグを持つ構造体フィールド名と一致する名前の属性、
//     または"名前,attr"の形式の構造体フィールドタグの明示的な名前を持つ場合、
//     Unmarshalはそのフィールドに属性値を記録します。
//
//   - XML要素が前のルールで処理されない属性を持ち、
//     構造体が",any,attr"を含む関連タグを持つフィールドを持つ場合、
//     Unmarshalは最初のそのようなフィールドに属性値を記録します。
//
//   - XML要素が文字データを含む場合、そのデータは
//     タグが",chardata"の最初の構造体フィールドに蓄積されます。
//     構造体フィールドは[]byte型またはstring型を持つことができます。
//     そのようなフィールドがない場合、文字データは破棄されます。
//
//   - XML要素がコメントを含む場合、それらは
//     タグが",comment"の最初の構造体フィールドに蓄積されます。
//     構造体フィールドは[]byte型またはstring型を持つことができます。
//     そのようなフィールドがない場合、コメントは破棄されます。
//
//   - XML要素が、タグが"a"または"a>b>c"の形式のプレフィックスと一致する名前のサブ要素を含む場合、
//     Unmarshalは指定された名前を持つ要素を探してXML構造に降りていき、
//     最も内側の要素をその構造体フィールドにマップします。
//     ">"で始まるタグは、フィールド名に続く">"で始まるタグと同等です。
//
//   - XML要素が、名前が構造体フィールドのXMLNameタグと一致し、
//     前のルールに従って明示的な名前タグを持たないサブ要素を含む場合、
//     Unmarshalはそのサブ要素をその構造体フィールドにマップします。
//
//   - XML要素が、モードフラグ（",attr", ",chardata"など）を持たないフィールド名と一致する
//     サブ要素を含む場合、Unmarshalはそのサブ要素をその構造体フィールドにマップします。
//
//   - XML要素が、上記のルールのいずれにも一致しないサブ要素を含み、
//     構造体がタグ",any"のフィールドを持つ場合、Unmarshalはそのサブ要素をその構造体フィールドにマップします。
//
//   - 匿名の構造体フィールドは、その値のフィールドが外部の構造体の一部であるかのように処理されます。
//
//   - タグ"-"を持つ構造体フィールドは、決してアンマーシャルされません。
//
// UnmarshalがUnmarshalerインターフェースを実装するフィールドタイプに遭遇した場合、
// UnmarshalはそのUnmarshalXMLメソッドを呼び出してXML要素から値を生成します。
// それ以外の場合、値が [encoding.TextUnmarshaler] を実装している場合、
// Unmarshalはその値のUnmarshalTextメソッドを呼び出します。
//
// Unmarshalは、XML要素をstringまたは[]byteにマップします。これは、
// その要素の文字データの連結をstringまたは[]byteに保存することで行います。
// 保存された[]byteは決してnilになりません。
//
// Unmarshalは、属性値をstringまたは[]byteにマップします。これは、
// 値をstringまたはスライスに保存することで行います。
//
// Unmarshalは、属性値を [Attr] にマップします。これは、
// 名前を含む属性をAttrに保存することで行います。
//
// Unmarshalは、スライスの長さを拡張し、要素または属性を新しく作成された値にマッピングすることで、
// XML要素または属性値をスライスにマッピングします。
//
// Unmarshalは、XML要素または属性値をboolにマッピングします。
// これは、文字列で表されるブール値に設定することで行います。空白はトリムされ、無視されます。
//
// Unmarshalは、フィールドを文字列値を10進数で解釈した結果に設定することで、
// XML要素または属性値を整数または浮動小数点フィールドにマッピングします。
// オーバーフローのチェックはありません。空白はトリムされ、無視されます。
//
// Unmarshalは、要素名を記録することで、XML要素をNameにマッピングします。
//
// Unmarshalは、ポインタを新しく割り当てられた値に設定し、その値に要素をマッピングすることで、
// XML要素をポインタにマッピングします。
//
// 要素が欠落しているか、属性値が空の場合、ゼロ値としてアンマーシャルされます。
// フィールドがスライスの場合、ゼロ値がフィールドに追加されます。それ以外の場合、
// フィールドはそのゼロ値に設定されます。
func Unmarshal(data []byte, v any) error

// Decodeは [Unmarshal] と同様に動作しますが、開始要素を見つけるためにデコーダストリームを読みます。
func (d *Decoder) Decode(v any) error

// DecodeElementは [Unmarshal] と同様に動作しますが、
// vにデコードする開始XML要素へのポインタを取ります。
// クライアントが自身でいくつかの生のXMLトークンを読み込むが、
// 一部の要素については [Unmarshal] に委ねたい場合に便利です。
func (d *Decoder) DecodeElement(v any, start *StartElement) error

// UnmarshalErrorは、アンマーシャル処理中のエラーを表します。
type UnmarshalError string

func (e UnmarshalError) Error() string

// Unmarshalerは、自分自身のXML要素の説明をアンマーシャルできるオブジェクトが実装するインターフェースです。
//
// UnmarshalXMLは、与えられた開始要素で始まる単一のXML要素をデコードします。
// エラーを返す場合、外部のUnmarshalへの呼び出しは停止し、
// そのエラーを返します。
// UnmarshalXMLは正確に一つのXML要素を消費しなければなりません。
// 一般的な実装戦略の一つは、期待されるXMLに一致するレイアウトを持つ
// 別の値にアンマーシャルし、そのデータをレシーバにコピーすることです。
// もう一つの一般的な戦略は、d.Tokenを使用してXMLオブジェクトを
// 一つずつトークンで処理することです。
// UnmarshalXMLはd.RawTokenを使用してはなりません。
type Unmarshaler interface {
	UnmarshalXML(d *Decoder, start StartElement) error
}

// UnmarshalerAttrは、自分自身のXML属性の説明をアンマーシャルできるオブジェクトが実装するインターフェースです。
//
// UnmarshalXMLAttrは単一のXML属性をデコードします。
// エラーを返す場合、外部の [Unmarshal] への呼び出しは停止し、
// そのエラーを返します。
// UnmarshalXMLAttrは、フィールドタグに"attr"オプションを持つ構造体フィールドのみで使用されます。
type UnmarshalerAttr interface {
	UnmarshalXMLAttr(attr Attr) error
}

// Skipは、最も最近消費された開始要素に一致する終了要素を消費するまでトークンを読み込みます。
// ネストされた構造はスキップされます。
// 開始要素に一致する終了要素を見つけた場合、nilを返します。
// それ以外の場合は、問題を説明するエラーを返します。
func (d *Decoder) Skip() error
