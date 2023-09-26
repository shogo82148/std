// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of
// goroutines to wait for.  Then each of the goroutines
// runs and calls Done when finished.  At the same time,
// Wait can be used to block until all goroutines have finished.
type WaitGroup struct {
	m       Mutex
	counter int32
	waiters int32
	sema    *uint32
}

// Add adds delta, which may be negative, to the WaitGroup counter.
// If the counter becomes zero, all goroutines blocked on Wait are released.
// If the counter goes negative, Add panics.
//
// Note that calls with positive delta must happen before the call to Wait,
// or else Wait may wait for too small a group. Typically this means the calls
// to Add should execute before the statement creating the goroutine or
// other event to be waited for. See the WaitGroup example.
func (wg *WaitGroup) Add(delta int)

// Done decrements the WaitGroup counter.
func (wg *WaitGroup) Done()

// Wait blocks until the WaitGroup counter is zero.
func (wg *WaitGroup) Wait()
