// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comment

// Docは解析されたGoドキュメントコメントです。
type Doc struct {
	// Contentはコメント内のコンテンツブロックのシーケンスです。
	Content []Block

	// Linksはコメント内のリンクの定義です。
	Links []*LinkDef
}

// LinkDefは単一のリンク定義です。
type LinkDef struct {
	Text string
	URL  string
	Used bool
}

// ブロックは、ドキュメントコメント内のブロックレベルのコンテンツであり、[*Code]、[*Heading]、[*List]、または[*Paragraph]のいずれかです。
type Block interface {
	block()
}

// ヘッディングはドキュメントコメントの見出しです。
type Heading struct {
	Text []Text
}

// リストは番号付きまたは箇条書きリストです。
// リストは常に空でないことが保証されます：len(Items) > 0。
// 番号付きのリストでは、Items[i].Numberは空ではない文字列です。
// 箇条書きリストでは、Items[i].Numberは空の文字列です。
type List struct {
	// Itemsはアイテムのリストです。
	Items []*ListItem

	// ForceBlankBeforeは、コメントの再フォーマット時に、
	// 通常の条件を無視してリストの前に空行が必要であることを示します。
	// BlankBeforeメソッドを参照してください。
	//
	// コメントパーサーは、リストの前に空行がある場合に
	// ForceBlankBeforeを設定し、空行が出力時に保持されるようにします。
	ForceBlankBefore bool

	// ForceBlankBetweenは、コメントを再フォーマットする際、リストの項目は必ず空白行で区切られる必要があることを示しています。通常の条件を上書きします。BlankBetweenメソッドを参照してください。
	//
	// コメントパーサーは、リスト内の任意の二つの項目の間に空白行がある場合、それをプリントする際に空白行が保持されることを確認するために、ForceBlankBetweenを設定します。
	ForceBlankBetween bool
}

// BlankBeforeはコメントのリフォーマットにおいて、リストの前に空行を含めるべきかを報告します。
// デフォルトのルールは[BlankBetween]と同じです：
// もしリストの項目の内容に空行が含まれている場合
// （つまり、少なくとも1つの項目に複数の段落がある場合）
// リスト自体は前に空行を置く必要があります。
// 先行する空行を強制するためには、[List].ForceBlankBeforeを設定することができます。
func (l *List) BlankBefore() bool

// BlankBetweenはコメントのリフォーマット時に、
// リストの項目間に空行を追加する必要があるかどうかを報告します。
// デフォルトのルールは、リストの項目の内容に空行が含まれている場合
// （つまり、少なくとも1つの項目が複数の段落を持っている場合）、
// リストの項目自体も空行で分ける必要があるということです。
// 空行の区切りは、[List].ForceBlankBetweenを設定することで強制することができます。
func (l *List) BlankBetween() bool

// ListItemは、番号付きまたは点付きリスト内の単一の項目です。
type ListItem struct {

	// Numberは、番号付きリストの10進数の文字列または箇条書きリストの空の文字列です。
	Number string

	// Contentはリストの内容です。
	// 現在、パーサーとプリンターの制限により、Contentの各要素は*Paragraphである必要があります。
	Content []Block
}

// パラグラフはテキストの段落です。
type Paragraph struct {
	Text []Text
}

// Code は、整形済みのコードブロックです。
type Code struct {

	// Textは事前に書式設定されたテキストで、改行文字で終わります。
	// 複数行になることがあり、それぞれの行は改行文字で終わります。
	// 空でなく、空白行で始まることも終わることもありません。
	Text string
}

// テキストはドキュメントコメント内のテキストレベルの内容であり、
// [プレーン]、[イタリック]、[*リンク]、または[*ドキュリンク]のいずれかです。
type Text interface {
	text()
}

// Plainはプレーンテキストとしてレンダリングされる文字列です（イタリック化されません）。
type Plain string

// イタリック体は、斜体のテキストとしてレンダリングされる文字列です。
type Italic string

// Linkは特定のURLへのリンクです。
type Link struct {
	Auto bool
	Text []Text
	URL  string
}

