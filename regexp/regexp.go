// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package regexpは正規表現の検索を実装します。
//
// 受け入れる正規表現の構文は、Perl、Python、および他の言語で使用される一般的な構文です。
// より正確には、RE2が受け入れる構文であり、以下で説明されています。
// https://golang.org/s/re2syntax（\Cを除く）
// 構文の概要については、 [regexp/syntax] パッケージを参照してください。
//
// このパッケージによって提供される正規表現の実装は、入力のサイズに比例して線形の時間で実行されることが保証されています。
// （これは、ほとんどのオープンソースの正規表現の実装が保証していない特性です。）この特性の詳細については、次を参照してください。
//
// https://swtch.com/~rsc/regexp/regexp1.html
//
// またはオートマトン理論に関する書籍を参照してください。
//
// すべての文字はUTF-8でエンコードされたコードポイントです。
// [utf8.DecodeRune] に従って、無効なUTF-8シーケンスの各バイトは、utf8.RuneError（U+FFFD）としてエンコードされたものとして扱われます。
//
// 正規表現に一致し、一致したテキストを識別する [Regexp] の16個のメソッドがあります。
// これらのメソッドの名前は、次の正規表現と一致します。
//
// Find(All)?(String)?(Submatch)?(Index)?
//
// 'All'が存在する場合、このルーチンは表現全体の連続する重複しない一致を見つけます。直前の一致と隣接する空の一致は無視されます。戻り値は、対応する非-'All'ルーチンの連続する戻り値を含むスライスです。これらのルーチンは、追加の整数引数nを受け取ります。ただし、n >= 0の場合、関数は最大n個の一致/サブマッチを返し、それ以外の場合はすべてを返します。
//
// 'String'が存在する場合、引数は文字列です。それ以外の場合はバイトのスライスです。返り値は適切に調整されます。
//
// 'Submatch'が存在する場合、返り値は式の連続するサブマッチを識別するスライスです。サブマッチは、正規表現内のパレンセシスで囲まれたサブ式（キャプチャグループとも呼ばれる）の一致です。左から右にかけて開くかっこの順に番号が付けられています。サブマッチ0は式全体の一致であり、サブマッチ1は最初のカッコで囲まれた部分式の一致です。
//
// 'Index'が存在する場合、一致とサブマッチは入力文字列内のバイトインデックスのペアで識別されます。
// result[2*n:2*n+2]はn番目のサブマッチのインデックスを識別します。n==0の場合のペアは、式全体の一致を識別します。'Index'が存在しない場合、一致/サブマッチのテキストで識別されます。
// インデックスが負数であるか、テキストがnilの場合、サブ式は入力文字列内で一致するテキストがないことを意味します。
// 'String'バージョンでは、空の文字列は一致がないか空の一致を意味します。
//
// RuneReaderから読み取られるテキストに適用できるメソッドのサブセットもあります：
//
// # MatchReader、FindReaderIndex、FindReaderSubmatchIndex
//
// このセットは増える可能性があります。正規表現の一致では、一致のために返されるテキストを超えたテキストを調べる必要がある場合があるため、RuneReaderからテキストを一致させるメソッドは、返される前に任意の深さまで入力を読み込む可能性があります。
//
// （このパターンに一致しないいくつかの他のメソッドもあります。）
package regexp

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/regexp/syntax"
)

// Regexpはコンパイルされた正規表現の表現です。
// Regexpは [Regexp.Longest] などの構成方法を除いて、複数のゴルーチンによる並行利用に安全です。
type Regexp struct {
	expr           string
	prog           *syntax.Prog
	onepass        *onePassProg
	numSubexp      int
	maxBitStateLen int
	subexpNames    []string
	prefix         string
	prefixBytes    []byte
	prefixRune     rune
	prefixEnd      uint32
	mpool          int
	matchcap       int
	prefixComplete bool
	cond           syntax.EmptyOp
	minInputLen    int

	// このフィールドは Longest メソッドによって変更可能ですが、それ以外では読み取り専用です。
	longest bool
}

