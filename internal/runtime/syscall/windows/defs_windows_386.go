// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

const CONTEXT_CONTROL = 0x10001

type FloatingSaveArea struct {
	ControlWord   uint32
	StatusWord    uint32
	TagWord       uint32
	ErrorOffset   uint32
	ErrorSelector uint32
	DataOffset    uint32
	DataSelector  uint32
	RegisterArea  [80]uint8
	Cr0NpxState   uint32
}

type Context struct {
	ContextFlags      uint32
	Dr0               uint32
	Dr1               uint32
	Dr2               uint32
	Dr3               uint32
	Dr6               uint32
	Dr7               uint32
	FloatingSaveArea  FloatingSaveArea
	SegGs             uint32
	SegFs             uint32
	SegEs             uint32
	SegDs             uint32
	Edi               uint32
	Esi               uint32
	Ebx               uint32
	Edx               uint32
	Ecx               uint32
	Eax               uint32
	Ebp               uint32
	Eip               uint32
	SegCs             uint32
	EFlags            uint32
	Esp               uint32
	SegSs             uint32
	ExtendedRegisters [512]uint8
}

func (c *Context) PC() uintptr
func (c *Context) SP() uintptr

// 386 does not have link register, so this returns 0.
func (c *Context) LR() uintptr
func (c *Context) SetLR(x uintptr)

func (c *Context) SetPC(x uintptr)
func (c *Context) SetSP(x uintptr)

// 386 does not have frame pointer register.
func (c *Context) SetFP(x uintptr)

func (c *Context) PushCall(targetPC, resumePC uintptr)

// DISPATCHER_CONTEXT is not defined on 386.
type DISPATCHER_CONTEXT struct{}

func (c *DISPATCHER_CONTEXT) Ctx() *Context
