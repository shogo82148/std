// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testtrace

import (
	"github.com/shogo82148/std/internal/trace"
)

// Validator is a type used for validating a stream of trace.Events.
type Validator struct {
	lastTs   trace.Time
	gs       map[trace.GoID]*goState
	ps       map[trace.ProcID]*procState
	ms       map[trace.ThreadID]*schedContext
	ranges   map[trace.ResourceID][]string
	tasks    map[trace.TaskID]string
	seenSync bool
	Go121    bool
}

// NewValidator creates a new Validator.
func NewValidator() *Validator

// Event validates ev as the next event in a stream of trace.Events.
//
// Returns an error if validation fails.
func (v *Validator) Event(ev trace.Event) error
