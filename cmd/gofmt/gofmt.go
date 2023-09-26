// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Keep these in sync with go/format/format.go.

// fdSem guards the number of concurrently-open file descriptors.
//
// For now, this is arbitrarily set to 200, based on the observation that many
// platforms default to a kernel limit of 256. Ideally, perhaps we should derive
// it from rlimit on platforms that support that system call.
//
// File descriptors opened from outside of this package are not tracked,
// so this limit may be approximate.

// A sequencer performs concurrent tasks that may write output, but emits that
// output in a deterministic order.

// exclusive is a weight that can be passed to a sequencer to cause
// a task to be executed without any other concurrent tasks.

// A reporter reports output, warnings, and errors.

// reporterState carries the state of a reporter instance.
//
// Only one reporter at a time may have access to a reporterState.
