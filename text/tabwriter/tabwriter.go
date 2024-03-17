// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// tabwriterパッケージは、入力のタブ区切りの列を適切に整列したテキストに変換する
// 書き込みフィルタ（tabwriter.Writer）を実装します。
//
// このパッケージは、http://nickgravgaard.com/elastictabstops/index.html で
// 説明されているElastic Tabstopsアルゴリズムを使用しています。
//
// text/tabwriterパッケージは凍結されており、新しい機能は受け入れていません。
package tabwriter

import (
	"github.com/shogo82148/std/io"
)

// Writerは、入力のタブ区切りの列の周囲にパディングを挿入して、
// 出力でそれらを整列させるフィルタです。
//
// Writerは、入力バイトを水平('\t')または垂直('\v')のタブ、
// 改行('\n')またはフォームフィード('\f')文字で終了するセルとして
// 扱います。改行とフォームフィードの両方が行の区切りとして機能します。
//
// 連続する行のタブで終了するセルは列を構成します。Writerは、
// 列内のすべてのセルが同じ幅になるように必要に応じてパディングを挿入し、
// 事実上、列を整列させます。すべての文字が同じ幅を持つと仮定していますが、
// タブについてはタブ幅を指定する必要があります。列のセルはタブで終了する必要があり、
// タブで区切られるべきではありません：行の終わりの非タブで終了する末尾のテキストは
// セルを形成しますが、そのセルは整列した列の一部ではありません。
// 例えば、この例では（ここで | は水平タブを表します）：
//
//	aaaa|bbb|d
//	aa  |b  |dd
//	a   |
//	aa  |cccc|eee
//
// bとcは別々の列にあります（b列は連続していません）。
// dとeは全く列にありません（終端のタブがなく、列も連続していません）。
//
// Writerは、すべてのUnicodeコードポイントが同じ幅を持つと仮定しています。
// これは、一部のフォントでは真ではないかもしれません、または文字列が結合文字を含んでいる場合。
//
<<<<<<< HEAD
// DiscardEmptyColumnsが設定されている場合、垂直（または「ソフト」）タブによって
// 完全に終了する空の列は破棄されます。水平（または「ハード」）タブで終了する列は
// このフラグの影響を受けません。
=======
// If [DiscardEmptyColumns] is set, empty columns that are terminated
// entirely by vertical (or "soft") tabs are discarded. Columns
// terminated by horizontal (or "hard") tabs are not affected by
// this flag.
>>>>>>> upstream/master
//
// WriterがHTMLをフィルタリングするように設定されている場合、HTMLタグとエンティティは
// そのまま通過します。タグとエンティティの幅は、フォーマットの目的でゼロ（タグ）と
// 一（エンティティ）とみなされます。
//
<<<<<<< HEAD
// テキストのセグメントは、Escape文字でそれを括ることでエスケープできます。
// tabwriterはエスケープされたテキストセグメントをそのまま通過させます。
// 特に、セグメント内のタブや改行は解釈しません。StripEscapeフラグが設定されている場合、
// Escape文字は出力から削除されます。それ以外の場合、それらもそのまま通過します。
// フォーマットの目的で、エスケープされたテキストの幅は常にEscape文字を除いて計算されます。
//
// フォームフィード文字は改行のように機能しますが、現在の行のすべての列も終了します
// （事実上Flushを呼び出します）。次の行のタブで終了するセルは新しい列を開始します。
// HTMLタグ内やエスケープされたテキストセグメント内で見つからない限り、
// フォームフィード文字は出力で改行として表示されます。
//
// Writerは、適切な行の間隔が将来の行のセルに依存する可能性があるため、
// 入力を内部的にバッファリングする必要があります。クライアントは、
// Writeの呼び出しが終了したらFlushを呼び出す必要があります。
=======
// A segment of text may be escaped by bracketing it with [Escape]
// characters. The tabwriter passes escaped text segments through
// unchanged. In particular, it does not interpret any tabs or line
// breaks within the segment. If the [StripEscape] flag is set, the
// Escape characters are stripped from the output; otherwise they
// are passed through as well. For the purpose of formatting, the
// width of the escaped text is always computed excluding the Escape
// characters.
//
// The formfeed character acts like a newline but it also terminates
// all columns in the current line (effectively calling [Writer.Flush]). Tab-
// terminated cells in the next line start new columns. Unless found
// inside an HTML tag or inside an escaped text segment, formfeed
// characters appear as newlines in the output.
//
// The Writer must buffer input internally, because proper spacing
// of one line may depend on the cells in future lines. Clients must
// call Flush when done calling [Writer.Write].
>>>>>>> upstream/master
type Writer struct {
	// configuration
	output   io.Writer
	minwidth int
	tabwidth int
	padding  int
	padbytes [8]byte
	flags    uint

	// current state
	buf     []byte
	pos     int
	cell    cell
	endChar byte
	lines   [][]cell
	widths  []int
}

