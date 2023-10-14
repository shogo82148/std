// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package regexpは正規表現の検索を実装します。
//
<<<<<<< HEAD
// The syntax of the regular expressions accepted is the same
// general syntax used by Perl, Python, and other languages.
// More precisely, it is the syntax accepted by RE2 and described at
// https://golang.org/s/re2syntax, except for \C.
// For an overview of the syntax, see the [regexp/syntax] package.
=======
// 受け入れる正規表現の構文は、Perl、Python、および他の言語で使用される一般的な構文です。
// より正確には、RE2が受け入れる構文であり、以下で説明されています。
// https://golang.org/s/re2syntax（\Cを除く）
// 構文の概要については、次を実行してください。
//
// go doc regexp/syntax
>>>>>>> release-branch.go1.21
//
// このパッケージによって提供される正規表現の実装は、入力のサイズに比例して線形の時間で実行されることが保証されています。
// （これは、ほとんどのオープンソースの正規表現の実装が保証していない特性です。）この特性の詳細については、次を参照してください。
//
// https://swtch.com/~rsc/regexp/regexp1.html
//
// またはオートマトン理論に関する書籍を参照してください。
//
<<<<<<< HEAD
// All characters are UTF-8-encoded code points.
// Following [utf8.DecodeRune], each byte of an invalid UTF-8 sequence
// is treated as if it encoded utf8.RuneError (U+FFFD).
//
// There are 16 methods of [Regexp] that match a regular expression and identify
// the matched text. Their names are matched by this regular expression:
=======
// すべての文字はUTF-8でエンコードされたコードポイントです。
// utf8.DecodeRuneに従って、無効なUTF-8シーケンスの各バイトは、utf8.RuneError（U+FFFD）としてエンコードされたものとして扱われます。
//
// 正規表現に一致し、一致したテキストを識別するRegexpの16個のメソッドがあります。
// これらのメソッドの名前は、次の正規表現と一致します。
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// Regexp is the representation of a compiled regular expression.
// A Regexp is safe for concurrent use by multiple goroutines,
// except for configuration methods, such as [Regexp.Longest].
=======
// Regexpはコンパイルされた正規表現の表現です。
// RegexpはLongestなどの構成方法を除いて、複数のゴルーチンによる並行利用に安全です。
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// Copy returns a new [Regexp] object copied from re.
// Calling [Regexp.Longest] on one copy does not affect another.
//
// Deprecated: In earlier releases, when using a [Regexp] in multiple goroutines,
// giving each goroutine its own copy helped to avoid lock contention.
// As of Go 1.12, using Copy is no longer necessary to avoid lock contention.
// Copy may still be appropriate if the reason for its use is to make
// two copies with different [Regexp.Longest] settings.
func (re *Regexp) Copy() *Regexp

// Compile parses a regular expression and returns, if successful,
// a [Regexp] object that can be used to match against text.
//
// When matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses the one that a backtracking search would have found first.
// This so-called leftmost-first matching is the same semantics
// that Perl, Python, and other implementations use, although this
// package implements it without the expense of backtracking.
// For POSIX leftmost-longest matching, see [CompilePOSIX].
func Compile(expr string) (*Regexp, error)

// CompilePOSIX is like [Compile] but restricts the regular expression
// to POSIX ERE (egrep) syntax and changes the match semantics to
// leftmost-longest.
//
// That is, when matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses a match that is as long as possible.
// This so-called leftmost-longest matching is the same semantics
// that early regular expression implementations used and that POSIX
// specifies.
//
// However, there can be multiple leftmost-longest matches, with different
// submatch choices, and here this package diverges from POSIX.
// Among the possible leftmost-longest matches, this package chooses
// the one that a backtracking search would have found first, while POSIX
// specifies that the match be chosen to maximize the length of the first
// subexpression, then the second, and so on from left to right.
// The POSIX rule is computationally prohibitive and not even well-defined.
// See https://swtch.com/~rsc/regexp/regexp2.html#posix for details.
func CompilePOSIX(expr string) (*Regexp, error)

// Longest makes future searches prefer the leftmost-longest match.
// That is, when matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses a match that is as long as possible.
// This method modifies the [Regexp] and may not be called concurrently
// with any other methods.
func (re *Regexp) Longest()

// MustCompile is like [Compile] but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
func MustCompile(str string) *Regexp

// MustCompilePOSIX is like [CompilePOSIX] but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
func MustCompilePOSIX(str string) *Regexp

// NumSubexp returns the number of parenthesized subexpressions in this [Regexp].
func (re *Regexp) NumSubexp() int

