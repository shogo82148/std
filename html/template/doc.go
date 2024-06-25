// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
パッケージテンプレート（html/template）は、コードインジェクションに対して安全なHTML出力を生成するための
データ駆動型テンプレートを実装します。それは [text/template] と同じインターフェースを提供し、出力がHTMLの場合は
[text/template] の代わりに使用すべきです。

ここでのドキュメンテーションは、パッケージのセキュリティ機能に焦点を当てています。
テンプレート自体のプログラミングについての情報は、[text/template] のドキュメンテーションを参照してください。

# Introduction

このパッケージは[text/template]をラップして、そのテンプレートAPIを共有して
HTMLテンプレートを安全に解析し実行できます。

	tmpl, err := template.New("name").Parse(...)
	// エラーチェックは省略
	err = tmpl.Execute(out, data)

成功した場合、tmplは今後、インジェクションに対して安全になります。それ以外の場合、errはErrorCodeのドキュメントで定義されたエラーです。

HTMLテンプレートは、データ値をHTMLドキュメントに安全に埋め込むためにエンコードするべきプレーンテキストとして扱います。エスケープは文脈に依存するため、JavaScript、CSS、URIの文脈内にアクションが現れることがあります。

このパッケージが使用するセキュリティモデルは、テンプレートの作者が信頼できると仮定し、
一方でExecuteのデータパラメータは信頼できないと仮定します。詳細は以下に提供されています。

例

	import "text/template"
	...
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")

は以下を生成します

	Hello, <script>alert('you have been pwned')</script>!

しかし、html/templateの文脈自動エスケープでは

	import "html/template"
	...
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")

は安全な、エスケープされたHTML出力を生成します

	Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!

# Contexts

このパッケージはHTML、CSS、JavaScript、およびURIを理解します。それは各単純なアクションパイプラインに
サニタイジング関数を追加するので、以下の抜粋が与えられた場合

	<a href="/search?q={{.}}">{{.}}</a>

パース時に、必要に応じてエスケープ関数を追加するため、各{{.}}が上書きされます。
この場合、それは次のようになります。

	<a href="/search?q={{. | urlescaper | attrescaper}}">{{. | htmlescaper}}</a>

ここで、urlescaper、attrescaper、およびhtmlescaperは、内部エスケープ関数のエイリアスです。

これらの内部エスケープ関数については、アクションパイプラインがnilインターフェース値を評価すると、
それは空の文字列であるかのように扱われます。

# Namespaced and data- attributes

名前空間を持つ属性は、名前空間がないかのように扱われます。
以下の抜粋が与えられた場合

	<a my:href="{{.}}"></a>

パース時に、属性はまるでそれがただの"href"であるかのように扱われます。
したがって、パース時にテンプレートは次のようになります：

	<a my:href="{{. | urlescaper | attrescaper}}"></a>

同様に、"data-"プレフィックスを持つ属性は、まるでそれらが"data-"プレフィックスを持っていないかのように扱われます。したがって、以下が与えられた場合

	<a data-href="{{.}}"></a>

パース時に、これは次のようになります。

	<a data-href="{{. | urlescaper | attrescaper}}"></a>

属性が名前空間と"data-"プレフィックスの両方を持っている場合、コンテキストを決定するときには名前空間のみが削除されます。例えば

	<a my:data-href="{{.}}"></a>

これは、"my:data-href"がただの"data-href"であるかのように、そして"href"であるかのように（"data-"プレフィックスも無視される場合）扱われます。したがって、パース時には次のようになります。

	<a my:data-href="{{. | attrescaper}}"></a>

特別なケースとして、"xmlns"名前空間を持つ属性は常にURLを含んでいるとして扱われます。以下の抜粋が与えられた場合

	<a xmlns:title="{{.}}"></a>
	<a xmlns:href="{{.}}"></a>
	<a xmlns:onclick="{{.}}"></a>

パース時に、それらは次のようになります。

	<a xmlns:title="{{. | urlescaper | attrescaper}}"></a>
	<a xmlns:href="{{. | urlescaper | attrescaper}}"></a>
	<a xmlns:onclick="{{. | urlescaper | attrescaper}}"></a>

# Errors

詳細はErrorCodeのドキュメンテーションを参照してください。

# A fuller picture

このパッケージのコメントの残りの部分は、最初の読み込み時にスキップしても構いません。これには、
エスケープの文脈とエラーメッセージを理解するために必要な詳細が含まれています。ほとんどのユーザーは
これらの詳細を理解する必要はありません。

# Contexts