// DocLinkはGoのパッケージまたはシンボルのドキュメントへのリンクです。
type DocLink struct {
	Text []Text

	// ImportPath、Recv、Nameは、Goのパッケージまたはシンボルを識別します。
	// 非空のフィールドの組み合わせは以下の通りです：
	//  - ImportPath：別のパッケージへのリンク
	//  - ImportPath、Name：別のパッケージ内のconst、func、type、varへのリンク
	//  - ImportPath、Recv、Name：別のパッケージ内のメソッドへのリンク
	//  - Name：このパッケージ内のconst、func、type、varへのリンク
	//  - Recv、Name：このパッケージ内のメソッドへのリンク
	ImportPath string
	Recv       string
	Name       string
}

<<<<<<< HEAD
// Parserはドキュメントコメントのパーサーです。
// 構造体のフィールドは、Parseを呼び出す前に埋めることで、
// パースプロセスの詳細をカスタマイズすることができます。
=======
// A Parser is a doc comment parser.
// The fields in the struct can be filled in before calling [Parser.Parse]
// in order to customize the details of the parsing process.
>>>>>>> upstream/master
type Parser struct {

	// WordsはGo言語の識別子に対応する単語のマップであり、斜体にする必要があり、
	// かつ可能であればリンクも作成します。
	// もしWords[w]が空の文字列であれば、単語wは斜体にのみなります。
	// そうでなければ、Words[w]をリンクターゲットとしてリンクが作成されます。
	// Wordsは[go/doc.ToHTML]のwordsパラメータに対応します。
	Words map[string]string

	// LookupPackageはパッケージ名をインポートパスに変換する。
	//
	// LookupPackage(name)がok == trueを返す場合、[name]
	//（または[name.Sym]または[name.Sym.Method]）
	//はimportPathのパッケージドキュメントへのリンクと見なされる。
	// 空文字列とtrueを返すことも有効であり、その場合はnameが現在のパッケージを参照していると見なされる。
	//
	// LookupPackage(name)がok == falseを返す場合、
	//[name]（または[name.Sym]または[name.Sym.Method]）
	//はドキュメントへのリンクと見なされないが、
	// nameが標準ライブラリ内のパッケージの完全な（ただし要素が1つだけの）インポートパスである場合、
	//[math]や[io.Reader]のように
	//例外として扱われる。 LookupPackageはそれらの名前に対しても呼び出されるため、
	//同じパッケージ名の他のパッケージのインポートを参照することができる。
	//
	// LookupPackageをnilに設定するのは、常に""とfalseを返す関数に設定するのと同じです。
	LookupPackage func(name string) (importPath string, ok bool)

	// LookupSymは、現在のパッケージにシンボル名またはメソッド名が存在するかどうかを報告します。
	//
	// LookupSym("", "Name")がtrueを返す場合、[Name]はconst、func、type、またはvarのドキュメントリンクと見なされます。
	//
	// 同様に、LookupSym("Recv", "Name")がtrueを返す場合、[Recv.Name]はtype RecvのメソッドNameのドキュメントリンクと見なされます。
	//
	// LookupSymをnilに設定することは、常にfalseを返す関数に設定することと同じです。
	LookupSym func(recv, name string) (ok bool)
}

<<<<<<< HEAD
// DefaultLookupPackageは、[Parser].LookupPackageがnilの場合に使用されるデフォルトのパッケージルックアップ関数です。
// これは、単一要素のインポートパスを持つ標準ライブラリのパッケージ名を認識します。
// それ以外の場合は、名前を付けることができません。
=======
// DefaultLookupPackage is the default package lookup
// function, used when [Parser.LookupPackage] is nil.
// It recognizes names of the packages from the standard
// library with single-element import paths, such as math,
// which would otherwise be impossible to name.
>>>>>>> upstream/master
//
// ただし、現在のパッケージで使用されているインポートに基づいたより洗練されたルックアップを提供するgo/docパッケージがあることに注意してください。
func DefaultLookupPackage(name string) (importPath string, ok bool)

<<<<<<< HEAD
=======
// Parse parses the doc comment text and returns the *[Doc] form.
// Comment markers (/* // and */) in the text must have already been removed.
>>>>>>> upstream/master
func (p *Parser) Parse(text string) *Doc

const (
	_ spanKind = iota
)
