// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

// v2への移行
//
// このパッケージ（つまり [encoding/json]）は、v2パッケージが[encoding/json/v2]に存在するため、正式にはv1パッケージと呼ばれるようになりました。
// v1パッケージのすべての動作は、v2パッケージの適切なオプションセットを指定することで、v1の歴史的な動作を維持した形で実装されています。
//
// [jsonv2.Marshal] 関数は、v1の [Marshal] の新しい同等関数です。
// [jsonv2.Unmarshal] 関数は、v1の [Unmarshal] の新しい同等関数です。
// v2の関数は、v1と同じ呼び出しシグネチャですが、可変長の [Options] 引数を受け取り、マーシャルやアンマーシャルの動作を変更できます。
// v1とv2は一般的に似た動作をしますが、いくつか顕著な違いがあります。
//
// 以下はv1とv2の違いの一覧です：
//
//   - v1では、JSONオブジェクトのメンバーはGo構造体のフィールド名と大文字小文字を区別しない一致でアンマーシャルされます。
//     一方、v2ではフィールド名の完全一致（大文字小文字区別）で一致させます。
//     [jsonv2.MatchCaseInsensitiveNames] や [MatchCaseSensitiveDelimiter] オプションでこの動作を制御できます。
//     Go構造体フィールドごとに一致方法を指定したい場合は、`case:ignore`や`case:strict`フィールドオプションを指定できます。
//     フィールド指定のオプションは呼び出し元指定のオプションより優先されます。
//
//   - v1では、Go構造体のフィールドに`omitempty`が付いている場合、値が「空」のGo値（false, 0, nilポインタ, nilインターフェース値、長さ0の配列・スライス・マップ・文字列）なら省略されます。
//     一方、v2では`omitempty`は「空」のJSON値（JSON null、空のJSON文字列・オブジェクト・配列）としてエンコードされる場合に省略されます。
//     [OmitEmptyWithLegacySemantics]オプションでこの動作を制御できます。
//     なお、`omitempty`はGoの配列・スライス・マップ・文字列についてはv1とv2で同じ動作です（ユーザー定義MarshalJSONメソッドがなければ）。
//     既存のGoのbool, number, pointer, interface値に対する`omitempty`は、`omitzero`に移行してください（v1/v2両方で同じ動作）。
//
//   - v1では、Go構造体フィールドに`string`を付けると、Goのstring, bool, numberをJSON文字列として引用できます。複合型には再帰的に適用されません。
//     一方、v2では`string`オプションはGoのnumberのみをJSON文字列として引用でき、複合型内のGo numberにも再帰的に適用されます。
//     [StringifyWithLegacySemantics] オプションでこの動作を制御できます。
//
//   - v1では、nilのGoスライスやGoマップはJSON nullとしてマーシャルされます。
//     一方、v2ではnilのGoスライスは空のJSON配列、nilのGoマップは空のJSONオブジェクトとしてマーシャルされます。
//     [jsonv2.FormatNilSliceAsNull] や [jsonv2.FormatNilMapAsNull] オプションでこの動作を制御できます。
//     Go構造体フィールドごとにnilの表現を指定したい場合は、`format:emitempty`や`format:emitnull`フィールドオプションを指定できます。
//     フィールド指定のオプションは呼び出し元指定のオプションより優先されます。
//
//   - v1では、Go配列は任意の長さのJSON配列からアンマーシャルできます。
//     一方、v2ではGo配列は同じ長さのJSON配列からのみアンマーシャルでき、長さが違うとエラーになります。
//     [UnmarshalArrayFromAnyLength] オプションでこの動作を制御できます。
//
//   - v1では、Goのバイト配列はJSONの数値配列として表現されます。
//     一方、v2ではGoのバイト配列はBase64エンコードされたJSON文字列として表現されます。
//     [FormatByteArrayAsArray] オプションでこの動作を制御できます。
//     Go構造体フィールドごとに表現を指定したい場合は、`format:array`や`format:base64`フィールドオプションを指定できます。
//     フィールド指定のオプションは呼び出し元指定のオプションより優先されます。
//
//   - v1では、ポインタレシーバで宣言されたMarshalJSONメソッドはGo値がアドレス可能な場合のみ呼び出されます。
//     一方、v2ではMarshalJSONメソッドはアドレス可能かどうかに関係なく常に呼び出されます。
//     [CallMethodsWithLegacySemantics] オプションでこの動作を制御できます。
//
//   - v1では、Goマップのキーに対してMarshalJSONやUnmarshalJSONメソッドは呼び出されません。
//     一方、v2ではGoマップのキーにもMarshalJSONやUnmarshalJSONメソッドが呼び出される可能性があります。
//     [CallMethodsWithLegacySemantics] オプションでこの動作を制御できます。
//
//   - v1では、Goマップは決定的な順序でマーシャルされます。
//     一方、v2ではGoマップは非決定的な順序でマーシャルされます。
//     [jsonv2.Deterministic] オプションでこの動作を制御できます。
//
//   - v1では、JSON文字列はHTMLやJavaScript固有の文字がエスケープされてエンコードされます。
//     一方、v2ではJSON文字列は最小限のエンコーディングとなり、JSON文法で必要な場合のみエスケープされます。
//     [jsontext.EscapeForHTML] や [jsontext.EscapeForJS] オプションでこの動作を制御できます。
//
//   - v1では、文字列内の無効なUTF-8バイトは黙ってUnicodeの置換文字に置き換えられます。
//     一方、v2では無効なUTF-8があるとエラーになります。[jsontext.AllowInvalidUTF8] オプションでこの動作を制御できます。
//
//   - v1では、重複した名前を持つJSONオブジェクトが許可されます。
//     一方、v2では重複した名前のJSONオブジェクトはエラーになります。[jsontext.AllowDuplicateNames] オプションでこの動作を制御できます。
//
//   - v1では、JSON nullを非空のGo値にアンマーシャルする場合、値をゼロクリアするか何もしないかが一貫しません。
//     一方、v2ではJSON nullをアンマーシャルすると常にGo値をゼロクリアします。[MergeWithLegacySemantics] オプションでこの動作を制御できます。
//
//   - v1では、JSON値を非ゼロのGo値にアンマーシャルする場合、配列要素・スライス要素・構造体フィールド（ただしマップ値は除く）・ポインタ値・インターフェース値（非nilポインタのみ）にマージされます。
//     一方、v2では構造体フィールド・マップ値・ポインタ値・インターフェース値にマージされます。
//     一般的に、v2のセマンティクスはJSONオブジェクトをアンマーシャルする場合にマージし、それ以外は値を置き換えます。[MergeWithLegacySemantics] オプションでこの動作を制御できます。
//
//   - v1では、[time.Duration] はナノ秒数のJSON数値として表現されます。
//     一方、v2では [time.Duration] はデフォルトの表現がなく、実行時エラーになります。[FormatDurationAsNano] オプションでこの動作を制御できます。
//     Go構造体フィールドごとに表現を指定したい場合は、`format:nano`や`format:units`フィールドオプションを指定できます。
//     フィールド指定のオプションは呼び出し元指定のオプションより優先されます。
//
//   - v1では、Go構造体型に構造的なエラー（例：タグオプションの不正）があっても実行時エラーは報告されません。
//     一方、v2ではJSONシリアライズに関連するGo型が不正な場合は実行時エラーが報告されます。例えば、エクスポートされていないフィールドのみのGo構造体はシリアライズできません。
//     [ReportErrorsWithLegacySemantics] オプションでこの動作を制御できます。
//
// 前述の通り、v1の全機能はv2を使って実装されており、オプションを指定することでレガシー動作に切り替えています。
// 例えば、[Marshal] は [jsonv2.Marshal] を [DefaultOptionsV1] 付きで直接呼び出します。
// 同様に、[Unmarshal] は [jsonv2.Unmarshal] を [DefaultOptionsV1] 付きで直接呼び出します。
// [DefaultOptionsV1] オプションはv1のデフォルト動作を指定するすべてのオプションセットです。
//
// 多くの動作の違いについては、Go型の作者がGo構造体フィールドオプションを指定することで、v1/v2どちらのセマンティクスでも同じJSON表現になるよう制御できます。
//
// [DefaultOptionsV1] と [jsonv2.DefaultOptionsV2] の両方を利用でき、後者のオプションが前者より優先されるため、v1からv2への段階的な移行が可能です。例：
//
//   - jsonv1.Marshal(v)
//     デフォルトのv1セマンティクスを使用します。
//
//   - jsonv2.Marshal(v, jsonv1.DefaultOptionsV1())
//     jsonv1.Marshalと同じ意味で、デフォルトのv1セマンティクスを使用します。
//
//   - jsonv2.Marshal(v, jsonv1.DefaultOptionsV1(), jsontext.AllowDuplicateNames(false))
//     ほぼv1セマンティクスですが、1つだけv2固有の動作に切り替えます。
//
//   - jsonv2.Marshal(v, jsonv1.CallMethodsWithLegacySemantics(true))
//     ほぼv2セマンティクスですが、1つだけv1固有の動作に切り替えます。
//
//   - jsonv2.Marshal(v, ..., jsonv2.DefaultOptionsV2())
//     jsonv2.Marshalと同じ意味で、jsonv2.DefaultOptionsV2がそれ以前のオプションを上書きし、デフォルトのv2セマンティクスを使用します。
//
//   - jsonv2.Marshal(v)
//     デフォルトのv2セマンティクスを使用します。
//
// Goで新しく"json"を使う場合はv2パッケージの利用を推奨しますが、v1パッケージも今後もサポートされ続けます。
package json

