// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates a priority queue built using the heap interface.
package heap_test

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string
	priority int

	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

// This example pushes 10 items into a PriorityQueue and takes them out in
// order of priority.
func Example() {
	const nItem = 10
	// Random priorities for the items (a permutation of 0..9, times 11)).
	priorities := [nItem]int{
		77, 22, 44, 55, 11, 88, 33, 99, 00, 66,
	}
	values := [nItem]string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	// Create a priority queue and put some items in it.
	pq := make(PriorityQueue, 0, nItem)
	for i := 0; i < cap(pq); i++ {
		item := &Item{
			value:    values[i],
			priority: priorities[i],
		}
		heap.Push(&pq, item)
	}
	// Take the items out; should arrive in decreasing priority order.
	// For example, the highest priority (99) is the seventh item, so output starts with 99:"seven".
	for i := 0; i < nItem; i++ {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
	// Output:
	// 99:seven 88:five 77:zero 66:nine 55:three 44:two 33:six 22:one 11:four 00:eight
}
