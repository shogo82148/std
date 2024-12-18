// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race

package runtime

import (
	"github.com/shogo82148/std/unsafe"
)

func RaceRead(addr unsafe.Pointer)

func RaceWrite(addr unsafe.Pointer)

func RaceReadRange(addr unsafe.Pointer, len int)

func RaceWriteRange(addr unsafe.Pointer, len int)

func RaceErrors() int

// RaceAcquire/RaceRelease/RaceReleaseMerge establish happens-before relations
// between goroutines. These inform the race detector about actual synchronization
// that it can't see for some reason (e.g. synchronization within RaceDisable/RaceEnable
// sections of code).
// RaceAcquire establishes a happens-before relation with the preceding
// RaceReleaseMerge on addr up to and including the last RaceRelease on addr.
// In terms of the C memory model (C11 §5.1.2.4, §7.17.3),
// RaceAcquire is equivalent to atomic_load(memory_order_acquire).
//
//go:nosplit
func RaceAcquire(addr unsafe.Pointer)

// RaceRelease performs a release operation on addr that
// can synchronize with a later RaceAcquire on addr.
//
// In terms of the C memory model, RaceRelease is equivalent to
// atomic_store(memory_order_release).
//
//go:nosplit
func RaceRelease(addr unsafe.Pointer)

// RaceReleaseMerge is like RaceRelease, but also establishes a happens-before
// relation with the preceding RaceRelease or RaceReleaseMerge on addr.
//
// In terms of the C memory model, RaceReleaseMerge is equivalent to
// atomic_exchange(memory_order_release).
//
//go:nosplit
func RaceReleaseMerge(addr unsafe.Pointer)

// RaceDisable disables handling of race synchronization events in the current goroutine.
// Handling is re-enabled with RaceEnable. RaceDisable/RaceEnable can be nested.
// Non-synchronization events (memory accesses, function entry/exit) still affect
// the race detector.
//
//go:nosplit
func RaceDisable()

// RaceEnable re-enables handling of race events in the current goroutine.
//
//go:nosplit
func RaceEnable()
