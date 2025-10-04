// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

// Atoi64 parses an int64 from a string s.
// The bool result reports whether s is a number
// representable by a value of type int64.
func Atoi64(s string) (int64, bool)

// Atoi is like Atoi64 but for integers
// that fit into an int.
func Atoi(s string) (int, bool)

// Atoi32 is like Atoi but for integers
// that fit into an int32.
func Atoi32(s string) (int32, bool)
