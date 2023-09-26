// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

// A patchList is a list of instruction pointers that need to be filled in (patched).
// Because the pointers haven't been filled in yet, we can reuse their storage
// to hold the list.  It's kind of sleazy, but works well in practice.
// See http://swtch.com/~rsc/regexp/regexp1.html for inspiration.
//
// These aren't really pointers: they're integers, so we can reinterpret them
// this way without using package unsafe.  A value l denotes
// p.inst[l>>1].Out (l&1==0) or .Arg (l&1==1).
// l == 0 denotes the empty list, okay because we start every program
// with a fail instruction, so we'll never want to point at its output link.

// A frag represents a compiled program fragment.

// Compile compiles the regexp into a program to be executed.
// The regexp should have been simplified already (returned from re.Simplify).
func Compile(re *Regexp) (*Prog, error)
