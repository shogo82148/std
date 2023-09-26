// Code generated by mklockrank.go; DO NOT EDIT.

package runtime

// Constants representing the ranks of all non-leaf runtime locks, in rank order.
// Locks with lower rank must be taken before locks with higher rank,
// in addition to satisfying the partial order in lockPartialOrder.
// A few ranks allow self-cycles, which are specified in lockPartialOrder.

// lockRankLeafRank is the rank of lock that does not have a declared rank,
// and hence is a leaf lock.

// lockNames gives the names associated with each of the above ranks.

// lockPartialOrder is the transitive closure of the lock rank graph.
// An entry for rank X lists all of the ranks that can already be held
// when rank X is acquired.
//
// Lock ranks that allow self-cycles list themselves.
