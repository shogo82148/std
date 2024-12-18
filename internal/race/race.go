// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race

package race

import (
	"github.com/shogo82148/std/internal/abi"
	"github.com/shogo82148/std/unsafe"
)

const Enabled = true

//go:linkname Acquire
func Acquire(addr unsafe.Pointer)

//go:linkname Release
func Release(addr unsafe.Pointer)

//go:linkname ReleaseMerge
func ReleaseMerge(addr unsafe.Pointer)

//go:linkname Disable
func Disable()

//go:linkname Enable
func Enable()

//go:linkname Read
func Read(addr unsafe.Pointer)

//go:linkname ReadPC
func ReadPC(addr unsafe.Pointer, callerpc, pc uintptr)

//go:linkname ReadObjectPC
func ReadObjectPC(t *abi.Type, addr unsafe.Pointer, callerpc, pc uintptr)

//go:linkname Write
func Write(addr unsafe.Pointer)

//go:linkname WritePC
func WritePC(addr unsafe.Pointer, callerpc, pc uintptr)

//go:linkname WriteObjectPC
func WriteObjectPC(t *abi.Type, addr unsafe.Pointer, callerpc, pc uintptr)

//go:linkname ReadRange
func ReadRange(addr unsafe.Pointer, len int)

//go:linkname WriteRange
func WriteRange(addr unsafe.Pointer, len int)

//go:linkname Errors
func Errors() int
