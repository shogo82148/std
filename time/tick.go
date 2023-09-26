// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

// A Ticker holds a channel that delivers `ticks' of a clock
// at intervals.
type Ticker struct {
	C <-chan Time
	r runtimeTimer
}

// NewTicker returns a new Ticker containing a channel that will send the
// time with a period specified by the duration argument.
// It adjusts the intervals or drops ticks to make up for slow receivers.
// The duration d must be greater than zero; if not, NewTicker will panic.
func NewTicker(d Duration) *Ticker

// Stop turns off a ticker.  After Stop, no more ticks will be sent.
// Stop does not close the channel, to prevent a read from the channel succeeding
// incorrectly.
func (t *Ticker) Stop()

// Tick is a convenience wrapper for NewTicker providing access to the ticking
// channel only.  Useful for clients that have no need to shut down the ticker.
func Tick(d Duration) <-chan Time
