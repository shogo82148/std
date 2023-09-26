// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Trace is a tool for viewing trace files.

Trace files can be generated with:
  - runtime/trace.Start
  - net/http/pprof package
  - go test -trace

Example usage:
Generate a trace file with 'go test':

	go test -trace trace.out pkg

View the trace in a web browser:

	go tool trace pkg.test trace.out
*/
package main
