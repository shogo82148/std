// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

// NOTE(rsc): CONTEXT_CONTROL is actually 0x400001 and should include PC, SP, and LR.
// However, empirically, LR doesn't come along on Windows 10
// unless you also set CONTEXT_INTEGER (0x400002).
// Without LR, we skip over the next-to-bottom function in profiles
// when the bottom function is frameless.
// So we set both here, to make a working CONTEXT_CONTROL.
const CONTEXT_CONTROL = 0x400003

type Neon128 struct {
	Low  uint64
	High int64
}

// See https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-arm64_nt_context
type Context struct {
	ContextFlags uint32
	Cpsr         uint32
	X            [31]uint64
	XSp          uint64
	Pc           uint64
	V            [32]Neon128
	Fpcr         uint32
	Fpsr         uint32
	Bcr          [8]uint32
	Bvr          [8]uint64
	Wcr          [2]uint32
	Wvr          [2]uint64
}

func (c *Context) PC() uintptr
func (c *Context) SP() uintptr
func (c *Context) LR() uintptr

func (c *Context) SetPC(x uintptr)
func (c *Context) SetSP(x uintptr)
func (c *Context) SetLR(x uintptr)
func (c *Context) SetFP(x uintptr)

func (c *Context) PushCall(targetPC, resumePC uintptr)

type DISPATCHER_CONTEXT struct {
	ControlPc        uint64
	ImageBase        uint64
	FunctionEntry    uintptr
	EstablisherFrame uint64
	TargetIp         uint64
	Context          *Context
	LanguageHandler  uintptr
	HandlerData      uintptr
}

func (c *DISPATCHER_CONTEXT) Ctx() *Context
