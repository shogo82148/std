// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race
// +build race

// Public race detection API, present iff build with -race.

package runtime

import (
	"github.com/shogo82148/std/unsafe"
)

func RaceRead(addr unsafe.Pointer)
func RaceWrite(addr unsafe.Pointer)
func RaceReadRange(addr unsafe.Pointer, len int)
func RaceWriteRange(addr unsafe.Pointer, len int)

func RaceErrors() int

// private interface for the runtime

// Race runtime functions called via runtime·racecall.
//go:linkname __tsan_init __tsan_init

//go:linkname __tsan_fini __tsan_fini

//go:linkname __tsan_proc_create __tsan_proc_create

//go:linkname __tsan_proc_destroy __tsan_proc_destroy

//go:linkname __tsan_map_shadow __tsan_map_shadow

//go:linkname __tsan_finalizer_goroutine __tsan_finalizer_goroutine

//go:linkname __tsan_go_start __tsan_go_start

//go:linkname __tsan_go_end __tsan_go_end

//go:linkname __tsan_malloc __tsan_malloc

//go:linkname __tsan_free __tsan_free

//go:linkname __tsan_acquire __tsan_acquire

//go:linkname __tsan_release __tsan_release

//go:linkname __tsan_release_merge __tsan_release_merge

//go:linkname __tsan_go_ignore_sync_begin __tsan_go_ignore_sync_begin

//go:linkname __tsan_go_ignore_sync_end __tsan_go_ignore_sync_end

//go:linkname __tsan_report_count __tsan_report_count

// start/end of global data (data+bss).

// start/end of heap for race_amd64.s

func RaceAcquire(addr unsafe.Pointer)

func RaceRelease(addr unsafe.Pointer)

func RaceReleaseMerge(addr unsafe.Pointer)

// RaceDisable disables handling of race events in the current goroutine.
func RaceDisable()

// RaceEnable re-enables handling of race events in the current goroutine.
func RaceEnable()
