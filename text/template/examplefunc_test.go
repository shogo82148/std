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

// This example demonstrates registering two custom template functions
// and how to overwite one of the functions after the template has been
// parsed. Overwriting can be used, for example, to alter the operation
// of cloned templates.
func ExampleTemplate_funcs() {

	// Define a simple template to test the functions.
	const tmpl = `{{ . | lower | repeat }}`

	// Define the template funcMap with two functions.
	var funcMap = template.FuncMap{
		"lower":  strings.ToLower,
		"repeat": func(s string) string { return strings.Repeat(s, 2) },
	}

	// Define a New template, add the funcMap using Funcs and then Parse
	// the template.
	parsedTmpl, err := template.New("t").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	if err := parsedTmpl.Execute(os.Stdout, "ABC\n"); err != nil {
		log.Fatal(err)
	}

	// [Funcs] must be called before a template is parsed to add
	// functions to the template. [Funcs] can also be used after a
	// template is parsed to overwrite template functions.
	//
	// Here the function identified by "repeat" is overwritten.
	parsedTmpl.Funcs(template.FuncMap{
		"repeat": func(s string) string { return strings.Repeat(s, 3) },
	})
	if err := parsedTmpl.Execute(os.Stdout, "DEF\n"); err != nil {
		log.Fatal(err)
	}
	// Output:
	// abc
	// abc
	// def
	// def
	// def
}

// This example demonstrates how to use "if".
func ExampleTemplate_if() {
	type book struct {
		Stars float32
		Name  string
	}

	tpl, err := template.New("book").Parse(`{{ if (gt .Stars 4.0) }}"{{.Name }}" is a great book.{{ else }}"{{.Name}}" is not a great book.{{ end }}`)
	if err != nil {
		log.Fatalf("failed to parse template: %s", err)
	}

	b := &book{
		Stars: 4.9,
		Name:  "Good Night, Gopher",
	}
	err = tpl.Execute(os.Stdout, b)
	if err != nil {
		log.Fatalf("failed to execute template: %s", err)
	}

	// Output:
	// "Good Night, Gopher" is a great book.
}
