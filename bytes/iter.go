// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes

import (
	"github.com/shogo82148/std/iter"
)

// Lines returns an iterator over the newline-terminated lines in the byte slice s.
// The lines yielded by the iterator include their terminating newlines.
// If s is empty, the iterator yields no lines at all.
// If s does not end in a newline, the final yielded line will not end in a newline.
// It returns a single-use iterator.
func Lines(s []byte) iter.Seq[[]byte]

// SplitSeq returns an iterator over all subslices of s separated by sep.
// The iterator yields the same subslices that would be returned by [Split](s, sep),
// but without constructing a new slice containing the subslices.
// It returns a single-use iterator.
func SplitSeq(s, sep []byte) iter.Seq[[]byte]

// SplitAfterSeq returns an iterator over subslices of s split after each instance of sep.
// The iterator yields the same subslices that would be returned by [SplitAfter](s, sep),
// but without constructing a new slice containing the subslices.
// It returns a single-use iterator.
func SplitAfterSeq(s, sep []byte) iter.Seq[[]byte]

// FieldsSeq returns an iterator over subslices of s split around runs of
// whitespace characters, as defined by [unicode.IsSpace].
// The iterator yields the same subslices that would be returned by [Fields](s),
// but without constructing a new slice containing the subslices.
func FieldsSeq(s []byte) iter.Seq[[]byte]

// FieldsFuncSeq returns an iterator over subslices of s split around runs of
// Unicode code points satisfying f(c).
// The iterator yields the same subslices that would be returned by [FieldsFunc](s),
// but without constructing a new slice containing the subslices.
func FieldsFuncSeq(s []byte, f func(rune) bool) iter.Seq[[]byte]
