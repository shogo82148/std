// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/time"
)

// StartSpan starts a trace event with the given name. The Span ends when its Done method is called.
func StartSpan(ctx context.Context, name string) (context.Context, *Span)

// StartGoroutine associates the context with a new Thread ID. The Chrome trace viewer associates each
// trace event with a thread, and doesn't expect events with the same thread id to happen at the
// same time.
func StartGoroutine(ctx context.Context) context.Context

// Flow marks a flow indicating that the 'to' span depends on the 'from' span.
// Flow should be called while the 'to' span is in progress.
func Flow(ctx context.Context, from *Span, to *Span)

type Span struct {
	t *tracer

	name  string
	tid   uint64
	start time.Time
	end   time.Time
}

func (s *Span) Done()

// Start starts a trace which writes to the given file.
func Start(ctx context.Context, file string) (context.Context, func() error, error)
