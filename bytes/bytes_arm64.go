// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
func Index(s, sep []byte) int

// Count counts the number of non-overlapping instances of sep in s.
// If sep is an empty slice, Count returns 1 + the number of UTF-8-encoded code points in s.
func Count(s, sep []byte) int
