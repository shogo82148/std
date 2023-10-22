// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package build

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
)

// Contextは、ビルドのサポートコンテキストを指定します。
type Context struct {
	GOARCH string
	GOOS   string
	GOROOT string
	GOPATH string

	// Dirは呼び出し元の作業ディレクトリです。空の文字列の場合、実行中のプロセスのカレントディレクトリが使用されます。モジュールモードでは、これはメインモジュールを特定するために使用されます。
	//
	// Dirが空でない場合、ImportとImportDirに渡されるディレクトリは絶対である必要があります。
	Dir string

	CgoEnabled  bool
	UseAllFiles bool
	Compiler    string

	// build、tool、およびrelease タグは、go:build 行の処理時に満たされていると考慮されるビルド制約を指定します。
	// 新しいコンテキストを作成するクライアントは、BuildTags をカスタマイズすることができます。デフォルトは空ですが、ToolTags や ReleaseTags をカスタマイズすることは通常エラーです。
	// ToolTags は、現在の Go ツールチェインの構成に適したビルドタグをデフォルトで持ちます。
	// ReleaseTags は、現在のリリースが互換性のある Go のリリースの一覧をデフォルトで持ちます。
	// BuildTags は、デフォルトのビルドコンテキストでは設定されません。
	// BuildTags、ToolTags、ReleaseTags に加えて、ビルド制約は GOARCH と GOOS の値を満たしたタグとして考慮されます。
	// ReleaseTags の最後の要素は現在のリリースと見なされます。
	BuildTags   []string
	ToolTags    []string
	ReleaseTags []string

	// InstallSuffixは、インストールディレクトリの名前に使用する接尾辞を指定します。
	// デフォルトでは空ですが、カスタムビルドでは出力を分離する必要がある場合にInstallSuffixを設定できます。
	// たとえば、レースディテクタを使用する場合、goコマンドはInstallSuffix = "race"を使用するため、
	// Linux/386システムでは、通常の「linux_386」の代わりに「linux_386_race」という名前のディレクトリにパッケージが書き込まれます。
	InstallSuffix string

	// JoinPathはパスフラグメントのシーケンスを1つのパスに結合します。
	// JoinPathがnilの場合、Importはfilepath.Joinを使用します。
	JoinPath func(elem ...string) string

	// SplitPathList はパスリストを個々のパスのスライスに分割します。
	// SplitPathList が nil の場合、Import は filepath.SplitList を使用します。
	SplitPathList func(list string) []string

	// IsAbsPathは、パスが絶対パスかどうかを報告します。
	// IsAbsPathがnilの場合、Importはfilepath.IsAbsを使用します。
	IsAbsPath func(path string) bool

	// IsDirはパスがディレクトリかどうかを報告します。
	// IsDirがnilの場合、Importはos.Statを呼び出し、結果のIsDirメソッドを使用します。
	IsDir func(path string) bool

	// HasSubdirは、dirがrootの下位ディレクトリであるかどうかを報告します。
	// 複数のレベル下位かもしれません。dirが存在するかどうかを確認しようとはしません。
	// もしそうであれば、HasSubdirはrelをスラッシュで区切られたパスとして設定し、
	// rootに結合してdirと同等のパスを生成することができます。
	// HasSubdirがnilの場合、Importはfilepath.EvalSymlinksを基に実装されたものを使用します。
	HasSubdir func(root, dir string) (rel string, ok bool)

	// ReadDirはディレクトリの内容を表すfs.FileInfoのスライスを名前でソートして返します。
	// ReadDirがnilの場合、Importはos.ReadDirを使用します。
	ReadDir func(dir string) ([]fs.FileInfo, error)

	// OpenFileは読み取り用にファイル（ディレクトリではありません）を開きます。
	// OpenFileがnilの場合、Importはos.Openを使用します。
	OpenFile func(path string) (io.ReadCloser, error)
}

// SrcDirsはパッケージソースルートディレクトリのリストを返します。
// 現在のGoルートとGoパスから引っ張りますが、存在しないディレクトリは省略します。
func (ctxt *Context) SrcDirs() []string

