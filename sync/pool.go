// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"github.com/shogo82148/std/unsafe"
)

// プールとは個別に保存および取り出しが可能な一時オブジェクトのセットです。
//
// プールに保存されたアイテムは、通知なしにいつでも自動的に削除される可能性があります。
// これが発生する際にプールが唯一の参照を保持している場合、そのアイテムは解放される可能性があります。
//
// プールは、複数のゴルーチンによる同時の使用に対して安全です。
//
// プールの目的は、割り当てされたが未使用のアイテムをキャッシュして後で再利用することで、
// ガベージコレクタにかかる負荷を軽減することです。つまり、効率的かつスレッドセーフな無料リストを構築することを簡単にします。
// ただし、プールはすべての無料リストに適しているわけではありません。
//
// プールの適切な使用例は、パッケージの独立した並行クライアント間で静に共有され、
// 潜在的に再利用される一時アイテムのグループを管理することです。
// プールは、多くのクライアント間で割り当てのオーバーヘッドを分散する方法を提供します。
//
// プールの良い使用例は、fmtパッケージにあります。これは一時的な出力バッファのサイズを動的に管理しています。
// このストアは、負荷がかかった場合（多くのゴルーチンがアクティブに印刷している場合）に拡大し、静かな場合には縮小します。
//
// 一方、短命なオブジェクトの一部として維持される無料リストは、プールには適していません。
// なぜなら、このシナリオではオーバーヘッドを適切に分散させることができないからです。
// このようなオブジェクトは、独自の無料リストを実装する方が効率的です。
//
// プールは、最初の使用後にコピーしないでください。
//
<<<<<<< HEAD
// Goのメモリモデルの用語では、Put(x)の呼び出しは、
// 同じ値xを返す [Pool.Get] の呼び出しよりも「先に同期します」。
// 同様に、Newがxを返す呼び出しは、
// 同じ値xを返すGetの呼び出しよりも「先に同期します」。
=======
// In the terminology of [the Go memory model], a call to Put(x) “synchronizes before”
// a call to [Pool.Get] returning that same value x.
// Similarly, a call to New returning x “synchronizes before”
// a call to Get returning that same value x.
//
// [the Go memory model]: https://go.dev/ref/mem
>>>>>>> 41b4a7d0008e48dd077e189fd86911de2b36d90d
type Pool struct {
	noCopy noCopy

	local     unsafe.Pointer
	localSize uintptr

	victim     unsafe.Pointer
	victimSize uintptr

	// Newは、Getがnilを返す場合に値を生成するための関数を指定するオプションです。
	// Getの呼び出しと同時に変更することはできません。
	New func() any
}

// xをプールに追加します。
func (p *Pool) Put(x any)

// GetはPoolからランダムなアイテムを選択し、[Pool] から削除して呼び出し元に返します。
// Getはプールを無視して空として扱うことを選択する場合があります。
// [Pool.Put] に渡された値とGetが返す値の間には、呼び出し元は何の関係も仮定すべきではありません。
//
// Getが通常nilを返す場合であり、p.Newがnilでない場合、Getはp.Newを呼び出した結果を返します。
func (p *Pool) Get() any
