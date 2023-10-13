// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

<<<<<<< HEAD
// Scannerは、改行で区切られたテキストの行のファイルなどのデータを読み取るための便利なインターフェースを提供します。Scanメソッドの連続した呼び出しにより、ファイルの「トークン」を順番にステップし、トークン間のバイトをスキップします。トークンの仕様はSplitFunc型の分割関数によって定義されます。デフォルトの分割関数は、入力を行に分割し、行末の文字を取り除きます。分割関数は、ファイルを行、バイト、UTF-8エンコードされたルーン、スペースで区切られた単語にスキャンするために、このパッケージで定義されています。クライアントは代わりにカスタムの分割関数を提供することもできます。
// スキャンは、EOF、最初のI/Oエラー、またはバッファに収まらない大きすぎるトークンで回復不可能に停止します。スキャンが停止すると、リーダーは最後のトークンを遥かに超えて進んでいる可能性があります。エラーハンドリングや大きなトークンの制御をより細かく制御する必要があるプログラム、またはリーダーで順次スキャンを実行する必要があるプログラムは、代わりにbufio.Readerを使用する必要があります。
=======
// Scanner provides a convenient interface for reading data such as
// a file of newline-delimited lines of text. Successive calls to
// the [Scanner.Scan] method will step through the 'tokens' of a file, skipping
// the bytes between the tokens. The specification of a token is
// defined by a split function of type [SplitFunc]; the default split
// function breaks the input into lines with line termination stripped. [Scanner.Split]
// functions are defined in this package for scanning a file into
// lines, bytes, UTF-8-encoded runes, and space-delimited words. The
// client may instead provide a custom split function.
//
// Scanning stops unrecoverably at EOF, the first I/O error, or a token too
// large to fit in the buffer. When a scan stops, the reader may have
// advanced arbitrarily far past the last token. Programs that need more
// control over error handling or large tokens, or must run sequential scans
// on a reader, should use [bufio.Reader] instead.
>>>>>>> upstream/master
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

<<<<<<< HEAD
// SplitFuncは、入力をトークン化するために使用されるsplit関数のシグネチャです。
// 引数は、残りの未処理データの初期部分のサブストリングと、Readerにもうデータがないことを報告するフラグであるatEOFです。
// 戻り値は、入力を進めるためのバイト数と、ユーザーに返す次のトークン（あれば）、およびエラー（あれば）です。
//
// 関数がエラーを返すと、スキャンは停止し、入力の一部が破棄される場合があります。エラーがErrFinalTokenである場合、スキャンはエラーなしで停止します。
//
// それ以外の場合、スキャナは入力を進めます。トークンがnilでない場合、スキャナはユーザーにそれを返します。トークンがnilの場合、スキャナはさらにデータを読み込んでスキャンを続けます。もしデータがもうない場合（つまり、atEOFがtrueの場合）、スキャナは終了します。データがまだ完全なトークンを保持していない場合、例えば改行がない場合は、SplitFuncは(0、nil、nil)を返してスキャナにデータをスライスに読み込んで再試行するように指示できます。
=======
// SplitFunc is the signature of the split function used to tokenize the
// input. The arguments are an initial substring of the remaining unprocessed
// data and a flag, atEOF, that reports whether the [Reader] has no more data
// to give. The return values are the number of bytes to advance the input
// and the next token to return to the user, if any, plus an error, if any.
//
// Scanning stops if the function returns an error, in which case some of
// the input may be discarded. If that error is [ErrFinalToken], scanning
// stops with no error.
//
// Otherwise, the [Scanner] advances the input. If the token is not nil,
// the [Scanner] returns it to the user. If the token is nil, the
// Scanner reads more data and continues scanning; if there is no more
// data--if atEOF was true--the [Scanner] returns. If the data does not
// yet hold a complete token, for instance if it has no newline while
// scanning lines, a [SplitFunc] can return (0, nil, nil) to signal the
// [Scanner] to read more data into the slice and try again with a
// longer slice starting at the same point in the input.
>>>>>>> upstream/master
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
<<<<<<< HEAD
	// MaxScanTokenSizeは、トークンをバッファリングするために使用される最大サイズです
	// ユーザーがScanner.Bufferで明示的なバッファを提供しない限り。
	// 実際の最大トークンサイズは、バッファに改行などを含める必要があるため、より小さくなる場合があります。
	MaxScanTokenSize = 64 * 1024
)

// NewScannerはrから読み取るための新しいScannerを返します。
// split関数はデフォルトでScanLinesになります。

func NewScanner(r io.Reader) *Scanner

// ErrはScannerによって遭遇した最初のEOF以外のエラーを返します。
func (s *Scanner) Err() error

// BytesはScanの呼び出しによって生成された最新のトークンを返します。
// 基になる配列は、後続のScanの呼び出しによって上書きされる可能性があります。
// メモリ確保は行われません。
func (s *Scanner) Bytes() []byte

// TextはScanの呼び出しで生成された最新のトークンを、そのバイトを保持する新しく割り当てられた文字列として返します。
=======
	// MaxScanTokenSize is the maximum size used to buffer a token
	// unless the user provides an explicit buffer with [Scanner.Buffer].
	// The actual maximum token size may be smaller as the buffer
	// may need to include, for instance, a newline.
	MaxScanTokenSize = 64 * 1024
)

// NewScanner returns a new [Scanner] to read from r.
// The split function defaults to [ScanLines].
func NewScanner(r io.Reader) *Scanner

// Err returns the first non-EOF error that was encountered by the [Scanner].
func (s *Scanner) Err() error

