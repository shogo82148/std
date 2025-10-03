// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pprof writes runtime profiling data in the format expected
// by the pprof visualization tool.
//
// # Profiling a Go program
//
// The first step to profiling a Go program is to enable profiling.
// Support for profiling benchmarks built with the standard testing
// package is built into go test. For example, the following command
// runs benchmarks in the current directory and writes the CPU and
// memory profiles to cpu.prof and mem.prof:
//
//	go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
//
// To add equivalent profiling support to a standalone program, add
// code like the following to your main function:
//
//	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
//	var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
//
//	func main() {
//	    flag.Parse()
//	    if *cpuprofile != "" {
//	        f, err := os.Create(*cpuprofile)
//	        if err != nil {
//	            log.Fatal("could not create CPU profile: ", err)
//	        }
//	        defer f.Close() // error handling omitted for example
//	        if err := pprof.StartCPUProfile(f); err != nil {
//	            log.Fatal("could not start CPU profile: ", err)
//	        }
//	        defer pprof.StopCPUProfile()
//	    }
//
//	    // ... rest of the program ...
//
//	    if *memprofile != "" {
//	        f, err := os.Create(*memprofile)
//	        if err != nil {
//	            log.Fatal("could not create memory profile: ", err)
//	        }
//	        defer f.Close() // error handling omitted for example
//	        runtime.GC() // get up-to-date statistics
//	        // Lookup("allocs") creates a profile similar to go test -memprofile.
//	        // Alternatively, use Lookup("heap") for a profile
//	        // that has inuse_space as the default index.
//	        if err := pprof.Lookup("allocs").WriteTo(f, 0); err != nil {
//	            log.Fatal("could not write memory profile: ", err)
//	        }
//	    }
//	}
//
// There is also a standard HTTP interface to profiling data. Adding
// the following line will install handlers under the /debug/pprof/
// URL to download live profiles:
//
//	import _ "net/http/pprof"
//
// See the net/http/pprof package for more details.
//
// Profiles can then be visualized with the pprof tool:
//
//	go tool pprof cpu.prof
//
// There are many commands available from the pprof command line.
// Commonly used commands include "top", which prints a summary of the
// top program hot-spots, and "web", which opens an interactive graph
// of hot-spots and their call graphs. Use "help" for information on
// all pprof commands.
//
// For more information about pprof, see
// https://github.com/google/pprof/blob/main/doc/README.md.
package pprof

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

type Profile struct {
	name  string
	mu    sync.Mutex
	m     map[any][]uintptr
	count func() int
	write func(io.Writer, int) error
}

// NewProfile creates a new profile with the given name.
// If a profile with that name already exists, NewProfile panics.
// The convention is to use a 'import/path.' prefix to create
// separate name spaces for each package.
// For compatibility with various tools that read pprof data,
// profile names should not contain spaces.
func NewProfile(name string) *Profile

// Lookup returns the profile with the given name, or nil if no such profile exists.
func Lookup(name string) *Profile

// Profiles returns a slice of all the known profiles, sorted by name.
func Profiles() []*Profile

// Name returns this profile's name, which can be passed to [Lookup] to reobtain the profile.
func (p *Profile) Name() string

// Count returns the number of execution stacks currently in the profile.
func (p *Profile) Count() int

// Add adds the current execution stack to the profile, associated with value.
// Add stores value in an internal map, so value must be suitable for use as
// a map key and will not be garbage collected until the corresponding
// call to [Profile.Remove]. Add panics if the profile already contains a stack for value.
//
// The skip parameter has the same meaning as [runtime.Caller]'s skip
// and controls where the stack trace begins. Passing skip=0 begins the
// trace in the function calling Add. For example, given this
// execution stack:
//
//	Add
//	called from rpc.NewClient
//	called from mypkg.Run
//	called from main.main
//
// Passing skip=0 begins the stack trace at the call to Add inside rpc.NewClient.
// Passing skip=1 begins the stack trace at the call to NewClient inside mypkg.Run.
func (p *Profile) Add(value any, skip int)

// Remove removes the execution stack associated with value from the profile.
// It is a no-op if the value is not in the profile.
func (p *Profile) Remove(value any)

// WriteTo writes a pprof-formatted snapshot of the profile to w.
// If a write to w returns an error, WriteTo returns that error.
// Otherwise, WriteTo returns nil.
//
// The debug parameter enables additional output.
// Passing debug=0 writes the gzip-compressed protocol buffer described
// in https://github.com/google/pprof/tree/main/proto#overview.
// Passing debug=1 writes the legacy text format with comments
// translating addresses to function names and line numbers, so that a
// programmer can read the profile without tools.
//
// The predefined profiles may assign meaning to other debug values;
// for example, when printing the "goroutine" profile, debug=2 means to
// print the goroutine stacks in the same form that a Go program uses
// when dying due to an unrecovered panic.
func (p *Profile) WriteTo(w io.Writer, debug int) error

// WriteHeapProfile is shorthand for [Lookup]("heap").WriteTo(w, 0).
// It is preserved for backwards compatibility.
func WriteHeapProfile(w io.Writer) error

// StartCPUProfile enables CPU profiling for the current process.
// While profiling, the profile will be buffered and written to w.
// StartCPUProfile returns an error if profiling is already enabled.
//
// On Unix-like systems, StartCPUProfile does not work by default for
// Go code built with -buildmode=c-archive or -buildmode=c-shared.
// StartCPUProfile relies on the SIGPROF signal, but that signal will
// be delivered to the main program's SIGPROF signal handler (if any)
// not to the one used by Go. To make it work, call [os/signal.Notify]
// for [syscall.SIGPROF], but note that doing so may break any profiling
// being done by the main program.
func StartCPUProfile(w io.Writer) error

// StopCPUProfile stops the current CPU profile, if any.
// StopCPUProfile only returns after all the writes for the
// profile have completed.
func StopCPUProfile()
