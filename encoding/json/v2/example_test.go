// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/math"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/net/netip"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/reflect"
	"github.com/shogo82148/std/strconv"
	"github.com/shogo82148/std/strings"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/time"

	"github.com/shogo82148/std/encoding/json/jsontext"
)

// 型が [encoding.TextMarshaler] や [encoding.TextUnmarshaler] を実装している場合、
// MarshalText および UnmarshalText メソッドが値のエンコード・デコードに
// JSON文字列として利用されます。
func Example_textMarshal() {
	// netip.Addr型が encoding.TextMarshaler と encoding.TextUnmarshaler の両方を
	// 実装しているホスト名マップをラウンドトリップでマーシャル・アンマーシャルします。
	want := map[netip.Addr]string{
		netip.MustParseAddr("192.168.0.100"): "carbonite",
		netip.MustParseAddr("192.168.0.101"): "obsidian",
		netip.MustParseAddr("192.168.0.102"): "diamond",
	}
	b, err := json.Marshal(&want, json.Deterministic(true))
	if err != nil {
		log.Fatal(err)
	}
	var got map[netip.Addr]string
	err = json.Unmarshal(b, &got)
	if err != nil {
		log.Fatal(err)
	}

	// 正常性チェック。
	if !reflect.DeepEqual(got, want) {
		log.Fatalf("roundtrip mismatch: got %v, want %v", got, want)
	}

	// シリアライズされたJSONオブジェクトを表示します。
	(*jsontext.Value)(&b).Indent() // 可読性のためインデント
	fmt.Println(string(b))

	// Output:
	// {
	// 	"192.168.0.100": "carbonite",
	// 	"192.168.0.101": "obsidian",
	// 	"192.168.0.102": "diamond"
	// }
}

// デフォルトでは、Go構造体フィールドのJSONオブジェクト名は
// Goフィールド名から導出されますが、`json`タグで指定することもできます。
// JSONがJavaScript由来であるため、JSONオブジェクト名の最も一般的な命名規則はcamelCaseです。
func Example_fieldNames() {
	var value struct {
		// このフィールドは特殊な"-"名で明示的に無視されます。
		Ignored any `json:"-"`
		// JSON名が指定されていないため、Goフィールド名が使われます。
		GoName any
		// 特殊文字なしでJSON名が指定されています。
		JSONName any `json:"jsonName"`
		// JSON名が指定されていないため、Goフィールド名が使われます。
		Option any `json:",case:ignore"`
		// 空のJSON名を単一引用符文字列リテラルで指定しています。
		Empty any `json:"''"`
		// ダッシュのJSON名を単一引用符文字列リテラルで指定しています。
		Dash any `json:"'-'"`
		// カンマのJSON名を単一引用符文字列リテラルで指定しています。
		Comma any `json:"','"`
		// 引用符付きのJSON名を単一引用符文字列リテラルで指定しています。
		Quote any `json:"'\"\\''"`
		// 非公開フィールドは常に無視されます。
		unexported any
	}

	b, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}
	(*jsontext.Value)(&b).Indent() // 可読性のためインデント
	fmt.Println(string(b))

	// Output:
	// {
	// 	"GoName": null,
	// 	"jsonName": null,
	// 	"Option": null,
	// 	"": null,
	// 	"-": null,
	// 	",": null,
	// 	"\"'": null
	// }
}