{{.}}が`O'Reilly: How are <i>you</i>?`と仮定すると、以下の表は
左側の文脈で{{.}}がどのように表示されるかを示しています。

	Context                          {{.}} After
	{{.}}                            O'Reilly: How are &lt;i&gt;you&lt;/i&gt;?
	<a title='{{.}}'>                O&#39;Reilly: How are you?
	<a href="/{{.}}">                O&#39;Reilly: How are %3ci%3eyou%3c/i%3e?
	<a href="?q={{.}}">              O&#39;Reilly%3a%20How%20are%3ci%3e...%3f
	<a onx='f("{{.}}")'>             O\x27Reilly: How are \x3ci\x3eyou...?
	<a onx='f({{.}})'>               "O\x27Reilly: How are \x3ci\x3eyou...?"
	<a onx='pattern = /{{.}}/;'>     O\x27Reilly: How are \x3ci\x3eyou...\x3f

安全でないコンテキストで使用された場合、その値はフィルタリングされる可能性があります：

	Context                          {{.}} After
	<a href="{{.}}">                 #ZgotmplZ

なぜなら "O'Reilly:" は "http:" のような許可されたプロトコルではないからです。

もし {{.}} が無害な単語、`left`であるなら、それはより広範に現れることができます。

	Context                              {{.}} After
	{{.}}                                left
	<a title='{{.}}'>                    left
	<a href='{{.}}'>                     left
	<a href='/{{.}}'>                    left
	<a href='?dir={{.}}'>                left
	<a style="border-{{.}}: 4px">        left
	<a style="align: {{.}}">             left
	<a style="background: '{{.}}'>       left
	<a style="background: url('{{.}}')>  left
	<style>p.{{.}} {color:red}</style>   left

非文字列の値はJavaScriptの文脈で使用できます。
もし {{.}} が

	struct{A,B string}{ "foo", "bar" }

エスケープされたテンプレート内で

	<script>var pair = {{.}};</script>

その後、テンプレートの出力は次のようになります。

	<script>var pair = {"A": "foo", "B": "bar"};</script>

JavaScriptの文脈で埋め込むために非文字列コンテンツがどのようにマーシャルされるかを理解するために、jsonパッケージを参照してください。

# Typed Strings

デフォルトでは、このパッケージはすべてのパイプラインがプレーンテキストの文字列を生成すると仮定します。
それは、そのプレーンテキスト文字列を適切な文脈で正しく安全に埋め込むために必要なエスケープパイプラインステージを追加します。

データ値がプレーンテキストでない場合、そのタイプでマークすることで、それが過度にエスケープされないようにすることができます。

Types HTML, JS, URL, and others from content.go can carry safe content that is
exempted from escaping.

テンプレート

	Hello, {{.}}!

は以下のように呼び出すことができます

	tmpl.Execute(out, template.HTML(`<b>World</b>`))

これにより

	Hello, <b>World</b>!

が生成されます。

これは、{{.}}が通常の文字列であった場合に生成される

	Hello, &lt;b&gt;World&lt;b&gt;!

とは異なります。

# Security Model

https://rawgit.com/mikesamuel/sanitized-jquery-templates/trunk/safetemplate.html#problem_definition は、このパッケージが使用する「安全」を定義しています。

このパッケージは、テンプレートの作者が信頼できると仮定し、Executeのデータパラメータは信頼できないと仮定し、信頼できないデータに対して以下のプロパティを保持しようとします：

Structure Preservation Property:
"... テンプレートの作者が安全なテンプレート言語でHTMLタグを書くとき、
ブラウザは出力の対応する部分を、信頼できないデータの値に関係なくタグとして解釈します。
同様に、属性の境界やJSとCSSの文字列の境界などの他の構造についても同様です。"

Code Effect Property:
"... テンプレートの出力をページに注入する結果として実行されるのは、
テンプレートの作者によって指定されたコードのみであり、
同じ結果として実行されるすべてのコードもテンプレートの作者によって指定されるべきです。"

Least Surprise Property:
"HTML、CSS、JavaScriptに精通し、コンテキストに応じた自動エスケープが行われることを知っている開発者（またはコードレビュアー）は、{{.}}を見て、どのようなサニタイゼーションが行われるかを正しく推測することができるべきです。"

以前は、ECMAScript 6のテンプレートリテラルはデフォルトで無効にされており、GODEBUG=jstmpllitinterp=1 環境変数で有効にすることができました。
テンプレートリテラルは現在デフォルトでサポートされており、jstmpllitinterpを設定しても効果はありません。
*/
package template