import (
	"github.com/shogo82148/std/encoding"

	"github.com/shogo82148/std/encoding/json/internal/jsonopts"
	"github.com/shogo82148/std/encoding/json/jsontext"
)

// Reference encoding, jsonv2, and jsontext packages to assist pkgsite
// in being able to hotlink references to those packages.
var (
	_ encoding.TextMarshaler
	_ encoding.TextUnmarshaler
	_ jsonv2.Options
	_ jsontext.Options
)

// Optionsは、v2の "json" パッケージを特定の機能についてv1のセマンティクスで動作させるためのオプションセットです。
// この型の値は、[jsonv2.Marshal] や [jsonv2.Unmarshal] などのv2関数に渡すことができます。
// この型を直接参照するのではなく、[jsonv2.Options] を使用してください。
//
// v1からv2への移行方法については「v2への移行」セクションを参照してください。
type Options = jsonopts.Options

// DefaultOptionsV1は、v1のセマンティクスを定義するすべてのオプションセットです。
// 以下のブールオプションがtrueに設定されているのと同等です：
//
//   - [CallMethodsWithLegacySemantics]
//   - [FormatByteArrayAsArray]
//   - [FormatBytesWithLegacySemantics]
//   - [FormatDurationAsNano]
//   - [MatchCaseSensitiveDelimiter]
//   - [MergeWithLegacySemantics]
//   - [OmitEmptyWithLegacySemantics]
//   - [ParseBytesWithLooseRFC4648]
//   - [ParseTimeWithLooseRFC3339]
//   - [ReportErrorsWithLegacySemantics]
//   - [StringifyWithLegacySemantics]
//   - [UnmarshalArrayFromAnyLength]
//   - [jsonv2.Deterministic]
//   - [jsonv2.FormatNilMapAsNull]
//   - [jsonv2.FormatNilSliceAsNull]
//   - [jsonv2.MatchCaseInsensitiveNames]
//   - [jsontext.AllowDuplicateNames]
//   - [jsontext.AllowInvalidUTF8]
//   - [jsontext.EscapeForHTML]
//   - [jsontext.EscapeForJS]
//   - [jsontext.PreserveRawStrings]
//
// その他のブールオプションはすべてfalseに設定されます。
// 非ブールオプションはすべてゼロ値に設定されますが、[jsontext.WithIndent] のみ"\t"がデフォルトです。
//
// このパッケージの [Marshal] および [Unmarshal] 関数は、v2の同等関数にこのオプションを指定して呼び出すのと同じ意味です：
//
//	jsonv2.Marshal(v, jsonv1.DefaultOptionsV1())
//	jsonv2.Unmarshal(b, v, jsonv1.DefaultOptionsV1())
func DefaultOptionsV1() Options

