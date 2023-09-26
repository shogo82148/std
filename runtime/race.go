// Copyright 2012 The Go Authors.  All rights reserved.
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

func RaceSemacquire(s *uint32)
func RaceSemrelease(s *uint32)

// private interface for the runtime