// SubexpNames returns the names of the parenthesized subexpressions
// in this [Regexp]. The name for the first sub-expression is names[1],
// so that if m is a match slice, the name for m[i] is SubexpNames()[i].
// Since the Regexp as a whole cannot be named, names[0] is always
// the empty string. The slice should not be modified.
=======
// Copyは、reからコピーされた新しいRegexpオブジェクトを返します。
// コピーを使用してLongestを呼び出しても他のコピーに影響を与えません。
//
// 廃止予定: 以前のリリースでは、複数のゴルーチンでRegexpを使用する場合、
// 各ゴルーチンに独自のコピーを与えることでロック競合を回避できました。
// Go 1.12以降は、ロック競合を回避するためにCopyを使用する必要はありません。
// Copyは、異なるLongest設定で2つのコピーを作成する必要がある場合には依然適切かもしれません。
func (re *Regexp) Copy() *Regexp

// Compileは正規表現をパースし、成功した場合には、
// テキストと照合するために使用できるRegexpオブジェクトを返します。
//
// テキストと照合する際、正規表現は入力のなるべく早い位置（最も左端）から一致し、
// その中からバックトラック検索が最初に見つけたものを選択します。
// これを左端優先マッチングと呼びますが、
// これはPerl、Pythonなどの実装と同じセマンティクスです。
// ただし、このパッケージはバックトラックのコストなしで実装されています。
// POSIXの左端最長一致マッチングについては、CompilePOSIXを参照してください。
func Compile(expr string) (*Regexp, error)

// CompilePOSIXはCompileと同様ですが、正規表現をPOSIX ERE（egrep）構文に制限し、マッチのセマンティクスをleftmost-longestに変更します。
// つまり、テキストに対してマッチングする際に、正規表現は入力（最も左側）で可能な限り早く開始するマッチを返し、その中でも可能な限り長いマッチを選択します。
// このleftmost-longestマッチングと呼ばれる手法は、かつての正規表現の実装やPOSIXが指定するセマンティクスと同じです。
// ただし、複数のleftmost-longestマッチが存在する場合、このパッケージはPOSIXとは異なる方法を採用します。
// 可能なleftmost-longestマッチの中から、このパッケージはバックトラッキング検索で最初に見つかるマッチを選択します。一方、POSIXでは最初のサブエクスプレッション、次に2番目のサブエクスプレッション、以降左から右へと長さを最大化するマッチを選択すると規定されています。
// POSIXのルールは計算上の制約があり、定義もされていません。
// 詳細については、https://swtch.com/~rsc/regexp/regexp2.html#posixを参照してください。
func CompilePOSIX(expr string) (*Regexp, error)

// Longestは将来の検索において、最も左にある最長一致を優先します。
// つまり、テキストに対して一致を探す場合、正規表現はできるだけ早く入力の最初に一致するものを返し、その中から最長の一致を選択します。
// このメソッドはRegexpを修正するため、他のメソッドと同時に呼び出すことはできません。
func (re *Regexp) Longest()

// MustCompileはCompileと似ていますが、式を解析できない場合はパニックします。
// これにより、コンパイルされた正規表現を保持するグローバル変数の安全な初期化が簡素化されます。
func MustCompile(str string) *Regexp

// MustCompilePOSIXは、CompilePOSIXと似ていますが、式が解析できない場合にはpanicを発生させます。
// これにより、コンパイルされた正規表現を保持するグローバル変数の安全な初期化を簡素化します。
func MustCompilePOSIX(str string) *Regexp

// NumSubexpはこのRegexp内のカッコで囲まれたサブ式の数を返します。
func (re *Regexp) NumSubexp() int

// SubexpNamesはこの正規表現の括弧付きの部分式の名前を返します。
// 最初の部分式の名前はnames[1]ですので、mがマッチスライスである場合、m[i]の名前はSubexpNames()[i]です。
// 正規表現全体には名前を付けることができないため、names[0]は常に空の文字列です。
// スライスは変更しないでください。
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// MatchReader reports whether the text returned by the [io.RuneReader]
// contains any match of the regular expression re.
=======
// MatchReaderはRuneReaderが返すテキストに、正規表現reの一致が含まれているかどうかを報告します。
>>>>>>> release-branch.go1.21
func (re *Regexp) MatchReader(r io.RuneReader) bool

// MatchStringは文字列sに正規表現reの一致があるかどうかを報告します。
func (re *Regexp) MatchString(s string) bool

// Matchは、バイトスライスbに正規表現reの一致が含まれているかどうかを報告します。
func (re *Regexp) Match(b []byte) bool