// Stringは正規表現をコンパイルするために使用されたソーステキストを返します。
func (re *Regexp) String() string

// Copyは、reからコピーされた新しい [Regexp] オブジェクトを返します。
// コピーを使用して [Regexp.Longest] を呼び出しても他のコピーに影響を与えません。
//
// Deprecated: 以前のリリースでは、複数のゴルーチンで [Regexp] を使用する場合、
// 各ゴルーチンに独自のコピーを与えることでロック競合を回避できました。
// Go 1.12以降は、ロック競合を回避するためにCopyを使用する必要はありません。
// Copyは、異なる [Regexp.Longest] 設定で2つのコピーを作成する必要がある場合には依然適切かもしれません。
func (re *Regexp) Copy() *Regexp

// Compileは正規表現をパースし、成功した場合には、
// テキストと照合するために使用できる [Regexp] オブジェクトを返します。
//
// テキストと照合する際、正規表現は入力のなるべく早い位置（最も左端）から一致し、
// その中からバックトラック検索が最初に見つけたものを選択します。
// これを左端優先マッチングと呼びますが、
// これはPerl、Pythonなどの実装と同じセマンティクスです。
// ただし、このパッケージはバックトラックのコストなしで実装されています。
// POSIXの左端最長一致マッチングについては、 [CompilePOSIX] を参照してください。
func Compile(expr string) (*Regexp, error)

// CompilePOSIXは [Compile] と同様ですが、正規表現をPOSIX ERE（egrep）構文に制限し、マッチのセマンティクスをleftmost-longestに変更します。
// つまり、テキストに対してマッチングする際に、正規表現は入力（最も左側）で可能な限り早く開始するマッチを返し、その中でも可能な限り長いマッチを選択します。
// このleftmost-longestマッチングと呼ばれる手法は、かつての正規表現の実装やPOSIXが指定するセマンティクスと同じです。
// ただし、複数のleftmost-longestマッチが存在する場合、このパッケージはPOSIXとは異なる方法を採用します。
// 可能なleftmost-longestマッチの中から、このパッケージはバックトラッキング検索で最初に見つかるマッチを選択します。一方、POSIXでは最初のサブエクスプレッション、次に2番目のサブエクスプレッション、以降左から右へと長さを最大化するマッチを選択すると規定されています。
// POSIXのルールは計算上の制約があり、定義もされていません。
// 詳細については、https://swtch.com/~rsc/regexp/regexp2.html#posixを参照してください。
func CompilePOSIX(expr string) (*Regexp, error)

// Longestは将来の検索において、最も左にある最長一致を優先します。
// つまり、テキストに対して一致を探す場合、正規表現はできるだけ早く入力の最初に一致するものを返し、その中から最長の一致を選択します。
// このメソッドは [Regexp] を修正するため、他のメソッドと同時に呼び出すことはできません。
func (re *Regexp) Longest()

// MustCompileは [Compile] と似ていますが、式を解析できない場合はパニックします。
// これにより、コンパイルされた正規表現を保持するグローバル変数の安全な初期化が簡素化されます。
func MustCompile(str string) *Regexp

// MustCompilePOSIXは、 [CompilePOSIX] と似ていますが、式が解析できない場合にはpanicを発生させます。
// これにより、コンパイルされた正規表現を保持するグローバル変数の安全な初期化を簡素化します。
func MustCompilePOSIX(str string) *Regexp

// NumSubexpはこの [Regexp] 内のカッコで囲まれたサブ式の数を返します。
func (re *Regexp) NumSubexp() int

// SubexpNamesはこの [Regexp] の括弧付きの部分式の名前を返します。
// 最初の部分式の名前はnames[1]ですので、mがマッチスライスである場合、m[i]の名前はSubexpNames()[i]です。
// 正規表現全体には名前を付けることができないため、names[0]は常に空の文字列です。
// スライスは変更しないでください。
func (re *Regexp) SubexpNames() []string

