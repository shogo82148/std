// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package regexp

// A queue is a 'sparse array' holding pending threads of execution.
// See http://research.swtch.com/2008/03/using-uninitialized-memory-for-fun-and.html

// A entry is an entry on a queue.
// It holds both the instruction pc and the actual thread.
// Some queue entries are just place holders so that the machine
// knows it has considered that pc.  Such entries have t == nil.

// A thread is the state of a single path through the machine:
// an instruction and a corresponding capture array.
// See http://swtch.com/~rsc/regexp/regexp2.html

// A machine holds all the state during an NFA simulation for p.

// empty is a non-nil 0-element slice,
// so doExecute can avoid an allocation
// when 0 captures are requested from a successful match.
