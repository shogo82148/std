// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// Scannerは、改行で区切られたテキストの行のファイルなどのデータを読み取るための便利なインターフェースを提供します。Scanメソッドの連続した呼び出しにより、ファイルの「トークン」を順番にステップし、トークン間のバイトをスキップします。トークンの仕様は [SplitFunc] 型の分割関数によって定義されます。デフォルトの分割関数は、入力を行に分割し、行末の文字を取り除きます。[Scanner.Split] 関数は、ファイルを行、バイト、UTF-8エンコードされたルーン、スペースで区切られた単語にスキャンするために、このパッケージで定義されています。クライアントは代わりにカスタムの分割関数を提供することもできます。
//
// スキャンは、EOF、最初のI/Oエラー、または [Scanner.Buffer] に収まりきらないトークンで不可回復的に停止します。
// スキャンが停止すると、リーダーは最後のトークンを超えて任意の距離を進めている可能性があります。
// エラー処理や大きなトークンの制御が必要なプログラム、またはリーダーで連続的なスキャンを実行する必要があるプログラムは、
// 代わりに [bufio.Reader] を使用する必要があります。
type Scanner struct {
	r            io.Reader
	split        SplitFunc
	maxTokenSize int
	token        []byte
	buf          []byte
	start        int
	end          int
	err          error
	empties      int
	scanCalled   bool
	done         bool
}

// SplitFuncは、入力をトークン化するために使用されるsplit関数のシグネチャです。
// 引数は、残りの未処理データの初期部分のサブストリングと、 [Reader] にもうデータがないことを報告するフラグであるatEOFです。
// 戻り値は、入力を進めるためのバイト数と、ユーザーに返す次のトークン（あれば）、およびエラー（あれば）です。
//
<<<<<<< HEAD
// 関数がエラーを返すと、スキャンは停止し、入力の一部が破棄される場合があります。エラーが [ErrFinalToken] である場合、スキャンはエラーなしで停止します。
=======
// Scanning stops if the function returns an error, in which case some of
// the input may be discarded. If that error is [ErrFinalToken], scanning
// stops with no error. A non-nil token delivered with [ErrFinalToken]
// will be the last token, and a nil token with [ErrFinalToken]
// immediately stops the scanning.
>>>>>>> upstream/master
//
// それ以外の場合、スキャナは入力を進めます。トークンがnilでない場合、スキャナはユーザーにそれを返します。トークンがnilの場合、スキャナはさらにデータを読み込んでスキャンを続けます。もしデータがもうない場合（つまり、atEOFがtrueの場合）、スキャナは終了します。データがまだ完全なトークンを保持していない場合、例えば改行がない場合は、 [SplitFunc] は(0、nil、nil)を返して [Scanner] にデータをスライスに読み込んで再試行するように指示できます。
//
// この関数は、空のデータスライスでは呼び出されません（ただし、atEOFがtrueの場合は除く）。ただし、atEOFがtrueの場合、データが空でなく、常に未処理のテキストを保持しているかもしれません。
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

// Scannerが返すエラー。
var (
	ErrTooLong         = errors.New("bufio.Scanner: token too long")
	ErrNegativeAdvance = errors.New("bufio.Scanner: SplitFunc returns negative advance count")
	ErrAdvanceTooFar   = errors.New("bufio.Scanner: SplitFunc returns advance count beyond input")
	ErrBadReadCount    = errors.New("bufio.Scanner: Read returned impossible count")
)

const (
	// MaxScanTokenSizeは、トークンをバッファリングするために使用される最大サイズです
	// ユーザーが [Scanner.Buffer] で明示的なバッファを提供しない限り。
	// 実際の最大トークンサイズは、バッファに改行などを含める必要があるため、より小さくなる場合があります。
	MaxScanTokenSize = 64 * 1024
)

// NewScannerはrから読み取るための新しい [Scanner] を返します。
// split関数はデフォルトで [ScanLines] になります。

func NewScanner(r io.Reader) *Scanner

// Errは [Scanner] によって遭遇した最初のEOF以外のエラーを返します。
func (s *Scanner) Err() error

// BytesはScanの呼び出しによって生成された最新のトークンを返します。
// 基になる配列は、後続の [Scanner.Scan] の呼び出しによって上書きされる可能性があります。
// メモリ確保は行われません。
func (s *Scanner) Bytes() []byte

