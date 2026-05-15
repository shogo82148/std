// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// regexpパッケージは正規表現の検索を実装します。
//
// 受け入れる正規表現の構文は、Perl、Python、および他の言語で使用される一般的な構文です。
// より正確には、RE2が受け入れる構文であり、以下で説明されています。
// https://golang.org/s/re2syntax（\Cを除く）
// 構文の概要については、 [regexp/syntax] パッケージを参照してください。
//
// このパッケージによって提供される正規表現の実装は、入力のサイズに比例して線形の時間で実行されることが保証されています。
// （これは、ほとんどのオープンソースの正規表現の実装が保証していない特性です。）この特性の詳細については、
// https://swtch.com/~rsc/regexp/regexp1.html を参照してください。
// またはオートマトン理論に関する書籍を参照してください。
//
// すべての文字はUTF-8でエンコードされたコードポイントです。
// [utf8.DecodeRune] に従って、無効なUTF-8シーケンスの各バイトは、utf8.RuneError（U+FFFD）としてエンコードされたものとして扱われます。
//
// [Regexp]には正規表現にマッチして、マッチしたテキストを識別する24個のメソッドがあります。
// これらのメソッド名は以下の正規表現にマッチします:
//
//	(All|Find|FindAll)(String)?(Submatch)?(Index)?
//
// 'All'バリアントは、式全体の連続した重複しないマッチに対するイテレーターを返します。
// 'FindAll'バリアントは、代わりにそれらのマッチのスライスを返します。
// 前のマッチに隣接する空のマッチは無視されます。'FindAll'バリアントは追加の整数引数nを取ります。
// n >= 0の場合、関数は最大でnのマッチ/サブマッチを返します。
// そうでない場合、すべてを返します。
//
// 'Find'バリアントは、AllまたはFindAllが返す最初のマッチのみを返します。
//
// 'String'が存在する場合、引数は文字列です。そうでない場合は[]byteです。
//
// デフォルトでは、返される各マッチは正規表現にマッチする部分文字列で表され、
// 引数の型に応じてstringまたは[]byteです。
// 'Submatch'が存在する場合、各マッチは代わりに
// 正規表現の括弧で囲まれた部分式（キャプチャグループとも呼ばれます）にマッチする
// 部分文字列のスライスで表され、開き括弧の順序で左から右へ番号付けされます。
// Submatch 0は式全体のマッチ、submatch 1は最初の括弧付き部分式のマッチなどです。
// 'Index'が存在する場合、各部分文字列は代わりに入力文字列内のバイトインデックスのペアで表されます。
// インデックスが負またはサブ文字列がnilの場合、その部分式は入力内のいかなる文字列にもマッチしていません。
// 'String'バージョンの場合、空文字列はマッチなしまたは空のマッチを意味します。
//
// また、[io.RuneReader]から読み込まれたテキストに適用できるメソッドのサブセットもあります:
// [Regexp.MatchReader]、[Regexp.FindReaderIndex]、
// [Regexp.FindReaderSubmatchIndex]。
// 正規表現のマッチはマッチとして返されたテキストを超えてテキストを
// 調べる必要がある場合があることに注意してください。
// そのため、[io.RuneReader]からテキストをマッチさせるメソッドは
// 返す前に任意の距離入力を読み込む可能性があります。
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

// MatchReaderは、[io.RuneReader] によって返されるテキストに正規表現パターンの一致があるかどうかを報告します。
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

// Findはbでreの最も左のマッチのテキストを返します。
// マッチがない場合、戻り値はnilです。
func (re *Regexp) Find(b []byte) []byte

// FindStringはsでreの最も左のマッチのテキストを返します。
// 戻り値は空のマッチとマッチがない場合の両方で空文字列です。
// これら2つのケースを区別するには、[Regexp.FindStringIndex]または[Regexp.FindStringSubmatch]を使用してください。
func (re *Regexp) FindString(s string) string