// Unmarshalは、JSONオブジェクト名とGo構造体フィールドを
// 大文字・小文字を区別して一致させますが、"case:ignore"オプションを使うことで
// 大文字・小文字を区別しない一致に設定できます。
// これにより、camelCase、snake_case、kebab-caseなどの命名規則を使った入力からアンマーシャルできます。
func Example_caseSensitivity() {
	// 様々な命名規則を使ったJSON入力。
	const input = `[
        {"firstname": true},
        {"firstName": true},
        {"FirstName": true},
        {"FIRSTNAME": true},
        {"first_name": true},
        {"FIRST_NAME": true},
        {"first-name": true},
        {"FIRST-NAME": true},
        {"unknown": true}
    ]`

	// "case:ignore"なしの場合、Unmarshalは完全一致のみを探します。
	var caseStrict []struct {
		X bool `json:"firstName"`
	}
	if err := json.Unmarshal([]byte(input), &caseStrict); err != nil {
		log.Fatal(err)
	}
	fmt.Println(caseStrict) // 完全一致が1つだけ見つかる

	// "case:ignore"ありの場合、Unmarshalはまず完全一致を探し、
	// 見つからなければ大文字・小文字を区別しない一致を探します。
	var caseIgnore []struct {
		X bool `json:"firstName,case:ignore"`
	}
	if err := json.Unmarshal([]byte(input), &caseIgnore); err != nil {
		log.Fatal(err)
	}
	fmt.Println(caseIgnore) // 8つ一致が見つかる

	// Output:
	// [{false} {true} {false} {false} {false} {false} {false} {false} {false}]
	// [{true} {true} {true} {true} {true} {true} {true} {true} {false}]
}