<<<<<<< HEAD
// MatchReader reports whether the text returned by the RuneReader
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)

// MatchString reports whether the string s
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
func MatchString(pattern string, s string) (matched bool, err error)

// Match reports whether the byte slice b
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
func Match(pattern string, b []byte) (matched bool, err error)

// ReplaceAllString returns a copy of src, replacing matches of the [Regexp]
// with the replacement string repl.
// Inside repl, $ signs are interpreted as in [Regexp.Expand].
func (re *Regexp) ReplaceAllString(src, repl string) string

// ReplaceAllLiteralString returns a copy of src, replacing matches of the [Regexp]
// with the replacement string repl. The replacement repl is substituted directly,
// without using [Regexp.Expand].
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string

// ReplaceAllStringFunc returns a copy of src in which all matches of the
// [Regexp] have been replaced by the return value of function repl applied
// to the matched substring. The replacement returned by repl is substituted
// directly, without using [Regexp.Expand].
func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string

// ReplaceAll returns a copy of src, replacing matches of the [Regexp]
// with the replacement text repl.
// Inside repl, $ signs are interpreted as in [Regexp.Expand].
func (re *Regexp) ReplaceAll(src, repl []byte) []byte

// ReplaceAllLiteral returns a copy of src, replacing matches of the [Regexp]
// with the replacement bytes repl. The replacement repl is substituted directly,
// without using [Regexp.Expand].
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte

// ReplaceAllFunc returns a copy of src in which all matches of the
// [Regexp] have been replaced by the return value of function repl applied
// to the matched byte slice. The replacement returned by repl is substituted
// directly, without using [Regexp.Expand].
=======
// MatchReaderは、RuneReaderによって返されるテキストに正規表現パターンの一致があるかどうかを報告します。
// より複雑なクエリにはCompileと完全なRegexpインターフェイスを使用する必要があります。
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)

// MatchStringは、文字列sが正規表現パターンに一致するものを含んでいるかどうかを報告します。
// より複雑なクエリを行う場合は、Compileと完全なRegexpインターフェースを使用する必要があります。
func MatchString(pattern string, s string) (matched bool, err error)

// Matchは、バイトスライス b が正規表現パターンのいずれかに一致するかどうかを報告します。
// より複雑なクエリには、Compileと完全なRegexpインターフェースを使用する必要があります。
func Match(pattern string, b []byte) (matched bool, err error)

// ReplaceAllStringは、srcの正規表現にマッチする箇所を置換文字列replで置き換えたsrcのコピーを返します。repl内では、Expandと同様に$記号が解釈されます。例えば、$1は最初のサブマッチのテキストを表します。
func (re *Regexp) ReplaceAllString(src, repl string) string

// ReplaceAllLiteralStringはsrcのコピーを返し、Regexpの一致部分を置換文字列replで置き換えます。置換replは直接代入され、Expandを使用しません。
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string

// ReplaceAllStringFuncは、srcのすべての正規表現の一致箇所を、関数replが適用された一致した部分文字列の返り値に置き換えたコピーを返します。replによって返される置換文字列は、Expandを使用せずに直接代入されます。
func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string

// ReplaceAllは、Regexpの一致した箇所を置換テキストreplで置き換えたsrcのコピーを返します。repl内の$記号はExpandと同様に解釈されます。つまり、$1は最初のサブマッチのテキストを表します。
func (re *Regexp) ReplaceAll(src, repl []byte) []byte

// ReplaceAllLiteralは、Regexpの一致する箇所を置換バイトreplで置換したsrcのコピーを返します。置換replは、Expandを使用せずに直接代入されます。
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte

// ReplaceAllFuncは、Regexpのすべての一致箇所を、
// マッチしたバイトスライスに対して適用した関数replの戻り値で置換したsrcのコピーを返します。
// replによって返される置換は、Expandを使用せずに直接代入されます。
>>>>>>> release-branch.go1.21
func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte

// QuoteMetaは、引数のテキスト内のすべての正規表現メタ文字をエスケープした文字列を返します。返された文字列は、リテラルテキストにマッチする正規表現です。
func QuoteMeta(s string) string

// Findは正規表現に一致する最も左側のテキストを含むスライスを返します。
// nilの返り値は一致なしを示します。
func (re *Regexp) Find(b []byte) []byte

// FindIndexは、正規表現の一致する最も左側の箇所を示す整数の2要素スライスを返します。一致部分はb[loc[0]:loc[1]]にあります。
// nilを返す場合は一致なしを示します。
func (re *Regexp) FindIndex(b []byte) (loc []int)

