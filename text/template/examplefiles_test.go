// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template_test

import (
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/path/filepath"
	"github.com/shogo82148/std/text/template"
)

// ここでは、ディレクトリから一連のテンプレートをロードする方法を示します。
func ExampleTemplate_glob() {
	// ここでは、一時的なディレクトリを作成し、それをサンプルの
	// テンプレート定義ファイルで満たします。通常、テンプレートファイルはすでに
	// プログラムが知っている何らかの場所に存在します。
	dir := createTestDir([]templateFile{
		// T0.tmplは、単にT1を呼び出すプレーンなテンプレートファイルです。
		{"T0.tmpl", `T0 invokes T1: ({{template "T1"}})`},
		// T1.tmplは、T2を呼び出すテンプレート、T1を定義します。
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
		// T2.tmplは、テンプレートT2を定義します。
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})
	// テスト後のクリーンアップ；これも例として実行する際の特性です。
	defer os.RemoveAll(dir)

	// patternは、すべてのテンプレートファイルを見つけるために使用されるglobパターンです。
	pattern := filepath.Join(dir, "*.tmpl")

	// ここからが実際の例です。
	// T0.tmplは最初にマッチした名前なので、それが開始テンプレートとなり、
	// ParseGlobによって返される値となります。
	tmpl := template.Must(template.ParseGlob(pattern))

	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
	// Output:
	// T0 invokes T1: (T1 invokes T2: (This is T2))
}

// この例は、いくつかのテンプレートを共有し、それらを異なるコンテキストで
// 使用する一つの方法を示しています。このバリアントでは、既存のテンプレートの
// バンドルに手動で複数のドライバーテンプレートを追加します。
func ExampleTemplate_helpers() {
	// ここでは、一時的なディレクトリを作成し、それをサンプルの
	// テンプレート定義ファイルで満たします。通常、テンプレートファイルはすでに
	// プログラムが知っている何らかの場所に存在します。
	dir := createTestDir([]templateFile{
		// T1.tmplは、T2を呼び出すテンプレート、T1を定義します。
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
		// T2.tmplは、テンプレートT2を定義します。
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})
	// テスト後のクリーンアップ；これも例として実行する際の特性です。
	defer os.RemoveAll(dir)

	// patternは、すべてのテンプレートファイルを見つけるために使用されるglobパターンです。
	pattern := filepath.Join(dir, "*.tmpl")

	// ここからが実際の例です。
	// ヘルパーをロードします。
	templates := template.Must(template.ParseGlob(pattern))
	// 明示的なテンプレート定義を使用して、一連のテンプレートに一つのドライバーテンプレートを追加します。
	_, err := templates.Parse("{{define `driver1`}}Driver 1 calls T1: ({{template `T1`}})\n{{end}}")
	if err != nil {
		log.Fatal("parsing driver1: ", err)
	}
	// 別のドライバーテンプレートを追加します。
	_, err = templates.Parse("{{define `driver2`}}Driver 2 calls T2: ({{template `T2`}})\n{{end}}")
	if err != nil {
		log.Fatal("parsing driver2: ", err)
	}
	// 実行前にすべてのテンプレートをロードします。このパッケージはそのような振る舞いを
	// 必要としませんが、html/templateのエスケープ処理はそれを必要とするので、それは良い習慣です。
	err = templates.ExecuteTemplate(os.Stdout, "driver1", nil)
	if err != nil {
		log.Fatalf("driver1 execution: %s", err)
	}
	err = templates.ExecuteTemplate(os.Stdout, "driver2", nil)
	if err != nil {
		log.Fatalf("driver2 execution: %s", err)
	}
	// Output:
	// Driver 1 calls T1: (T1 invokes T2: (This is T2))
	// Driver 2 calls T2: (This is T2)
}

// この例は、一つのグループのドライバーテンプレートを、異なるセットのヘルパーテンプレートと
// どのように使用するかを示しています。
func ExampleTemplate_share() {
	// ここでは、一時的なディレクトリを作成し、それをサンプルの
	// テンプレート定義ファイルで満たします。通常、テンプレートファイルはすでに
	// プログラムが知っている何らかの場所に存在します。
	dir := createTestDir([]templateFile{
		// T0.tmplは、単にT1を呼び出すプレーンなテンプレートファイルです。
		{"T0.tmpl", "T0 ({{.}} version) invokes T1: ({{template `T1`}})\n"},
		// T1.tmplは、T2を呼び出すテンプレート、T1を定義します。注意：T2は定義されていません
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
	})
	// テスト後のクリーンアップ；これも例として実行する際の特性です。
	defer os.RemoveAll(dir)

	// patternは、すべてのテンプレートファイルを見つけるために使用されるglobパターンです。
	pattern := filepath.Join(dir, "*.tmpl")

	// ここからが実際の例です。
	// ドライバーをロードします。
	drivers := template.Must(template.ParseGlob(pattern))

	// 我々はT2テンプレートの実装を定義しなければなりません。まず、
	// ドライバーをクローンし、その後、テンプレート名前空間にT2の定義を追加します。

	// 1. ヘルパーセットをクローンして、それらを実行するための新しい名前空間を作成します。
	first, err := drivers.Clone()
	if err != nil {
		log.Fatal("cloning helpers: ", err)
	}
	// 2. T2、バージョンAを定義し、解析します。
	_, err = first.Parse("{{define `T2`}}T2, version A{{end}}")
	if err != nil {
		log.Fatal("parsing T2: ", err)
	}

	// 今度は全体を繰り返しますが、T2の別のバージョンを使用します。
	// 1. ドライバーをクローンします。
	second, err := drivers.Clone()
	if err != nil {
		log.Fatal("cloning drivers: ", err)
	}
	// 2. T2、バージョンBを定義し、解析します。
	_, err = second.Parse("{{define `T2`}}T2, version B{{end}}")
	if err != nil {
		log.Fatal("parsing T2: ", err)
	}

	// テンプレートを逆の順序で実行して、
	// 最初のものが二番目のものに影響を受けないことを確認します。
	err = second.ExecuteTemplate(os.Stdout, "T0.tmpl", "second")
	if err != nil {
		log.Fatalf("second execution: %s", err)
	}
	err = first.ExecuteTemplate(os.Stdout, "T0.tmpl", "first")
	if err != nil {
		log.Fatalf("first: execution: %s", err)
	}

	// Output:
	// T0 (second version) invokes T1: (T1 invokes T2: (T2, version B))
	// T0 (first version) invokes T1: (T1 invokes T2: (T2, version A))
}
