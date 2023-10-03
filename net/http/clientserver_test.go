// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Tests that use both the client & server, in both HTTP/1 and HTTP/2 mode.

package http_test

import (
	"testing"
)

type TBRun[T any] interface {
	testing.TB
	Run(string, func(T)) bool
}

// h12Compare is a test that compares HTTP/1 and HTTP/2 behavior
// against each other.
