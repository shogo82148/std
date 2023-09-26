// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Benchmark for accessing Value values.

package slog

import (
	"time"
)

// Problem: adding a type means adding a method, which is a breaking change.
// Using an unexported method to force embedding will make programs compile,
// But they will panic at runtime when we call the new method.
type Visitor interface {
	String(string)
	Int64(int64)
	Uint64(uint64)
	Float64(float64)
	Bool(bool)
	Duration(time.Duration)
	Any(any)
}
