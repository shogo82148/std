// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package signal implements access to incoming signals.
package signal

import (
	"github.com/shogo82148/std/os"
)

// Notify causes package signal to relay incoming signals to c.
// If no signals are listed, all incoming signals will be relayed to c.
// Otherwise, just the listed signals will.
//
// Package signal will not block sending to c: the caller must ensure
// that c has sufficient buffer space to keep up with the expected
// signal rate.  For a channel used for notification of just one signal value,
// a buffer of size 1 is sufficient.
func Notify(c chan<- os.Signal, sig ...os.Signal)
