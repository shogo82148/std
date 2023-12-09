// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

// 信頼できるソースからのコンテンツの文字列。
type (
	// CSSは、以下のいずれかに一致する既知の安全なコンテンツをカプセル化します：
	//   1. CSS3のスタイルシートの生成、例えば `p { color: purple }`。
	//   2. CSS3のルールの生成、例えば `a[href=~"https:"].foo#bar`。
	//   3. CSS3の宣言の生成、例えば `color: red; margin: 2px`。
	//   4. CSS3の値の生成、例えば `rgba(0, 0, 255, 127)`。
	// https://www.w3.org/TR/css3-syntax/#parsing および
	// https://web.archive.org/web/20090211114933/http://w3.org/TR/css3-syntax#style を参照してください。
	//
	// このタイプの使用はセキュリティリスクを伴います：
	// カプセル化されたコンテンツは信頼できるソースから来るべきであり、
	// それはテンプレートの出力にそのまま含まれます。
	CSS string

	// HTMLは、既知の安全なHTMLドキュメントフラグメントをカプセル化します。
	// それは、第三者からのHTMLや、閉じられていないタグやコメントが含まれるHTMLには使用すべきではありません。
	// 信頼できるHTMLサニタイザの出力と、このパッケージによってエスケープされたテンプレートは、HTMLでの使用に適しています。
	//
	// このタイプの使用はセキュリティリスクを伴います：
	// カプセル化されたコンテンツは信頼できるソースから来るべきであり、
	// それはテンプレートの出力にそのまま含まれます。
	HTML string

	// HTMLAttrは、信頼できるソースからのHTML属性をカプセル化します。
	// 例えば、` dir="ltr"`。
	//
	// このタイプの使用はセキュリティリスクを伴います：
	// カプセル化されたコンテンツは信頼できるソースから来るべきであり、
	// それはテンプレートの出力にそのまま含まれます。
	HTMLAttr string

	// JSは、例えば `(x + y * z())` のような、既知の安全なEcmaScript5の式をカプセル化します。
	// テンプレートの作者は、型付けされた式が意図した優先順位を壊さないこと、そして
	// "{ foo: bar() }\n['foo']()" のような式を渡すときのように、
	// ステートメント/式の曖昧性がないことを確認する責任があります。
	// これは、非常に異なる意味を持つ有効な式と有効なプログラムの両方です。
	//
	// このタイプの使用はセキュリティリスクを伴います：
	// カプセル化されたコンテンツは信頼できるソースから来るべきであり、
	// それはテンプレートの出力にそのまま含まれます。
	//
	// 有効だが信頼できないJSONを含めるためにJSを使用することは安全ではありません。
	// 安全な代替手段は、json.UnmarshalでJSONを解析し、
	// 結果のオブジェクトをテンプレートに渡すことです。これは、JavaScriptのコンテキストで提示されるときに、
	// サニタイズされたJSONに変換されます。
	JS string

	// JSStrは、JavaScriptの式のクォートの間に埋め込むことを意図した一連の文字をカプセル化します。
	// 文字列は一連のStringCharactersに一致しなければなりません：
	//   StringCharacter :: SourceCharacter ただし `\` または LineTerminator は除く
	//                    | EscapeSequence
	// LineContinuationsは許可されていません。
	// JSStr("foo\\nbar")は問題ありませんが、JSStr("foo\\\nbar")は許可されていません。
	//
	// このタイプの使用はセキュリティリスクを伴います：
	// カプセル化されたコンテンツは信頼できるソースから来るべきであり、
	// それはテンプレートの出力にそのまま含まれます。
	JSStr string

	// URLは、既知の安全なURLまたはURL部分文字列（RFC 3986を参照）をカプセル化します。
	// 信頼できるソースからの`javascript:checkThatFormNotEditedBeforeLeavingPage()`のようなURLは
	// ページに含まれるべきですが、デフォルトでは動的な`javascript:` URLは、
	// 頻繁に悪用されるインジェクションベクトルであるためフィルタリングされます。
	//
	// このタイプの使用はセキュリティリスクを伴います：
	// カプセル化されたコンテンツは信頼できるソースから来るべきであり、
	// それはテンプレートの出力にそのまま含まれます。
	URL string

	// Srcsetは、既知の安全なsrcset属性をカプセル化します
	// (https://w3c.github.io/html/semantics-embedded-content.html#element-attrdef-img-srcset を参照)。
	//
	// このタイプの使用はセキュリティリスクを伴います：
	// カプセル化されたコンテンツは信頼できるソースから来るべきであり、
	// それはテンプレートの出力にそのまま含まれます。
	Srcset string
)
