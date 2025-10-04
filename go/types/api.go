// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// typesパッケージは、Goパッケージの型チェックのためのアルゴリズムを実装し、
// データ型を宣言します。[Config.Check] を使用してパッケージの型チェッカーを呼び出すか、
// 代わりに [NewChecker] で新しい型チェッカーを作成し、[Checker.Files] を呼び出して
// インクリメンタルに呼び出すことができます。
//
// 型チェックは、いくつかの相互依存するフェーズで構成されています。
//
// 名前解決は、プログラム内の各識別子([ast.Ident])を、それが示すシンボル([Object])にマッピングします。
// 識別子のシンボルを見つけるには、[Info] のDefsフィールドとUsesフィールド、または [Info.ObjectOf] メソッドを使用します。
// そして、特定の他の種類の構文ノードのシンボルを見つけるには、[Info] のImplicitsフィールドを使用します。
//
// 定数畳み込みは、コンパイル時定数であるすべての式([ast.Expr])の正確な定数値
// ([constant.Value])を計算します。式の定数畳み込みの結果を見つけるには、[Info] のTypesフィールドを使用します。
//
// 型推論は、すべての式([ast.Expr])の型([Type])を計算し、言語仕様に準拠しているかをチェックします。
// 型推論の結果については、[Info]のTypesフィールドを使用してください。
//
<<<<<<< HEAD
// チュートリアルについては、https://go.dev/s/types-tutorial を参照してください。
=======
// Applications that need to type-check one or more complete packages
// of Go source code may find it more convenient not to invoke the
// type checker directly but instead to use the Load function in
// package [golang.org/x/tools/go/packages].
//
// For a tutorial, see https://go.dev/s/types-tutorial.
>>>>>>> upstream/release-branch.go1.25
package types

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/constant"
	"github.com/shogo82148/std/go/token"
)

// エラーは型チェックのエラーを示します。エラーインターフェースを実装します。
// "ソフト"エラーは、パッケージの有効な解釈を許容するエラーです（例：「未使用の変数」）。
// "ハード"エラーは無視した場合に予測不可能な動作につながる可能性があります。
type Error struct {
	Fset *token.FileSet
	Pos  token.Pos
	Msg  string
	Soft bool

	// go116codeは将来のAPIであり、エラーコードのセットが大きいため、公開されていません。また、実験中に変更される可能性も高いです。この機能をプレビューしたいツールは、リフレクションを使用してgo116codeを読むことができます（errorcodes_test.goを参照）。ただし、将来の互換性は保証されていません。
	go116code  Code
	go116start token.Pos
	go116end   token.Pos
}

// Errorは以下の形式でフォーマットされたエラー文字列を返します:
// filename:line:column: message
func (err Error) Error() string

// ArgumentErrorは引数のインデックスに関連するエラーを保持します。
type ArgumentError struct {
	Index int
	Err   error
}

func (e *ArgumentError) Error() string
func (e *ArgumentError) Unwrap() error

// Importerは、インポートパスをパッケージに解決します。
//
// 注意: このインターフェースは、ローカルにvendoredされたパッケージのインポートには対応していません。
// 詳細は、https://golang.org/s/go15vendor を参照してください。
// 可能であれば、外部の実装では [ImporterFrom] を実装するべきです。
type Importer interface {
	Import(path string) (*Package, error)
}

// ImportMode は将来の利用用途のために予約されています。
type ImportMode int

// ImporterFromは、インポートパスをパッケージに解決します。
// https://golang.org/s/go15vendorに従い、ベンダリングをサポートします。
// ImporterFromの実装を取得するには、go/importerを使用してください。
type ImporterFrom interface {
	Importer

	ImportFrom(path, dir string, mode ImportMode) (*Package, error)
}

