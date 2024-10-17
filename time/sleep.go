// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

// Sleep pauses the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func Sleep(d Duration)

// The Timer type represents a single event.
// When the Timer expires, the current time will be sent on C,
// unless the Timer was created by [AfterFunc].
// A Timer must be created with [NewTimer] or AfterFunc.
type Timer struct {
	C         <-chan Time
	initTimer bool
}

// Stop prevents the [Timer] from firing.
// It returns true if the call stops the timer, false if the timer has already
// expired or been stopped.
//
// For a func-based timer created with [AfterFunc](d, f),
// if t.Stop returns false, then the timer has already expired
// and the function f has been started in its own goroutine;
// Stop does not wait for f to complete before returning.
// If the caller needs to know whether f is completed,
// it must coordinate with f explicitly.
//
// For a chan-based timer created with NewTimer(d), as of Go 1.23,
// any receive from t.C after Stop has returned is guaranteed to block
// rather than receive a stale time value from before the Stop;
// if the program has not received from t.C already and the timer is
// running, Stop is guaranteed to return true.
// Before Go 1.23, the only safe way to use Stop was insert an extra
// <-t.C if Stop returned false to drain a potential stale value.
// See the [NewTimer] documentation for more details.
func (t *Timer) Stop() bool

// NewTimer creates a new Timer that will send
// the current time on its channel after at least duration d.
//
// Before Go 1.23, the garbage collector did not recover
// timers that had not yet expired or been stopped, so code often
// immediately deferred t.Stop after calling NewTimer, to make
// the timer recoverable when it was no longer needed.
// As of Go 1.23, the garbage collector can recover unreferenced
// timers, even if they haven't expired or been stopped.
// The Stop method is no longer necessary to help the garbage collector.
// (Code may of course still want to call Stop to stop the timer for other reasons.)
//
// Before Go 1.23, the channel associated with a Timer was
// asynchronous (buffered, capacity 1), which meant that
// stale time values could be received even after [Timer.Stop]
// or [Timer.Reset] returned.
// As of Go 1.23, the channel is synchronous (unbuffered, capacity 0),
// eliminating the possibility of those stale values.
//
// The GODEBUG setting asynctimerchan=1 restores both pre-Go 1.23
// behaviors: when set, unexpired timers won't be garbage collected, and
// channels will have buffered capacity. This setting may be removed
// in Go 1.27 or later.
func NewTimer(d Duration) *Timer

// Reset changes the timer to expire after duration d.
// It returns true if the timer had been active, false if the timer had
// expired or been stopped.
//
// For a func-based timer created with [AfterFunc](d, f), Reset either reschedules
// when f will run, in which case Reset returns true, or schedules f
// to run again, in which case it returns false.
// When Reset returns false, Reset neither waits for the prior f to
// complete before returning nor does it guarantee that the subsequent
// goroutine running f does not run concurrently with the prior
// one. If the caller needs to know whether the prior execution of
// f is completed, it must coordinate with f explicitly.
//
// For a chan-based timer created with NewTimer, as of Go 1.23,
// any receive from t.C after Reset has returned is guaranteed not
// to receive a time value corresponding to the previous timer settings;
// if the program has not received from t.C already and the timer is
// running, Reset is guaranteed to return true.
// Before Go 1.23, the only safe way to use Reset was to call [Timer.Stop]
// and explicitly drain the timer first.
// See the [NewTimer] documentation for more details.
func (t *Timer) Reset(d Duration) bool

// After waits for the duration to elapse and then sends the current time
// on the returned channel.
// It is equivalent to [NewTimer](d).C.
//
// Before Go 1.23, this documentation warned that the underlying
// [Timer] would not be recovered by the garbage collector until the
// timer fired, and that if efficiency was a concern, code should use
// NewTimer instead and call [Timer.Stop] if the timer is no longer needed.
// As of Go 1.23, the garbage collector can recover unreferenced,
// unstopped timers. There is no reason to prefer NewTimer when After will do.
func After(d Duration) <-chan Time

// AfterFunc waits for the duration to elapse and then calls f
// in its own goroutine. It returns a [Timer] that can
// be used to cancel the call using its Stop method.
// The returned Timer's C field is not used and will be nil.
func AfterFunc(d Duration, f func()) *Timer
