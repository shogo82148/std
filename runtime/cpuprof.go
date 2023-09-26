// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// CPU profiling.
//
// The signal handler for the profiling clock tick adds a new stack trace
// to a log of recent traces. The log is read by a user goroutine that
// turns it into formatted profile data. If the reader does not keep up
// with the log, those writes will be recorded as a count of lost records.
// The actual profile buffer is in profbuf.go.

package runtime

// SetCPUProfileRate sets the CPU profiling rate to hz samples per second.
// If hz <= 0, SetCPUProfileRate turns off profiling.
// If the profiler is on, the rate cannot be changed without first turning it off.
//
// Most clients should use the runtime/pprof package or
// the testing package's -test.cpuprofile flag instead of calling
// SetCPUProfileRate directly.
func SetCPUProfileRate(hz int)

// CPUProfile panics.
// It formerly provided raw access to chunks of
// a pprof-format profile generated by the runtime.
// The details of generating that format have changed,
// so this functionality has been removed.
//
// Deprecated: Use the runtime/pprof package,
// or the handlers in the net/http/pprof package,
// or the testing package's -test.cpuprofile flag instead.
func CPUProfile() []byte
