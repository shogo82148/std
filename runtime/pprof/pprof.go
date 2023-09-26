// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pprof writes runtime profiling data in the format expected
// by the pprof visualization tool.
// For more information about pprof, see
// http://code.google.com/p/google-perftools/.
package pprof

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

// A Profile is a collection of stack traces showing the call sequences
// that led to instances of a particular event, such as allocation.
// Packages can create and maintain their own profiles; the most common
// use is for tracking resources that must be explicitly closed, such as files
// or network connections.
//
// A Profile's methods can be called from multiple goroutines simultaneously.
//
// Each Profile has a unique name.  A few profiles are predefined:
//
//	goroutine    - stack traces of all current goroutines
//	heap         - a sampling of all heap allocations
//	threadcreate - stack traces that led to the creation of new OS threads
//	block        - stack traces that led to blocking on synchronization primitives
//
// These predefined profiles maintain themselves and panic on an explicit
// Add or Remove method call.
//
// The CPU profile is not available as a Profile.  It has a special API,
// the StartCPUProfile and StopCPUProfile functions, because it streams
// output to a writer during profiling.
type Profile struct {
	name  string
	mu    sync.Mutex
	m     map[interface{}][]uintptr
	count func() int
	write func(io.Writer, int) error
}

// profiles records all registered profiles.

// NewProfile creates a new profile with the given name.
// If a profile with that name already exists, NewProfile panics.
// The convention is to use a 'import/path.' prefix to create
// separate name spaces for each package.
func NewProfile(name string) *Profile

// Lookup returns the profile with the given name, or nil if no such profile exists.
func Lookup(name string) *Profile

// Profiles returns a slice of all the known profiles, sorted by name.
func Profiles() []*Profile

// Name returns this profile's name, which can be passed to Lookup to reobtain the profile.
func (p *Profile) Name() string

// Count returns the number of execution stacks currently in the profile.
func (p *Profile) Count() int

// Add adds the current execution stack to the profile, associated with value.
// Add stores value in an internal map, so value must be suitable for use as
// a map key and will not be garbage collected until the corresponding
// call to Remove.  Add panics if the profile already contains a stack for value.
//
// The skip parameter has the same meaning as runtime.Caller's skip
// and controls where the stack trace begins.  Passing skip=0 begins the
// trace in the function calling Add.  For example, given this
// execution stack:
//
//	Add
//	called from rpc.NewClient
//	called from mypkg.Run
//	called from main.main
//
// Passing skip=0 begins the stack trace at the call to Add inside rpc.NewClient.
// Passing skip=1 begins the stack trace at the call to NewClient inside mypkg.Run.
func (p *Profile) Add(value interface{}, skip int)

// Remove removes the execution stack associated with value from the profile.
// It is a no-op if the value is not in the profile.
func (p *Profile) Remove(value interface{})

// WriteTo writes a pprof-formatted snapshot of the profile to w.
// If a write to w returns an error, WriteTo returns that error.
// Otherwise, WriteTo returns nil.
//
// The debug parameter enables additional output.
// Passing debug=0 prints only the hexadecimal addresses that pprof needs.
// Passing debug=1 adds comments translating addresses to function names
// and line numbers, so that a programmer can read the profile without tools.
//
// The predefined profiles may assign meaning to other debug values;
// for example, when printing the "goroutine" profile, debug=2 means to
// print the goroutine stacks in the same form that a Go program uses
// when dying due to an unrecovered panic.
func (p *Profile) WriteTo(w io.Writer, debug int) error

// A countProfile is a set of stack traces to be printed as counts
// grouped by stack trace.  There are multiple implementations:
// all that matters is that we can find out how many traces there are
// and obtain each trace in turn.

// WriteHeapProfile is shorthand for Lookup("heap").WriteTo(w, 0).
// It is preserved for backwards compatibility.
func WriteHeapProfile(w io.Writer) error

// StartCPUProfile enables CPU profiling for the current process.
// While profiling, the profile will be buffered and written to w.
// StartCPUProfile returns an error if profiling is already enabled.
func StartCPUProfile(w io.Writer) error

// StopCPUProfile stops the current CPU profile, if any.
// StopCPUProfile only returns after all the writes for the
// profile have completed.
func StopCPUProfile()
