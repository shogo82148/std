// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/time"
)

// A Record holds information about a log event.
// Copies of a Record share state.
// Do not modify a Record after handing out a copy to it.
// Call [NewRecord] to create a new Record.
// Use [Record.Clone] to create a copy with no shared state.
type Record struct {
	Time time.Time

	Message string

	Level Level

	PC uintptr

	front [nAttrsInline]Attr

	nFront int

	back []Attr
}

// NewRecord creates a Record from the given arguments.
// Use [Record.AddAttrs] to add attributes to the Record.
//
// NewRecord is intended for logging APIs that want to support a [Handler] as
// a backend.
func NewRecord(t time.Time, level Level, msg string, pc uintptr) Record

// Clone returns a copy of the record with no shared state.
// The original record and the clone can both be modified
// without interfering with each other.
func (r Record) Clone() Record

// NumAttrs returns the number of attributes in the Record.
func (r Record) NumAttrs() int

// Attrs calls f on each Attr in the Record.
// Iteration stops if f returns false.
func (r Record) Attrs(f func(Attr) bool)

// AddAttrs appends the given Attrs to the Record's list of Attrs.
// It omits empty groups.
func (r *Record) AddAttrs(attrs ...Attr)

// Add converts the args to Attrs as described in [Logger.Log],
// then appends the Attrs to the Record's list of Attrs.
// It omits empty groups.
func (r *Record) Add(args ...any)

// Source describes the location of a line of source code.
type Source struct {
	Function string `json:"function"`

	File string `json:"file"`
	Line int    `json:"line"`
}
