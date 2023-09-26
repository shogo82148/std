// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

var OrigZoneSources = zoneSources

var Interrupt = interrupt
var DaysIn = daysIn

var (
	MinMonoTime = Time{wall: 1 << 63, ext: -1 << 63, loc: UTC}
	MaxMonoTime = Time{wall: 1 << 63, ext: 1<<63 - 1, loc: UTC}

	NotMonoNegativeTime = Time{wall: 0, ext: -1<<63 + 50}
)
