// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"github.com/shogo82148/std/text/template/parse"
)

// Errorは、テンプレートのエスケープ処理中に遭遇した問題を説明します。
type Error struct {
	// ErrorCodeはエラーの種類を説明します。
	ErrorCode ErrorCode
	// Nodeは問題を引き起こしたノードです（もし分かる場合）。
	// nilでない場合、NameとLineを上書きします。
	Node parse.Node
	// Nameはエラーが発生したテンプレートの名前です。
	Name string
	// Lineはテンプレートソース内のエラーの行番号、または0です。
	Line int
	// Descriptionは問題の人間が読める説明です。
	Description string
}

// ErrorCodeはエラーの種類を表すコードです。
type ErrorCode int

// テンプレートをエスケープする際に現れる各エラーに対してコードを定義していますが、
// エスケープされたテンプレートは実行時にも失敗する可能性があります。
//
// 出力: "ZgotmplZ"
// 例：
//
//	<img src="{{.X}}">
//	ここで {{.X}} は `javascript:...` に評価されます
//
// 議論：
//
//	"ZgotmplZ" は、実行時に安全でないコンテンツがCSSまたはURLのコンテキストに到達したことを示す特別な値です。
//	例の出力は
//	  <img src="#ZgotmplZ">
//	になります。
//	データが信頼できるソースから来る場合は、フィルタリングから免除するためにコンテンツタイプを使用します：URL(`javascript:...`)。
const (
	// OKはエラーがないことを示します。
	OK ErrorCode = iota

	// ErrAmbigContext: "...はURL内の曖昧なコンテキストに現れます"
	// 例：
	//   <a href="
	//      {{if .C}}
	//        /path/
	//      {{else}}
	//        /search?q=
	//      {{end}}
	//      {{.X}}
	//   ">
	// 議論：
	//   {{.X}}は曖昧なURLコンテキストにあります。なぜなら、{{.C}}によって、
	//   URLの接尾辞かクエリパラメータのどちらかになる可能性があるからです。
	//   {{.X}}を条件の中に移動すると曖昧さがなくなります：
	//   <a href="{{if .C}}/path/{{.X}}{{else}}/search?q={{.X}}">
	ErrAmbigContext

	// ErrBadHTML: "スペース、属性名、またはタグの終わりを期待していましたが、...が得られました",
	//   "...は引用符で囲まれていない属性内にあります", "...は属性名内にあります"
	// 例：
	//   <a href = /search?q=foo>
	//   <href=foo>
	//   <form na<e=...>
	//   <option selected<
	// 議論：
	//   これは、HTML要素のタイプミスが原因であることが多いですが、一部のルーンは、
	//   パーサーの曖昧さを引き起こす可能性があるため、タグ名、属性名、引用符で囲まれていない属性値で禁止されています。
	//   すべての属性を引用符で囲むのが最善の方針です。
	ErrBadHTML

	// ErrBranchEnd: "{{if}}の分岐が異なるコンテキストで終わります"
	// 例：
	//   {{if .C}}<a href="{{end}}{{.X}}
	// 議論：
	//   パッケージhtml/templateは、{{if}}、{{range}}、または{{with}}を通じて各パスを静的に調べ、
	//   その後のパイプラインをエスケープします。例は曖昧です。なぜなら、{{.X}}はHTMLテキストノードであるか、
	//   HTML属性のURLプレフィックスである可能性があるからです。{{.X}}のコンテキストは、それをどのようにエスケープするかを
	//   理解するために使用されますが、そのコンテキストは実行時の{{.C}}の値に依存し、それは静的には知られていません。
	//
	//   問題は通常、引用符や角括弧が欠けているなどの問題であり、または、2つのコンテキストをif、range、withの
	//   異なる分岐にリファクタリングすることで回避できます。問題が空であるべきではないコレクションに対する{{range}}にある場合、
	//   ダミーの{{else}}を追加すると役立つことがあります。
	ErrBranchEnd

	// ErrEndContext: "...は非テキストコンテキストで終わります: ..."
	// 例：
	//   <div
	//   <div title="閉じ引用符なし>
	//   <script>f()
	// 議論：
	//   実行されたテンプレートはHTMLのDocumentFragmentを生成するべきです。
	//   閉じタグなしで終わるテンプレートはこのエラーを引き起こします。
	//   HTMLコンテキストで使用すべきでないテンプレート、または不完全なFragmentを生成するテンプレートは、
	//   直接実行すべきではありません。
	//
	//   {{define "main"}} <script>{{template "helper"}}</script> {{end}}
	//   {{define "helper"}} document.write(' <div title=" ') {{end}}
	//
	//   "helper"は有効なドキュメントフラグメントを生成しないため、直接実行すべきではありません。
	ErrEndContext

	// ErrNoSuchTemplate: "そのようなテンプレートは存在しません ..."
	// 例：
	//   {{define "main"}}<div {{template "attrs"}}>{{end}}
	//   {{define "attrs"}}href="{{.URL}}"{{end}}
	// 議論：
	//   パッケージhtml/templateはテンプレート呼び出しを見てコンテキストを計算します。
	//   ここでは、"attrs"の{{.URL}}は"main"から呼び出されたときにURLとして扱われなければなりませんが、
	//   "main"が解析されたときに"attrs"が定義されていない場合、このエラーが発生します。
	ErrNoSuchTemplate

	// ErrOutputContext: "テンプレート...の出力コンテキストを計算できません"
	// 例：
	//   {{define "t"}}{{if .T}}{{template "t" .T}}{{end}}{{.H}}",{{end}}
	// 議論：
	//   再帰的なテンプレートは、開始したときと同じコンテキストで終わらないため、
	//   信頼性のある出力コンテキストを計算することはできません。
	//   名前付きテンプレートのタイプミスを探してみてください。
	//   もしテンプレートが名前付きの開始コンテキストで呼び出されるべきでないなら、
	//   予期しないコンテキストでそのテンプレートへの呼び出しを探してみてください。
	//   再帰的なテンプレートを再帰的でないようにリファクタリングすることも考えてみてください。
	ErrOutputContext

	// ErrPartialCharset: "未完成のJS正規表現文字セットが...に存在します"
	// 例：
	//     <script>var pattern = /foo[{{.Chars}}]/</script>
	// 議論：
	//   パッケージhtml/templateは、正規表現リテラルの文字セットへの補間をサポートしていません。
	ErrPartialCharset

	// ErrPartialEscape: "未完成のエスケープシーケンスが...に存在します"
	// 例：
	//   <script>alert("\{{.X}}")</script>
	// 議論：
	//   パッケージhtml/templateは、バックスラッシュの後に続くアクションをサポートしていません。
	//   これは通常、エラーであり、より良い解決策があります。例えば、
	//     <script>alert("{{.X}}")</script>
	//   は動作するはずで、もし{{.X}}が"xA0"のような部分的なエスケープシーケンスであれば、
	//   全体を安全なコンテンツとしてマークします：JSStr(`\xA0`)
	ErrPartialEscape

	// ErrRangeLoopReentry: "範囲ループの再入時に: ..."
	// 例：
	//   <script>var x = [{{range .}}'{{.}},{{end}}]</script>
	// 議論：
	//   範囲を通じた反復が、以前のパスと異なるコンテキストで終わるような場合、単一のコンテキストは存在しません。
	//   例では、引用符が欠けているため、{{.}}がJS文字列の内部にあるのか、JS値のコンテキストにあるのかが明確ではありません。
	//   2回目の反復では、次のようなものが生成されます。
	//
	//     <script>var x = ['firstValue,'secondValue]</script>
	ErrRangeLoopReentry

	// ErrSlashAmbig: "'/'は除算または正規表現を開始する可能性があります"
	// 例：
	//   <script>
	//     {{if .C}}var x = 1{{end}}
	//     /-{{.N}}/i.test(x) ? doThis : doThat();
	//   </script>
	// 議論：
	//   上記の例では、最初の'/'が数学的な除算演算子である`var x = 1/-2/i.test(s)...`を生成するか、
	//   最初の'/'が正規表現リテラルを開始する`/-2/i.test(s)`を生成する可能性があります。
	//   分岐内のセミコロンが欠けていないか確認し、どちらの解釈を意図しているか明確にするために
	//   括弧を追加することを検討してみてください。
	ErrSlashAmbig

	// ErrPredefinedEscaper: "テンプレートで禁止されている事前定義されたエスケーパー..."
	// 例：
	//   <div class={{. | html}}>Hello<div>
	// 議論：
	//   パッケージhtml/templateは、すべてのパイプラインをコンテキストに応じてエスケープして、
	//   コードインジェクションに対して安全なHTML出力を生成します。事前定義されたエスケーパー"html"または"urlquery"を
	//   使用してパイプライン出力を手動でエスケープすることは不要であり、Go 1.8以前ではエスケープされたパイプライン出力の
	//   正確さや安全性に影響を与える可能性があります。
	//
	//   ほとんどの場合、例えば上記の例のような場合、このエラーはパイプラインから事前定義されたエスケーパーを単純に削除し、
	//   コンテキスト自動エスケーパーがパイプラインのエスケープを処理することで解決できます。他の場合、事前定義されたエスケーパーが
	//   パイプラインの中間に存在し、後続のコマンドがエスケープされた入力を期待する場合、例えば
	//     {{.X | html | makeALink}}
	//   ここでmakeALinkは
	//     return `<a href="`+input+`">link</a>`
	//   を行う場合、周囲のテンプレートをリファクタリングしてコンテキスト自動エスケーパーを利用するように考えてみてください。つまり、
	//     <a href="{{.X}}">link</a>
	//
	//   Go 1.9以降への移行を容易にするために、"html"と"urlquery"はパイプラインの最後のコマンドとして引き続き許可されます。
	//   ただし、パイプラインが引用符で囲まれていない属性値のコンテキストで発生する場合、"html"は禁止されます。
	//   新しいテンプレートでは"html"と"urlquery"を全く使用しないようにしてください。
	ErrPredefinedEscaper

	// ErrJSTemplate: "...はJSテンプレートリテラル内に存在します"
	// 例：
	//     <script>var tmpl = `{{.Interp}}`</script>
	// 議論:
	//   パッケージhtml/templateは、JSテンプレートリテラル内のアクションをサポートしていません。
	//
	// Deprecated: JSテンプレートリテラル内にアクションが存在する場合、ErrJSTemplateはもはや返されません。
	// JSテンプレートリテラル内のアクションは、現在予想通りにエスケープされます。
	ErrJSTemplate
)

func (e *Error) Error() string
