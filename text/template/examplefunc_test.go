// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template_test

import (
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/strings"
	"github.com/shogo82148/std/text/template"
)

// This example demonstrates a custom function to process template text.
// It installs the strings.Title function and uses it to
// Make Title Text Look Good In Our Template's Output.
func ExampleTemplate_func() {
	// 最初に、関数を登録するためのFuncMapを作成します。
	funcMap := template.FuncMap{
		// 名前 "title" は、テンプレートテキスト内で関数が呼ばれる名前です。
		"title": strings.Title,
	}

	// 関数をテストするためのシンプルなテンプレート定義。
	// 入力テキストをいくつかの方法で出力します：
	// - オリジナル
	// - タイトルケース
	// - タイトルケースにした後に %q で出力
	// - %q で出力した後にタイトルケースにします。
	const templateText = `
Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
`

	// テンプレートを作成し、関数マップを追加し、テキストを解析します。
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// テンプレートを実行して出力を確認します。
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	// Output:
	// Input: "the go programming language"
	// Output 0: The Go Programming Language
	// Output 1: "The Go Programming Language"
	// Output 2: "The Go Programming Language"
}
