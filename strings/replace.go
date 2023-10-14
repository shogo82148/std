// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

// Replacer replaces a list of strings with replacements.
// It is safe for concurrent use by multiple goroutines.
type Replacer struct {
	once   sync.Once
	r      replacer
	oldnew []string
}

// NewReplacer returns a new [Replacer] from a list of old, new string
// pairs. Replacements are performed in the order they appear in the
// target string, without overlapping matches. The old string
// comparisons are done in argument order.
//
// NewReplacer panics if given an odd number of arguments.
func NewReplacer(oldnew ...string) *Replacer

// Replace returns a copy of s with all replacements performed.
func (r *Replacer) Replace(s string) string

// WriteString writes s to w with all replacements performed.
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
