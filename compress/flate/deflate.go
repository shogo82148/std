// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flate

import (
	"github.com/shogo82148/std/io"
)

const (
	NoCompression      = 0
	BestSpeed          = 1
	BestCompression    = 9
	DefaultCompression = -1

	// HuffmanOnlyは、Lempel-Zivマッチ検索を無効化し、Huffmanエントロピー符号化のみを実行します。このモードは、すでにLZスタイルのアルゴリズム（例：SnappyやLZ4）で圧縮されたデータを圧縮する際に便利ですが、それにはエントロピー符号化が欠けています。特定のバイトが入力ストリームで他のバイトよりも頻繁に発生する場合、圧縮率が向上します。
	// HuffmanOnlyは、圧縮された出力をRFC 1951に準拠して生成します。つまり、有効なDEFLATEデコンプレッサーは、この出力を引き続き解凍できるようになります。
	HuffmanOnly = -2
)

<<<<<<< HEAD
// NewWriter は指定されたレベルでデータを圧縮する新しい Writer を返します。
// zlib に従って、レベルは 1 (BestSpeed) から 9 (BestCompression) の範囲です。
// より高いレベルでは一般的に圧縮がより効率的ですが、速度は遅くなります。
// レベル 0 (NoCompression) では圧縮は試みず、必要な DEFLATE フレーミングのみが追加されます。
// レベル -1 (DefaultCompression) はデフォルトの圧縮レベルを使用します。
// レベル -2 (HuffmanOnly) は Huffman 圧縮のみを使用し、全ての入力の圧縮を非常に高速化しますが、
// 圧縮効率を犠牲にします。
=======
// NewWriter returns a new [Writer] compressing data at the given level.
// Following zlib, levels range from 1 ([BestSpeed]) to 9 ([BestCompression]);
// higher levels typically run slower but compress more. Level 0
// ([NoCompression]) does not attempt any compression; it only adds the
// necessary DEFLATE framing.
// Level -1 ([DefaultCompression]) uses the default compression level.
// Level -2 ([HuffmanOnly]) will use Huffman compression only, giving
// a very fast compression for all types of input, but sacrificing considerable
// compression efficiency.
>>>>>>> upstream/master
//
// もし level が [-2, 9] の範囲にある場合、返されたエラーは nil になります。
// そうでない場合、返されたエラーは nil ではありません。
func NewWriter(w io.Writer, level int) (*Writer, error)

<<<<<<< HEAD
// NewWriterDictはNewWriterと似ていますが、新しいWriterをプリセット辞書で初期化します。返されたWriterは、圧縮された出力を生成せずに、辞書が書き込まれたかのように振る舞います。wに書き込まれた圧縮データは、同じ辞書で初期化されたReaderでのみ解凍することができます。
func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)

// Writerは、書き込まれたデータを受け取り、そのデータの圧縮された形式を基になるWriterに書き込む。 (NewWriterを参照してください)。
=======
// NewWriterDict is like [NewWriter] but initializes the new
// [Writer] with a preset dictionary. The returned [Writer] behaves
// as if the dictionary had been written to it without producing
// any compressed output. The compressed data written to w
// can only be decompressed by a [Reader] initialized with the
// same dictionary.
func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)

// A Writer takes data written to it and writes the compressed
// form of that data to an underlying writer (see [NewWriter]).
>>>>>>> upstream/master
type Writer struct {
	d    compressor
	dict []byte
}

// Writeは、wにデータを書き込み、その後データの圧縮形式を基になるライターに書き込みます。
func (w *Writer) Write(data []byte) (n int, err error)

<<<<<<< HEAD
// Flushは、保留中のデータを基礎となるライターにフラッシュします。
// 主に圧縮されたネットワークプロトコルで有用であり、リモートのリーダーがパケットを再構築するのに十分なデータがあることを確保します。
// データが書き込まれるまで、Flushは処理を返しません。
// パンディングのないデータでFlushを呼び出すと、ライターは少なくとも4バイトの同期マーカーを出力します。
// 基礎となるライターがエラーを返す場合、Flushはそのエラーを返します。
=======
// Flush flushes any pending data to the underlying writer.
// It is useful mainly in compressed network protocols, to ensure that
// a remote reader has enough data to reconstruct a packet.
// Flush does not return until the data has been written.
// Calling Flush when there is no pending data still causes the [Writer]
// to emit a sync marker of at least 4 bytes.
// If the underlying writer returns an error, Flush returns that error.
>>>>>>> upstream/master
//
// zlibライブラリの用語では、FlushはZ_SYNC_FLUSHと等価です。
func (w *Writer) Flush() error

// Close は書き込みバッファをフラッシュしてクローズします。
func (w *Writer) Close() error

<<<<<<< HEAD
// Resetは、ライターの状態を破棄し、dstとwのレベルと辞書を使用してNewWriterまたはNewWriterDictが呼び出された結果と同じ状態にします。
=======
// Reset discards the writer's state and makes it equivalent to
// the result of [NewWriter] or [NewWriterDict] called with dst
// and w's level and dictionary.
>>>>>>> upstream/master
func (w *Writer) Reset(dst io.Writer)