// Defaultはビルド用のデフォルトのContextです。
// もし設定している場合は、GOARCH、GOOS、GOROOT、およびGOPATHの環境変数を使用します。
// 設定されていない場合は、コンパイルされたコードのGOARCH、GOOS、およびGOROOTが使用されます。
var Default Context = defaultContext()

// ImportModeはImportメソッドの動作を制御します。
type ImportMode uint

const (

	// FindOnlyが設定されている場合、Importはパッケージのソースが含まれるディレクトリを見つけた後に停止します。ディレクトリ内のファイルは読み込まれません。
	FindOnly ImportMode = 1 << iota

	// AllowBinaryが設定されている場合、対応するソースコードなしでコンパイル済みのパッケージオブジェクトでImportを満たすことができます。
	//
	// 廃止予定：
	// コンパイルのみのパッケージを作成するサポートされる方法は、ファイルの先頭に//go:binary-only-packageコメントを含むソースコードを書くことです。このようなパッケージは、このフラグの設定に関係なく認識されます（ソースコードを持っているので）し、戻り値のパッケージにBinaryOnlyフラグがtrueで設定されます。
	AllowBinary

	// ImportCommentが設定されている場合、パッケージ文のimportコメントを解析します。
	// もし理解できないコメントが見つかるか、複数のソースファイルで矛盾するコメントが見つかった場合、エラーが返されます。
	// 詳細はgolang.org/s/go14customimportを参照してください。
	ImportComment

	// デフォルトでは、Importは指定されたソースディレクトリ内のベンダーディレクトリを検索します。
	// GOROOTとGOPATHのルートを検索する前に適用されるベンダーディレクトリを検索します。
	// もしImportがベンダーディレクトリを見つけてパッケージを返す場合、返されるImportPathは完全なパスです。
	// "vendor"までのパス要素を含み、"vendor"自体を含むパッケージのパスです。
	// 例えば、Import("y", "x/subdir", 0)が"x/vendor/y"を見つけた場合、返されるパッケージのImportPathは"x/vendor/y"です。
	// 単純に"y"ではありません。
	// 詳細については、golang.org/s/go15vendorを参照してください。
	//
	// IgnoreVendorを設定すると、ベンダーディレクトリは無視されます。
	//
	// パッケージのImportPathとは異なり、返されるパッケージのImports、TestImports、およびXTestImportsは常にソースファイルからの正確なインポートパスです。
	// Importはこれらのパスを解決したり確認したりする試みはしません。
	IgnoreVendor
)

// パッケージは、ディレクトリにあるGoパッケージを説明します。
type Package struct {
	Dir           string
	Name          string
	ImportComment string
	Doc           string
	ImportPath    string
	Root          string
	SrcRoot       string
	PkgRoot       string
	PkgTargetRoot string
	BinDir        string
	Goroot        bool
	PkgObj        string
	AllTags       []string
	ConflictDir   string
	BinaryOnly    bool

	// ソースファイル
	GoFiles           []string
	CgoFiles          []string
	IgnoredGoFiles    []string
	InvalidGoFiles    []string
	IgnoredOtherFiles []string
	CFiles            []string
	CXXFiles          []string
	MFiles            []string
	HFiles            []string
	FFiles            []string
	SFiles            []string
	SwigFiles         []string
	SwigCXXFiles      []string
	SysoFiles         []string

	// Cgoの指令
	CgoCFLAGS    []string
	CgoCPPFLAGS  []string
	CgoCXXFLAGS  []string
	CgoFFLAGS    []string
	CgoLDFLAGS   []string
	CgoPkgConfig []string

	// テスト情報
	TestGoFiles  []string
	XTestGoFiles []string

	// ソースファイル内で見つかったGoディレクティブコメント（//go:zzz...）。
	Directives      []Directive
	TestDirectives  []Directive
	XTestDirectives []Directive

	// 依存関係情報
	Imports        []string
	ImportPos      map[string][]token.Position
	TestImports    []string
	TestImportPos  map[string][]token.Position
	XTestImports   []string
	XTestImportPos map[string][]token.Position

	// //go:embed はGoのソースファイル内で見つかるパターンです。
	// 例えば、ソースファイルに次のように記述されている場合、
	//	//go:embed a* b.c
	// このリストはそれぞれの文字列を別々のエントリとして含んでいます。
	// （//go:embedについての詳細はパッケージembedを参照してください。）
	EmbedPatterns        []string
	EmbedPatternPos      map[string][]token.Position
	TestEmbedPatterns    []string
	TestEmbedPatternPos  map[string][]token.Position
	XTestEmbedPatterns   []string
	XTestEmbedPatternPos map[string][]token.Position
}