// Bytes returns the most recent token generated by a call to [Scanner.Scan].
// The underlying array may point to data that will be overwritten
// by a subsequent call to Scan. It does no allocation.
func (s *Scanner) Bytes() []byte

// Text returns the most recent token generated by a call to [Scanner.Scan]
// as a newly allocated string holding its bytes.
>>>>>>> upstream/master
func (s *Scanner) Text() string

// ErrFinalTokenは特別なシグナルエラー値です。これはSplit関数によって返され、エラーと一緒に配信されるトークンが最後のトークンであり、スキャンはこれ以上停止する必要があることを示します。
// ErrFinalTokenがScanによって受け取られた後、エラーなしでスキャンが停止します。
// この値は、処理を早めに停止したい場合や最終の空のトークンを配信する必要がある場合に役立ちます。同様の動作をカスタムエラー値で実現することもできますが、ここで提供することで整理された方法が提供されます。
// この値の使用例については、emptyFinalTokenの例を参照してください。
var ErrFinalToken = errors.New("final token")

<<<<<<< HEAD
// ScanはScannerを次のトークンに進め、それをBytesメソッドまたはTextメソッドで利用できるようにします。入力の終わりまたはエラーによりスキャンが停止すると、falseを返します。
// Scanがfalseを返した後、Errメソッドはスキャン中に発生したエラーを返しますが、もしio.EOFだった場合は、Errはnilを返します。
// スプリット関数が入力を進めずに多くの空のトークンを返すと、Scanはパニックを起こします。これはスキャナーの共通のエラーモードです。
func (s *Scanner) Scan() bool

// Bufferは、スキャン中に使用する初期バッファと、割り当て可能な最大バッファサイズを設定します。
// 最大トークンサイズは、maxとcap(buf)の大きい方よりも小さくなければなりません。
// max <= cap(buf)の場合、Scanはこのバッファのみを使用し、割り当てを行いません。
//
// デフォルトでは、Scanは内部バッファを使用し、最大トークンサイズをMaxScanTokenSizeに設定します。
=======
// Scan advances the [Scanner] to the next token, which will then be
// available through the [Scanner.Bytes] or [Scanner.Text] method. It returns false when the
// scan stops, either by reaching the end of the input or an error.
// After Scan returns false, the [Scanner.Err] method will return any error that
// occurred during scanning, except that if it was [io.EOF], [Scanner.Err]
// will return nil.
// Scan panics if the split function returns too many empty
// tokens without advancing the input. This is a common error mode for
// scanners.
func (s *Scanner) Scan() bool

// Buffer sets the initial buffer to use when scanning
// and the maximum size of buffer that may be allocated during scanning.
// The maximum token size must be less than the larger of max and cap(buf).
// If max <= cap(buf), [Scanner.Scan] will use this buffer only and do no allocation.
//
// By default, [Scanner.Scan] uses an internal buffer and sets the
// maximum token size to [MaxScanTokenSize].
>>>>>>> upstream/master
//
// Bufferはスキャンの開始後に呼び出された場合、パニックを発生させます。
func (s *Scanner) Buffer(buf []byte, max int)

<<<<<<< HEAD
// SplitはScannerの分割関数を設定します。
// デフォルトの分割関数はScanLinesです。
=======
// Split sets the split function for the [Scanner].
// The default split function is [ScanLines].
>>>>>>> upstream/master
//
// Splitはスキャンが開始された後に呼び出された場合、パニックを引き起こします。
func (s *Scanner) Split(split SplitFunc)

<<<<<<< HEAD
// ScanBytesは、各バイトをトークンとして返すScannerの分割関数です。
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)

// ScanRunesは、Scannerに対して使用される分割関数で、それぞれのUTF-8エンコードされたルーンをトークンとして返します。返されるルーンのシーケンスは、文字列としての入力を範囲ループで処理した場合と同等です。つまり、誤ったUTF-8エンコーディングはU+FFFD = "\xef\xbf\xbd"に変換されます。Scanインターフェースのため、クライアントは正しくエンコードされた置換ルーンとエンコーディングエラーを区別することが不可能です。
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)

// ScanLinesは、改行マーカー以降のテキストが除かれた、Scannerのスプリット関数です。
// 返される行は空になる場合もあります。
// 改行マーカーは1つのオプションのキャリッジリターンに続いて1つの必須の改行からなります。
// 正規表現の表記では、 `\r?\n`です。
// 入力の最後の非空行は改行がない場合でも返されます。
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)

// ScanWordsは、スペースで区切られたテキストのScannerのための分割関数です。周囲のスペースは削除され、空の文字列は返されません。スペースの定義はunicode.IsSpaceによって設定されます。
=======
// ScanBytes is a split function for a [Scanner] that returns each byte as a token.
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)

// ScanRunes is a split function for a [Scanner] that returns each
// UTF-8-encoded rune as a token. The sequence of runes returned is
// equivalent to that from a range loop over the input as a string, which
// means that erroneous UTF-8 encodings translate to U+FFFD = "\xef\xbf\xbd".
// Because of the Scan interface, this makes it impossible for the client to
// distinguish correctly encoded replacement runes from encoding errors.
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)

// ScanLines is a split function for a [Scanner] that returns each line of
// text, stripped of any trailing end-of-line marker. The returned line may
// be empty. The end-of-line marker is one optional carriage return followed
// by one mandatory newline. In regular expression notation, it is `\r?\n`.
// The last non-empty line of input will be returned even if it has no
// newline.
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)

// ScanWords is a split function for a [Scanner] that returns each
// space-separated word of text, with surrounding spaces deleted. It will
// never return an empty string. The definition of space is set by
// unicode.IsSpace.
>>>>>>> upstream/master
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
