// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージparseは、text/templateおよびhtml/templateで定義されているテンプレートのパースツリーを構築します。
// クライアントは、一般的な使用を目的としていない共有内部データ構造を提供するこのパッケージではなく、
// それらのパッケージを使用してテンプレートを構築する必要があります。
package parse

// Tree is the representation of a single parsed template.
type Tree struct {
	Name      string
	ParseName string
	Root      *ListNode
	Mode      Mode
	text      string
	// Parsing only; cleared after parse.
	funcs      []map[string]any
	lex        *lexer
	token      [3]item
	peekCount  int
	vars       []string
	treeSet    map[string]*Tree
	actionLine int
	rangeDepth int
}

// mode値はフラグのセット（または0）です。モードはパーサの動作を制御します。
type Mode uint

const (
	ParseComments Mode = 1 << iota
	SkipFuncCheck
)

// CopyはTreeのコピーを返します。パース状態は破棄されます。
func (t *Tree) Copy() *Tree

// Parseは、引数の文字列で記述されたテンプレートを解析することで作成された、
// テンプレート名からparse.Treeへのマップを返します。トップレベルのテンプレートには
// 指定された名前が付けられます。エラーが発生した場合、解析は停止し、
// エラーと共に空のマップが返されます。
func Parse(name, text, leftDelim, rightDelim string, funcs ...map[string]any) (map[string]*Tree, error)

// Newは、指定された名前を持つ新しいパースツリーを割り当てます。
func New(name string, funcs ...map[string]any) *Tree

// ErrorContextは、入力テキスト内のノードの位置のテキスト表現を返します。
// 受信者は、ノードが内部にツリーへのポインタを持っていない場合にのみ使用されます。
// これは古いコードで発生する可能性があります。
func (t *Tree) ErrorContext(n Node) (location, context string)

// Parseは、テンプレート定義文字列を解析して、テンプレートの実行用の表現を構築します。
// アクションデリミタ文字列のいずれかが空の場合、デフォルト（"{{"または"}}"）が使用されます。
// 埋め込まれたテンプレート定義は、treeSetマップに追加されます。
func (t *Tree) Parse(text, leftDelim, rightDelim string, treeSet map[string]*Tree, funcs ...map[string]any) (tree *Tree, err error)

// IsEmptyTreeは、このツリー（ノード）がスペースまたはコメント以外のすべてが空であるかどうかを報告します。
func IsEmptyTree(n Node) bool
