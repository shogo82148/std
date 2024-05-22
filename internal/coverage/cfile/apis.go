// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cfile

import (
	"github.com/shogo82148/std/io"
)

// WriteMetaDir implements [runtime/coverage.WriteMetaDir].
func WriteMetaDir(dir string) error

// WriteMeta implements [runtime/coverage.WriteMeta].
func WriteMeta(w io.Writer) error

// WriteCountersDir implements [runtime/coverage.WriteCountersDir].
func WriteCountersDir(dir string) error

// WriteCounters implements [runtime/coverage.WriteCounters].
func WriteCounters(w io.Writer) error

// ClearCounters implements [runtime/coverage.ClearCounters].
func ClearCounters() error
