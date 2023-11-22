// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Goroutine-related profiles.

package trace

import (
	"github.com/shogo82148/std/internal/trace"

	"github.com/shogo82148/std/net/http"
)

// GoroutinesHandlerFunc returns a HandlerFunc that serves list of goroutine groups.
func GoroutinesHandlerFunc(summaries map[tracev2.GoID]*trace.GoroutineSummary) http.HandlerFunc

// GoroutineHandler creates a handler that serves information about
// goroutines in a particular group.
func GoroutineHandler(summaries map[tracev2.GoID]*trace.GoroutineSummary) http.HandlerFunc
