// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

import (
	"github.com/shogo82148/std/io"
)

// Decoderは生のJSONトークンや値をストリームでデコードするためのデコーダです。
// トップレベルのJSON値のストリームを読み取るために使用されます。
// 各値は任意の空白文字で区切られています。
//
// [Decoder.ReadToken]と[Decoder.ReadValue]の呼び出しは交互に行うことができます。
// 例えば、次のJSON値：
//
//	{"name":"value","array":[null,false,true,3.14159],"object":{"k":"v"}}
//
// は、以下の呼び出しでパースできます（簡略化のためエラー処理は省略しています）:
//
//	d.ReadToken() // {
//	d.ReadToken() // "name"
//	d.ReadToken() // "value"
//	d.ReadValue() // "array"
//	d.ReadToken() // [
//	d.ReadToken() // null
//	d.ReadToken() // false
//	d.ReadValue() // true
//	d.ReadToken() // 3.14159
//	d.ReadToken() // ]
//	d.ReadValue() // "object"
//	d.ReadValue() // {"k":"v"}
//	d.ReadToken() // }
//
// 上記は呼び出しの一例であり、
// 任意のトークンや値に対して最も適切な呼び出し方法を示すものではありません。
// 例えば、オブジェクト名の文字列トークンを取得するために[Decoder.ReadToken]を呼び出す方が一般的です。
type Decoder struct {
	s decoderState
}

// NewDecoderは、rから読み込む新しいストリーミングデコーダを構築します。
//
// rが [bytes.Buffer] の場合、デコーダは中間バッファに内容をコピーせず、直接バッファからパースします。
// デコーダの使用中にバッファへ追加の書き込みを行ってはいけません。
func NewDecoder(r io.Reader, opts ...Options) *Decoder

// Resetはデコーダをリセットし、新たにrから読み込み、指定されたオプションで構成します。
// Resetは [encoding/json/v2.UnmarshalerFrom.UnmarshalJSONFrom] メソッドや
// [encoding/json/v2.UnmarshalFromFunc] 関数に渡されたDecoderに対して呼び出してはいけません。
func (d *Decoder) Reset(r io.Reader, opts ...Options)

// Optionsはエンコーダの構築に使用されたオプションを返します。
// また、[encoding/json/v2.UnmarshalDecode]呼び出しに渡されたセマンティックオプションを含む場合があります。
//
// [encoding/json/v2.UnmarshalerFrom.UnmarshalJSONFrom]メソッド呼び出しや
// [encoding/json/v2.UnmarshalFromFunc]関数呼び出しの中で動作している場合、
// 返されるオプションはその呼び出しの間のみ有効です。
func (d *Decoder) Options() Options

// PeekKindは次のトークン種別を取得しますが、読み取り位置は進めません。
//
// エラーが発生した場合は0を返します。エラーは次の読み取り呼び出しまでキャッシュされ、
// 呼び出し側は最終的にPeekKindの後に読み取り呼び出しを行う責任があります。
func (d *Decoder) PeekKind() Kind

// SkipValueは、[Decoder.ReadValue] を呼び出して結果を破棄するのと意味的に同等ですが、
// 結果全体を保持するためにメモリを無駄に消費しません。
func (d *Decoder) SkipValue() error

// ReadTokenは次の [Token] を読み取り、読み取り位置を進めます。
// 返されるトークンは次のPeek、Read、Skip呼び出しまでのみ有効です。
// トークンがこれ以上ない場合は [io.EOF] を返します。
func (d *Decoder) ReadToken() (Token, error)

// ReadValueは次の生のJSON値を返し、読み取り位置を進めます。
// 値は先頭および末尾の空白が除去され、入力の正確なバイト列が含まれます。
// [AllowInvalidUTF8] が指定されている場合、不正なUTF-8を含むことがあります。
//
// 返される値は次のPeek、Read、Skip呼び出しまでのみ有効であり、
// Decoderの使用中に値を変更してはいけません。
// デコーダが現在オブジェクトや配列の終了トークン位置にある場合、
// [SyntacticError] を報告し、内部状態は変更されません。
// これ以上値がない場合は [io.EOF] を返します。
func (d *Decoder) ReadValue() (Value, error)

// InputOffsetは現在の入力バイトオフセットを返します。これは直前に返されたトークンまたは値の直後の次のバイト位置を示します。
// 実際に基礎となる [io.Reader] から読み込まれたバイト数は、内部バッファリングの影響でこのオフセットより多い場合があります。
func (d *Decoder) InputOffset() int64

// UnreadBufferは未読バッファに残っているデータを返します。
// それは0バイト以上を含む場合があります。
// 返されたバッファはDecoderの使用中に変更してはいけません。
// バッファの内容は次のPeek、Read、Skip呼び出しまで有効です。
func (d *Decoder) UnreadBuffer() []byte

// StackDepthは、読み取ったJSONデータに対する状態マシンの深さを返します。
// スタックの各レベルは入れ子になったJSONオブジェクトまたは配列を表します。
// [BeginObject] または [BeginArray] トークンが現れるたびに増加し、
// [EndObject] または [EndArray] トークンが現れるたびに減少します。
// 深さはゼロ始まりで、ゼロはトップレベルのJSON値を表します。
func (d *Decoder) StackDepth() int

// StackIndexは指定されたスタックレベルの情報を返します。
// 0から [Decoder.StackDepth] までの数値でなければなりません。
// 各レベルについて、その種別を報告します:
//
//   - 0 はゼロレベルを表します。
//   - '{' はJSONオブジェクトのレベルを表します。
//   - '[' はJSON配列のレベルを表します。
//
// また、そのJSONオブジェクトや配列の長さも報告します。
// JSONオブジェクト内の各名前と値は個別にカウントされるため、
// 実際のメンバー数は長さの半分になります。
// 完全なJSONオブジェクトは偶数の長さでなければなりません。
func (d *Decoder) StackIndex(i int) (Kind, int64)

// StackPointerは、直近に読み取った値へのJSONポインタ（RFC 6901）を返します。
func (d *Decoder) StackPointer() Pointer