// Go構造体フィールドは、入力Go値または出力JSONエンコーディングのいずれかに応じて
// 出力から省略される場合があります。
// "omitzero"オプションは、フィールドがGoのゼロ値である場合や
// "IsZero() bool"メソッドがtrueを返す場合にフィールドを省略します。
// "omitempty"オプションは、フィールドが空のJSON値としてエンコードされる場合に省略します。
// ここで空のJSON値とは、JSON null、空文字列、空オブジェクト、空配列を指します。
// 多くの場合、"omitzero"と"omitempty"の挙動は同等です。
// 両方で期待する効果が得られる場合は、"omitzero"の使用が推奨されます。
func Example_omitFields() {
	type MyStruct struct {
		Foo string `json:",omitzero"`
		Bar []int  `json:",omitempty"`
		// "omitzero"と"omitempty"は両方同時に指定できます。
		// その場合、どちらかの条件を満たせばフィールドは省略されます。
		// この場合、Bazフィールドはnilポインタであるか、
		// 空のJSONオブジェクトとしてエンコードされる場合に省略されます。
		Baz *MyStruct `json:",omitzero,omitempty"`
	}

	// "omitzero"の挙動を示します。
	b, err := json.Marshal(struct {
		Bool         bool        `json:",omitzero"`
		Int          int         `json:",omitzero"`
		String       string      `json:",omitzero"`
		Time         time.Time   `json:",omitzero"`
		Addr         netip.Addr  `json:",omitzero"`
		Struct       MyStruct    `json:",omitzero"`
		SliceNil     []int       `json:",omitzero"`
		Slice        []int       `json:",omitzero"`
		MapNil       map[int]int `json:",omitzero"`
		Map          map[int]int `json:",omitzero"`
		PointerNil   *string     `json:",omitzero"`
		Pointer      *string     `json:",omitzero"`
		InterfaceNil any         `json:",omitzero"`
		Interface    any         `json:",omitzero"`
	}{
		// Boolは、falseがGoのbool型のゼロ値であるため省略されます。
		Bool: false,
		// Intは、0がGoのint型のゼロ値であるため省略されます。
		Int: 0,
		// Stringは、""がGoのstring型のゼロ値であるため省略されます。
		String: "",
		// Timeは、time.Time.IsZeroがtrueを返すため省略されます。
		Time: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		// Addrは、netip.Addr{}がGoの構造体のゼロ値であるため省略されます。
		Addr: netip.Addr{},
		// Structは、Goの構造体のゼロ値ではないため省略されません。
		Struct: MyStruct{Bar: []int{}, Baz: new(MyStruct)},
		// SliceNilは、nilがGoのスライスのゼロ値であるため省略されます。
		SliceNil: nil,
		// Sliceは、[]int{}がGoのスライスのゼロ値ではないため省略されません。
		Slice: []int{},
		// MapNilは、nilがGoのマップのゼロ値であるため省略されます。
		MapNil: nil,
		// Mapは、map[int]int{}がGoのマップのゼロ値ではないため省略されません。
		Map: map[int]int{},
		// PointerNilは、nilがGoのポインタのゼロ値であるため省略されます。
		PointerNil: nil,
		// Pointerは、new(string)がGoのポインタのゼロ値ではないため省略されません。
		Pointer: new(string),
		// InterfaceNilは、nilがGoのインターフェースのゼロ値であるため省略されます。
		InterfaceNil: nil,
		// Interfaceは、(*string)(nil)がGoのインターフェースのゼロ値ではないため省略されません。
		Interface: (*string)(nil),
	})
	if err != nil {
		log.Fatal(err)
	}
	(*jsontext.Value)(&b).Indent()      // 可読性のためインデント
	fmt.Println("OmitZero:", string(b)) // "Struct", "Slice", "Map", "Pointer", "Interface"が出力されます

	// "omitempty"の挙動を示します。
	b, err = json.Marshal(struct {
		Bool         bool        `json:",omitempty"`
		Int          int         `json:",omitempty"`
		String       string      `json:",omitempty"`
		Time         time.Time   `json:",omitempty"`
		Addr         netip.Addr  `json:",omitempty"`
		Struct       MyStruct    `json:",omitempty"`
		Slice        []int       `json:",omitempty"`
		Map          map[int]int `json:",omitempty"`
		PointerNil   *string     `json:",omitempty"`
		Pointer      *string     `json:",omitempty"`
		InterfaceNil any         `json:",omitempty"`
		Interface    any         `json:",omitempty"`
	}{
		// Boolは、省略されません。falseは空のJSON値ではないためです。
		Bool: false,
		// Intは、省略されません。0は空のJSON値ではないためです。
		Int: 0,
		// Stringは、省略されます。""は空のJSON文字列だからです。
		String: "",
		// Timeは、省略されません。これは空でないJSON文字列としてエンコードされるためです。
		Time: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		// Addrは、省略されます。これは空のJSON文字列としてエンコードされるためです。
		Addr: netip.Addr{},
		// Structは、省略されます。{}は空のJSONオブジェクトだからです。
		Struct: MyStruct{Bar: []int{}, Baz: new(MyStruct)},
		// Sliceは、省略されます。[]は空のJSON配列だからです。
		Slice: []int{},
		// Mapは、省略されます。{}は空のJSONオブジェクトだからです。
		Map: map[int]int{},
		// PointerNilは、省略されます。nullは空のJSON値だからです。
		PointerNil: nil,
		// Pointerは、省略されます。""は空のJSON文字列だからです。
		Pointer: new(string),
		// InterfaceNilは、省略されます。nullは空のJSON値だからです。
		InterfaceNil: nil,
		// Interfaceは、省略されます。nullは空のJSON値だからです。
		Interface: (*string)(nil),
	})
	if err != nil {
		log.Fatal(err)
	}
	(*jsontext.Value)(&b).Indent()       // indent for readability
	fmt.Println("OmitEmpty:", string(b)) // outputs "Bool", "Int", and "Time"

	// Output:
	// OmitZero: {
	// 	"Struct": {},
	// 	"Slice": [],
	// 	"Map": {},
	// 	"Pointer": "",
	// 	"Interface": null
	// }
	// OmitEmpty: {
	// 	"Bool": false,
	// 	"Int": 0,
	// 	"Time": "0001-01-01T00:00:00Z"
	// }
}

