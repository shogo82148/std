// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package iotest implements Readers and Writers useful mainly for testing.
package iotest

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// OneByteReader returns a Reader that implements
// each non-empty Read by reading one byte from r.
func OneByteReader(r io.Reader) io.Reader

// HalfReader returns a Reader that implements Read
// by reading half as many requested bytes from r.
func HalfReader(r io.Reader) io.Reader

// DataErrReader returns a Reader that returns the final
// error with the last data read, instead of by itself with
// zero bytes of data.
func DataErrReader(r io.Reader) io.Reader

var ErrTimeout = errors.New("timeout")

// TimeoutReader returns ErrTimeout on the second read
// with no data.  Subsequent calls to read succeed.
func TimeoutReader(r io.Reader) io.Reader
