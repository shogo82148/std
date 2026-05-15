// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/encoding"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"

	"github.com/shogo82148/std/encoding/json/jsontext"
)

// Reference encoding and time packages to assist pkgsite
// in being able to hotlink references to those packages.
var (
	_ encoding.TextMarshaler
	_ encoding.TextAppender
	_ encoding.TextUnmarshaler
	_ time.Time
	_ time.Duration
)

// Marshalは、指定されたマーシャルおよびエンコードオプションに従って
// Go値を[]byteとしてシリアライズします（アンマーシャルやデコードオプションは無視されます）。
// 出力の末尾に改行は付加しません。
//
// 型固有のマーシャル関数やメソッドは、値のデフォルト表現よりも優先されます。
// *Tを操作する関数やメソッドは、T型の値（アドレスを取得）や非nilの*T値をエンコードする場合のみ呼び出されます。
// Marshalは、値が常にアドレス可能であることを保証します
// （必要に応じてヒープ上にボックス化）ので、これらの関数やメソッドを一貫して呼び出せます。
// パフォーマンスのため、Marshalには非nilポインタ値を渡すことを推奨します。
//
// 入力値は以下のルールに従ってJSONとしてエンコードされます：
//
//   - [WithMarshalers]オプション内の型固有関数が値の型に一致する場合、
//     それらの関数が値をエンコードするために呼び出されます。
//     適用可能なすべての関数が[errors.ErrUnsupported]を返した場合、
//     値は後続のルールに従ってエンコードされます。
//
//   - 値の型が [MarshalerTo] を実装している場合、
//     MarshalJSONToメソッドが呼び出されて値がエンコードされます。
//     メソッドが [errors.ErrUnsupported] を返す場合、
//     入力は後続のルールに従ってエンコードされます。
//
//   - 値の型が [Marshaler] を実装している場合、
//     MarshalJSONメソッドが呼び出されます。
//
//   - 値の型が [encoding.TextAppender] を実装している場合、
//     AppendTextメソッドが呼び出され、その結果がJSON文字列としてエンコードされます。
//
//   - 値の型が [encoding.TextMarshaler] を実装している場合、
//     MarshalTextメソッドが呼び出され、その結果がJSON文字列としてエンコードされます。
//
//   - それ以外の場合、値の型に応じて以下の詳細なルールでエンコードされます。
//
// ほとんどの Go 型には、次のようなデフォルトの JSON 表現があります。
//
//   - Go の真偽値は JSON の真偽値（例: true や false）としてエンコードされます。
//
//   - Go の文字列は JSON 文字列としてエンコードされます。
//
//   - Go の []byte または [N]byte は、RFC 4648 の 4 節に従う Base64 エンコーディングを用いた
//     バイナリ値を含む JSON 文字列としてエンコードされます。
//
//   - Go の整数は、小数部や指数部を持たない JSON 数値としてエンコードされます。
//     [StringifyNumbers] が指定されている場合、または JSON オブジェクト名を
//     エンコードしている場合は、その JSON 数値は JSON 文字列の中にエンコードされます。
//
//   - Go の浮動小数点数は JSON 数値としてエンコードされます。
//     [StringifyNumbers] が指定されている場合、または JSON オブジェクト名を
//     エンコードしている場合は、その JSON 数値は JSON 文字列の中にエンコードされます。
//
//   - Go の map は JSON オブジェクトとしてエンコードされ、各 Go の map のキーと値は
//     JSON オブジェクト内の名前と値の組として再帰的にエンコードされます。
//     Go の map のキーは JSON 文字列としてエンコードできなければならず、そうでなければ
//     [SemanticError] になります。Go の map は非決定的な順序で走査されます。
//     決定的なエンコードが必要な場合は、[Deterministic] オプションの使用を検討してください。
//     デフォルトでは、nil の map は空の JSON オブジェクトとしてエンコードされます。
//     ただし、[FormatNilMapAsNull] オプションが指定されている場合を除きます。
//
//   - Goの構造体はJSONオブジェクトとしてエンコードされます。
//     詳細はパッケージレベルの「Go構造体のJSON表現」セクション参照。
//
//   - Go のスライスは JSON 配列としてエンコードされ、各 Go スライス要素は
//     JSON 配列の要素として再帰的に JSON エンコードされます。
//     デフォルトでは、nil スライスは空の JSON 配列としてエンコードされます。
//     ただし、[FormatNilSliceAsNull] オプションが指定されている場合を除きます。
//
//   - Go の配列は JSON 配列としてエンコードされ、各 Go 配列要素は
//     JSON 配列の要素として再帰的に JSON エンコードされます。
//     JSON 配列の長さは常に Go 配列の長さと一致します。
//
//   - Go のポインタは、nil なら JSON null としてエンコードされ、
//     そうでなければ基底値を再帰的に JSON エンコードした表現になります。
//
//   - Go のインターフェースは、nil なら JSON null としてエンコードされ、
//     そうでなければ基底値を再帰的に JSON エンコードした表現になります。
//
//   - Go の [time.Time] は、RFC 3339 に従いナノ秒精度でフォーマットされた
//     タイムスタンプを含む JSON 文字列としてエンコードされます。
//
//   - Go の [time.Duration] は現在デフォルト表現を持たず、
//     [encoding/json.FormatDurationAsNano] オプションが指定されていない限り、
//     [SemanticError] になります。そのオプションが指定されている場合は、
//     継続時間をナノ秒で表す、小数部も指数部も持たない JSON 数値として
//     エンコードされます。
//
//   - その他のGo型（複素数、チャネル、関数など）はデフォルト表現がなく、[SemanticError] になります。
//
// JSONは循環データ構造を表現できず、Marshalはそれらを扱いません。
// 循環構造を渡すとエラーになります。
func Marshal(in any, opts ...Options) (out []byte, err error)