// FindIndexはbでreの最も左のマッチの位置を返します。
// マッチそのものはb[m[0]:m[1]]にあります。
// マッチがない場合、戻り値はnilです。
func (re *Regexp) FindIndex(b []byte) (m []int)

// FindStringIndexはsでreの最も左のマッチの位置を返します。
// マッチそのものはs[m[0]:m[1]]にあります。
// マッチがない場合、戻り値はnilです。
func (re *Regexp) FindStringIndex(s string) (m []int)

// FindReaderIndexはrでreの最も左のマッチの位置を返します。
// マッチはバイトインデックスm[0]から始まり、バイトインデックスm[1]の直前で終わります。
// マッチがない場合、戻り値はnilです。
//
// FindReaderIndexはrから任意の距離読み込む可能性があります。
// これは返されたマッチを超えて読み込むことも含まれます。
func (re *Regexp) FindReaderIndex(r io.RuneReader) (m []int)

// FindSubmatchはbでreの最初のマッチをサブマッチを含めて返します。
// 全体のマッチはm[0]、最初のサブマッチはm[1]以降です。
// マッチがない場合、戻り値はnilです。
func (re *Regexp) FindSubmatch(b []byte) [][]byte

// FindStringSubmatchはsでreの最初のマッチをサブマッチを含めて返します。
// 全体のマッチはs[0]、最初のサブマッチはs[1]以降です。
// マッチがない場合、戻り値はnilです。
func (re *Regexp) FindStringSubmatch(s string) []string

// FindSubmatchIndexはbでreの最初のマッチをサブマッチを含めて返します。
// 全体のマッチはb[m[0]:m[1]]、最初のサブマッチはb[m[2]:m[3]]以降です。
// マッチがない場合、戻り値はnilです。
func (re *Regexp) FindSubmatchIndex(b []byte) []int

// FindStringSubmatchIndexはsでreの最初のマッチをサブマッチを含めて返します。
// 全体のマッチはs[m[0]:m[1]]、最初のサブマッチはs[m[2]:m[3]]以降です。
// マッチがない場合、戻り値はnilです。
func (re *Regexp) FindStringSubmatchIndex(s string) []int

// FindReaderSubmatchIndexはrでreの最初のマッチをサブマッチを含めて返します。
// 全体のマッチはバイトインデックスm[0]からm[1]まで、
// 最初のサブマッチはバイトインデックスm[2]からm[3]までなどです。
// マッチがない場合、戻り値はnilです。
//
// FindReaderSubmatchIndexはrから任意の距離読み込む可能性があります。
// これは返されたマッチを超えて読み込むことも含まれます。
func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int

// FindAllはbでreのすべてのマッチを返します。
// n >= 0の場合、FindAllは最大でnのマッチを返します。
// 同等のイテレーター形式については、[Regexp.All]を参照してください。
func (re *Regexp) FindAll(b []byte, n int) [][]byte

// FindAllStringはsでreのすべてのマッチを返します。
// n >= 0の場合、FindAllStringは最大でnのマッチを返します。
// 同等のイテレーター形式については、[Regexp.AllString]を参照してください。
func (re *Regexp) FindAllString(s string, n int) []string

// FindAllIndexはbでreのすべてのマッチの位置を返します。
// n >= 0の場合、FindAllIndexは最大でnのマッチを返します。
// 同等のイテレーター形式については、[Regexp.AllIndex]を参照してください。
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int

// FindAllStringIndexはsでreのすべてのマッチの位置を返します。
// n >= 0の場合、FindAllStringIndexは最大でnのマッチを返します。
// 同等のイテレーター形式については、[Regexp.AllStringIndex]を参照してください。
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int

// FindAllSubmatchはbでreのすべてのマッチをサブマッチの位置を含めて返します。
// 返された各マッチmでは、全体のマッチはm[0]、最初のサブマッチはm[1]以降です。
// n >= 0の場合、FindAllSubmatchは最大でnのマッチを返します。
// 同等のイテレーター形式については、[Regexp.AllSubmatch]を参照してください。
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte

