// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"github.com/shogo82148/std/sync/atomic"
)

// A WaitGroup is a counting semaphore typically used to wait
// for a group of goroutines to finish.
//
// The main goroutine calls [WaitGroup.Add] to set (or increase) the number of
// goroutines to wait for. Then each of the goroutines
// runs and calls [WaitGroup.Done] when finished. At the same time,
// [WaitGroup.Wait] can be used to block until all goroutines have finished.
//
// This is a typical pattern of WaitGroup usage to
// synchronize 3 goroutines, each calling the function f:
//
//	var wg sync.WaitGroup
//	for range 3 {
//	   wg.Add(1)
//	   go func() {
//	       defer wg.Done()
//	       f()
//	   }()
//	}
//	wg.Wait()
//
// For convenience, the [WaitGroup.Go] method simplifies this pattern to:
//
//	var wg sync.WaitGroup
//	for range 3 {
//	   wg.Go(f)
//	}
//	wg.Wait()
//
// A WaitGroup must not be copied after first use.
//
// In the terminology of [the Go memory model], a call to [WaitGroup.Done]
// “synchronizes before” the return of any Wait call that it unblocks.
//
// [the Go memory model]: https://go.dev/ref/mem
type WaitGroup struct {
	noCopy noCopy

	state atomic.Uint64
	sema  uint32
}

// Add adds delta, which may be negative, to the [WaitGroup] counter.
// If the counter becomes zero, all goroutines blocked on [WaitGroup.Wait] are released.
// If the counter goes negative, Add panics.
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

// Done decrements the [WaitGroup] counter by one.
func (wg *WaitGroup) Done()

// Wait blocks until the [WaitGroup] counter is zero.
func (wg *WaitGroup) Wait()

// Go calls f in a new goroutine and adds that task to the WaitGroup.
// When f returns, the task is removed from the WaitGroup.
//
// If the WaitGroup is empty, Go must happen before a [WaitGroup.Wait].
// Typically, this simply means Go is called to start tasks before Wait is called.
// If the WaitGroup is not empty, Go may happen at any time.
// This means a goroutine started by Go may itself call Go.
// If a WaitGroup is reused to wait for several independent sets of tasks,
// new Go calls must happen after all previous Wait calls have returned.
//
// In the terminology of [the Go memory model](https://go.dev/ref/mem),
// the return from f "synchronizes before" the return of any Wait call that it unblocks.
func (wg *WaitGroup) Go(f func())
