// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// Scannerは、改行で区切られたテキストの行のファイルなどのデータを読み取るための便利なインターフェースを提供します。Scanメソッドの連続した呼び出しにより、ファイルの「トークン」を順番にステップし、トークン間のバイトをスキップします。トークンの仕様はSplitFunc型の分割関数によって定義されます。デフォルトの分割関数は、入力を行に分割し、行末の文字を取り除きます。分割関数は、ファイルを行、バイト、UTF-8エンコードされたルーン、スペースで区切られた単語にスキャンするために、このパッケージで定義されています。クライアントは代わりにカスタムの分割関数を提供することもできます。
// スキャンは、EOF、最初のI/Oエラー、またはバッファに収まらない大きすぎるトークンで回復不可能に停止します。スキャンが停止すると、リーダーは最後のトークンを遥かに超えて進んでいる可能性があります。エラーハンドリングや大きなトークンの制御をより細かく制御する必要があるプログラム、またはリーダーで順次スキャンを実行する必要があるプログラムは、代わりにbufio.Readerを使用する必要があります。
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
// 引数は、残りの未処理データの初期部分のサブストリングと、Readerにもうデータがないことを報告するフラグであるatEOFです。
// 戻り値は、入力を進めるためのバイト数と、ユーザーに返す次のトークン（あれば）、およびエラー（あれば）です。
//
// 関数がエラーを返すと、スキャンは停止し、入力の一部が破棄される場合があります。エラーがErrFinalTokenである場合、スキャンはエラーなしで停止します。
//
// それ以外の場合、スキャナは入力を進めます。トークンがnilでない場合、スキャナはユーザーにそれを返します。トークンがnilの場合、スキャナはさらにデータを読み込んでスキャンを続けます。もしデータがもうない場合（つまり、atEOFがtrueの場合）、スキャナは終了します。データがまだ完全なトークンを保持していない場合、例えば改行がない場合は、SplitFuncは(0、nil、nil)を返してスキャナにデータをスライスに読み込んで再試行するように指示できます。
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
func (s *Scanner) Text() string

// ErrFinalTokenは特別なシグナルエラー値です。これはSplit関数によって返され、エラーと一緒に配信されるトークンが最後のトークンであり、スキャンはこれ以上停止する必要があることを示します。
// ErrFinalTokenがScanによって受け取られた後、エラーなしでスキャンが停止します。
// この値は、処理を早めに停止したい場合や最終の空のトークンを配信する必要がある場合に役立ちます。同様の動作をカスタムエラー値で実現することもできますが、ここで提供することで整理された方法が提供されます。
// この値の使用例については、emptyFinalTokenの例を参照してください。
var ErrFinalToken = errors.New("final token")

// ScanはScannerを次のトークンに進め、それをBytesメソッドまたはTextメソッドで利用できるようにします。入力の終わりまたはエラーによりスキャンが停止すると、falseを返します。
// Scanがfalseを返した後、Errメソッドはスキャン中に発生したエラーを返しますが、もしio.EOFだった場合は、Errはnilを返します。
// スプリット関数が入力を進めずに多くの空のトークンを返すと、Scanはパニックを起こします。これはスキャナーの共通のエラーモードです。
func (s *Scanner) Scan() bool

// Bufferはスキャン時に使用する初期バッファと、スキャン中に割り当て可能なバッファの最大サイズを設定します。最大トークンサイズはmaxとcap(buf)のうち大きい方です。max <= cap(buf)の場合、Scanはこのバッファのみを使用し、割り当ては行いません。
//
// デフォルトでは、Scanは内部バッファを使用し、最大トークンサイズをMaxScanTokenSizeに設定します。
//
// Bufferはスキャンの開始後に呼び出された場合、パニックを発生させます。
func (s *Scanner) Buffer(buf []byte, max int)

// SplitはScannerの分割関数を設定します。
// デフォルトの分割関数はScanLinesです。
//
// Splitはスキャンが開始された後に呼び出された場合、パニックを引き起こします。
func (s *Scanner) Split(split SplitFunc)

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
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
