// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// typesパッケージは、Goパッケージの型チェックのためのアルゴリズムを実装し、
// データ型を宣言します。Config.Checkを使用してパッケージの型チェッカーを呼び出すか、
// 代わりにNewCheckerで新しい型チェッカーを作成し、Checker.Filesを呼び出して
// インクリメンタルに呼び出すことができます。
//
// 型チェックは、いくつかの相互依存するフェーズで構成されています。
//
// 名前解決は、プログラム内の各識別子（ast.Ident）を、それが示す言語オブジェクト（Object）にマップします。
// 名前解決の結果には、Info.{Defs,Uses,Implicits}を使用します。
//
// 定数畳み込みは、コンパイル時定数であるすべての式（ast.Expr）の正確な定数値（constant.Value）を計算します。
// 定数畳み込みの結果には、Info.Types[expr].Valueを使用します。
//
// 型推論は、すべての式（ast.Expr）の型（Type）を計算し、言語仕様に準拠しているかどうかをチェックします。
// 型推論の結果には、Info.Types[expr].Typeを使用します。
//
// For a tutorial, see https://golang.org/s/types-tutorial.
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
// 可能であれば、外部の実装ではImporterFromを実装するべきです。
type Importer interface {
	// Import returns the imported package for the given import path.
	// The semantics is like for ImporterFrom.ImportFrom except that
	// dir and mode are ignored (since they are not present).
	Import(path string) (*Package, error)
}

// ImportMode は将来の利用用途のために予約されています。
type ImportMode int

// ImporterFromは、インポートパスをパッケージに解決します。
// https://golang.org/s/go15vendorに従い、ベンダリングをサポートします。
// ImporterFromの実装を取得するには、go/importerを使用してください。
type ImporterFrom interface {
	// Importer is present for backward-compatibility. Calling
	// Import(path) is the same as calling ImportFrom(path, "", 0);
	// i.e., locally vendored packages may not be found.
	// The types package does not call Import if an ImporterFrom
	// is present.
	Importer

	// ImportFrom returns the imported package for the given import
	// path when imported by a package file located in dir.
	// If the import failed, besides returning an error, ImportFrom
	// is encouraged to cache and return a package anyway, if one
	// was created. This will reduce package inconsistencies and
	// follow-on type checker errors due to the missing package.
	// The mode value must be 0; it is reserved for future use.
	// Two calls to ImportFrom with the same path and dir must
	// return the same package.
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

	// もし空ではない_ErrorURLフォーマット文字列が提供された場合、それはエラーメッセージの最初の行に追加されるエラーURLリンクのフォーマットに使用されます。ErrorURLは、正確に1つの"%s"フォーマットを含むフォーマット文字列でなければなりません。例："[go.dev/e/%s]"。
	_ErrorURL string
}

// Infoは型チェック済みパッケージの結果タイプ情報を保持します。
// マップが提供された情報のみ収集されます。
// パッケージに型エラーがある場合、収集された情報は不完全である場合があります。
type Info struct {

	// Typesは式をその型にマップし、定数式の値もマップします。無効な式は省略されます。
	// （カッコで囲まれた）組み込み関数を示す識別子に対して、記録されるシグネチャは呼び出し側に特化されます：
	// 呼び出し結果が定数でない場合、記録される型は引数ごとのシグネチャとなります。それ以外の場合、記録される型は無効です。
	//
	// Typesマップはすべての識別子の型を記録するのではなく、任意の式が許される場所に現れる識別子のみを記録します。
	// たとえば、セレクタ式x.fの中の識別子fはSelectionsマップにのみ存在し、変数宣言 'var z int' の中の識別子zはDefsマップにのみ存在し、
	// 限定識別子内のパッケージを示す識別子はUsesマップに収集されます。
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
	// 不変条件: Defs[id] == nil || Defs[id].Pos() == id.Pos()
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

	// Scopesはast.Nodeをその定義するスコープにマップします。パッケージスコープは特定のノードに関連付けられていませんが、パッケージに所属するすべてのファイルに関連付けられています。したがって、パッケージスコープは型チェックされたPackageオブジェクトに見つけることができます。スコープはネストされ、宇宙スコープが最も外側のスコープで、パッケージスコープを囲みます。パッケージスコープには（1つ以上の）ファイルスコープが含まれ、それらは関数スコープを囲みます。関数スコープはステートメントと関数リテラルのスコープを囲みます。注意すべきは、パッケージレベルの関数はパッケージスコープで宣言されますが、関数スコープは関数宣言を含むファイルスコープに埋め込まれているということです。
	// 以下のノードのタイプがScopesに表示される可能性があります：
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

	// _FileVersions maps a file to the file's Go version string.
	// If the file doesn't specify a version and Config.GoVersion
	// is not given, the reported version is the empty string.
	// TODO(gri) should this be "go0.0" instead in that case?
	_FileVersions map[*ast.File]string
}