// Directiveは、ソースファイル内で見つかるGoのディレクティブコメント（//go:zzz...）です。
type Directive struct {
	Text string
	Pos  token.Position
}

// IsCommandは、パッケージがインストールされるコマンド（単なるライブラリではない）として
// 考えられるかどうかを報告します。
// "main"という名前のパッケージはコマンドとして扱われます。
func (p *Package) IsCommand() bool

// ImportDirは、名前付きディレクトリで見つかったGoパッケージを処理する [Import] のようなものです。
func (ctxt *Context) ImportDir(dir string, mode ImportMode) (*Package, error)

// NoGoErrorは、ビルド可能なGoソースファイルが含まれていないディレクトリを説明するために [Import] で使用されるエラーです。（テストファイル、ビルドタグによって非表示にされたファイルなどは含まれる可能性があります。）
type NoGoError struct {
	Dir string
}

func (e *NoGoError) Error() string

type MultiplePackageError struct {
	Dir      string
	Packages []string
	Files    []string
}

func (e *MultiplePackageError) Error() string

// Importはimportパスによって指定されたGoパッケージについての詳細を返します。
// srcDirディレクトリを基準として、ローカルなimportパスを解釈します。
// もしパッケージが標準のimportパスを使用してインポート可能なローカルなパッケージの場合、返されるパッケージにはp.ImportPathがそのパスに設定されます。
//
// パッケージを含むディレクトリでは、.go、.c、.h、および.sファイルはパッケージの一部と見なされますが、以下のものは除外されます：
//
//   - パッケージドキュメント内の.goファイル
//   - _または.で始まるファイル（おそらくエディタの一時ファイル）
//   - コンテキストで満たされないビルド制約を持つファイル
//
// エラーが発生した場合、Importは非nilのエラーと部分的な情報を含む非nilの*[Package] を返します。
func (ctxt *Context) Import(path string, srcDir string, mode ImportMode) (*Package, error)

// MatchFileは、指定されたディレクトリ内の指定された名前のファイルがコンテキストに一致し、そのディレクトリで [ImportDir] によって作成される [Package] に含まれるかどうかを報告します。
//
// MatchFileは、ファイルの名前を考慮し、ctxt.OpenFileを使用してファイルの内容の一部または全部を読み取ることがあります。
func (ctxt *Context) MatchFile(dir, name string) (match bool, err error)

// Import は Default.Import の略記法です。
func Import(path, srcDir string, mode ImportMode) (*Package, error)

// ImportDir は Default.ImportDir の省略形です。
func ImportDir(dir string, mode ImportMode) (*Package, error)

// ToolDirはビルドツールを含むディレクトリです。
var ToolDir = getToolDir()

// IsLocalImportは、インポートパスがローカルなインポートパス（"."、".."、"./foo"、または"../foo"など）であるかどうかを報告します。
func IsLocalImport(path string) bool

// ArchCharは"?"とエラーを返します。
// Goの以前のバージョンでは、返された文字列はコンパイラとリンカのツール名、デフォルトのオブジェクトファイルの接尾辞、
// およびデフォルトのリンカの出力名に使用されました。Go 1.5以降、これらの文字列はアーキテクチャによって異なりません。
// それらの文字列は、compile、link、.o、およびa.outです。
func ArchChar(goarch string) (string, error)
