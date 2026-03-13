// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

// NewRandomWriteScheduler constructs a WriteScheduler that ignores HTTP/2
// priorities. Control frames like SETTINGS and PING are written before DATA
// frames, but if no control frames are queued and multiple streams have queued
// HEADERS or DATA frames, Pop selects a ready stream arbitrarily.
func NewRandomWriteScheduler() WriteScheduler