// Configは型チェックの設定を指定します。
// Configのゼロ値は、すぐに使用できるデフォルトの設定です。
type Config struct {

	// Contextはグローバルな識別子を解決するために使用されるコンテキストです。nilの場合、
	// 型チェッカーはこのフィールドを新たに作成されたコンテキストで初期化します。
	Context *Context

	// GoVersionは、受け入れられるGo言語のバージョンを説明します。文字列は、"go%d.%d"の形式の接頭辞で始まる必要があります（例："go1.20"、"go1.21rc1"、または"go1.21.0"）。または空である必要があります。空の文字列はGo言語のバージョンチェックを無効にします。フォーマットが無効な場合、型チェッカーを呼び出すとエラーが発生します。
	GoVersion string

	// IgnoreFuncBodies が設定されている場合、関数の本体は型チェックされません。
	IgnoreFuncBodies bool

	// If FakeImportC is set, `import "C"` (for packages requiring Cgo)
	// declares an empty "C" package and errors are omitted for qualified
	// identifiers referring to package C (which won't find an object).
	// This feature is intended for the standard library cmd/api tool.
	//
	// Caution: Effects may be unpredictable due to follow-on errors.
	//          Do not use casually!
	FakeImportC bool

	// go115UsesCgoが設定されている場合、型チェッカーはcmd/cgoを実行して生成された_cgo_gotypes.goファイルを
	// パッケージのソースファイルとして提供することを期待します。パッケージCを参照する修飾子の識別子は、
	// _cgo_gotypes.go内のcgo提供の宣言に解決されます。
	//
	// FakeImportCとgo115UsesCgoの両方を設定することはエラーです。
	go115UsesCgo bool

	// もし _Trace が設定されている場合、デバッグトレースが標準出力に表示されます。
	_Trace bool

	// もしError != nilなら、エラーが見つかるたびに呼び出されます
	// タイプチェック中に見つかったエラーの動的な型はErrorです
	// 二次エラー（例：無効な再帰型宣言に関わるすべての型を列挙する）は'\t'文字で始まるエラー文字列を持ちます
	// もしError == nilなら、最初に見つかったエラーでタイプチェックは停止します。
	Error func(err error)

	// インポート宣言から参照されるパッケージをインポートするためにインポーターが使われます。
	// インストールされたインポーターがImporterFromを実装している場合、型チェッカーはImportFromを呼び出します。
	// インポーターが必要ですがインストールされていない場合、型チェッカーはエラーを報告します。
	Importer Importer

	// Sizesがnilでない場合、unsafeパッケージのサイズ計算関数が提供されます。
	// それ以外の場合はSizesFor("gc", "amd64")が代わりに使用されます。
	Sizes Sizes

	// DisableUnusedImportCheckが設定されている場合、パッケージは未使用のインポートについてチェックされません。
	DisableUnusedImportCheck bool

	// If a non-empty _ErrorURL format string is provided, it is used
	// to format an error URL link that is appended to the first line
	// of an error message. ErrorURL must be a format string containing
	// exactly one "%s" format, e.g. "[go.dev/e/%s]".
	_ErrorURL string

	// If EnableAlias is set, alias declarations produce an Alias type. Otherwise
	// the alias information is only in the type name, which points directly to
	// the actual (aliased) type.
	//
	// This setting must not differ among concurrent type-checking operations,
	// since it affects the behavior of Universe.Lookup("any").
	//
	// This flag will eventually be removed (with Go 1.24 at the earliest).
	_EnableAlias bool
}