// TypeOfは式eの型を返します。見つからない場合はnilを返します。
// 前提条件：Types、Uses、Defsのマップが入力されていることが前提です。
func (info *Info) TypeOf(e ast.Expr) Type

// ObjectOfは、指定したidによって指示されたオブジェクトを返します。
// 存在しない場合はnilを返します。
//
// idが埋め込み構造体フィールドである場合、ObjectOfはフィールド(*Var)を返します
// それが定義する特定のフィールド(*TypeName)ではありません。
//
// 前提条件：UsesおよびDefsマップが入力されています。
func (info *Info) ObjectOf(id *ast.Ident) Object

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

// Instanceは、型と関数のインスタンス化のための型引数とインスタンス化された型を報告します。型のインスタンス化では、Typeは動的型*Namedになります。関数のインスタンス化では、Typeは動的型*Signatureになります。
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

// Checkはパッケージの型チェックを行い、結果のパッケージオブジェクトと初めのエラー（もし存在すれば）を返します。さらに、infoがnilでない場合、CheckはInfo構造体の非nilのマップそれぞれを埋めます。
// エラーが発生しなかった場合、パッケージは完全であるとマークされます。そうでなければ不完全です。エラーの存在に応じた動作の制御については、Config.Errorを参照してください。
// パッケージはast.Filesのリストと対応するファイルセット、およびパッケージが識別されるパッケージパスで指定されます。クリーンパスは空またはドット（"."）ではないでしょう。
func (conf *Config) Check(path string, fset *token.FileSet, files []*ast.File, info *Info) (*Package, error)

// AssertableToは、型Vの値が型Tにアサートされることができるかどうかを報告します。
//
// AssertableToの動作は、3つのケースで未指定です：
//   - TがTyp[Invalid]である場合
//   - Vが一般化されたインタフェースである場合。つまり、Goコードで型制約としてのみ使用されるインタフェースである場合
//   - Tが未実体化のジェネリック型である場合
func AssertableTo(V *Interface, T Type) bool

// AssignableToは、型Vの値が型Tの変数に代入可能かどうかを報告します。
//
// AssignableToの動作は、VまたはTがTyp[Invalid]またはインスタンス化されていないジェネリック型の場合、指定されていません。
func AssignableTo(V, T Type) bool

// ConvertibleToは、型Vの値が型Tの値に変換可能かどうかを報告します。
//
// ConvertibleToの動作は、VまたはTがTyp[Invalid]またはインスタンス化されていないジェネリック型である場合、指定されていません。
func ConvertibleTo(V, T Type) bool

// Implementsは、型VがインターフェースTを実装しているかどうかを報告します。
//
// VがTyp[Invalid]やインスタンス化されていないジェネリック型の場合、Implementsの動作は未指定です。
func Implements(V Type, T *Interface) bool

// Satisfiesは型Vが制約Tを満たすかどうかを報告します。
//
// VがTyp[Invalid]またはインスタンス化されていないジェネリック型である場合、Satisfiesの動作は指定されていません。
func Satisfies(V Type, T *Interface) bool

// Identicalはxとyが同じ型であるかどうかを返します。
// Signature型のレシーバは無視されます。
func Identical(x, y Type) bool

// IdenticalIgnoreTagsは、タグを無視した場合にxとyが同じ型であるかどうかを報告します。
// Signature型のレシーバーは無視されます。
func IdenticalIgnoreTags(x, y Type) bool
