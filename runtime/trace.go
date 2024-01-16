// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !goexperiment.exectracer2

// Go execution tracer.
// The tracer captures a wide range of execution events like goroutine
// creation/blocking/unblocking, syscall enter/exit/block, GC-related events,
// changes of heap size, processor start/stop, etc and writes them to a buffer
// in a compact form. A precise nanosecond-precision timestamp and a stack
// trace is captured for most events.
// See https://golang.org/s/go15trace for more info.

package runtime

// StartTrace enables tracing for the current process.
// While tracing, the data will be buffered and available via [ReadTrace].
// StartTrace returns an error if tracing is already enabled.
// Most clients should use the [runtime/trace] package or the [testing] package's
// -test.trace flag instead of calling StartTrace directly.
func StartTrace() error

// StopTrace stops tracing, if it was previously enabled.
// StopTrace only returns after all the reads for the trace have completed.
func StopTrace()

// ReadTrace returns the next chunk of binary tracing data, blocking until data
// is available. If tracing is turned off and all the data accumulated while it
// was on has been returned, ReadTrace returns nil. The caller must copy the
// returned data before calling ReadTrace again.
// ReadTrace must be called from one goroutine at a time.
func ReadTrace() []byte
