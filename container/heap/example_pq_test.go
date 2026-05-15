// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates a priority queue built using the heap interface.
package heap_test

import (
	"github.com/shogo82148/std/container/heap"
	"github.com/shogo82148/std/fmt"
)

// この例では、いくつかの要素を持つPriorityQueueを作成し、要素を追加して操作し、
// その後、優先度順に要素を取り出します。
func Example_priorityQueue() {
	// いくつかの要素とその優先度。
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// 優先度付きキューを作成し、要素を入れて、
	// 優先度付きキュー（ヒープ）の不変条件を確立します。
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// 新しい要素を挿入してから、その優先度を変更します。
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// 要素を取り出します。優先度の高い順に取り出されます。
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple
}