// JSONオブジェクトは、Go構造体が親構造体に埋め込まれるのと同様に
// 親オブジェクト内にインライン化できます。
// インライン化のルールはGoの埋め込みと似ていますが、JSONの名前空間上で動作します。
func Example_inlinedFields() {
	// BaseはContainerに埋め込まれています。
	type Base struct {
		// IDはContainerのJSONオブジェクトに昇格されます。
		ID string
		// TypeはContainer.Typeが存在するため無視されます。
		Type string
		// TimeはContainer.Inlined.Timeと打ち消し合います。
		Time time.Time
	}
	// OtherはContainerに埋め込まれています。
	type Other struct{ Cost float64 }
	// ContainerはBaseとOtherを埋め込みます。
	type Container struct {
		// Baseは埋め込み構造体であり、暗黙的にJSONインライン化されます。
		Base
		// TypeはBase.Typeより優先されます。
		Type int
		// Inlinedは名前付きGoフィールドですが、明示的にJSONインライン化されています。
		Inlined struct {
			// UserはContainerのJSONオブジェクトに昇格されます。
			User string
			// TimeはBase.Timeと打ち消し合います。
			Time string
		} `json:",inline"`
		// IDはJSON名が異なるため、Base.IDと競合しません。
		ID string `json:"uuid"`
		// Otherは明示的なJSON名があるためJSONインライン化されません。
		Other `json:"other"`
	}

	// 空のContainerをフォーマットして、どのフィールドがJSONシリアライズ可能かを表示します。
	var input Container
	b, err := json.Marshal(&input)
	if err != nil {
		log.Fatal(err)
	}
	(*jsontext.Value)(&b).Indent() // 可読性のためインデント
	fmt.Println(string(b))

	// Output:
	// {
	// 	"ID": "",
	// 	"Type": 0,
	// 	"User": "",
	// 	"uuid": "",
	// 	"other": {
	// 		"Cost": 0
	// 	}
	// }
}

// バージョンの違いにより、コンパイル時に既知のJSONオブジェクトメンバー集合と
// 実行時に遭遇するメンバー集合が異なる場合があります。
// そのため、未知のメンバーを細かく制御できると便利です。
// このパッケージは、未知のメンバーの保持・拒否・破棄をサポートします。
func Example_unknownMembers() {
	const input = `{
		"Name": "Teal",
		"Value": "#008080",
		"WebSafe": false
	}`
	type Color struct {
		Name  string
		Value string

		// Unknownは未知のJSONオブジェクトメンバーを保持するGo構造体フィールドです。
		// "unknown"タグオプションを指定することでこの挙動になります。
		//
		// 型はjsontext.Valueまたはmap[string]Tが利用できます。
		Unknown jsontext.Value `json:",unknown"`
	}

	// デフォルトでは、未知のメンバーは "unknown" とマークされたGoフィールドに格納されます。
	// そのようなフィールドが存在しない場合は無視されます。
	var color Color
	err := json.Unmarshal([]byte(input), &color)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unknown members:", string(color.Unknown))

	// RejectUnknownMembersを指定すると、Unmarshalは
	// 未知のメンバーが存在する場合に拒否します。
	err = json.Unmarshal([]byte(input), new(Color), json.RejectUnknownMembers(true))
	var serr *json.SemanticError
	if errors.As(err, &serr) && serr.Err == json.ErrUnknownName {
		fmt.Println("Unmarshal error:", serr.Err, strconv.Quote(serr.JSONPointer.LastToken()))
	}

	// デフォルトでは、Marshalは "unknown" とマークされた
	// Go構造体フィールドに格納された未知のメンバーを保持します。
	b, err := json.Marshal(color)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Output with unknown members:   ", string(b))

	// DiscardUnknownMembersを指定すると、Marshalは
	// 未知のメンバーを破棄します。
	b, err = json.Marshal(color, json.DiscardUnknownMembers(true))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Output without unknown members:", string(b))

	// Output:
	// Unknown members: {"WebSafe":false}
	// Unmarshal error: unknown object member name "WebSafe"
	// Output with unknown members:    {"Name":"Teal","Value":"#008080","WebSafe":false}
	// Output without unknown members: {"Name":"Teal","Value":"#008080"}
}