// CallMethodsWithLegacySemanticsは、型が提供するマーシャル・アンマーシャルメソッドの呼び出しをレガシーセマンティクスで行うことを指定します:
//
//   - マーシャル時、ポインタレシーバで宣言されたマーシャルメソッドはGo値がアドレス可能な場合のみ呼び出されます。
//     インターフェースやマップ要素から取得した値はアドレス不可です。
//     ポインタやスライス要素から取得した値はアドレス可能です。
//     配列要素や構造体フィールドから取得した値は親のアドレス可能性を継承します。
//     v2のセマンティクスではアドレス可能かどうかに関係なく常にメソッドを呼び出します。
//
//   - マーシャル・アンマーシャル時、マップキーに対して [Marshaler] や [Unmarshaler] メソッドは無視されます。
//     ただし [encoding.TextMarshaler] や [encoding.TextUnmarshaler] は呼び出されます。
//     v2のセマンティクスではマップキーも他の値と同様にメソッドを呼び出してシリアライズします。
//     実装側がGo値をJSON文字列として表現する責任があります（JSONオブジェクト名として必要）。
//
//   - マーシャル時、マップキー値がマーシャルメソッドを実装していてnilポインタの場合、空のJSON文字列としてシリアライズされます。
//     v2のセマンティクスではエラーになります。
//
//   - マーシャル時、インターフェース型がマーシャルメソッドを実装していてインターフェース値が具体型へのnilポインタの場合、常にマーシャルメソッドが呼び出されます。
//     v2のセマンティクスではインターフェース値に直接メソッドを呼び出さず、基底の具体値に基づいて評価を遅延します。
//     非インターフェース値と同様、nilポインタにはメソッドを呼び出さず、JSON nullとしてシリアライズされます。
//
// このオプションはマーシャル・アンマーシャルの両方に影響します。
// v1のデフォルトはtrueです。
func CallMethodsWithLegacySemantics(v bool) Options

