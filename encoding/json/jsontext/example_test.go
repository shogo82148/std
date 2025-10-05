// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/encoding/json/jsontext"
	"github.com/shogo82148/std/encoding/json/v2"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/strings"
)

// この例は、[Encoder] と [Decoder] を使って、
// JSONを具体的なGo型にアンマーシャルせずにパース・修正する方法を示します。
func Example_stringReplace() {
	// 「Go」の代わりに非慣用的な「Golang」を使った例の入力。
	const input = `{
		"title": "Golang version 1 is released",
		"author": "Andrew Gerrand",
		"date": "2012-03-28",
		"text": "Today marks a major milestone in the development of the Golang programming language.",
		"otherArticles": [
			"Twelve Years of Golang",
			"The Laws of Reflection",
			"Learn Golang from your browser"
		]
	}`

	// DecoderとEncoderを使うことで、すべてのトークンをパースし、
	// 必要に応じてトークンをチェック・修正し、
	// 出力にトークンを書き込むことができます。
	var replacements []jsontext.Pointer
	in := strings.NewReader(input)
	dec := jsontext.NewDecoder(in)
	out := new(bytes.Buffer)
	enc := jsontext.NewEncoder(out, jsontext.Multiline(true)) // 可読性向上のため展開
	for {
		// 入力からトークンを読み取ります。
		tok, err := dec.ReadToken()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		// トークンに "Golang" という文字列が含まれているかをチェックし、
		// 見つかった箇所をすべて "Go" に置き換えます。
		if tok.Kind() == '"' && strings.Contains(tok.String(), "Golang") {
			replacements = append(replacements, dec.StackPointer())
			tok = jsontext.String(strings.ReplaceAll(tok.String(), "Golang", "Go"))
		}

		// （場合によっては修正された）トークンを出力に書き込みます。
		if err := enc.WriteToken(tok); err != nil {
			log.Fatal(err)
		}
	}

	// 置換箇所の一覧と修正後のJSON出力を表示します。
	if len(replacements) > 0 {
		fmt.Println(`Replaced "Golang" with "Go" in:`)
		for _, where := range replacements {
			fmt.Println("\t" + where)
		}
		fmt.Println()
	}
	fmt.Println("Result:", out.String())

	// Output:
	// Replaced "Golang" with "Go" in:
	// 	/title
	// 	/text
	// 	/otherArticles/0
	// 	/otherArticles/2
	//
	// Result: {
	// 	"title": "Go version 1 is released",
	// 	"author": "Andrew Gerrand",
	// 	"date": "2012-03-28",
	// 	"text": "Today marks a major milestone in the development of the Go programming language.",
	// 	"otherArticles": [
	// 		"Twelve Years of Go",
	// 		"The Laws of Reflection",
	// 		"Learn Go from your browser"
	// 	]
	// }
}

// HTML内にJSONを直接埋め込む場合は、安全性のため特別な処理が必要です。
// JSONがHTMLとして直接扱われた際に<script>インジェクションなどができないよう、特定のルーンをエスケープします。
//
// この例は、v1の [encoding/json] パッケージで提供されていた同等の動作を得る方法を示しますが、
// このパッケージでは直接サポートされなくなっています。
// 新しくJSONとHTMLを混在させるコードを書く場合は、安全のため [github.com/google/safehtml] モジュールを使用してください。
func ExampleEscapeForHTML() {
	page := struct {
		Title string
		Body  string
	}{
		Title: "Example Embedded Javascript",
		Body:  `<script> console.log("Hello, world!"); </script>`,
	}

	b, err := json.Marshal(&page,
		// JSON文字列内の特定のルーンをエスケープすることで、
		// JSONをHTML内に直接埋め込んでも安全になるようにします。
		jsontext.EscapeForHTML(true),
		jsontext.EscapeForJS(true),
		jsontext.Multiline(true)) // 可読性向上のため展開
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	// Output:
	// {
	// 	"Title": "Example Embedded Javascript",
	// 	"Body": "\u003cscript\u003e console.log(\"Hello, world!\"); \u003c/script\u003e"
	// }
}
