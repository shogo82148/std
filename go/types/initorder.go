// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types
<<<<<<< HEAD

// A dependency is an object that may be a dependency in an initialization
// expression. Only constants, variables, and functions can be dependencies.
// Constants are here because constant expression cycles are reported during
// initialization order computation.

// A graphNode represents a node in the object dependency graph.
// Each node p in n.pred represents an edge p->n, and each node
// s in n.succ represents an edge n->s; with a->b indicating that
// a depends on b.

// nodeQueue implements the container/heap interface;
// a nodeQueue may be used as a priority queue.
=======
>>>>>>> upstream/master