<<<<<<< HEAD
// FindString returns a string holding the text of the leftmost match in s of the regular
// expression. If there is no match, the return value is an empty string,
// but it will also be empty if the regular expression successfully matches
// an empty string. Use [Regexp.FindStringIndex] or [Regexp.FindStringSubmatch] if it is
// necessary to distinguish these cases.
=======
// FindStringは、正規表現の左端と一致する最初のテキストを保持する文字列を返します。
// 一致がない場合、返り値は空の文字列になりますが、正規表現が空の文字列と一致する場合も同様に空になります。
// これらのケースを区別する必要がある場合は、FindStringIndexまたはFindStringSubmatchを使用してください。
>>>>>>> release-branch.go1.21
func (re *Regexp) FindString(s string) string

// FindStringIndexは、正規表現のsにおける最も左にマッチする部分の位置を定義する、整数の2要素のスライスを返します。マッチはs[loc[0]:loc[1]]にあります。
// nilの返り値は、マッチが見つからなかったことを示します。
func (re *Regexp) FindStringIndex(s string) (loc []int)

<<<<<<< HEAD
// FindReaderIndex returns a two-element slice of integers defining the
// location of the leftmost match of the regular expression in text read from
// the [io.RuneReader]. The match text was found in the input stream at
// byte offset loc[0] through loc[1]-1.
// A return value of nil indicates no match.
=======
// FindReaderIndexは、RuneReaderから読み込まれたテキスト内で正規表現の最左一致の位置を示す整数の2要素スライスを返します。マッチしたテキストは、入力ストリームのバイトオフセットloc[0]からloc[1]-1までで見つかりました。
// nilの戻り値は一致がないことを示します。
>>>>>>> release-branch.go1.21
func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)

// FindSubmatchは、正規表現でb内で最も左にマッチするテキストと、そのサブエクスプレッション（'Submatch'のパッケージの説明による）のマッチ（あれば）を保持するスライスのスライスを返します。
// nilの戻り値は、マッチがないことを示します。
func (re *Regexp) FindSubmatch(b []byte) [][]byte

<<<<<<< HEAD
// Expand appends template to dst and returns the result; during the
// append, Expand replaces variables in the template with corresponding
// matches drawn from src. The match slice should have been returned by
// [Regexp.FindSubmatchIndex].
//
// In the template, a variable is denoted by a substring of the form
// $name or ${name}, where name is a non-empty sequence of letters,
// digits, and underscores. A purely numeric name like $1 refers to
// the submatch with the corresponding index; other names refer to
// capturing parentheses named with the (?P<name>...) syntax. A
// reference to an out of range or unmatched index or a name that is not
// present in the regular expression is replaced with an empty slice.
//
// In the $name form, name is taken to be as long as possible: $1x is
// equivalent to ${1x}, not ${1}x, and, $10 is equivalent to ${10}, not ${1}0.
//
// To insert a literal $ in the output, use $$ in the template.
func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte

// ExpandString is like [Regexp.Expand] but the template and source are strings.
// It appends to and returns a byte slice in order to give the calling
// code control over allocation.
=======
// Expandはテンプレートをdstに追加し、結果を返します。追加の過程で、Expandはテンプレート内の変数をsrcから引っ張った対応する一致で置き換えます。一致スライスはFindSubmatchIndexによって返されるべきです。
// テンプレート中では、変数は$nameまたは${name}の形式の部分文字列で示されます。nameは非空の文字、数字、アンダースコアの連続です。$1のような純粋な数字の名前は、対応するインデックスのサブマッチを参照します。その他の名前は、(?P<name>...)構文で名前付きのキャプチャ括弧を参照します。範囲外またはマッチしないインデックスの参照または正規表現に存在しない名前は、空のスライスで置き換えられます。
// $name形式では、nameは可能な限り長くなります：$1xは${1x}ではなく${1}xと等価です。また、$10は${10}ではなく${1}0と等価です。
// 出力にリテラルの$を挿入するには、テンプレートで$$を使用してください。
func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte

// ExpandStringは、Expandと同様にテンプレートとソースが文字列の場合に使用します。
// 割り当てに対する制御を呼び出し元のコードに提供するために、バイトスライスに追加して返します。
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// FindReaderSubmatchIndex returns a slice holding the index pairs
// identifying the leftmost match of the regular expression of text read by
// the [io.RuneReader], and the matches, if any, of its subexpressions, as defined
// by the 'Submatch' and 'Index' descriptions in the package comment. A
// return value of nil indicates no match.
=======
// FindReaderSubmatchIndexは、RuneReaderによって読み取られたテキストの正規表現の最も左の一致と、そのサブエクスプレッションの一致（ある場合）を識別するインデックスのペアを保持するスライスを返します。パッケージコメントの'Submatch'と'Index'の説明で定義されています。nilの返り値は一致がないことを示します。
>>>>>>> release-branch.go1.21
func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int