// FindAllStringSubmatchはsでreのすべてのマッチをサブマッチの位置を含めて返します。
// 返された各マッチmでは、m[0]が全体のマッチ、m[1]が最初のサブマッチ以降です。
// n >= 0の場合、FindAllStringSubmatchは最大でnのマッチを返します。
// 同等のイテレーター形式については、[Regexp.AllStringSubmatch]を参照してください。
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string

// FindAllSubmatchIndexはbでreのすべてのマッチをサブマッチの位置を含めて返します。
// 返された各マッチmでは、全体のマッチはb[m[0]:m[1]]、最初のサブマッチはb[m[2]:m[3]]以降です。
// n >= 0の場合、FindAllSubmatchIndexは最大でnのマッチを返します。
// 同等のイテレーター形式については、[Regexp.AllSubmatchIndex]を参照してください。
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int

// FindAllStringSubmatchIndexはsでreのすべてのマッチをサブマッチの位置を含めて返します。
// 返された各マッチmでは、全体のマッチはs[m[0]:m[1]]、最初のサブマッチはs[m[2]:m[3]]以降です。
// n >= 0の場合、FindAllStringSubmatchIndexは最大でnのマッチを返します。
// 同等のイテレーター形式については、[Regexp.AllStringSubmatchIndex]を参照してください。
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int

// Expandはtemplateをdstに追加して結果を返します。追加時に、Expandはテンプレート内の
// 変数をsrcから取得された対応するマッチに置き換えます。マッチスライスは
// [Regexp.FindSubmatchIndex]によって返されたものである必要があります。
//
// テンプレート中では、変数は$nameまたは${name}の形式の部分文字列で示されます。nameは非空の文字、数字、アンダースコアの連続です。$1のような純粋な数字の名前は、対応するインデックスのサブマッチを参照します。その他の名前は、(?P<name>...)構文で名前付きのキャプチャ括弧を参照します。範囲外またはマッチしないインデックスの参照または正規表現に存在しない名前は、空のスライスで置き換えられます。
//
// $name形式では、nameは可能な限り長くなります：$1xは${1x}ではなく${1}xと等価です。また、$10は${10}ではなく${1}0と等価です。
// 出力にリテラルの$を挿入するには、テンプレートで$$を使用してください。
func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte

// ExpandStringは、 [Regexp.Expand] と同様にテンプレートとソースが文字列の場合に使用します。
// 割り当てに対する制御を呼び出し元のコードに提供するために、バイトスライスに追加して返します。
func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte

// Splitは式で区切られたs内の部分文字列にスライスし、
// それらの式のマッチ間の部分文字列のスライスを返します。
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
//   - n > 0: 最大でn個の部分文字列を返します。最後の部分文字列は分割されなかった残りの部分です。
//   - n == 0: 結果はnil（部分文字列なし）です。
//   - n < 0: 全ての部分文字列を返します。
func (re *Regexp) Split(s string, n int) []string

// AppendTextは [encoding.TextAppender] を実装します。出力は [Regexp.String] メソッドを呼び出した場合と同じ内容になります。
//
// この出力は場合によっては情報が失われることに注意してください：このメソッドはPOSIX正規表現（[CompilePOSIX] でコンパイルされたもの）や、[Regexp.Longest] メソッドが呼ばれたものを示しません。
func (re *Regexp) AppendText(b []byte) ([]byte, error)

// MarshalTextは [encoding.TextMarshaler] を実装します。出力は [Regexp.AppendText] メソッドを呼び出した場合と同じ内容になります。
//
// 詳細は [Regexp.AppendText] を参照してください。
func (re *Regexp) MarshalText() ([]byte, error)

// UnmarshalTextは、エンコードされた値に対して[Compile]を呼び出すことで、[encoding.TextUnmarshaler]を実装します。
func (re *Regexp) UnmarshalText(text []byte) error
