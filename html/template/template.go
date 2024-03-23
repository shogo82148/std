// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/text/template"
	"github.com/shogo82148/std/text/template/parse"
)

// Templateは、安全なHTMLドキュメントフラグメントを生成する"text/template"からの特化したTemplateです。
type Template struct {
	// Sticky error if escaping fails, or escapeOK if succeeded.
	escapeErr error
	// We could embed the text/template field, but it's safer not to because
	// we need to keep our version of the name space and the underlying
	// template's in sync.
	text *template.Template
	// The underlying template's parse tree, updated to be HTML-safe.
	Tree *parse.Tree
	*nameSpace
}

// Templatesは、t自体を含む、tに関連付けられたテンプレートのスライスを返します。
func (t *Template) Templates() []*Template

// Optionは、テンプレートのオプションを設定します。オプションは
// 文字列で記述され、単純な文字列または "key=value" の形式を取ります。オプション文字列には
// 最大で一つの等号が含まれます。オプション文字列が認識できない、または無効な場合、
// Optionはパニックを起こします。
//
// 既知のオプション:
//
// missingkey: マップが存在しないキーでインデックス付けされた場合の、実行中の振る舞いを制御します。
//
//	"missingkey=default" または "missingkey=invalid"
//		デフォルトの振る舞い: 何もせずに実行を続けます。
//		印刷される場合、インデックス操作の結果は文字列
//		"<no value>" です。
//	"missingkey=zero"
//		操作はマップタイプの要素のゼロ値を返します。
//	"missingkey=error"
//		エラーで直ちに実行が停止します。
func (t *Template) Option(opt ...string) *Template

// Executeは、解析されたテンプレートを指定されたデータオブジェクトに適用し、
// 出力をwrに書き込みます。
// テンプレートの実行中またはその出力の書き込み中にエラーが発生した場合、
// 実行は停止しますが、部分的な結果はすでに出力ライターに書き込まれている可能性があります。
// テンプレートは並行して安全に実行できますが、並行実行がWriterを共有する場合、
// 出力が交互になる可能性があります。
func (t *Template) Execute(wr io.Writer, data any) error

// ExecuteTemplateは、指定された名前を持つtに関連付けられたテンプレートを
// 指定されたデータオブジェクトに適用し、出力をwrに書き込みます。
// テンプレートの実行中またはその出力の書き込み中にエラーが発生した場合、
// 実行は停止しますが、部分的な結果はすでに出力ライターに書き込まれている可能性があります。
// テンプレートは並行して安全に実行できますが、並行実行がWriterを共有する場合、
// 出力が交互になる可能性があります。
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error

// DefinedTemplatesは、定義されたテンプレートのリストを返します。
// それは文字列 "; defined templates are: " で始まります。もし定義されたテンプレートがなければ、
// 空の文字列を返します。エラーメッセージを生成するために使用されます。
func (t *Template) DefinedTemplates() string

// Parseは、tのテンプレートボディとしてテキストを解析します。
// テキスト内の名前付きテンプレート定義 ({{define ...}}または{{block ...}}ステートメント) は、
// tに関連付けられた追加のテンプレートを定義し、t自体の定義からは削除されます。
//
// テンプレートは、tまたは関連するテンプレートの [Template.Execute] が初めて使用される前に、
// Parseを連続して呼び出すことで再定義できます。
// ボディが空白とコメントのみで構成されるテンプレート定義は空とみなされ、
// 既存のテンプレートのボディを置き換えません。
// これにより、Parseを使用して新しい名前付きテンプレート定義を追加することができますが、
// メインのテンプレートボディを上書きすることはありません。
func (t *Template) Parse(text string) (*Template, error)

// AddParseTreeは、名前とパースツリーを持つ新しいテンプレートを作成し、
// それをtに関連付けます。
//
// tまたは関連するテンプレートがすでに実行されている場合、エラーを返します。
func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)

// Cloneは、テンプレートの複製を返します。これには、すべての関連テンプレートも含まれます。
// 実際の表現はコピーされませんが、関連テンプレートの名前空間はコピーされるため、
// コピーでの [Template.Parse] へのさらなる呼び出しは、コピーにテンプレートを追加しますが、元のテンプレートには追加しません。
// [Template.Clone] は、共通のテンプレートを準備し、それらを他のテンプレートのバリアント定義とともに使用するために使用できます。
// バリアントは、クローンが作成された後に追加します。
//
// tがすでに実行されている場合、エラーを返します。
func (t *Template) Clone() (*Template, error)

// Newは、指定された名前を持つ新しいHTMLテンプレートを割り当てます。
func New(name string) *Template

// Newは、指定された名前を持つ新しいHTMLテンプレートを割り当て、
// それを与えられたテンプレートと同じデリミタと関連付けます。この関連付けは推移的で、
// 一つのテンプレートが{{template}}アクションで別のテンプレートを呼び出すことを可能にします。
//
// 指定された名前を持つテンプレートがすでに存在する場合、新しいHTMLテンプレートは
// それを置き換えます。既存のテンプレートはリセットされ、tとの関連付けが解除されます。
func (t *Template) New(name string) *Template

