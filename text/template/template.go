// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"github.com/shogo82148/std/text/template/parse"
)

// Templateは、解析されたテンプレートの表現です。*parse.Tree
// フィールドは、html/templateによる使用のためだけにエクスポートされており、
// 他のすべてのクライアントによって未エクスポートとして扱われるべきです。
type Template struct {
	name string
	*parse.Tree
	*common
	leftDelim  string
	rightDelim string
}

// Newは、指定された名前を持つ新しい未定義のテンプレートを割り当てます。
func New(name string) *Template

// Nameはテンプレートの名前を返します。
func (t *Template) Name() string

// Newは、与えられたテンプレートと同じデリミタを持つ新しい未定義のテンプレートを割り当てます。
// この関連付けは推移的で、一つのテンプレートが{{template}}アクションで別のテンプレートを
// 呼び出すことを可能にします。
//
// 関連付けられたテンプレートは基礎となるデータを共有するため、テンプレートの構築は
// 並行して安全に行うことはできません。テンプレートが構築されたら、それらは並行して
// 実行することができます。
func (t *Template) New(name string) *Template

// Cloneは、関連付けられたすべてのテンプレートを含むテンプレートの複製を返します。
// 実際の表現はコピーされませんが、関連付けられたテンプレートの名前空間はコピーされるため、
// コピーでのさらなるParseの呼び出しは、コピーにテンプレートを追加しますが、元のテンプレートには追加しません。
// Cloneは、共通のテンプレートを準備し、それらを他のテンプレートのバリアント定義とともに使用するために、
// クローン作成後にバリアントを追加することで使用できます。
func (t *Template) Clone() (*Template, error)

// AddParseTreeは、引数のパースツリーをテンプレートtに関連付け、
// それに指定された名前を付けます。テンプレートがまだ定義されていない場合、
// このツリーがその定義となります。すでに定義されていてその名前を持っている場合、
// 既存の定義が置き換えられます。それ以外の場合は、新しいテンプレートが作成され、
// 定義され、返されます。
func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)

// Templatesは、tに関連付けられた定義済みテンプレートのスライスを返します。
func (t *Template) Templates() []*Template

// Delimsは、指定された文字列にアクションデリミタを設定します。これは、
// その後のParse、ParseFiles、またはParseGlobへの呼び出しで使用されます。
// ネストしたテンプレート定義はこの設定を継承します。空のデリミタは、
// 対応するデフォルト（{{または}}）を表します。
// 戻り値はテンプレートなので、呼び出しはチェーンできます。
func (t *Template) Delims(left, right string) *Template

// Funcsは、引数のマップの要素をテンプレートの関数マップに追加します。
// これはテンプレートが解析される前に呼び出す必要があります。
// マップの値が適切な戻り値型を持つ関数でない場合、または名前がテンプレート内の関数として
// 文法的に使用できない場合、パニックを起こします。
// マップの要素を上書きすることは合法です。戻り値はテンプレートなので、呼び出しはチェーンできます。
func (t *Template) Funcs(funcMap FuncMap) *Template

// Lookupは、tに関連付けられた指定された名前のテンプレートを返します。
// そのようなテンプレートがないか、テンプレートが定義を持っていない場合はnilを返します。
func (t *Template) Lookup(name string) *Template

// Parseは、テキストをtのテンプレートボディとして解析します。
// テキスト内の名前付きテンプレート定義（{{define ...}}または{{block ...}}ステートメント）は、
// tに関連付けられた追加のテンプレートを定義し、t自体の定義からは削除されます。
//
// テンプレートは、Parseへの連続した呼び出しで再定義することができます。
// 本文が空白とコメントのみで構成されるテンプレート定義は空とみなされ、
// 既存のテンプレートの本文を置き換えません。
// これにより、Parseを使用して新しい名前付きテンプレート定義を追加し、
// メインのテンプレート本文を上書きすることなく行うことができます。
func (t *Template) Parse(text string) (*Template, error)
