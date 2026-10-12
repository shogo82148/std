// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// If true, check poset integrity after every mutation
var DebugPoset = false

// Poset is a union-find data structure that can represent a partially ordered set
// of SSA values. Given a binary relation that creates a partial order (eg: '<'),
// clients can record relations between SSA values using SetOrder, and later
// check relations (in the transitive closure) with Ordered. For instance,
// if SetOrder is called to record that A<B and B<C, Ordered will later confirm
// that A<C.
//
// It is possible to record equality relations between SSA values with SetEqual and check
// equality with Equal. Equality propagates into the transitive closure for the partial
// order so that if we know that A<B<C and later learn that A==D, Ordered will return
// true for D<C.
//
// It is also possible to record inequality relations between nodes with SetNonEqual;
// non-equality relations are not transitive, but they can still be useful: for instance
// if we know that A<=B and later we learn that A!=B, we can deduce that A<B.
// NonEqual can be used to check whether it is known that the nodes are different, either
// because SetNonEqual was called before, or because we know that they are strictly ordered.
//
// Poset will refuse to record new relations that contradict existing relations:
// for instance if A<B<C, calling SetOrder for C<A will fail returning false; also
// calling SetEqual for C==A will fail.
//
// Poset is implemented as a forest of DAGs; in each DAG, if there is a path (directed)
// from node A to B, it means that A<B (or A<=B). Equality is represented by mapping
// two SSA values to the same DAG node; when a new equality relation is recorded
// between two existing nodes, the nodes are merged, adjusting incoming and outgoing edges.
//
// Poset is designed to be memory efficient and do little allocations during normal usage.
// Most internal data structures are pre-allocated and flat, so for instance adding a
// new relation does not cause any allocation. For performance reasons,
// each node has only up to two outgoing edges (like a binary tree), so intermediate
// "extra" nodes are required to represent more than two relations. For instance,
// to record that A<I, A<J, A<K (with no known relation between I,J,K), we create the
// following DAG:
//
//	  A
//	 / \
//	I  extra
//	    /  \
//	   J    K
type Poset struct {
	lastidx uint32
	values  map[ID]uint32
	nodes   []posetNode
	roots   []uint32
	noneq   map[uint32]bitset
	undo    []posetUndo
}

// CheckIntegrity verifies internal integrity of a poset. It is intended
// for debugging purposes.
func (po *Poset) CheckIntegrity()

// CheckEmpty checks that a poset is completely empty.
// It can be used for debugging purposes, as a poset is supposed to
// be empty after it's fully rolled back through Undo.
func (po *Poset) CheckEmpty() error

// DotDump dumps the poset in graphviz format to file fn, with the specified title.
func (po *Poset) DotDump(fn string, title string) error

// Ordered reports whether n1<n2. It returns false either when it is
// certain that n1<n2 is false, or if there is not enough information
// to tell.
// Complexity is O(n).
func (po *Poset) Ordered(n1, n2 *Value) bool

// OrderedOrEqual reports whether n1<=n2. It returns false either when it is
// certain that n1<=n2 is false, or if there is not enough information
// to tell.
// Complexity is O(n).
func (po *Poset) OrderedOrEqual(n1, n2 *Value) bool

// Equal reports whether n1==n2. It returns false either when it is
// certain that n1==n2 is false, or if there is not enough information
// to tell.
// Complexity is O(1).
func (po *Poset) Equal(n1, n2 *Value) bool

// NonEqual reports whether n1!=n2. It returns false either when it is
// certain that n1!=n2 is false, or if there is not enough information
// to tell.
// Complexity is O(n) (because it internally calls Ordered to see if we
// can infer n1!=n2 from n1<n2 or n2<n1).
func (po *Poset) NonEqual(n1, n2 *Value) bool

// SetOrder records that n1<n2. Returns false if this is a contradiction
// Complexity is O(1) if n2 was never seen before, or O(n) otherwise.
func (po *Poset) SetOrder(n1, n2 *Value) bool

// SetOrderOrEqual records that n1<=n2. Returns false if this is a contradiction
// Complexity is O(1) if n2 was never seen before, or O(n) otherwise.
func (po *Poset) SetOrderOrEqual(n1, n2 *Value) bool

// SetEqual records that n1==n2. Returns false if this is a contradiction
// (that is, if it is already recorded that n1<n2 or n2<n1).
// Complexity is O(1) if n2 was never seen before, or O(n) otherwise.
func (po *Poset) SetEqual(n1, n2 *Value) bool

// SetNonEqual records that n1!=n2. Returns false if this is a contradiction
// (that is, if it is already recorded that n1==n2).
// Complexity is O(n).
func (po *Poset) SetNonEqual(n1, n2 *Value) bool

// Checkpoint saves the current state of the DAG so that it's possible
// to later undo this state.
// Complexity is O(1).
func (po *Poset) Checkpoint()

// Undo restores the state of the poset to the previous checkpoint.
// Complexity depends on the type of operations that were performed
// since the last checkpoint; each Set* operation creates an undo
// pass which Undo has to revert with a worst-case complexity of O(n).
func (po *Poset) Undo()