// SubexpIndexは指定された名前を持つ最初のサブ式のインデックスを返します。
// もし指定した名前のサブ式が存在しない場合は-1を返します。
//
// 複数のサブ式は同じ名前で書くこともできます。たとえば、(?P<bob>a+)(?P<bob>b+)のように、
// "bob"という名前で2つのサブ式を宣言することができます。
// この場合、SubexpIndexは正規表現内で最も左にあるサブ式のインデックスを返します。
func (re *Regexp) SubexpIndex(name string) int

// LiteralPrefixは、正規表現reの一致の開始部分である必要があるリテラル文字列を返します。もしリテラル文字列が正規表現全体を構成している場合、真を返します。
func (re *Regexp) LiteralPrefix() (prefix string, complete bool)

// MatchReaderは [io.RuneReader] が返すテキストに、正規表現reの一致が含まれているかどうかを報告します。
func (re *Regexp) MatchReader(r io.RuneReader) bool

// MatchStringは文字列sに正規表現reの一致があるかどうかを報告します。
func (re *Regexp) MatchString(s string) bool

// Matchは、バイトスライスbに正規表現reの一致が含まれているかどうかを報告します。
func (re *Regexp) Match(b []byte) bool

// MatchReaderは、RuneReaderによって返されるテキストに正規表現パターンの一致があるかどうかを報告します。
// より複雑なクエリには [Compile] と完全な [Regexp] インターフェイスを使用する必要があります。
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)

// MatchStringは、文字列sが正規表現パターンに一致するものを含んでいるかどうかを報告します。
// より複雑なクエリを行う場合は、 [Compile] と完全な [Regexp] インターフェースを使用する必要があります。
func MatchString(pattern string, s string) (matched bool, err error)

// Matchは、バイトスライス b が正規表現パターンのいずれかに一致するかどうかを報告します。
// より複雑なクエリには、 [Compile] と完全な [Regexp] インターフェースを使用する必要があります。
func Match(pattern string, b []byte) (matched bool, err error)

// ReplaceAllStringは、srcの [Regexp] にマッチする箇所を置換文字列replで置き換えたsrcのコピーを返します。repl内では、 [Regexp.Expand] と同様に$記号が解釈されます。例えば、$1は最初のサブマッチのテキストを表します。
func (re *Regexp) ReplaceAllString(src, repl string) string

// ReplaceAllLiteralStringはsrcのコピーを返し、 [Regexp] の一致部分を置換文字列replで置き換えます。置換replは直接代入され、 [Regexp.Expand] を使用しません。
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string

// ReplaceAllStringFuncは、srcのすべての [Regexp] の一致箇所を、関数replが適用された一致した部分文字列の返り値に置き換えたコピーを返します。replによって返される置換文字列は、 [Regexp.Expand] を使用せずに直接代入されます。
func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string

// ReplaceAllは、 [Regexp] の一致した箇所を置換テキストreplで置き換えたsrcのコピーを返します。repl内の$記号は [Regexp.Expand] と同様に解釈されます。つまり、$1は最初のサブマッチのテキストを表します。
func (re *Regexp) ReplaceAll(src, repl []byte) []byte

// ReplaceAllLiteralは、 [Regexp] の一致する箇所を置換バイトreplで置換したsrcのコピーを返します。置換replは、 [Regexp.Expand] を使用せずに直接代入されます。
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte

// ReplaceAllFuncは、 [Regexp] のすべての一致箇所を、
// マッチしたバイトスライスに対して適用した関数replの戻り値で置換したsrcのコピーを返します。
// replによって返される置換は、 [Regexp.Expand] を使用せずに直接代入されます。
func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte

// QuoteMetaは、引数のテキスト内のすべての正規表現メタ文字をエスケープした文字列を返します。返された文字列は、リテラルテキストにマッチする正規表現です。
func QuoteMeta(s string) string

// Findは正規表現に一致する最も左側のテキストを含むスライスを返します。
// nilの返り値は一致なしを示します。
func (re *Regexp) Find(b []byte) []byte