// Infoは型チェック済みパッケージの結果タイプ情報を保持します。
// マップが提供された情報のみ収集されます。
// パッケージに型エラーがある場合、収集された情報は不完全である場合があります。
type Info struct {

	// Typesは式をその型にマップし、定数式の値もマップします。無効な式は省略されます。
	// （カッコで囲まれた）組み込み関数を示す識別子に対して、記録されるシグネチャは呼び出し側に特化されます：
	// 呼び出し結果が定数でない場合、記録される型は引数ごとのシグネチャとなります。それ以外の場合、記録される型は無効です。
	//
<<<<<<< HEAD
	// Typesマップはすべての識別子の型を記録するのではなく、任意の式が許される場所に現れる識別子のみを記録します。
	// たとえば、セレクタ式x.fの中の識別子fはSelectionsマップにのみ存在し、変数宣言 'var z int' の中の識別子zはDefsマップにのみ存在し、
	// 限定識別子内のパッケージを示す識別子はUsesマップに収集されます。
=======
	// For (possibly parenthesized) identifiers denoting built-in
	// functions, the recorded signatures are call-site specific:
	// if the call result is not a constant, the recorded type is
	// an argument-specific signature. Otherwise, the recorded type
	// is invalid.
	//
	// The Types map does not record the type of every identifier,
	// only those that appear where an arbitrary expression is
	// permitted. For instance:
	// - an identifier f in a selector expression x.f is found
	//   only in the Selections map;
	// - an identifier z in a variable declaration 'var z int'
	//   is found only in the Defs map;
	// - an identifier p denoting a package in a qualified
	//   identifier p.X is found only in the Uses map.
	//
	// Similarly, no type is recorded for the (synthetic) FuncType
	// node in a FuncDecl.Type field, since there is no corresponding
	// syntactic function type expression in the source in this case
	// Instead, the function type is found in the Defs map entry for
	// the corresponding function declaration.
>>>>>>> upstream/release-branch.go1.25
	Types map[ast.Expr]TypeAndValue

	// インスタンス変数は、ジェネリックな型や関数を指定する識別子を、その型引数とインスタンス化された型とのマッピングする。
	//
	// 例えば、T[int, string]という型のインスタンス化において、Tという識別子を[int, string]という型引数とインスタンス化された*Named型とのマッピングを行う。
	// ジェネリックな関数func F[A any](A)が与えられた場合、Fという呼び出し式における識別子を推定された型引数[int]とインスタンス化された*Signatureにマッピングする。
	//
	// 不変条件: Instantiating Uses[id].Type() with Instances[id].TypeArgs は Instances[id].Type と等価な結果を返す。
	Instances map[*ast.Ident]Instance

	// Defsは識別子を定義するオブジェクト（パッケージ名、ドットインポートの"."、ブランクの"_"識別子を含む）にマップします。
	// オブジェクトを示さない識別子（例：パッケージ節のパッケージ名、または型スイッチヘッダーのt := x.(type)における記号変数t）の場合、対応するオブジェクトはnilです。
	//
	// 埋め込まれたフィールドの場合、Defsはフィールド*Varを返します。
	//
<<<<<<< HEAD
	// 不変条件: Defs[id] == nil || Defs[id].Pos() == id.Pos()
=======
	// In ill-typed code, such as a duplicate declaration of the
	// same name, Defs may lack an entry for a declaring identifier.
	//
	// Invariant: Defs[id] == nil || Defs[id].Pos() == id.Pos()
>>>>>>> upstream/release-branch.go1.25
	Defs map[*ast.Ident]Object

	// マップは識別子をその指すオブジェクトに使用します。
	//
	// 埋め込みフィールドの場合、Usesは指定された*TypeNameを返します。
	//
	// 不変条件: Uses[id].Pos() != id.Pos()
	Uses map[*ast.Ident]Object

	// インプリシットは、ノードを暗黙的に宣言されたオブジェクトにマッピングします（オブジェクトがある場合）。
	// 次のノードとオブジェクトのタイプが表示される場合があります：
	//
	//     ノード                   宣言されたオブジェクト
	//
	//     *ast.ImportSpec    名前変更のないimportの場合は*PkgName
	//     *ast.CaseClause    typeスイッチの各case節（デフォルトを含む）ごとのtype固有の*Var
	//     *ast.Field         匿名パラメータ *Var（無名の結果を含む）
	Implicits map[ast.Node]Object

	// Selectionsはセレクタ式（修飾子を除く）をその対応する選択肢にマッピングします。
	Selections map[*ast.SelectorExpr]*Selection

	// Scopesは、ast.Nodesをそれらが定義するスコープにマッピングします。パッケージスコープは特定のノードではなく、
	// パッケージに属するすべてのファイルに関連付けられています。
	// したがって、パッケージスコープは型チェックされたPackageオブジェクトで見つけることができます。
	// スコープはネストし、Universeスコープが最も外側のスコープで、パッケージスコープを囲み、
	// それには（1つ以上の）ファイルスコープが含まれ、それらは関数スコープを囲み、
	// その順番でステートメントと関数リテラルスコープを囲みます。
	// パッケージレベルの関数はパッケージスコープで宣言されますが、関数スコープは関数宣言を含むファイルの
	// ファイルスコープに埋め込まれています。
	//
	// 関数のScopeには、型パラメータ、パラメータ、名前付き結果の宣言、および
	// ボディブロックのローカル宣言が含まれます。
	// それは関数の構文（[*ast.FuncDecl]または[*ast.FuncLit]）の完全な範囲と一致します。
	// Scopesマッピングには、関数ボディ（[*ast.BlockStmt]）のエントリは含まれていません。
	// 関数のスコープは[*ast.FuncType]に関連付けられています。
	//
	// 以下のノードタイプがScopesに表示される可能性があります：
	//
	//     *ast.File
	//     *ast.FuncType
	//     *ast.TypeSpec
	//     *ast.BlockStmt
	//     *ast.IfStmt
	//     *ast.SwitchStmt
	//     *ast.TypeSwitchStmt
	//     *ast.CaseClause
	//     *ast.CommClause
	//     *ast.ForStmt
	//     *ast.RangeStmt
	Scopes map[ast.Node]*Scope

	// InitOrderはパッケージレベルの初期化子のリストであり、実行する必要がある順序で並んでいます。初期化依存関係に関連する変数を参照する初期化子は、トポロジカル順序で表示されます。他の初期化子はソース順序で表示されます。初期化式を持たない変数は、このリストに表示されません。
	InitOrder []*Initializer

	// FileVersions maps a file to its Go version string.
	// If the file doesn't specify a version, the reported
	// string is Config.GoVersion.
	// Version strings begin with “go”, like “go1.21”, and
	// are suitable for use with the [go/version] package.
	FileVersions map[*ast.File]string
}

