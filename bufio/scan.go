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
// 関数がエラーを返した場合、スキャンは停止します。この場合、入力の一部が破棄される可能性があります。
// エラーが [ErrFinalToken] である場合、スキャンはエラーなしで停止します。
// [ErrFinalToken] と一緒に非nilトークンが配信される場合、最後のトークンとなり、
// [ErrFinalToken] と一緒にnilトークンが配信される場合、スキャンが直ちに停止します。
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

// ErrFinalTokenは、特別なセンチネルエラー値です。
// スキャンをエラーなしで停止することを示すために、Split関数によって返されることを意図しています。
// このエラーと一緒に配信されるトークンがnilでない場合、トークンは最後のトークンです。
//
// この値は、処理を早期に停止する必要がある場合や、最後の空のトークン（nilトークンとは異なる）を配信する必要がある場合に役立ちます。
// カスタムエラー値で同じ動作を実現できますが、ここで提供することで整理されたコードになります。
// この値の使用例については、emptyFinalTokenの例を参照してください。
var ErrFinalToken = errors.New("final token")

// Scanは、 [Scanner.Bytes] または [Scanner.Text] メソッドを介して利用可能な次のトークンに [Scanner] を進めます。
// 入力の終わりに到達するかエラーが発生すると、falseを返します。
// Scanがfalseを返した後、 [Scanner.Err] メソッドはスキャン中に発生したエラーを返しますが、
// エラーが[io.EOF]の場合、 [Scanner.Err] はnilを返します。
// スキャンが進まずに空のトークンを多数返す場合、スキャナーはpanicします。
// これは、スキャナーの一般的なエラーモードです。
func (s *Scanner) Scan() bool

<<<<<<< HEAD
// Bufferは、スキャン中に使用する初期バッファと、割り当て可能な最大バッファサイズを設定します。
// 最大トークンサイズは、maxとcap(buf)の大きい方よりも小さくなければなりません。
// max <= cap(buf)の場合、 [Scanner.Scan] はこのバッファのみを使用し、割り当てを行いません。
=======
// Buffer controls memory allocation by the Scanner.
// It sets the initial buffer to use when scanning
// and the maximum size of buffer that may be allocated during scanning.
// The contents of the buffer are ignored.
//
// The maximum token size must be less than the larger of max and cap(buf).
// If max <= cap(buf), [Scanner.Scan] will use this buffer only and do no allocation.
>>>>>>> upstream/release-branch.go1.25
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