// FindIndexは、正規表現の一致する最も左側の箇所を示す整数の2要素スライスを返します。一致部分はb[loc[0]:loc[1]]にあります。
// nilを返す場合は一致なしを示します。
func (re *Regexp) FindIndex(b []byte) (loc []int)

// FindStringは、正規表現の左端と一致する最初のテキストを保持する文字列を返します。
// 一致がない場合、返り値は空の文字列になりますが、正規表現が空の文字列と一致する場合も同様に空になります。
// これらのケースを区別する必要がある場合は、 [Regexp.FindStringIndex] または [Regexp.FindStringSubmatch] を使用してください。
func (re *Regexp) FindString(s string) string

// FindStringIndexは、正規表現のsにおける最も左にマッチする部分の位置を定義する、整数の2要素のスライスを返します。マッチはs[loc[0]:loc[1]]にあります。
// nilの返り値は、マッチが見つからなかったことを示します。
func (re *Regexp) FindStringIndex(s string) (loc []int)

// FindReaderIndexは、 [io.RuneReader] から読み込まれたテキスト内で正規表現の最左一致の位置を示す整数の2要素スライスを返します。マッチしたテキストは、入力ストリームのバイトオフセットloc[0]からloc[1]-1までで見つかりました。
// nilの戻り値は一致がないことを示します。
func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)

// FindSubmatchは、正規表現でb内で最も左にマッチするテキストと、そのサブエクスプレッション（'Submatch'のパッケージの説明による）のマッチ（あれば）を保持するスライスのスライスを返します。
// nilの戻り値は、マッチがないことを示します。
func (re *Regexp) FindSubmatch(b []byte) [][]byte

// Expandはテンプレートをdstに追加し、結果を返します。追加の過程で、Expandはテンプレート内の変数をsrcから引っ張った対応する一致で置き換えます。一致スライスは [Regexp.FindSubmatchIndex] によって返されるべきです。
//
// テンプレート中では、変数は$nameまたは${name}の形式の部分文字列で示されます。nameは非空の文字、数字、アンダースコアの連続です。$1のような純粋な数字の名前は、対応するインデックスのサブマッチを参照します。その他の名前は、(?P<name>...)構文で名前付きのキャプチャ括弧を参照します。範囲外またはマッチしないインデックスの参照または正規表現に存在しない名前は、空のスライスで置き換えられます。
//
// $name形式では、nameは可能な限り長くなります：$1xは${1x}ではなく${1}xと等価です。また、$10は${10}ではなく${1}0と等価です。
// 出力にリテラルの$を挿入するには、テンプレートで$$を使用してください。
func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte

// ExpandStringは、 [Regexp.Expand] と同様にテンプレートとソースが文字列の場合に使用します。
// 割り当てに対する制御を呼び出し元のコードに提供するために、バイトスライスに追加して返します。
func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte

// FindSubmatchIndexは、正規表現の最も左側の一致と、'Submatch'および'Index'の説明で定義される、必要に応じてそのサブ式のマッチを示すインデックスのペアを保持するスライスを返します。
// nilの返り値は、一致が見つからないことを示します。
func (re *Regexp) FindSubmatchIndex(b []byte) []int

// FindStringSubmatchは、正規表現の最も左にマッチするテキストと、そのサブエクスプレッションにマッチするテキスト（あれば）を保持する文字列のスライスを返します。パッケージのコメントにある'Submatch'の説明によって定義されます。
// nilの返り値は、マッチがないことを示します。
func (re *Regexp) FindStringSubmatch(s string) []string

// FindStringSubmatchIndexは、正規表現の最も左にある一致と、
// パッケージコメントで定義された'Submatch'および'Index'の説明によって決まる、
// サブ式の一致（ある場合）を特定するインデックスのペアを保持するスライスを返します。
// nilの返り値は一致なしを示します。
func (re *Regexp) FindStringSubmatchIndex(s string) []int

