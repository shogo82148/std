// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// An objNode represents a node in the object dependency graph.
// Each node b in a.out represents an edge a->b indicating that
// b depends on a.
// Nodes may be marked for cycle detection. A node n is marked
// if n.mark corresponds to the current mark value.

// nodeQueue implements the container/heap interface;
// a nodeQueue may be used as a priority queue.
