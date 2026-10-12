// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package par

// Queue manages a set of work items to be executed in parallel. The number of
// active work items is limited, and excess items are queued sequentially.
type Queue struct {
	maxActive int
	st        chan queueState
}

// NewQueue returns a Queue that executes up to maxActive items in parallel.
//
// maxActive must be positive.
func NewQueue(maxActive int) *Queue

// Add adds f as a work item in the queue.
//
// Add returns immediately, but the queue will be marked as non-idle until after
// f (and any subsequently-added work) has completed.
func (q *Queue) Add(f func())

// Idle returns a channel that will be closed when q has no (active or enqueued)
// work outstanding.
func (q *Queue) Idle() <-chan struct{}
