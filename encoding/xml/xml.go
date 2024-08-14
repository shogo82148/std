// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// xmlパッケージは、XML名前空間を理解する
// シンプルなXML 1.0パーサーを実装します。
package xml

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/io"
)

// SyntaxErrorは、XML入力ストリームの構文エラーを表します。
type SyntaxError struct {
	Msg  string
	Line int
}

func (e *SyntaxError) Error() string

// Nameは、名前空間識別子（Space）で注釈付けされたXML名（Local）を表します。
// [Decoder.Token] によって返されるトークンでは、Space識別子は
// パースされるドキュメントで使用される短いプレフィックスではなく、
// 正規のURLとして与えられます。
type Name struct {
	Space, Local string
}

// Attrは、XML要素内の属性（Name=Value）を表します。
type Attr struct {
	Name  Name
	Value string
}

// Tokenは、次のトークンタイプのいずれかを保持するインターフェースです：
// [StartElement]、[EndElement]、[CharData]、[Comment]、[ProcInst]、または [Directive]。
type Token any

// StartElementは、XMLの開始要素を表します。
type StartElement struct {
	Name Name
	Attr []Attr
}

// Copyは、StartElementの新しいコピーを作成します。
func (e StartElement) Copy() StartElement

// Endは、対応するXML終了要素を返します。
func (e StartElement) End() EndElement

// EndElementは、XMLの終了要素を表します。
type EndElement struct {
	Name Name
}

// CharDataは、XMLエスケープシーケンスがそれらが表す文字に置き換えられた
// XML文字データ（生テキスト）を表します。
type CharData []byte

// Copyは、CharDataの新しいコピーを作成します。
func (c CharData) Copy() CharData

// Commentは、<!--comment-->の形式のXMLコメントを表します。
// バイトには、<!-- および --> のコメントマーカーは含まれません。
type Comment []byte

// Copyは、Commentの新しいコピーを作成します。
func (c Comment) Copy() Comment

// ProcInstは、<?target inst?>の形式のXML処理命令を表します。
type ProcInst struct {
	Target string
	Inst   []byte
}

// Copyは、ProcInstの新しいコピーを作成します。
func (p ProcInst) Copy() ProcInst

// Directiveは、<!text>形式のXML指示を表します。
// バイトには、<! および > のマーカーは含まれません。
type Directive []byte

// Copyは、Directiveの新しいコピーを作成します。
func (d Directive) Copy() Directive

// CopyTokenは、Tokenのコピーを返します。
func CopyToken(t Token) Token

// TokenReaderは、XMLトークンのストリームをデコードできるものを指します。
// これには、[Decoder] も含まれます。
//
// Tokenがトークンの読み取りに成功した後にエラーまたはファイル終了の状態に遭遇した場合、
// それはそのトークンを返します。それは同じ呼び出しから（非nilの）エラーを返すか、
// 次の呼び出しからエラー（とnilトークン）を返すかもしれません。
// この一般的なケースの一例は、トークンストリームの終わりで非nilのトークンを返すTokenReaderが、
// io.EOFまたはnilエラーのどちらかを返す可能性があるということです。
// 次のReadはnil, [io.EOF] を返すべきです。
//
// Tokenの実装は、nilトークンとnilエラーを返すことを推奨されていません。
// 呼び出し元はnil, nilの返り値を何も起こらなかったことを示すものとして扱うべきです。
// 特に、これはEOFを示すものではありません。
type TokenReader interface {
	Token() (Token, error)
}