// Nameはテンプレートの名前を返します。
func (t *Template) Name() string

type FuncMap = template.FuncMap

// Funcsは引数のマップの要素をテンプレートの関数マップに追加します。
// これはテンプレートが解析される前に呼び出す必要があります。
// マップの値が適切な戻り値型を持つ関数でない場合、パニックを起こします。ただし、
// マップの要素を上書きすることは合法です。戻り値はテンプレートなので、
// 呼び出しはチェーンできます。
func (t *Template) Funcs(funcMap FuncMap) *Template

// Delimsは、アクションのデリミタを指定された文字列に設定します。これは、
// その後の [Template.Parse]、[ParseFiles]、または [ParseGlob] への呼び出しで使用されます。ネストしたテンプレート
// 定義はこの設定を継承します。空のデリミタは、対応するデフォルトを表します: {{ または }}。
// 戻り値はテンプレートなので、呼び出しはチェーンできます。
func (t *Template) Delims(left, right string) *Template

// Lookupは、tに関連付けられた指定された名前のテンプレートを返します。
// もし該当するテンプレートがなければ、nilを返します。
func (t *Template) Lookup(name string) *Template

// Mustは、([*Template], error)を返す関数への呼び出しをラップし、
// エラーが非nilの場合にパニックを起こすヘルパーです。これは変数の初期化での使用を意図しています。
// 例えば、
//
//	var t = template.Must(template.New("name").Parse("html"))
func Must(t *Template, err error) *Template

// ParseFilesは新しい [Template] を作成し、
// 指定されたファイルからテンプレート定義を解析します。返されるテンプレートの名前は、
// 最初のファイルの（ベース）名と（解析された）内容になります。少なくとも一つのファイルが必要です。
// エラーが発生した場合、解析は停止し、返される [*Template] はnilになります。
//
// 異なるディレクトリにある同じ名前の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
// 例えば、ParseFiles("a/foo", "b/foo")は "b/foo" を "foo" という名前のテンプレートとして保存し、
// "a/foo" は利用できません。
func ParseFiles(filenames ...string) (*Template, error)

// ParseFilesは指定されたファイルを解析し、結果として得られるテンプレートを
// tに関連付けます。エラーが発生した場合、解析は停止し、返されるテンプレートはnilになります。
// それ以外の場合、それはtです。少なくとも一つのファイルが必要です。
//
// 異なるディレクトリにある同じ名前の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
//
// ParseFilesは、tまたは関連するテンプレートがすでに実行されている場合、エラーを返します。
func (t *Template) ParseFiles(filenames ...string) (*Template, error)

// ParseGlobは新しい [Template] を作成し、パターンによって識別されたファイルから
// テンプレート定義を解析します。ファイルはfilepath.Matchのセマンティクスに従ってマッチし、
// パターンは少なくとも一つのファイルとマッチしなければなりません。
// 返されるテンプレートの名前は、パターンによって最初にマッチしたファイルの（ベース）名と
// （解析された）内容になります。ParseGlobは、パターンにマッチしたファイルのリストで
// [ParseFiles] を呼び出すのと同等です。
//
// 異なるディレクトリにある同じ名前の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
func ParseGlob(pattern string) (*Template, error)

// ParseGlobは、パターンによって識別されたファイルのテンプレート定義を解析し、
// 結果として得られるテンプレートをtに関連付けます。ファイルはfilepath.Matchのセマンティクスに従ってマッチし、
// パターンは少なくとも一つのファイルとマッチしなければなりません。
// ParseGlobは、パターンにマッチしたファイルのリストでt.ParseFilesを呼び出すのと同等です。
//
// 異なるディレクトリにある同じ名前の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
//
// ParseGlobは、tまたは関連するテンプレートがすでに実行されている場合、エラーを返します。
func (t *Template) ParseGlob(pattern string) (*Template, error)

// IsTrueは、値がその型のゼロでない「真」であるか、
// そして値が意味のある真偽値を持っているかどうかを報告します。
// これはifやその他のアクションで使用される真実の定義です。
func IsTrue(val any) (truth, ok bool)

// ParseFSは [ParseFiles] や [ParseGlob] と似ていますが、ホストのオペレーティングシステムのファイルシステムではなく、
// ファイルシステムfsから読み取ります。
// それはグロブパターンのリストを受け入れます。
// （ほとんどのファイル名は、自分自身のみにマッチするグロブパターンとして機能することに注意してください。）
func ParseFS(fs fs.FS, patterns ...string) (*Template, error)

// ParseFSは [Template.ParseFiles] や [Template.ParseGlob] と似ていますが、ホストのオペレーティングシステムのファイルシステムではなく、
// ファイルシステムfsから読み取ります。
// それはグロブパターンのリストを受け入れます。
// （ほとんどのファイル名は、自分自身のみにマッチするグロブパターンとして機能することに注意してください。）
func (t *Template) ParseFS(fs fs.FS, patterns ...string) (*Template, error)
