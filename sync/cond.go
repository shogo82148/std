// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

// Condはイベントの発生を待つまたは宣言するための待機ポイントである条件変数を実装します。
//
// 各Condには関連付けられたLocker L（通常は [*Mutex] または [*RWMutex]）があり、条件を変更するときや [Cond.Wait] メソッドを呼び出すときに保持する必要があります。
//
// 最初の使用後にCondをコピーしてはいけません。
//
<<<<<<< HEAD
// Goのメモリモデルの用語では、Condは
// [Cond.Broadcast] または [Cond.Signal] の呼び出しが、それによってブロック解除される任意のWait呼び出しよりも
// 「先に同期する」ように配置します。
=======
// In the terminology of [the Go memory model], Cond arranges that
// a call to [Cond.Broadcast] or [Cond.Signal] “synchronizes before” any Wait call
// that it unblocks.
>>>>>>> 41b4a7d0008e48dd077e189fd86911de2b36d90d
//
// 単純な使用例では、ユーザーはチャネルを使用する方がCondよりも優れています（Broadcastはチャネルを閉じることに対応し、Signalはチャネルに送信することに対応します）。
//
// [sync.Cond] の代わりに他のものについては、[Roberto Clapis's series on
// advanced concurrency patterns] と [Bryan Mills's talk on concurrency
// patterns] を参照してください。
//
// [the Go memory model]: https://go.dev/ref/mem
// [Roberto Clapis's series on advanced concurrency patterns]: https://blogtitle.github.io/categories/concurrency/
// [Bryan Mills's talk on concurrency patterns]: https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view
type Cond struct {
	noCopy noCopy

	// 条件を観察または変更する間は、Lを保持します
	L Locker

	notify  notifyList
	checker copyChecker
}

// NewCondはLocker lを持つ新しいCondを返します。
func NewCond(l Locker) *Cond

// Waitはc.Lのロックを解除して、呼び出し元のゴルーチンの実行を一時停止します。
// 後で再開すると、Waitは戻る前にc.Lをロックします。他のシステムとは異なり、
// Waitは [Cond.Broadcast] または [Cond.Signal] によって起こされない限り戻りません。
//
// Waitが待機している間、c.Lはロックされていないため、呼び出し元は
// 待機が返るときに条件が真であることを前提とすることはできません。代わりに、
// 呼び出し元はループ内でWaitを使用する必要があります：
//
//	c.L.Lock()
//	for !condition() {
//	    c.Wait()
//	}
//	... 条件を活用する ...
//	c.L.Unlock()
func (c *Cond) Wait()

// cに待機しているゴルーチンがあれば、Signalは1つのゴルーチンを起こします。
//
// 呼び出し元がc.Lを保持していることは必須ではありませんが、許可されています。
//
// Signal()はゴルーチンのスケジューリングの優先順位に影響を与えません。他のゴルーチンがc.Lをロックしようとしている場合、"待機中"のゴルーチンよりも先に起きる場合があります。
func (c *Cond) Signal()

// Broadcastは、cで待機しているすべてのゴルーチンを起こします。
//
// 呼び出し元がc.Lを保持していることは許可されていますが、必須ではありません。
func (c *Cond) Broadcast()
