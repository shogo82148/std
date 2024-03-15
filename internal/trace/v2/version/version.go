// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version

import (
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/internal/trace/v2/event"
)

// Version represents the version of a trace file.
type Version uint32

const (
	Go111   Version = 11
	Go119   Version = 19
	Go121   Version = 21
	Go122   Version = 22
	Current         = Go122
)

// Specs returns the set of event.Specs for this version.
func (v Version) Specs() []event.Spec

func (v Version) Valid() bool

// ReadHeader reads the version of the trace out of the trace file's
// header, whose prefix must be present in v.
func ReadHeader(r io.Reader) (Version, error)

// WriteHeader writes a header for a trace version v to w.
func WriteHeader(w io.Writer, v Version) (int, error)