// TypeOfは式eの型を返します。見つからない場合はnilを返します。
// 前提条件：Types、Uses、Defsのマップが入力されていることが前提です。
func (info *Info) TypeOf(e ast.Expr) Type

// ObjectOfは、指定したidによって指示されたオブジェクトを返します。
// 存在しない場合はnilを返します。
//
// idが埋め込み構造体フィールドである場合、[Info.ObjectOf] はフィールド(*[Var])を返します
// それが定義する特定のフィールド(*[TypeName])ではありません。
//
// 前提条件：UsesおよびDefsマップが入力されています。
func (info *Info) ObjectOf(id *ast.Ident) Object

// PkgNameOfは、インポートによって定義されたローカルパッケージ名を返します。
// 見つからない場合はnilを返します。
//
// ドットインポートの場合、パッケージ名は"."です。
//
// Precondition: DefsとImplictsのマップが設定されています。
func (info *Info) PkgNameOf(imp *ast.ImportSpec) *PkgName

// TypeAndValueは対応する式の型と値（定数の場合）を報告します。
type TypeAndValue struct {
	mode  operandMode
	Type  Type
	Value constant.Value
}

// IsVoid は、対応する式が結果のない関数呼び出しであるかどうかを報告します。
func (tv TypeAndValue) IsVoid() bool

// IsTypeは、対応する式が型を指定しているかどうかを報告します。
func (tv TypeAndValue) IsType() bool

// IsBuiltinは、対応する式が（たぶん括弧で囲まれた）組み込み関数を示しているかどうかを報告します。
func (tv TypeAndValue) IsBuiltin() bool

// IsValueは、対応する式が値かどうかを報告します。
// 組み込み関数は値とは見なされません。定数値はnon-nilのValueを持ちます。
func (tv TypeAndValue) IsValue() bool

// IsNilは、対応する式が事前宣言された値nilを示しているかどうかを報告します。
func (tv TypeAndValue) IsNil() bool

// Addressableは、対応する式がアドレス指定可能であるかどうかを報告します（https://golang.org/ref/spec#Address_operators）。
func (tv TypeAndValue) Addressable() bool

// Assignableは、対応する式が（適切な型の値が提供された場合に）代入可能かどうかを報告します。
func (tv TypeAndValue) Assignable() bool

// HasOkは、対応する式がコンマOK代入の右辺に使用できるかどうかを報告します。
func (tv TypeAndValue) HasOk() bool

// Instanceは、型と関数のインスタンス化のための型引数とインスタンス化された型を報告します。型のインスタンス化では、[Type] は動的型*[Named] になります。関数のインスタンス化では、[Type] は動的型*Signatureになります。
type Instance struct {
	TypeArgs *TypeList
	Type     Type
}

// イニシャライザは、パッケージレベルの変数、または複数の値を持つ初期化式の場合、変数のリストと対応する初期化式を表します。
type Initializer struct {
	Lhs []*Var
	Rhs ast.Expr
}

func (init *Initializer) String() string

// Checkはパッケージの型チェックを行い、結果のパッケージオブジェクトと初めのエラー（もし存在すれば）を返します。さらに、infoがnilでない場合、Checkは [Info] 構造体の非nilのマップそれぞれを埋めます。
// エラーが発生しなかった場合、パッケージは完全であるとマークされます。そうでなければ不完全です。エラーの存在に応じた動作の制御については、[Config.Error] を参照してください。
// パッケージはast.Filesのリストと対応するファイルセット、およびパッケージが識別されるパッケージパスで指定されます。クリーンパスは空またはドット（"."）ではないでしょう。
func (conf *Config) Check(path string, fset *token.FileSet, files []*ast.File, info *Info) (*Package, error)