// これらのフラグを使用して、フォーマットを制御できます。
const (
	// HTMLタグを無視し、エンティティ（'&'で始まり';'で終わる）を単一の文字（幅=1）として扱います。
	FilterHTML uint = 1 << iota

	// エスケープされたテキストセグメントを括るエスケープ文字を削除します。
	// テキストと一緒に変更せずにそれらを通過させる代わりに。
	StripEscape

	// セルの内容を右揃えに強制します。
	// デフォルトは左揃えです。
	AlignRight

	// 空の列を、最初から入力に存在しなかったかのように扱います。
	DiscardEmptyColumns

	// 常にタブをインデント列（つまり、左側の先頭の空セルのパディング）に使用します。
	// padcharに関係なく。
	TabIndent

	// 列の間に垂直バー ('|') を印刷します（フォーマット後）。
	// 破棄された列はゼロ幅の列として表示されます ("||")。
	Debug
)

<<<<<<< HEAD
// Writerは、Initへの呼び出しで初期化する必要があります。最初のパラメータ（output）は
// フィルタ出力を指定します。残りのパラメータはフォーマットを制御します：
=======
// A [Writer] must be initialized with a call to Init. The first parameter (output)
// specifies the filter output. The remaining parameters control the formatting:
>>>>>>> upstream/master
//
//	minwidth	パディングを含む最小セル幅
//	tabwidth	タブ文字の幅（相当するスペースの数）
//	padding		セルの幅を計算する前にセルに追加されるパディング
//	padchar		パディングに使用されるASCII文字
//			もし padchar == '\t' なら、Writerはフォーマットされた出力の
//			'\t'の幅がtabwidthであると仮定し、align_leftに関係なく
//			セルは左揃えになります
//			（正確に見える結果のために、tabwidthは結果を表示するビューアの
//			タブ幅に対応している必要があります）
//	flags		フォーマット制御
func (b *Writer) Init(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer

// テキストセグメントをエスケープするには、Escape文字でそれを括ります。
// 例えば、この文字列 "Ignore this tab: \xff\t\xff" のタブはセルを終了せず、
// フォーマットの目的で幅一の単一文字を構成します。
//
// 値0xffは、有効なUTF-8シーケンスには現れないため選ばれました。
const Escape = '\xff'

<<<<<<< HEAD
// Writeの最後の呼び出し後にFlushを呼び出す必要があります。これにより、
// Writerにバッファリングされたデータがすべて出力に書き込まれます。終了時に不完全な
// エスケープシーケンスは、フォーマットの目的で完全と見なされます。
=======
// Flush should be called after the last call to [Writer.Write] to ensure
// that any data buffered in the [Writer] is written to output. Any
// incomplete escape sequence at the end is considered
// complete for formatting purposes.
>>>>>>> upstream/master
func (b *Writer) Flush() error

// Writeは、bufをライターbに書き込みます。
// 返されるエラーは、基礎となる出力ストリームへの書き込み中に遭遇したものだけです。
func (b *Writer) Write(buf []byte) (n int, err error)

<<<<<<< HEAD
// NewWriterは新しいtabwriter.Writerを割り当てて初期化します。
// パラメータはInit関数と同じです。
=======
// NewWriter allocates and initializes a new [Writer].
// The parameters are the same as for the Init function.
>>>>>>> upstream/master
func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