// FormatByteArrayAsArrayは、Goの [N]byte 型を通常のGo配列としてフォーマットすることを指定します。
// v2のデフォルトでは [N]byte 型はバイナリデータエンコーディング（RFC 4648）としてフォーマットされます。
// 構造体フィールドに `format` タグオプションが指定されている場合は、そのフォーマットが優先されます。
//
// このオプションはマーシャル・アンマーシャルの両方に影響します。
// v1のデフォルトはtrueです。
func FormatByteArrayAsArray(v bool) Options

// FormatBytesWithLegacySemanticsは、[]~byte型および[N]~byte型の扱いをレガシーセマンティクスに従うことを指定します:
//
//   - Goの[]~byte型は、v2のデフォルトである[]byte型のみをバイナリデータとして扱うのとは異なり、何らかのバイナリデータエンコーディング（RFC 4648）として扱われます。特に、v2では名前付きbyte型のスライスはバイナリデータとして扱われません。
//
//   - マーシャル時、名前付きbyte型がマーシャルメソッドを実装している場合、スライスは各要素ごとにマーシャルメソッドを呼び出してJSON配列としてシリアライズされます。
//
//   - アンマーシャル時、入力がJSON配列の場合、通常のGoスライスとして[]~byte型にアンマーシャルされます。対して、v2のデフォルトではバイナリデータエンコーディングを期待している場合にJSON配列のアンマーシャルはエラーとなります。
//
// このオプションはマーシャル・アンマーシャルの両方に影響します。
// v1のデフォルトはtrueです。
func FormatBytesWithLegacySemantics(v bool) Options

