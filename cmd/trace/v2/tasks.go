// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/net/http"
)

// UserTasksHandlerFunc returns a HandlerFunc that reports all tasks found in the trace.
func UserTasksHandlerFunc(t *parsedTrace) http.HandlerFunc

// UserTaskHandlerFunc returns a HandlerFunc that presents the details of the selected tasks.
func UserTaskHandlerFunc(t *parsedTrace) http.HandlerFunc
