// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/time"
)

// An Attr is a key-value pair.
type Attr struct {
	Key   string
	Value Value
}

// String returns an Attr for a string value.
func String(key, value string) Attr

// Int64 returns an Attr for an int64.
func Int64(key string, value int64) Attr

// Int converts an int to an int64 and returns
// an Attr with that value.
func Int(key string, value int) Attr

// Uint64 returns an Attr for a uint64.
func Uint64(key string, v uint64) Attr

// Float64 returns an Attr for a floating-point number.
func Float64(key string, v float64) Attr

// Bool returns an Attr for a bool.
func Bool(key string, v bool) Attr

// Time returns an Attr for a time.Time.
// It discards the monotonic portion.
func Time(key string, v time.Time) Attr

// Duration returns an Attr for a time.Duration.
func Duration(key string, v time.Duration) Attr

// Group returns an Attr for a Group Value.
// The first argument is the key; the remaining arguments
// are converted to Attrs as in [Logger.Log].
//
// Use Group to collect several key-value pairs under a single
// key on a log line, or as the result of LogValue
// in order to log a single value as multiple Attrs.
func Group(key string, args ...any) Attr

// Any returns an Attr for the supplied value.
// See [AnyValue] for how values are treated.
func Any(key string, value any) Attr

// Equal reports whether a and b have equal keys and values.
func (a Attr) Equal(b Attr) bool

func (a Attr) String() string
