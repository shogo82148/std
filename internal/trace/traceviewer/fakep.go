// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package traceviewer

const (
	// Special P identifiers:
	FakeP    = 1000000 + iota
	TimerP
	NetpollP
	SyscallP
	GCP
	ProfileP
)