// FindReaderSubmatchIndexは、 [io.RuneReader] によって読み取られたテキストの正規表現の最も左の一致と、そのサブエクスプレッションの一致（ある場合）を識別するインデックスのペアを保持するスライスを返します。パッケージコメントの'Submatch'と'Index'の説明で定義されています。nilの返り値は一致がないことを示します。
func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int

// FindAllは [Regexp.Find] の 'All' バージョンであり、パッケージコメントで定義されている 'All' の説明に従って、
// 式の全ての連続するマッチのスライスを返します。
// nilの返り値はマッチがないことを示します。
func (re *Regexp) FindAll(b []byte, n int) [][]byte

// FindAllIndexは [Regexp.FindIndex] の「All」バージョンであり、
// パッケージコメントで定義されている「All」の説明に従って、
// 式のすべての連続する一致のスライスを返します。
// nilの返り値は一致がないことを示します。
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int

// FindAllStringは [Regexp.FindString] の'All'バージョンです。式によって定義されるように、
// 'All'の説明に従って、連続する全ての一致する部分文字列のスライスを返します。
// nilの返り値は一致なしを示します。
func (re *Regexp) FindAllString(s string, n int) []string

// FindAllStringIndexは [Regexp.FindStringIndex] の「All」バージョンです。式によって定義されるすべての連続したマッチのスライスを返します。「All」の説明によってパッケージのコメントで定義されます。
// nilの返り値はマッチがないことを示します。
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int

// FindAllSubmatchは、 [Regexp.FindSubmatch] の 'All' バージョンです。この関数は、'All' 説明によって定義された通り、式に連続するすべての一致部分をスライスとして返します。
// nilの返り値は、マッチが見つからなかったことを示します。
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte

// FindAllSubmatchIndexはFindSubmatchIndexの'All'バージョンであり、
// パッケージコメントの'All'の説明に従って、式に対するすべての連続した一致結果のスライスを返します。
// nilの返り値は一致なしを示します。
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int

// FindAllStringSubmatchは、 [Regexp.FindSubmatchIndex] の「All」バージョンであり、式によって定義されたすべての連続した一致のスライスを返します。パッケージコメントの「All」の説明に従います。
// nilの戻り値は一致がないことを示します。
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string

// FindAllStringSubmatchIndexは [Regexp.FindStringSubmatchIndex] の「All」バージョンであり、
// 式によって定義されるすべての連続した一致のスライスを返します。
// パッケージコメントの「All」の説明で定義されているように、
// nilの戻り値は一致がないことを示します。
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int

// Splitメソッドは、文字列sを指定の表現によって区切り、それらの表現にマッチする部分文字列のスライスを返します。
//
// このメソッドによって返されるスライスは、sのうちFindAllStringで返されるスライスに含まれていない
// 全ての部分文字列からなります。メタ文字を含まない表現に対しては、 [strings.SplitN] と同等の動作になります。
//
// 例:
//
//	s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
//	// s: ["", "b", "b", "c", "cadaaae"]
//
// countパラメータによって返す部分文字列の数が決まります:
//
//	n > 0: 最大でn個の部分文字列を返します。最後の部分文字列は分割されなかった残りの部分です。
//	n == 0: 結果はnil（部分文字列なし）です。
//	n < 0: 全ての部分文字列を返します。
func (re *Regexp) Split(s string, n int) []string

// MarshalTextは[encoding.TextMarshaler]を実装します。出力は
// [Regexp.String]メソッドを呼び出した場合と一致します。
//
// 注意：このメソッドはいくつかの場合において情報の損失があります。POSIX
// 正規表現（つまり、[CompilePOSIX]を呼び出してコンパイルされたもの）や、
// [Regexp.Longest]メソッドが呼び出された正規表現については示しません。
func (re *Regexp) MarshalText() ([]byte, error)

// UnmarshalTextは、エンコードされた値に対して[Compile]を呼び出すことで、[encoding.TextUnmarshaler]を実装します。
func (re *Regexp) UnmarshalText(text []byte) error