// Decoderは、特定の入力ストリームを読み取るXMLパーサーを表します。
// パーサーは、その入力がUTF-8でエンコードされていると仮定します。
type Decoder struct {
	// Strictはデフォルトでtrueで、XML仕様の要件を強制します。
	// falseに設定すると、パーサーは一般的な間違いを含む入力を許可します：
	//	* 要素が終了タグを欠いている場合、パーサーは必要に応じて
	//	  終了タグを発明して、Tokenからの戻り値を適切にバランスさせます。
	//	* 属性値とキャラクターデータでは、未知または不正な
	//	  キャラクターエンティティ（&で始まるシーケンス）はそのままにされます。
	//
	// 設定：
	//
	//	d.Strict = false
	//	d.AutoClose = xml.HTMLAutoClose
	//	d.Entity = xml.HTMLEntity
	//
	// これにより、一般的なHTMLを処理できるパーサーが作成されます。
	//
	// 厳格モードでは、XML名前空間TRの要件は強制されません。
	// 特に、未定義のプレフィックスを使用する名前空間タグは拒否されません。
	// そのようなタグは、未知のプレフィックスを名前空間URLとして記録します。
	Strict bool

	// Strict == falseの場合、AutoCloseは、開かれた直後に閉じるとみなす要素のセットを示します。
	// これは、終了要素が存在するかどうかに関係なく適用されます。
	AutoClose []string

	// Entityは、非標準のエンティティ名を文字列の置換にマッピングするために使用できます。
	// パーサーは、実際のマップの内容に関係なく、これらの標準マッピングがマップに存在するかのように動作します：
	//
	//	"lt": "<",
	//	"gt": ">",
	//	"amp": "&",
	//	"apos": "'",
	//	"quot": `"`,
	Entity map[string]string

	// CharsetReaderがnilでない場合、提供された非UTF-8文字セットからUTF-8に変換する
	// 文字セット変換リーダーを生成する関数を定義します。CharsetReaderがnilであるか、
	// エラーを返す場合、パースはエラーで停止します。CharsetReaderの結果値のうちの
	// 一つは非nilでなければなりません。
	CharsetReader func(charset string, input io.Reader) (io.Reader, error)

	// DefaultSpaceは、飾り気のないタグに使用されるデフォルトの名前空間を設定します。
	// まるでXMLストリーム全体が、属性xmlns="DefaultSpace"を含む要素で
	// ラップされているかのように動作します。
	DefaultSpace string

	r              io.ByteReader
	t              TokenReader
	buf            bytes.Buffer
	saved          *bytes.Buffer
	stk            *stack
	free           *stack
	needClose      bool
	toClose        Name
	nextToken      Token
	nextByte       int
	ns             map[string]string
	err            error
	line           int
	linestart      int64
	offset         int64
	unmarshalDepth int
}

// NewDecoderは、rから読み取る新しいXMLパーサーを作成します。
// もしrが [io.ByteReader] を実装していない場合、NewDecoderは
// 自身でバッファリングを行います。
func NewDecoder(r io.Reader) *Decoder

// NewTokenDecoderは、基礎となるトークンストリームを使用して新しいXMLパーサーを作成します。
func NewTokenDecoder(t TokenReader) *Decoder

// Tokenは、入力ストリームの次のXMLトークンを返します。
// 入力ストリームの終わりでは、Tokenはnil, [io.EOF] を返します。
//
// 返されたトークンデータのバイトスライスは、パーサーの内部バッファを参照し、
// 次のTokenへの呼び出しまでのみ有効です。バイトのコピーを取得するには、
// [CopyToken] を呼び出すか、トークンのCopyメソッドを呼び出します。
//
// Tokenは、<br>のような自己閉鎖要素を展開し、
// 連続した呼び出しで返される別々の開始要素と終了要素にします。
//
// Tokenは、返される [StartElement] と [EndElement] トークンが適切にネストされ、
// マッチしていることを保証します：もしTokenが予期しない終了要素や、
// すべての予期される終了要素の前にEOFに遭遇した場合、エラーを返します。
//
// [Decoder.CharsetReader] が呼び出され、エラーを返す場合、
// そのエラーはラップされて返されます。
//
// Tokenは、https://www.w3.org/TR/REC-xml-names/ で説明されているような
// XML名前空間を実装します。Tokenに含まれる各 [Name] 構造体は、その名前空間を
// 識別するURLがわかっている場合にSpaceに設定されます。
// もしTokenが認識できない名前空間プレフィックスに遭遇した場合、
// エラーを報告する代わりにプレフィックスをSpaceとして使用します。
func (d *Decoder) Token() (Token, error)

// RawTokenはTokenと同様ですが、開始要素と終了要素が一致することを検証せず、
// 名前空間のプレフィックスを対応するURLに変換しません。
func (d *Decoder) RawToken() (Token, error)

// InputOffsetは、現在のデコーダ位置の入力ストリームバイトオフセットを返します。
// オフセットは、最近返されたトークンの終わりと次のトークンの始まりの位置を示します。
func (d *Decoder) InputOffset() int64

// InputPosは、現在のデコーダ位置の行と、行の1ベースの入力位置を返します。
// 位置は、最近返されたトークンの終わりの位置を示します。
func (d *Decoder) InputPos() (line, column int)

// HTMLEntityは、標準的なHTMLエンティティ文字の変換を含むエンティティマップです。
//
// [Decoder.Strict] と [Decoder.Entity] フィールドのドキュメンテーションを参照してください。
var HTMLEntity map[string]string = htmlEntity

// HTMLAutoCloseは、自動的に閉じるとみなすべきHTML要素のセットです。
//
// [Decoder.Strict] と [Decoder.Entity] フィールドのドキュメンテーションを参照してください。
var HTMLAutoClose []string = htmlAutoClose

// EscapeTextは、プレーンテキストデータsの適切にエスケープされたXML相当物をwに書き込みます。
func EscapeText(w io.Writer, s []byte) error

// Escapeは [EscapeText] と同様ですが、エラーの戻り値を省略します。
// これはGo 1.0との後方互換性のために提供されています。
// Go 1.1以降を対象とするコードは [EscapeText] を使用するべきです。
func Escape(w io.Writer, s []byte)
