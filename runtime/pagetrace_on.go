// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.pagetrace

// Page tracer.
//
// This file contains an implementation of page trace instrumentation for tracking
// the way the Go runtime manages pages of memory. The trace may be enabled at program
// startup with the GODEBUG option pagetrace.
//
// Each page trace event is either 8 or 16 bytes wide. The first
// 8 bytes follow this format for non-sync events:
//
//     [16 timestamp delta][35 base address][10 npages][1 isLarge][2 pageTraceEventType]
//
// If the "large" bit is set then the event is 16 bytes wide with the second 8 byte word
// containing the full npages value (the npages bitfield is 0).
//
// The base address's bottom pageShift bits are always zero hence why we can pack other
// data in there. We ignore the top 16 bits, assuming a 48 bit address space for the
// heap.
//
// The timestamp delta is computed from the difference between the current nanotime
// timestamp and the last sync event's timestamp. The bottom pageTraceTimeLostBits of
// this delta is removed and only the next pageTraceTimeDeltaBits are kept.
//
// A sync event is emitted at the beginning of each trace buffer and whenever the
// timestamp delta would not fit in an event.
//
// Sync events have the following structure:
//
//    [61 timestamp or P ID][1 isPID][2 pageTraceSyncEvent]
//
// In essence, the "large" bit repurposed to indicate whether it's a timestamp or a P ID
// (these are typically uint32). Note that we only have 61 bits for the 64-bit timestamp,
// but like for the delta we drop the bottom pageTraceTimeLostBits here as well.

package runtime