// MarshalWriteは、指定されたマーシャルおよびエンコードオプションに従って
// Go値を [io.Writer] にシリアライズします（アンマーシャルやデコードオプションは無視されます）。
// 出力の末尾に改行は付加しません。
// Go値をJSONへ変換する詳細は [Marshal] を参照してください。
func MarshalWrite(out io.Writer, in any, opts ...Options) (err error)

// MarshalEncodeは、指定されたマーシャルオプションに従って
// Go値を [jsontext.Encoder] にシリアライズします（アンマーシャル、エンコード、デコードオプションは無視されます）。
// [jsontext.Encoder] に既に指定されているマーシャル関連のオプションは、呼び出し元が指定したオプションよりも優先度が低くなります。
// [Marshal] や [MarshalWrite] と異なり、エンコードオプションは無視されます。
// これは、エンコードオプションが既に指定済みである必要があるためです。
//
// Go値をJSONへ変換する詳細は [Marshal] を参照してください。
func MarshalEncode(out *jsontext.Encoder, in any, opts ...Options) (err error)

// Unmarshalは、指定されたアンマーシャルおよびデコードオプションに従って
// []byte入力をGo値へデコードします（マーシャルやエンコードオプションは無視されます）。
// 入力は、空白を含んでもよい単一のJSON値でなければなりません。
// 出力は非nilのポインタでなければなりません。
//
// 型固有のアンマーシャル関数やメソッドは、値のデフォルト表現よりも優先されます。
// *Tを操作する関数やメソッドは、T型の値（アドレスを取得）や非nilの*T値をデコードする場合のみ呼び出されます。
// Unmarshalは、値が常にアドレス可能であることを保証します
// （必要に応じてヒープ上にボックス化）ので、これらの関数やメソッドを一貫して呼び出せます。
//
// 入力は以下のルールに従って出力へデコードされます：
//
//   - [WithUnmarshalers]オプション内の型固有関数が値の型に一致する場合、
//     それらの関数がJSON値をデコードするために呼び出されます。
//     適用可能なすべての関数が[errors.ErrUnsupported]を返した場合、
//     入力は後続のルールに従ってデコードされます。
//
//   - 値の型が [UnmarshalerFrom] を実装している場合、
//     UnmarshalJSONFromメソッドが呼び出されてJSON値がデコードされます。
//     メソッドが [errors.ErrUnsupported] を返す場合、
//     入力は後続のルールに従ってデコードされます。
//
//   - 値の型が [Unmarshaler] を実装している場合、
//     UnmarshalJSONメソッドが呼び出されます。
//
//   - 値の型が [encoding.TextUnmarshaler] を実装している場合、
//     入力はJSON文字列としてデコードされ、
//     デコードされた文字列値でUnmarshalTextメソッドが呼び出されます。
//     入力がJSON文字列でない場合は [SemanticError] になります。
//
//   - それ以外の場合、値の型に応じて以下の詳細なルールでデコードされます。
//
// ほとんどの Go 型にはデフォルトの JSON 表現があります。
// JSON null は、サポートされるすべての Go 値に対してデコードでき、
// それはその Go 値のゼロ値を格納するのと同等です。
// 入力 JSON の種類が現在の Go 値型で扱えない場合、
// [SemanticError] になります。特に指定がない限り、
// デコードされた値は既存の値を置き換えます。
//
// 各型の表現は以下の通りです：
//
//   - Go の真偽値は JSON の真偽値（例: true や false）からデコードされます。
//
//   - Go の文字列は JSON 文字列からデコードされます。
//
//   - Go の []byte または [N]byte は、RFC 4648 の 4 節に従う Base64 エンコーディングを用いた
//     バイナリ値を含む JSON 文字列からデコードされます。
//     非 nil の []byte へデコードする場合、スライス長はゼロにリセットされ、
//     デコードされた入力がそこへ追加されます。
//     [N]byte へデコードする場合、入力はちょうど N バイトにデコードされなければならず、
//     そうでなければ [SemanticError] になります。
//
//   - Go の整数は JSON 数値からデコードされます。
//     [StringifyNumbers] が指定されている場合や JSON オブジェクト名をデコードする場合、
//     JSON 数値を含む JSON 文字列からデコードされなければなりません。
//     JSON 数値に小数部や指数部がある場合は [SemanticError] になります。
//     Go 整数型の表現をオーバーフローした場合も失敗します。
//
//   - Go の浮動小数点数は JSON 数値からデコードされます。
//     [StringifyNumbers] が指定されている場合や JSON オブジェクト名をデコードする場合、
//     JSON 数値を含む JSON 文字列からデコードされなければなりません。
//     Go 浮動小数点型の表現をオーバーフローした場合は失敗します。
//
//   - Go の map は JSON オブジェクトからデコードされ、
//     各 JSON オブジェクト名と値の組が Go map のキーと値として再帰的にデコードされます。
//     map はクリアされません。
//     Go map が nil の場合、新しい map が割り当てられてデコード先になります。
//     デコードされたキーが既存の Go map エントリと一致する場合、
//     エントリ値は再利用され、JSON オブジェクト値がそこへデコードされます。
//
//   - Goの構造体はJSONオブジェクトからデコードされます。
//     詳細はパッケージレベルの「Go構造体のJSON表現」セクション参照。
//
//   - Go のスライスは JSON 配列からデコードされ、各 JSON 要素は再帰的にデコードされて
//     Go スライスへ追加されます。
//     Go スライスへ追加する前に、それが nil なら新しいスライスが割り当てられ、
//     そうでなければスライス長はゼロにリセットされます。
//
//   - Go の配列は JSON 配列からデコードされ、各 JSON 配列要素は対応する
//     Go 配列要素として再帰的にデコードされます。
//     各 Go 配列要素はデコード前にゼロ化されます。
//     JSON 配列が Go 配列とまったく同じ数の要素を含まない場合、
//     [SemanticError] になります。
//
//   - Go のポインタは、JSON の種類と基底 Go 型に基づいてデコードされます。
//     入力が JSON null の場合、nil ポインタが格納されます。
//     それ以外の場合、ポインタが nil なら新しい基底値が割り当てられ、
//     その基底値へ再帰的に JSON デコードされます。
//
//   - Go のインターフェースは、JSON の種類と基底 Go 型に基づいてデコードされます。
//     入力が JSON null の場合、nil インターフェース値が格納されます。
//     それ以外の場合、空インターフェース型の nil インターフェース値は、入力が
//     JSON の真偽値、文字列、数値、オブジェクト、配列であれば、それぞれ
//     ゼロ値の Go bool、string、float64、map[string]any、[]any で初期化されます。
//     それでもインターフェース値が nil のままである場合、適切な Go 型を
//     デコード先として決定できなかったため、[SemanticError] になります。
//     たとえば、nil の io.Reader へアンマーシャルすると、
//     インターフェース値に設定する具体型が存在しないため失敗します。
//     それ以外では基底値が存在し、JSON 入力はそこへ再帰的にデコードされます。
//
//   - Go の [time.Time] は、RFC 3339 に従いナノ秒精度でフォーマットされた時刻を含む
//     JSON 文字列からデコードされます。
//
//   - Go の [time.Duration] は現在デフォルト表現を持たず、
//     [encoding/json.FormatDurationAsNano] オプションが指定されていない限り、
//     [SemanticError] になります。そのオプションが指定されている場合は、
//     継続時間をナノ秒で表す、小数部も指数部も持たない JSON 数値として
//     デコードされます。
//
//   - その他のGo型（複素数、チャネル、関数など）はデフォルト表現がなく、[SemanticError] になります。
//
// 一般に、アンマーシャルはマージセマンティクス（RFC 7396に類似）に従い、
// デコードされたGo値はJSONオブジェクト以外の種類では出力値を置き換えます。
// JSONオブジェクトの場合、入力オブジェクトは出力値にマージされ、
// 一致するメンバーは再帰的にマージセマンティクスが適用されます。
func Unmarshal(in []byte, out any, opts ...Options) (err error)

