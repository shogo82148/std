// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file provides an internal debug logging facility. The debug
// log is a lightweight, in-memory, per-M ring buffer. By default, the
// runtime prints the debug log on panic.
//
// To print something to the debug log, call dlog to obtain a dlogger
// and use the methods on that to add values. The values will be
// space-separated in the output (much like println).
//
// This facility can be enabled by passing -tags debuglog when
// building. Without this tag, dlog calls compile to nothing.

package runtime

// debugLogBytes is the size of each per-M ring buffer. This is
// allocated off-heap to avoid blowing up the M and hence the GC'd
// heap size.

// debugLogStringLimit is the maximum number of bytes in a string.
// Above this, the string will be truncated with "..(n more bytes).."

// A dlogger writes to the debug log.
//
// To obtain a dlogger, call dlog(). When done with the dlogger, call
// end().

// allDloggers is a list of all dloggers, linked through
// dlogger.allLink. This is accessed atomically. This is prepend only,
// so it doesn't need to protect against ABA races.

// A debugLogWriter is a ring buffer of binary debug log records.
//
// A log record consists of a 2-byte framing header and a sequence of
// fields. The framing header gives the size of the record as a little
// endian 16-bit value. Each field starts with a byte indicating its
// type, followed by type-specific data. If the size in the framing
// header is 0, it's a sync record consisting of two little endian
// 64-bit values giving a new time base.
//
// Because this is a ring buffer, new records will eventually
// overwrite old records. Hence, it maintains a reader that consumes
// the log as it gets overwritten. That reader state is where an
// actual log reader would start.
