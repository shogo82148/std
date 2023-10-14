// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lzw

import (
	"github.com/shogo82148/std/io"
)

// WriterはLZWコンプレッサーです。データの圧縮形式を
// 下層のライターに書き込みます（ [NewWriter] を参照）。
type Writer struct {
	// wは圧縮バイトが書き込まれるライターです。
	w writer

	// order, write, bits, nBits, and widthは、
	// コードストリームをバイトストリームに変換するための状態です。
	order Order
	write func(*Writer, uint32) error
	bits  uint32
	nBits uint
	width uint
	// litWidthはリテラルコードのビット幅です。
	litWidth uint

	// hiは次のコード生成によって暗示されるコードです。
	// overflowはhiがコードの幅を超えるコードです。
	hi, overflow uint32

	// savedCodeは最新のWrite呼び出しの終わりに蓄積されるコードです。
	// もしWrite呼び出しがなかった場合、invalidCodeと等しいです。
	savedCode uint32

	// err は書き込み中に最初に発生したエラーです。ライターをクローズすると、
	// 以降の書き込み呼び出しは errClosed を返します。
	err error

	// tableは20ビットのキーから12ビットの値へのハッシュテーブルです。各テーブルエントリにはkey<<12|valが含まれており、衝突は線形探査法によって解決されます。キーは12ビットのコード接頭辞と8ビットのバイト接尾辞で構成されます。値は12ビットのコードです。
	table [tableSize]uint32
}

// Writeはpの圧縮された表現をwの基になるライターに書き込みます。
func (w *Writer) Write(p []byte) (n int, err error)

// Closeは [Writer] を閉じ、保留中の出力をフラッシュします。wの基になるライターは閉じません。
func (w *Writer) Close() error

<<<<<<< HEAD
// Resetは [Writer] の状態をクリアし、新しい [Writer] として再利用できるようにします。
=======
// Reset clears the [Writer]'s state and allows it to be reused again
// as a new [Writer].
>>>>>>> upstream/master
func (w *Writer) Reset(dst io.Writer, order Order, litWidth int)

// NewWriterは新しい [io.WriteCloser] を作成します。
// 返された [io.WriteCloser] に書き込まれたデータは圧縮され、wに書き込まれます。
// 書き込みが完了した場合、呼び出し元の責任で [io.WriteCloser] をCloseする必要があります。
// リテラルコードに使用するビット数であるlitWidthは、範囲[2,8]内でなければなりませんが、通常は8です。
// 入力バイトは1<<litWidth未満でなければなりません。
//
// 返された [io.WriteCloser] の基になる型が [*Writer] であることが保証されます。
func NewWriter(w io.Writer, order Order, litWidth int) io.WriteCloser