// UnmarshalReadは、指定されたアンマーシャルおよびデコードオプションに従って
// [io.Reader] からGo値をデシリアライズします（マーシャルやエンコードオプションは無視されます）。
// 入力は空白を含んでもよい単一のJSON値でなければなりません。
// [io.Reader] 全体を [io.EOF] に達するまで消費し、EOFでエラーを報告しません。
// 出力は非nilのポインタでなければなりません。
// JSONからGo値への変換の詳細は [Unmarshal] を参照してください。
func UnmarshalRead(in io.Reader, out any, opts ...Options) (err error)

// UnmarshalDecodeは、指定されたアンマーシャルオプションに従って
// [jsontext.Decoder] からGo値をデシリアライズします（マーシャル、エンコード、デコードオプションは無視されます）。
// [jsontext.Decoder] に既に指定されているアンマーシャルオプションは、呼び出し元が指定したオプションよりも優先度が低くなります。
// [Unmarshal] や [UnmarshalRead] と異なり、デコードオプションは無視されます。
// これは、デコードオプションが既に指定済みである必要があるためです。
//
// 入力は0個以上のJSON値のストリームであっても構いませんが、
// これはストリーム内の次のJSON値のみをアンマーシャルします。
// トップレベルのJSON値がもうない場合は、[io.EOF] を報告します。
// 出力は非nilのポインタでなければなりません。
// JSONからGo値への変換の詳細は [Unmarshal] を参照してください。
func UnmarshalDecode(in *jsontext.Decoder, out any, opts ...Options) (err error)