// "format"タグオプションを使うことで、特定の型の書式を変更できます。
func Example_formatFlags() {
	value := struct {
		BytesBase64     []byte         `json:",format:base64"`
		BytesHex        [8]byte        `json:",format:hex"`
		BytesArray      []byte         `json:",format:array"`
		FloatNonFinite  float64        `json:",format:nonfinite"`
		MapEmitNull     map[string]any `json:",format:emitnull"`
		SliceEmitNull   []any          `json:",format:emitnull"`
		TimeDateOnly    time.Time      `json:",format:'2006-01-02'"`
		TimeUnixSec     time.Time      `json:",format:unix"`
		DurationSecs    time.Duration  `json:",format:sec"`
		DurationNanos   time.Duration  `json:",format:nano"`
		DurationISO8601 time.Duration  `json:",format:iso8601"`
	}{
		BytesBase64:     []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		BytesHex:        [8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		BytesArray:      []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		FloatNonFinite:  math.NaN(),
		MapEmitNull:     nil,
		SliceEmitNull:   nil,
		TimeDateOnly:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		TimeUnixSec:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		DurationSecs:    12*time.Hour + 34*time.Minute + 56*time.Second + 7*time.Millisecond + 8*time.Microsecond + 9*time.Nanosecond,
		DurationNanos:   12*time.Hour + 34*time.Minute + 56*time.Second + 7*time.Millisecond + 8*time.Microsecond + 9*time.Nanosecond,
		DurationISO8601: 12*time.Hour + 34*time.Minute + 56*time.Second + 7*time.Millisecond + 8*time.Microsecond + 9*time.Nanosecond,
	}

	b, err := json.Marshal(&value)
	if err != nil {
		log.Fatal(err)
	}
	(*jsontext.Value)(&b).Indent() // 可読性のためインデント
	fmt.Println(string(b))

	// Output:
	// {
	// 	"BytesBase64": "ASNFZ4mrze8=",
	// 	"BytesHex": "0123456789abcdef",
	// 	"BytesArray": [
	// 		1,
	// 		35,
	// 		69,
	// 		103,
	// 		137,
	// 		171,
	// 		205,
	// 		239
	// 	],
	// 	"FloatNonFinite": "NaN",
	// 	"MapEmitNull": null,
	// 	"SliceEmitNull": null,
	//	"TimeDateOnly": "2000-01-01",
	//	"TimeUnixSec": 946684800,
	//	"DurationSecs": 45296.007008009,
	//	"DurationNanos": 45296007008009,
	//	"DurationISO8601": "PT12H34M56.007008009S"
	// }
}

// HTTPエンドポイントを実装する際、[io.Reader] や [io.Writer] を扱うことが一般的です。
// [MarshalWrite] と [UnmarshalRead] 関数は、こうした入出力型の操作を補助します。
// [UnmarshalRead] は [io.Reader] 全体を読み込み、トップレベルのJSON値の後に
// 予期しないバイトがないことを確認します。
func Example_serveHTTP() {
	// サーバーが保持するグローバルな状態。
	var n int64

	// "add"エンドポイントは、JSONオブジェクトを含むPOSTリクエストを受け付け、
	// サーバーのグローバルカウンターに数値をアトミックに加算します。
	// 更新されたカウンターの値を返します。
	http.HandleFunc("/api/add", func(w http.ResponseWriter, r *http.Request) {
		// クライアントからのリクエストをアンマーシャルします。
		var val struct{ N int64 }
		if err := json.UnmarshalRead(r.Body, &val); err != nil {
			// 入力のアンマーシャルに失敗した場合は、クライアント側の問題を示唆します。
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// サーバーからのレスポンスをマーシャルします。
		val.N = atomic.AddInt64(&n, val.N)
		if err := json.MarshalWrite(w, &val); err != nil {
			// 出力のマーシャルに失敗した場合は、サーバー側の問題を示唆します。
			// このエラーは、json.MarshalWriteがすでに出力に書き込んでいる場合、
			// クライアントからは常に観測できるとは限りません。
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

// 一部のGo型は、独自のJSON表現を外部パッケージに委譲しています。
// そのため、"json"パッケージは外部実装の使い方を知りません。
// 例えば、[google.golang.org/protobuf/encoding/protojson] パッケージは
// すべての [google.golang.org/protobuf/proto.Message] 型のJSONを実装しています。
// [WithMarshalers] や [WithUnmarshalers] を使うことで
// "json"と"protojson"を連携させることができます。
func Example_protoJSON() {
	// protoMessage を "google.golang.org/protobuf/proto".Message とします。
	type protoMessage interface{ ProtoReflect() }
	// foopbMyMessage を proto.Message の具体的な実装とします。
	type foopbMyMessage struct{ protoMessage }
	// protojson を "google.golang.org/protobuf/encoding/protojson" のインポートとします。
	var protojson struct {
		Marshal   func(protoMessage) ([]byte, error)
		Unmarshal func([]byte, protoMessage) error
	}

	// この値は、非proto.Message型とproto.Message型の両方を混在させています。
	// 非proto.Message型には "json" パッケージを、
	// proto.Message型には "protojson" パッケージを使うべきです。
	var value struct {
		// GoStructはproto.Messageを実装していないため、
		// "json"パッケージのデフォルト動作を使うべきです。
		GoStruct struct {
			Name string
			Age  int
		}

		// ProtoMessageはproto.Messageを実装しているため、
		// protojson.Marshalで処理すべきです。
		ProtoMessage *foopbMyMessage
	}

	// proto.Message型にはprotojson.Marshalを使ってマーシャルします。
	b, err := json.Marshal(&value,
		// protojson.Marshalを型固有のマーシャラーとして使います。
		json.WithMarshalers(json.MarshalFunc(protojson.Marshal)))
	if err != nil {
		log.Fatal(err)
	}

	// proto.Message型にはprotojson.Unmarshalを使ってアンマーシャルします。
	err = json.Unmarshal(b, &value,
		// protojson.Unmarshalを型固有のアンマーシャラーとして使います。
		json.WithUnmarshalers(json.UnmarshalFunc(protojson.Unmarshal)))
	if err != nil {
		log.Fatal(err)
	}
}

// 多くのエラー型は、Goの構造体で公開フィールドを持たないためシリアライズできません
// （例: [errors.New] で生成されたエラーなど）。
// 一部のアプリケーションでは、アンマーシャルできないエラーであっても
// エラーをJSON文字列としてマーシャルしたい場合があります。
func ExampleWithMarshalers_errors() {
	// シリアライズ時にGoエラーがいくつか発生したレスポンス。
	response := []struct {
		Result string `json:",omitzero"`
		Error  error  `json:",omitzero"`
	}{
		{Result: "Oranges are a good source of Vitamin C."},
		{Error: &strconv.NumError{Func: "ParseUint", Num: "-1234", Err: strconv.ErrSyntax}},
		{Error: &os.PathError{Op: "ReadFile", Path: "/path/to/secret/file", Err: os.ErrPermission}},
	}

	b, err := json.Marshal(&response,
		// エラー型をマーシャルするすべての試みをインターセプトします。
		json.WithMarshalers(json.JoinMarshalers(
			// 例えば、strconv.NumErrorは安全にシリアライズできると仮定します:
			// この型固有のマーシャル関数はこの型をインターセプトし、
			// エラーメッセージをJSON文字列としてエンコードします。
			json.MarshalToFunc(func(enc *jsontext.Encoder, err *strconv.NumError) error {
				return enc.WriteToken(jsontext.String(err.Error()))
			}),
			// エラーメッセージには機密情報が含まれる場合があり、
			// シリアライズに適さないことがあります。上記で処理されなかった
			// すべてのエラーには汎用的なエラーメッセージを返します。
			json.MarshalFunc(func(error) ([]byte, error) {
				return []byte(`"internal server error"`), nil
			}),
		)),
		jsontext.Multiline(true)) // 可読性のため展開
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	// Output:
	// [
	// 	{
	// 		"Result": "Oranges are a good source of Vitamin C."
	// 	},
	// 	{
	// 		"Error": "strconv.ParseUint: parsing \"-1234\": invalid syntax"
	// 	},
	// 	{
	// 		"Error": "internal server error"
	// 	}
	// ]
}

// 一部のアプリケーションでは、JSON数値の正確な精度を
// アンマーシャル時に保持する必要があります。これは型固有の
// アンマーシャル関数を使い、any型へのすべてのアンマーシャルをインターセプトして
// インターフェース値を [jsontext.Value] で事前に埋めることで実現できます。
// [jsontext.Value] はJSON数値を正確に表現できます。
func ExampleWithUnmarshalers_rawNumber() {
	// float64の表現を超えるJSON数値を含む入力。
	const input = `[false, 1e-1000, 3.141592653589793238462643383279, 1e+1000, true]`

	var value any
	err := json.Unmarshal([]byte(input), &value,
		// any型へのすべてのアンマーシャル試行をインターセプトします。
		json.WithUnmarshalers(
			json.UnmarshalFromFunc(func(dec *jsontext.Decoder, val *any) error {
				// 次にデコードされる値がJSON数値の場合、
				// アンマーシャル先のGo型を具体的に指定します。
				if dec.PeekKind() == '0' {
					*val = jsontext.Value(nil)
				}
				// SkipFuncを返してデフォルトのアンマーシャル動作にフォールバックします。
				return json.SkipFunc
			}),
		))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

	// 正常性チェック。
	want := []any{false, jsontext.Value("1e-1000"), jsontext.Value("3.141592653589793238462643383279"), jsontext.Value("1e+1000"), true}
	if !reflect.DeepEqual(value, want) {
		log.Fatalf("value mismatch:\ngot  %v\nwant %v", value, want)
	}

	// Output:
	// [false 1e-1000 3.141592653589793238462643383279 1e+1000 true]
}

// JSONで設定ファイルをパースする場合、
// パース処理は入力のどこでエラーが発生したかを示すために
// 行番号や列番号付きでエラーを報告する必要があることがよくあります。
func ExampleWithUnmarshalers_recordOffsets() {
	// 仮想的な設定ファイル。
	const input = `[
					{"Source": "192.168.0.100:1234", "Destination": "192.168.0.1:80"},
					{"Source": "192.168.0.251:4004"},
					{"Source": "192.168.0.165:8080", "Destination": "0.0.0.0:80"}
			]`
	type Tunnel struct {
		Source      netip.AddrPort
		Destination netip.AddrPort

		// ByteOffsetはアンマーシャル時に、このGo構造体のJSONオブジェクトが
		// JSON入力内のどのバイトオフセットにあるかを記録します。
		ByteOffset int64 `json:"-"` // JSONシリアライズ時は無視されるメタデータ
	}

	var tunnels []Tunnel
	err := json.Unmarshal([]byte(input), &tunnels,
		// Tunnel型へのすべてのアンマーシャル試行をインターセプトします。
		json.WithUnmarshalers(
			json.UnmarshalFromFunc(func(dec *jsontext.Decoder, tunnel *Tunnel) error {
				// Decoder.InputOffsetは直前のトークンの後のオフセットを報告しますが、
				// 次のトークンの前のオフセットを記録したい場合があります。
				//
				// Decoder.PeekKindを呼び出して、次のトークンまで十分にバッファリングします。
				// 先頭の空白、カンマ、コロンの数を加算して
				// 次のトークンの開始位置を特定します。
				dec.PeekKind()
				unread := dec.UnreadBuffer()
				n := len(unread) - len(bytes.TrimLeft(unread, " \n\r\t,:"))
				tunnel.ByteOffset = dec.InputOffset() + int64(n)

				// SkipFuncを返してデフォルトのアンマーシャル動作にフォールバックします。
				return json.SkipFunc
			}),
		))
	if err != nil {
		log.Fatal(err)
	}

	// lineColumnはバイトオフセットを1始まりの行・列番号に変換します。
	// オフセットはinputの範囲内である必要があります。
	lineColumn := func(input string, offset int) (line, column int) {
		line = 1 + strings.Count(input[:offset], "\n")
		column = 1 + offset - (strings.LastIndex(input[:offset], "\n") + len("\n"))
		return line, column
	}

	// 設定ファイルが有効かどうかを検証します。
	for _, tunnel := range tunnels {
		if !tunnel.Source.IsValid() || !tunnel.Destination.IsValid() {
			line, column := lineColumn(input, int(tunnel.ByteOffset))
			fmt.Printf("%d:%d: source and destination must both be specified", line, column)
		}
	}

	// Output:
	// 3:3: source and destination must both be specified
}
