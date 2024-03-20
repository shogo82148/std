// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージscannerは、UTF-8エンコードされたテキストのスキャナとトークナイザを提供します。
// それはソースを提供するio.Readerを取り、その後、Scan関数を繰り返し呼び出すことでトークン化できます。
// 既存のツールとの互換性のため、NUL文字は許可されていません。ソースの最初の文字が
// UTF-8エンコードされたバイトオーダーマーク（BOM）である場合、それは破棄されます。
//
// デフォルトでは、[Scanner] は空白とGoのコメントをスキップし、Go言語仕様によって定義されたすべての
// リテラルを認識します。それは、それらのリテラルの一部のみを認識し、異なる識別子と空白文字を認識するように
// カスタマイズすることができます。
package scanner

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/io"
)

// Positionはソース位置を表す値です。
// Line > 0 の場合、位置は有効です。
type Position struct {
	Filename string
	Offset   int
	Line     int
	Column   int
}

// IsValidは位置が有効かどうかを報告します。
func (pos *Position) IsValid() bool

func (pos Position) String() string

// トークンの認識を制御するための事前定義されたモードビット。
// 例えば、(Goの)識別子と整数のみを認識し、コメントをスキップするように
// [Scanner] を設定するには、ScannerのModeフィールドを次のように設定します:
//
//	ScanIdents | ScanInts | SkipComments
//
// SkipCommentsが設定されている場合、コメントはスキップされるを除いて、
// 認識されないトークンは無視されません。代わりに、スキャナは単に
// それぞれの個々の文字（または可能性のあるサブトークン）を返します。
// 例えば、モードがScanIdents（ScanStringsではない）の場合、文字列
// "foo"はトークンシーケンス '"' [Ident] '"'としてスキャンされます。
//
// GoTokensを使用してScannerを設定すると、Goの識別子を含むすべてのGoリテラルトークンが受け入れられます。
// コメントはスキップされます。
const (
	ScanIdents     = 1 << -Ident
	ScanInts       = 1 << -Int
	ScanFloats     = 1 << -Float
	ScanChars      = 1 << -Char
	ScanStrings    = 1 << -String
	ScanRawStrings = 1 << -RawString
	ScanComments   = 1 << -Comment
	SkipComments   = 1 << -skipComment
	GoTokens       = ScanIdents | ScanFloats | ScanChars | ScanStrings | ScanRawStrings | ScanComments | SkipComments
)

// Scanの結果は、これらのトークンのいずれか、またはUnicode文字です。
const (
	EOF = -(iota + 1)
	Ident
	Int
	Float
	Char
	String
	RawString
	Comment
)

// TokenStringは、トークンまたはUnicode文字の印刷可能な文字列を返します。
func TokenString(tok rune) string

// GoWhitespaceは、[Scanner] のWhitespaceフィールドのデフォルト値です。
// その値はGoの空白文字を選択します。
const GoWhitespace = 1<<'\t' | 1<<'\n' | 1<<'\r' | 1<<' '

// Scannerは、io.ReaderからのUnicode文字とトークンの読み取りを実装します。
type Scanner struct {
	// Input
	src io.Reader

	// Source buffer
	srcBuf [bufLen + 1]byte
	srcPos int
	srcEnd int

	// Source position
	srcBufOffset int
	line         int
	column       int
	lastLineLen  int
	lastCharLen  int

	// Token text buffer
	// Typically, token text is stored completely in srcBuf, but in general
	// the token text's head may be buffered in tokBuf while the token text's
	// tail is stored in srcBuf.
	tokBuf bytes.Buffer
	tokPos int
	tokEnd int

	// One character look-ahead
	ch rune

	// Errorは、発生した各エラーに対して呼び出されます。Error
	// 関数が設定されていない場合、エラーはos.Stderrに報告されます。
	Error func(s *Scanner, msg string)

	// ErrorCountは、発生した各エラーごとに1ずつ増加します。
	ErrorCount int

	// Modeフィールドは、どのトークンが認識されるかを制御します。例えば、
	// Intsを認識するには、ModeのScanIntsビットを設定します。このフィールドは
	// いつでも変更することができます。
	Mode uint

	// Whitespaceフィールドは、どの文字が空白として認識されるかを制御します。
	// 文字ch <= ' 'を空白として認識するには、Whitespaceのch番目のビットを設定します
	// （ch > ' 'の値に対するScannerの挙動は未定義です）。このフィールドはいつでも変更できます。
	Whitespace uint64

	// IsIdentRuneは、識別子のithルーンとして受け入れられる文字を制御する述語です。
	// 有効な文字のセットは、空白文字のセットと交差してはなりません。
	// IsIdentRune関数が設定されていない場合、代わりに通常のGoの識別子が受け入れられます。
	// このフィールドはいつでも変更することができます。
	IsIdentRune func(ch rune, i int) bool

	// 最近スキャンされたトークンの開始位置。Scanによって設定されます。
	// InitまたはNextを呼び出すと、位置が無効になります（Line == 0）。
	// Filenameフィールドは常にScannerによって untouched のままです。
	// エラーが報告され（Error経由で）かつPositionが無効な場合、
	// スキャナはトークンの内部にはありません。その場合、または最近スキャンされたトークンの
	// 直後の位置を取得するには、Posを呼び出します。
	Position
}

// Initは新しいソースで [Scanner] を初期化し、sを返します。
// [Scanner.Error] はnilに設定され、[Scanner.ErrorCount] は0に設定され、[Scanner.Mode] は [GoTokens] に設定され、
// [Scanner.Whitespace] は [GoWhitespace] に設定されます。
func (s *Scanner) Init(src io.Reader) *Scanner

// Nextは次のUnicode文字を読み取り、返します。
// ソースの終わりで [EOF] を返します。読み取りエラーが発生した場合、
// s.Errorがnilでない場合はs.Errorを呼び出して報告します。それ以外の場合は
// [os.Stderr] にエラーメッセージを出力します。Nextは [Scanner.Position] フィールドを
// 更新しません。現在の位置を取得するには、[Scanner.Pos]()を使用します。
func (s *Scanner) Next() rune

// Peekは、スキャナを進めずにソースの次のUnicode文字を返します。
// スキャナの位置がソースの最後の文字にある場合、[EOF] を返します。
func (s *Scanner) Peek() rune

// Scanは、ソースから次のトークンまたはUnicode文字を読み取り、それを返します。
// それは、それぞれの [Scanner.Mode] ビット（1<<-t）が設定されているトークンtのみを認識します。
// ソースの終わりで [EOF] を返します。スキャナのエラー（読み取りエラーとトークンエラー）は、
// s.Errorがnilでない場合にはs.Errorを呼び出すことで報告します。それ以外の場合は、
// [os.Stderr] にエラーメッセージを出力します。
func (s *Scanner) Scan() rune

// Posは、最後の [Scanner.Next] または [Scanner.Scan] の呼び出しによって返された文字またはトークンの直後の文字の位置を返します。
// 最近スキャンされたトークンの開始位置には、Scannerの [Scanner.Position] フィールドを使用します。
func (s *Scanner) Pos() (pos Position)

// TokenTextは、最近スキャンされたトークンに対応する文字列を返します。
// [Scanner.Scan] の呼び出し後、および [Scanner.Error] の呼び出し中に有効です。
func (s *Scanner) TokenText() string
