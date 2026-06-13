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
// 上記は多数ある呼び出し順の一例であり、
// 任意のトークン/値に対して最も適切な呼び出し方法を示すものではありません。
// 例えば、オブジェクト名の文字列トークンを取得するには、
// [Decoder.ReadToken] を呼び出すほうが一般的です。
type Decoder struct {
	s decoderState
}

// NewDecoderは、rから読み込む新しいストリーミングデコーダを構築します。
//
// rが [bytes.Buffer] の場合、デコーダは中間バッファに内容をコピーせず、直接バッファからパースします。
// デコーダの使用中にバッファへ追加の書き込みを行ってはいけません。
func NewDecoder(r io.Reader, opts ...Options) *Decoder

// Resetはデコーダをリセットし、新たにrから読み込み、
// 指定されたオプションで構成します。Resetは
// [encoding/json/v2.UnmarshalerFrom.UnmarshalJSONFrom] メソッドに渡された
// Decoderや [encoding/json/v2.UnmarshalFromFunc] 関数に渡されたDecoderに対して呼び出してはいけません。
func (d *Decoder) Reset(r io.Reader, opts ...Options)

// Options は、デコーダの構築に使用されたオプションを返します。
// さらに、[encoding/json/v2.UnmarshalDecode] 呼び出しに渡された
// セマンティックなオプションを含む場合があります。
//
// [encoding/json/v2.UnmarshalerFrom.UnmarshalJSONFrom]メソッド呼び出しや
// [encoding/json/v2.UnmarshalFromFunc]関数呼び出しの中で動作している場合、
// 返されるオプションはその呼び出しの間のみ有効です。
func (d *Decoder) Options() Options

// PeekKindは次のトークン種別を取得しますが、読み取り位置は進めません。
//
// エラーが発生した場合は [KindInvalid] を返します。そのようなエラーは次の読み取り呼び出しまで
// キャッシュされ、最終的にPeekKind呼び出しの後に読み取り呼び出しを続けることは呼び出し側の責任です。
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

// UnreadBufferは、未読バッファに残っているデータを返します。
// このデータには0バイト以上が含まれる可能性があります。
// これは入力 [io.Reader] からすでに消費されたデータですが、
// まだ [Decoder.ReadToken] または [Decoder.ReadValue] の呼び出しでは読み取られていません。
// JSON文法に従ってまだ検証されていないため、
// 有効なJSONを構成しないバイトを含む可能性があります。
// バッファリングされるデータ量の正確な値はDecoderの実装詳細であり、
// 時間とともに変わる可能性があります。
//
// 最後に読み取られたJSONトークンまたは値の後に続くバイト列全体を得るには、
// このバッファと入力Readerの残りを連結する責任は呼び出し側にあります。
//
// 返されたバッファは、Decoderを使い続けている間は変更してはいけません。
// バッファ内容は次のPeek、Read、Skip呼び出しまで有効です。
func (d *Decoder) UnreadBuffer() []byte

// StackDepthは、読み取ったJSONデータに対する状態機械の深さを返します。
// スタック上の各レベルは、ネストしたJSONオブジェクトまたは配列を表します。
// [BeginObject] または [BeginArray] トークンに遭遇するたびにインクリメントされ、
// [EndObject] または [EndArray] トークンに遭遇するたびにデクリメントされます。
//
// オブジェクトや配列の内側にいない場合、StackDepthは0を返します。
// 具体的には、まだトークンを何も読んでいないとき、
// トップレベル値を1つ読み終えたあと、
// およびトップレベル値のストリーム（例: NDJSON）をデコードしているときの値間では
// 0を返します。
// StackDepthは、トップレベルのオブジェクトや配列の内側では1、
// ネストしたオブジェクトや配列の内側では2、というように増えていきます。
//
// 例として、次のJSONをデコードすることを考えます:
//
//	{"a": [1, 2], "b": {"c": 3}}
//
// デコード中、StackDepthは次のように報告されます:
//
//   - 開始時、StackDepthは0です。
//   - 外側の '{' をデコードした後、StackDepthは1です。
//   - 内側の '[' をデコードした後、StackDepthは2です。
//   - 内側の ']' をデコードした後、StackDepthは1です。
//   - 外側の '}' をデコードした後、StackDepthは0です。
func (d *Decoder) StackDepth() int

// StackIndexは指定されたスタックレベルの情報を返します。
// 0から [Decoder.StackDepth] までの数値でなければなりません。
// 各レベルについて、その種別を報告します:
//
//   - ゼロレベルの場合は [KindInvalid]
//   - JSONオブジェクトを表すレベルの場合は [KindBeginObject]
//   - JSON配列を表すレベルの場合は [KindBeginArray]
//
// また、これまでにデコードされたそのJSONオブジェクトまたは配列の長さも報告します。
// JSONオブジェクト内の各名前と値は個別に数えられるため、
// 実際のメンバー数は長さの半分になります。
// 完全なJSONオブジェクトは偶数の長さでなければなりません。
func (d *Decoder) StackIndex(i int) (Kind, int64)

// StackPointerは、直近に読み取った値へのJSONポインタ（RFC 6901）を返します。
func (d *Decoder) StackPointer() Pointer