// FormatDurationAsNanoは、[time.Duration] 型をJSON数値（ナノ秒数）としてフォーマットすることを指定します。
// v2のデフォルトではエラーとなります。
// フィールドに`format`タグオプションが指定されている場合は、そのフォーマットが優先されます。
//
// このオプションはマーシャル・アンマーシャルの両方に影響します。
// v1のデフォルトはtrueです。
func FormatDurationAsNano(v bool) Options

// MatchCaseSensitiveDelimiterは、大文字小文字を区別しない名前一致（[jsonv2.MatchCaseInsensitiveNames] や `case:ignore` タグオプション使用時）において、アンダースコアやハイフンを無視しないことを指定します。
// そのため、大文字小文字を区別しない名前一致は [strings.EqualFold] と同じ動作になります。
// このオプションを使用すると、大文字小文字を区別しない一致で一般的なケースバリアント（例："foo_bar" と "fooBar"）を一致させる能力が低下します。
//
// このオプションはマーシャル・アンマーシャルのどちらにも影響します。
// v1のデフォルトはtrueです。
func MatchCaseSensitiveDelimiter(v bool) Options

// MergeWithLegacySemanticsは、非ゼロのGo値へのアンマーシャル時にレガシーセマンティクスで動作することを指定します:
//
//   - JSON nullをアンマーシャルする場合、Go値の型がbool, int, uint, float, string, array, structであれば元の値を保持します。
//     それ以外の場合はGo値をゼロクリアします。
//     対して、v2のデフォルト動作ではJSON nullをアンマーシャルすると常にGo値をゼロクリアします。
//
//   - JSON null以外の値をアンマーシャルする場合、配列要素・スライス要素・構造体フィールド（ただしマップ値は除く）・ポインタ値・インターフェース値（非nilポインタのみ）にマージします。
//     対して、v2のデフォルト動作では構造体フィールド・マップ値・ポインタ値・インターフェース値にマージします。
//     一般的に、v2のセマンティクスではJSONオブジェクトをアンマーシャルする場合にマージし、それ以外は値を置き換えます。
//
// このオプションはアンマーシャル時のみ影響し、マーシャル時は無視されます。
// v1のデフォルトはtrueです。
func MergeWithLegacySemantics(v bool) Options

// OmitEmptyWithLegacySemanticsは、`omitempty`タグオプションが「空」の定義に従うことを指定します。
// この定義では、Go値がfalse、0、nilポインタ、nilインターフェース値、または空の配列・スライス・マップ・文字列の場合にフィールドが省略されます。
// この動作は、値がJSON nullや空のJSON文字列・オブジェクト・配列としてマーシャルされる場合にフィールドを省略するv2のセマンティクスを上書きします。
//
// v1とv2の`omitempty`の定義は、Goの文字列、スライス、配列、マップについてはほぼ同じです。
// Goのbool、int、uint、float、ポインタ、インターフェースに対する`omitempty`の利用は、ゼロ値の場合にフィールドを省略する`omitzero`タグオプションへの移行が推奨されます。
//
// このオプションはマーシャル時のみ影響し、アンマーシャル時は無視されます。
// v1のデフォルトはtrueです。
func OmitEmptyWithLegacySemantics(v bool) Options

// ParseBytesWithLooseRFC4648は、"base32"や"base64"でエンコードされたバイナリデータをパースする際に、'\r'や'\n'文字の存在を無視することを指定します。
// 対して、v2のデフォルトではRFC 4648の厳密な準拠のためエラーを報告します（RFC 4648セクション3.3では非アルファベット文字は拒否する必要があります）。
//
// このオプションはアンマーシャル時のみ影響し、マーシャル時は無視されます。
// v1のデフォルトはtrueです。
func ParseBytesWithLooseRFC4648(v bool) Options

// ParseTimeWithLooseRFC3339は、[time.Time] 型のパースをRFC 3339に緩やかに準拠して行うことを指定します。
// 特に、過去の誤った表現（時間のフォーマット、秒以下の区切り文字、タイムゾーン表現の揺れ）も許容します。
// 一方、v2のデフォルト動作ではRFC 3339で定められた文法に厳密に従います。
//
// このオプションはアンマーシャル時のみ影響し、マーシャル時は無視されます。
// v1のデフォルトはtrueです。
func ParseTimeWithLooseRFC3339(v bool) Options