// FindAllはFindの 'All' バージョンであり、パッケージコメントで定義されている 'All' の説明に従って、
// 式の全ての連続するマッチのスライスを返します。
// nilの返り値はマッチがないことを示します。
func (re *Regexp) FindAll(b []byte, n int) [][]byte

<<<<<<< HEAD
// FindAllIndex is the 'All' version of [Regexp.FindIndex]; it returns a slice of all
// successive matches of the expression, as defined by the 'All' description
// in the package comment.
// A return value of nil indicates no match.
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int

// FindAllString is the 'All' version of [Regexp.FindString]; it returns a slice of all
// successive matches of the expression, as defined by the 'All' description
// in the package comment.
// A return value of nil indicates no match.
func (re *Regexp) FindAllString(s string, n int) []string

// FindAllStringIndex is the 'All' version of [Regexp.FindStringIndex]; it returns a
// slice of all successive matches of the expression, as defined by the 'All'
// description in the package comment.
// A return value of nil indicates no match.
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int

// FindAllSubmatch is the 'All' version of [Regexp.FindSubmatch]; it returns a slice
// of all successive matches of the expression, as defined by the 'All'
// description in the package comment.
// A return value of nil indicates no match.
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte

// FindAllSubmatchIndex is the 'All' version of [Regexp.FindSubmatchIndex]; it returns
// a slice of all successive matches of the expression, as defined by the
// 'All' description in the package comment.
// A return value of nil indicates no match.
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int

// FindAllStringSubmatch is the 'All' version of [Regexp.FindStringSubmatch]; it
// returns a slice of all successive matches of the expression, as defined by
// the 'All' description in the package comment.
// A return value of nil indicates no match.
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string

// FindAllStringSubmatchIndex is the 'All' version of
// [Regexp.FindStringSubmatchIndex]; it returns a slice of all successive matches of
// the expression, as defined by the 'All' description in the package
// comment.
// A return value of nil indicates no match.
=======
// FindAllIndexはFindIndexの「All」バージョンであり、
// パッケージコメントで定義されている「All」の説明に従って、
// 式のすべての連続する一致のスライスを返します。
// nilの返り値は一致がないことを示します。
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int

// FindAllStringはFindStringの'All'バージョンです。式によって定義されるように、
// 'All'の説明に従って、連続する全ての一致する部分文字列のスライスを返します。
// nilの返り値は一致なしを示します。
func (re *Regexp) FindAllString(s string, n int) []string

// FindAllStringIndexはFindStringIndexの「All」バージョンです。式によって定義されるすべての連続したマッチのスライスを返します。「All」の説明によってパッケージのコメントで定義されます。
// nilの返り値はマッチがないことを示します。
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int

// FindAllSubmatchは、FindSubmatchの 'All' バージョンです。この関数は、'All' 説明によって定義された通り、式に連続するすべての一致部分をスライスとして返します。
// nilの返り値は、マッチが見つからなかったことを示します。
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte

// FindAllSubmatchIndexはFindSubmatchIndexの'All'バージョンであり、
// パッケージコメントの'All'の説明に従って、式に対するすべての連続した一致結果のスライスを返します。
// nilの返り値は一致なしを示します。
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int

// FindAllStringSubmatchは、FindStringSubmatchの「All」バージョンであり、式によって定義されたすべての連続した一致のスライスを返します。パッケージコメントの「All」の説明に従います。
// nilの戻り値は一致がないことを示します。
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string

// FindAllStringSubmatchIndexはFindStringSubmatchIndexのバージョンであり、
// 式によって定義されるすべての連続した一致のスライスを返します。
// パッケージコメントの「All」の説明で定義されているように、
// nilの戻り値は一致がないことを示します。
>>>>>>> release-branch.go1.21
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int

// Splitメソッドは、文字列sを指定の表現によって区切り、それらの表現にマッチする部分文字列のスライスを返します。
//
<<<<<<< HEAD
// The slice returned by this method consists of all the substrings of s
// not contained in the slice returned by [Regexp.FindAllString]. When called on an expression
// that contains no metacharacters, it is equivalent to [strings.SplitN].
=======
// このメソッドによって返されるスライスは、sのうちFindAllStringで返されるスライスに含まれていない
// 全ての部分文字列からなります。メタ文字を含まない表現に対しては、strings.SplitNと同等の動作になります。
>>>>>>> release-branch.go1.21
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
