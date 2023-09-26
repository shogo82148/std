// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Global pool of spans that have free stacks.
// Stacks are assigned an order according to size.
//     order = log_2(size/FixedStack)
// There is a free list for each order.
// TODO: one lock per order?

// List of stack spans to be freed at the end of GC. Protected by
// stackpoolmu.

// Cached value of haveexperiment("framepointer")

// Information from the compiler about the layout of stack frames.
