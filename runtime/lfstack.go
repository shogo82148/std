// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Lock-free stack.
// Initialize head to 0, compare with 0 to test for emptiness.
// The stack does not keep pointers to nodes,
// so they can be garbage collected if there are no other pointers to nodes.
// The following code runs only in non-preemptible contexts.

package runtime