// ReportErrorsWithLegacySemanticsは、MarshalおよびUnmarshalがレガシーセマンティクスでエラーを報告することを指定します:
//
//   - マーシャルまたはアンマーシャル時、返されるエラー値は通常 [SyntaxError]、[MarshalerError]、[UnsupportedTypeError]、[UnsupportedValueError]、[InvalidUnmarshalError]、[UnmarshalTypeError] などの型になります。
//     一方、v2のセマンティクスでは、常に [jsonv2.SemanticError] または [jsontext.SyntacticError] のいずれかとしてエラーを返します。
//
//   - マーシャル時、ユーザー定義のマーシャルメソッドがエラーを報告した場合、エラー自体がすでに [MarshalerError] であっても必ず [MarshalerError] でラップされ、冗長なラップが複数重なることがあります。
//     一方、v2のセマンティクスでは、すでにセマンティックエラーでない限り、常に [jsonv2.SemanticError] でラップします。
//
//   - アンマーシャル時、ユーザー定義のアンマーシャルメソッドがエラーを報告した場合、ラップせずそのまま報告します。
//     一方、v2のセマンティクスでは、すでにセマンティックエラーでない限り、常に [jsonv2.SemanticError] でラップします。
//
//   - マーシャルまたはアンマーシャル時、Go構造体に型エラー（例：名前の競合やフィールドタグの不正）がある場合、それらのエラーは無視され、Go構造体はベストエフォートで表現されます。
//     一方、v2のセマンティクスでは、ランタイムエラーとして報告します。
//
//   - アンマーシャル時、JSON入力の構文構造は、JSONデータをGo値にセマンティックアンマーシャルする前に完全に検証されます。
//     実際には、構文エラーのあるJSON入力はGo値の変更を一切引き起こしません。
//     一方、v2のセマンティクスではストリーミングデコードを行い、JSON入力を段階的にGo値へアンマーシャルするため、構文エラーが発生した場合でもGo値が部分的に変更される可能性があります。
//
//   - アンマーシャル時、セマンティックエラーが発生してもすぐに処理を終了せず、評価を継続します。
//     Unmarshalが返る際、最初のセマンティックエラーのみが報告されます。
//     一方、v2のセマンティクスでは、エラーが発生した時点でアンマーシャル処理を終了します。
//
// このオプションはマーシャルまたはアンマーシャルのどちらにも影響します。
// v1のデフォルトはtrueです。
func ReportErrorsWithLegacySemantics(v bool) Options

// StringifyWithLegacySemanticsは、`string`タグオプションがbool型やstring型の値を文字列化できることを指定します。
// このオプションは、フィールドのトップレベルの型がbool、string、数値型、またはそれらへのポインタの場合のみ有効です。
// 特に、`string`は複合型（配列、スライス、構造体、マップ、インターフェース）の内部にあるbool、string、数値型には適用されません。
//
// マーシャル時、これらのGo値は通常のJSON表現でシリアライズされますが、JSON文字列として引用されます。
// アンマーシャル時、これらのGo値は通常のJSON表現が含まれるJSON文字列からデシリアライズされなければなりません。
// JSON文字列内で引用されたJSON nullは、`string`が有効なGo値へのアンマーシャル時にJSON nullの代用として認められます。
//
// このオプションはマーシャルまたはアンマーシャルのどちらにも影響します。
// v1のデフォルトはtrueです。
func StringifyWithLegacySemantics(v bool) Options

// UnmarshalArrayFromAnyLengthは、Go配列が入力JSON配列の任意の長さからアンマーシャルできることを指定します。
// JSON配列が短すぎる場合、残りのGo配列要素はゼロクリアされます。
// JSON配列が長すぎる場合、余分なJSON配列要素はスキップされます。
//
// このオプションはアンマーシャル時のみ影響し、マーシャル時は無視されます。
// v1のデフォルトはtrueです。
func UnmarshalArrayFromAnyLength(v bool) Options
