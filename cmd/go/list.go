// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/bufio"
)

// TrackingWriter tracks the last byte written on every write so
// we can avoid printing a newline if one was already written or
// if there is no output at all.
type TrackingWriter struct {
	w    *bufio.Writer
	last byte
}

func (t *TrackingWriter) Write(p []byte) (n int, err error)

func (t *TrackingWriter) Flush()

func (t *TrackingWriter) NeedNL() bool
