// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"github.com/shogo82148/std/sync/atomic"
)

// WaitGroupは、主に複数のゴルーチンやタスクの終了を待つために使われるカウント型セマフォです。
//
// 通常、メインゴルーチンは各タスクを新しいゴルーチンで起動する際に [WaitGroup.Go] を呼び出し、
// すべてのタスクが完了するまで [WaitGroup.Wait] を呼び出して待機します。例：
//
//	var wg sync.WaitGroup
//	wg.Go(task1)
//	wg.Go(task2)
//	wg.Wait()
//
// WaitGroupは、Goを使わずにタスクの追跡にも利用でき、[WaitGroup.Add] と [WaitGroup.Done] を使います。
//
// 前述の例は、AddとDoneを使って明示的にゴルーチンを作成する形でも書き換えられます：
//
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		task1()
//	}()
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		task2()
//	}()
//	wg.Wait()
//
// このパターンは [WaitGroup.Go] より前のコードでよく使われます。
//
// WaitGroupは最初に使用した後でコピーしてはいけません。
type WaitGroup struct {
	noCopy noCopy

	// Bits (high to low):
	//   bits[0:32]  counter
	//   bits[32]    flag: synctest bubble membership
	//   bits[33:64] wait count
	state atomic.Uint64
	sema  uint32
}

// Addは、[WaitGroup] のタスクカウンターにdelta（負でも可）を加算します。
// カウンターがゼロになると、[WaitGroup.Wait] でブロックされているすべてのゴルーチンが解放されます。
// カウンターが負になると、Addはパニックを起こします。
//
// 呼び出し側は [WaitGroup.Go]を優先して使用すべきです。
//
// カウンターがゼロのときに正のdeltaで呼び出す場合、Waitの前に実行する必要があります。
// カウンターがゼロより大きいときに正のdeltaで呼び出す場合や、負のdeltaで呼び出す場合は、いつでも実行できます。
// 通常、Addの呼び出しは、ゴルーチンや待機対象のイベントを作成する文の前に実行します。
// WaitGroupを複数の独立したイベントセットの待機に再利用する場合は、
// 新しいAddの呼び出しは、すべての前回のWait呼び出しが返った後に実行する必要があります。
// WaitGroupの例を参照してください。
func (wg *WaitGroup) Add(delta int)

// Doneは、[WaitGroup] のタスクカウンターを1減算します。
// Add(-1)と同じです。
//
// 呼び出し側は [WaitGroup.Go] を優先して使用すべきです。
//
// [the Go memory model] の用語では、Doneの呼び出しは、
// それによって解除されるWait呼び出しの返り値より「先に同期」します。
//
// [the Go memory model]: https://go.dev/ref/mem
func (wg *WaitGroup) Done()

// Waitは、[WaitGroup]のタスクカウンターがゼロになるまでブロックします。
func (wg *WaitGroup) Wait()

// Goはfを新しいゴルーチンで呼び出し、そのタスクを [WaitGroup] に追加します。
// fが返ると、そのタスクはWaitGroupから削除されます。
//
// fはパニックを起こしてはいけません。
//
// WaitGroupが空の場合、Goは [WaitGroup.Wait] より前に実行されなければなりません。
// 通常は、Goでタスクを開始してからWaitを呼び出します。
// WaitGroupが空でない場合、Goはいつでも実行できます。
// Goで開始されたゴルーチン自身がGoを呼び出すこともできます。
// WaitGroupを複数の独立したタスクセットの待機に再利用する場合は、
// 新しいGoの呼び出しは、すべての前回のWait呼び出しが返った後に実行する必要があります。
//
// [the Go memory model] の用語では、fの返り値は、
// それによって解除されるWait呼び出しの返り値より「先に同期」します。
//
// [the Go memory model]: https://go.dev/ref/mem
func (wg *WaitGroup) Go(f func())
