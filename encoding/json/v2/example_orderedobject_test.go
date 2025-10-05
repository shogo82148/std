// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json_test

import (
	"encoding/json"

	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/reflect"

	"github.com/shogo82148/std/encoding/json/jsontext"
)

// JSONオブジェクトの正確な順序は、[MarshalerTo]と[UnmarshalerFrom]を実装した
// 専用型を使うことで保持できます。
func Example_orderedObject() {
	// 順序付きオブジェクトをマーシャル・アンマーシャルしてラウンドトリップします。
	// JSONオブジェクトメンバーの順序と重複が保持されることを期待します。
	// このオブジェクトには"fizz"が2回含まれるため、jsontext.AllowDuplicateNamesを指定します。
	want := OrderedObject[string]{
		{"fizz", "buzz"},
		{"hello", "world"},
		{"fizz", "wuzz"},
	}
	b, err := json.Marshal(&want, jsontext.AllowDuplicateNames(true))
	if err != nil {
		log.Fatal(err)
	}
	var got OrderedObject[string]
	err = json.Unmarshal(b, &got, jsontext.AllowDuplicateNames(true))
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
	// 	"fizz": "buzz",
	// 	"hello": "world",
	// 	"fizz": "wuzz"
	// }
}
