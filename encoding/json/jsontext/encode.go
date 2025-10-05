// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

import (
	"github.com/shogo82148/std/io"
)

// Encoderは生のJSONトークンや値をストリームでエンコードするためのエンコーダです。
// トップレベルのJSON値のストリームを書き込むために使用され、
// 各値は改行文字で区切られます。
//
// [Encoder.WriteToken]と[Encoder.WriteValue]の呼び出しは交互に行うことができます。
// 例えば、次のJSON値：
//
//	{"name":"value","array":[null,false,true,3.14159],"object":{"k":"v"}}
//
// は、以下の呼び出しで構成できます（簡略化のためエラー処理は省略しています）:
//
//	e.WriteToken(BeginObject)        // {
//	e.WriteToken(String("name"))     // "name"
//	e.WriteToken(String("value"))    // "value"
//	e.WriteValue(Value(`"array"`))   // "array"
//	e.WriteToken(BeginArray)         // [
//	e.WriteToken(Null)               // null
//	e.WriteToken(False)              // false
//	e.WriteValue(Value("true"))      // true
//	e.WriteToken(Float(3.14159))     // 3.14159
//	e.WriteToken(EndArray)           // ]
//	e.WriteValue(Value(`"object"`))  // "object"
//	e.WriteValue(Value(`{"k":"v"}`)) // {"k":"v"}
//	e.WriteToken(EndObject)          // }
//
// 上記は呼び出しの一例であり、
// 任意のトークンや値に対して最も適切な呼び出し方法を示すものではありません。
// 例えば、オブジェクト名の文字列トークンを取得するために[Encoder.WriteToken]を呼び出す方が一般的です。
type Encoder struct {
	s encoderState
}

// NewEncoderは、wに書き込む新しいストリーミングエンコーダを構築し、
// 指定されたオプションで構成します。
// バッファが十分に満杯になったときや、トップレベルの値が書き込まれたときに
// 内部バッファをフラッシュします。
//
// wが [bytes.Buffer] の場合、エンコーダは中間バッファから内容をコピーせず、
// 直接バッファに追加します。
func NewEncoder(w io.Writer, opts ...Options) *Encoder

// Resetはエンコーダをリセットし、新たにwに書き込み、指定されたオプションで構成します。
// Resetは [encoding/json/v2.MarshalerTo.MarshalJSONTo] メソッドや
// [encoding/json/v2.MarshalToFunc] 関数に渡されたEncoderに対して呼び出してはいけません。
func (e *Encoder) Reset(w io.Writer, opts ...Options)

// Optionsはデコーダの構築に使用されたオプションを返します。
// また、[encoding/json/v2.MarshalEncode] 呼び出しに渡されたセマンティックオプションを含む場合があります。
//
// [encoding/json/v2.MarshalerTo.MarshalJSONTo] メソッド呼び出しや
// [encoding/json/v2.MarshalToFunc] 関数呼び出しの中で動作している場合、
// 返されるオプションはその呼び出しの間のみ有効です。
func (e *Encoder) Options() Options

// WriteTokenは次のトークンを書き込み、内部の書き込みオフセットを進めます。
//
// 渡されたトークン種別はJSON文法と一致していなければなりません。
// 例えば、エンコーダがオブジェクト名（常に文字列）を期待しているときに数値を渡したり、
// 配列の終了処理中にオブジェクトの終了デリミタを渡すとエラーになります。
// 無効なトークンが渡された場合、[SyntacticError] を報告し、
// 内部状態は変更されません。[SyntacticError] で報告されるオフセットは
// [Encoder.OutputOffset] を基準とします。
func (e *Encoder) WriteToken(t Token) error

// WriteValueは次の生の値を書き込み、内部の書き込みオフセットを進めます。
// Encoderは渡された値を単純にそのままコピーするのではなく、
// 構文的に有効であることを確認するためにパースし、
// 空白や文字列のフォーマット方法に従って再フォーマットします。
// [AllowInvalidUTF8] が指定されている場合、不正なUTF-8はUnicodeの置換文字U+FFFDに変換されます。
//
// 渡された値の種別はJSON文法と一致していなければなりません
// （[Encoder.WriteToken] の例を参照）。値が無効な場合、[SyntacticError] を報告し、
// 内部状態は変更されません。[SyntacticError] で報告されるオフセットは
// [Encoder.OutputOffset] に加えて、v内で構文エラーが発生した位置となります。
func (e *Encoder) WriteValue(v Value) error

// OutputOffsetは現在の出力バイトオフセットを返します。これは直近に書き込まれたトークンまたは値の直後の次のバイト位置を示します。
// 実際に基礎となる [io.Writer] に書き込まれたバイト数は、内部バッファリングの影響でこのオフセットより少ない場合があります。
func (e *Encoder) OutputOffset() int64

// AvailableBufferは長さ0で容量が0でない可能性があるバッファを返します。
// このバッファは、直後に [Encoder.WriteValue] 呼び出しで渡す [Value] の生成に使用することを意図しています。
//
// 使用例:
//
//	b := d.AvailableBuffer()
//	b = append(b, '"')
//	b = appendString(b, v) // vの文字列フォーマットを追加
//	b = append(b, '"')
//	... := d.WriteValue(b)
//
// 値が有効なJSONであることは利用者の責任です。
func (e *Encoder) AvailableBuffer() []byte

// StackDepthは、書き込まれたJSONデータに対する状態マシンの深さを返します。
// スタックの各レベルは入れ子になったJSONオブジェクトまたは配列を表します。
// [BeginObject] または [BeginArray] トークンが現れるたびに増加し、
// [EndObject] または [EndArray] トークンが現れるたびに減少します。
// 深さはゼロ始まりで、ゼロはトップレベルのJSON値を表します。
func (e *Encoder) StackDepth() int

// StackIndexは指定されたスタックレベルの情報を返します。
// 0から [Encoder.StackDepth] までの数値でなければなりません。
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
func (e *Encoder) StackIndex(i int) (Kind, int64)

// StackPointerは、直近に書き込まれた値へのJSONポインタ（RFC 6901）を返します。
func (e *Encoder) StackPointer() Pointer
