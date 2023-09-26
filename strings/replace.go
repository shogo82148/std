// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import "github.com/shogo82148/std/io"

// A Replacer replaces a list of strings with replacements.
type Replacer struct {
	r replacer
}

// replacer is the interface that a replacement algorithm needs to implement.

// byteBitmap represents bytes which are sought for replacement.
// byteBitmap is 256 bits wide, with a bit set for each old byte to be
// replaced.

// NewReplacer returns a new Replacer from a list of old, new string pairs.
// Replacements are performed in order, without overlapping matches.
func NewReplacer(oldnew ...string) *Replacer

// Replace returns a copy of s with all replacements performed.
func (r *Replacer) Replace(s string) string

// WriteString writes s to w with all replacements performed.
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)

// genericReplacer is the fully generic (and least optimized) algorithm.
// It's used as a fallback when nothing faster can be used.

// byteReplacer is the implementation that's used when all the "old"
// and "new" values are single ASCII bytes.

// byteStringReplacer is the implementation that's used when all the
// "old" values are single ASCII bytes but the "new" values vary in
// size.

// strings is too low-level to import io/ioutil
