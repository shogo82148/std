// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"github.com/shogo82148/std/sync/atomic"
)

<<<<<<< HEAD
// WaitGroupは、一連のゴルーチンの完了を待機します。
// メインゴルーチンは、待機するゴルーチンの数を設定するために [WaitGroup.Add] を呼び出します。
// それぞれのゴルーチンは実行され、終了時に [WaitGroup.Done] を呼び出します。
// 同時に、全てのゴルーチンが終了するまでブロックするために [WaitGroup.Wait] を使用できます。
//
// WaitGroupは、初回使用後にコピーしてはいけません。
//
// [the Go memory model] の用語である、[WaitGroup.Done] への呼び出しは、それによってブロックが解除される任意のWait呼び出しの前に「同期します」。
//
// [the Go memory model]: https://go.dev/ref/mem
=======
// A WaitGroup is a counting semaphore typically used to wait
// for a group of goroutines or tasks to finish.
//
// Typically, a main goroutine will start tasks, each in a new
// goroutine, by calling [WaitGroup.Go] and then wait for all tasks to
// complete by calling [WaitGroup.Wait]. For example:
//
//	var wg sync.WaitGroup
//	wg.Go(task1)
//	wg.Go(task2)
//	wg.Wait()
//
// A WaitGroup may also be used for tracking tasks without using Go to
// start new goroutines by using [WaitGroup.Add] and [WaitGroup.Done].
//
// The previous example can be rewritten using explicitly created
// goroutines along with Add and Done:
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
// This pattern is common in code that predates [WaitGroup.Go].
//
// A WaitGroup must not be copied after first use.
>>>>>>> upstream/release-branch.go1.25
type WaitGroup struct {
	noCopy noCopy

	// Bits (high to low):
	//   bits[0:32]  counter
	//   bits[32]    flag: synctest bubble membership
	//   bits[33:64] wait count
	state atomic.Uint64
	sema  uint32
}

<<<<<<< HEAD
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
=======
// Add adds delta, which may be negative, to the [WaitGroup] task counter.
// If the counter becomes zero, all goroutines blocked on [WaitGroup.Wait] are released.
// If the counter goes negative, Add panics.
//
// Callers should prefer [WaitGroup.Go].
//
// Note that calls with a positive delta that occur when the counter is zero
// must happen before a Wait. Calls with a negative delta, or calls with a
// positive delta that start when the counter is greater than zero, may happen
// at any time.
// Typically this means the calls to Add should execute before the statement
// creating the goroutine or other event to be waited for.
// If a WaitGroup is reused to wait for several independent sets of events,
// new Add calls must happen after all previous Wait calls have returned.
// See the WaitGroup example.
func (wg *WaitGroup) Add(delta int)

// Done decrements the [WaitGroup] task counter by one.
// It is equivalent to Add(-1).
//
// Callers should prefer [WaitGroup.Go].
//
// In the terminology of [the Go memory model], a call to Done
// "synchronizes before" the return of any Wait call that it unblocks.
//
// [the Go memory model]: https://go.dev/ref/mem
func (wg *WaitGroup) Done()

// Wait blocks until the [WaitGroup] task counter is zero.
>>>>>>> upstream/release-branch.go1.25
func (wg *WaitGroup) Wait()

// Go calls f in a new goroutine and adds that task to the [WaitGroup].
// When f returns, the task is removed from the WaitGroup.
//
// The function f must not panic.
//
// If the WaitGroup is empty, Go must happen before a [WaitGroup.Wait].
// Typically, this simply means Go is called to start tasks before Wait is called.
// If the WaitGroup is not empty, Go may happen at any time.
// This means a goroutine started by Go may itself call Go.
// If a WaitGroup is reused to wait for several independent sets of tasks,
// new Go calls must happen after all previous Wait calls have returned.
//
// In the terminology of [the Go memory model], the return from f
// "synchronizes before" the return of any Wait call that it unblocks.
//
// [the Go memory model]: https://go.dev/ref/mem
func (wg *WaitGroup) Go(f func())
