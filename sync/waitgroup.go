// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"github.com/shogo82148/std/sync/atomic"
)

// WaitGroupは、一連のゴルーチンの完了を待機します。
// メインゴルーチンは、待機するゴルーチンの数を設定するために [WaitGroup.Add] を呼び出します。
// それぞれのゴルーチンは実行され、終了時に [WaitGroup.Done] を呼び出します。
// 同時に、全てのゴルーチンが終了するまでブロックするために [WaitGroup.Wait] を使用できます。
//
// WaitGroupは、初回使用後にコピーしてはいけません。
//
<<<<<<< HEAD
// Goのメモリモデルの用語である、[WaitGroup.Done] への呼び出しは、それによってブロックが解除される任意のWait呼び出しの前に「同期します」。
=======
// In the terminology of [the Go memory model], a call to [WaitGroup.Done]
// “synchronizes before” the return of any Wait call that it unblocks.
//
// [the Go memory model]: https://go.dev/ref/mem
>>>>>>> 41b4a7d0008e48dd077e189fd86911de2b36d90d
type WaitGroup struct {
	noCopy noCopy

	state atomic.Uint64
	sema  uint32
}

// Addは [WaitGroup] のカウンターにデルタを追加します。デルタは負であることもあります。
// カウンターがゼロになると、[WaitGroup.Wait] でブロックされているすべてのゴルーチンが解放されます。
// カウンターが負になると、Addはパニックを発生させます。
//
// カウンターがゼロの状態で正のデルタで呼び出される場合は、Waitより前に実行される必要があることに注意してください。
// 負のデルタで呼び出される場合や、カウンターがゼロより大きい状態で正のデルタで呼ばれる場合は、任意のタイミングで発生する場合があります。
// 通常、これはAddの呼び出しは、ゴルーチンの作成または待機する他のイベントの直前に実行されるべきことを意味します。
// WaitGroupが複数の独立したイベントセットを待機するために再利用される場合、新しいAdd呼び出しは以前のすべてのWait呼び出しが返された後に行われる必要があります。
// WaitGroupの例を参照してください。
func (wg *WaitGroup) Add(delta int)

// Doneは [WaitGroup] のカウンターを1つ減らします。
func (wg *WaitGroup) Done()

// Wait は [WaitGroup] カウンタがゼロになるまでブロックします。
func (wg *WaitGroup) Wait()