// Textは [Scanner.Scan] の呼び出しで生成された最新のトークンを、そのバイトを保持する新しく割り当てられた文字列として返します。
func (s *Scanner) Text() string

<<<<<<< HEAD
// ErrFinalTokenは特別なシグナルエラー値です。これはSplit関数によって返され、エラーと一緒に配信されるトークンが最後のトークンであり、スキャンはこれ以上停止する必要があることを示します。
// ErrFinalTokenがScanによって受け取られた後、エラーなしでスキャンが停止します。
// この値は、処理を早めに停止したい場合や最終の空のトークンを配信する必要がある場合に役立ちます。同様の動作をカスタムエラー値で実現することもできますが、ここで提供することで整理された方法が提供されます。
// この値の使用例については、emptyFinalTokenの例を参照してください。
var ErrFinalToken = errors.New("final token")

// Scanは [Scanner] を次のトークンに進め、それを [Scanner.Bytes] メソッドまたは [Scanner.Text] メソッドで利用できるようにします。入力の終わりまたはエラーによりスキャンが停止すると、falseを返します。
// Scanがfalseを返した後、 [Scanner.Err] メソッドはスキャン中に発生したエラーを返しますが、もし [io.EOF] だった場合は、 [Scanner.Err] はnilを返します。
// スプリット関数が入力を進めずに多くの空のトークンを返すと、Scanはパニックを起こします。これはスキャナーの共通のエラーモードです。
=======
// ErrFinalToken is a special sentinel error value. It is intended to be
// returned by a Split function to indicate that the scanning should stop
// with no error. If the token being delivered with this error is not nil,
// the token is the last token.
//
// The value is useful to stop processing early or when it is necessary to
// deliver a final empty token (which is different from a nil token).
// One could achieve the same behavior with a custom error value but
// providing one here is tidier.
// See the emptyFinalToken example for a use of this value.
var ErrFinalToken = errors.New("final token")

// Scan advances the [Scanner] to the next token, which will then be
// available through the [Scanner.Bytes] or [Scanner.Text] method. It returns false when
// there are no more tokens, either by reaching the end of the input or an error.
// After Scan returns false, the [Scanner.Err] method will return any error that
// occurred during scanning, except that if it was [io.EOF], [Scanner.Err]
// will return nil.
// Scan panics if the split function returns too many empty
// tokens without advancing the input. This is a common error mode for
// scanners.
>>>>>>> upstream/master
func (s *Scanner) Scan() bool

// Bufferは、スキャン中に使用する初期バッファと、割り当て可能な最大バッファサイズを設定します。
// 最大トークンサイズは、maxとcap(buf)の大きい方よりも小さくなければなりません。
// max <= cap(buf)の場合、 [Scanner.Scan] はこのバッファのみを使用し、割り当てを行いません。
//
// デフォルトでは、[Scanner.Scan] は内部バッファを使用し、最大トークンサイズを [MaxScanTokenSize] に設定します。
//
// Bufferはスキャンの開始後に呼び出された場合、パニックを発生させます。
func (s *Scanner) Buffer(buf []byte, max int)

// Splitは [Scanner] の分割関数を設定します。
// デフォルトの分割関数は [ScanLines] です。
//
// Splitはスキャンが開始された後に呼び出された場合、パニックを引き起こします。
func (s *Scanner) Split(split SplitFunc)

// ScanBytesは、各バイトをトークンとして返す [Scanner] の分割関数です。
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)

// ScanRunesは、 [Scanner] に対して使用される分割関数で、それぞれのUTF-8エンコードされたルーンをトークンとして返します。返されるルーンのシーケンスは、文字列としての入力を範囲ループで処理した場合と同等です。つまり、誤ったUTF-8エンコーディングはU+FFFD = "\xef\xbf\xbd"に変換されます。Scanインターフェースのため、クライアントは正しくエンコードされた置換ルーンとエンコーディングエラーを区別することが不可能です。
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)

// ScanLinesは、改行マーカー以降のテキストが除かれた、 [Scanner] のスプリット関数です。
// 返される行は空になる場合もあります。
// 改行マーカーは1つのオプションのキャリッジリターンに続いて1つの必須の改行からなります。
// 正規表現の表記では、 `\r?\n`です。
// 入力の最後の非空行は改行がない場合でも返されます。
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)

// ScanWordsは、スペースで区切られたテキストの [Scanner] のための分割関数です。周囲のスペースは削除され、空の文字列は返されません。スペースの定義はunicode.IsSpaceによって設定されます。
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
