// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import (
	"github.com/shogo82148/std/sync"
)

// DLLError describes reasons for DLL load failures.
type DLLError struct {
	Err     error
	ObjName string
	Msg     string
}

func (e *DLLError) Error() string

func (e *DLLError) Unwrap() error

// Deprecated: Use [SyscallN] instead.
//
//go:nosplit
//go:uintptrkeepalive
func Syscall(trap, nargs, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)

// Deprecated: Use [SyscallN] instead.
//
//go:nosplit
//go:uintptrkeepalive
func Syscall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

// Deprecated: Use [SyscallN] instead.
//
//go:nosplit
//go:uintptrkeepalive
func Syscall9(trap, nargs, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)

// Deprecated: Use [SyscallN] instead.
//
//go:nosplit
//go:uintptrkeepalive
func Syscall12(trap, nargs, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12 uintptr) (r1, r2 uintptr, err Errno)

// Deprecated: Use [SyscallN] instead.
//
//go:nosplit
//go:uintptrkeepalive
func Syscall15(trap, nargs, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15 uintptr) (r1, r2 uintptr, err Errno)

// Deprecated: Use [SyscallN] instead.
//
//go:nosplit
//go:uintptrkeepalive
func Syscall18(trap, nargs, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18 uintptr) (r1, r2 uintptr, err Errno)

// SyscallN executes procedure p with arguments args.
//
// See [Proc.Call] for more information.
//
//go:nosplit
//go:uintptrkeepalive
func SyscallN(p uintptr, args ...uintptr) (r1, r2 uintptr, err Errno)

// A DLL implements access to a single DLL.
type DLL struct {
	Name   string
	Handle Handle
}

// LoadDLL loads the named DLL file into memory.
//
// If name is not an absolute path and is not a known system DLL used by
// Go, Windows will search for the named DLL in many locations, causing
// potential DLL preloading attacks.
//
// Use [LazyDLL] in golang.org/x/sys/windows for a secure way to
// load system DLLs.
func LoadDLL(name string) (*DLL, error)

// MustLoadDLL is like [LoadDLL] but panics if load operation fails.
func MustLoadDLL(name string) *DLL

// FindProc searches [DLL] d for procedure named name and returns [*Proc]
// if found. It returns an error if search fails.
func (d *DLL) FindProc(name string) (proc *Proc, err error)

// MustFindProc is like [DLL.FindProc] but panics if search fails.
func (d *DLL) MustFindProc(name string) *Proc

// Release unloads [DLL] d from memory.
func (d *DLL) Release() (err error)

// A Proc implements access to a procedure inside a [DLL].
type Proc struct {
	Dll  *DLL
	Name string
	addr uintptr
}

// Addr returns the address of the procedure represented by p.
// The return value can be passed to Syscall to run the procedure.
func (p *Proc) Addr() uintptr

// Call executes procedure p with arguments a.
//
// The returned error is always non-nil, constructed from the result of GetLastError.
// Callers must inspect the primary return value to decide whether an error occurred
// (according to the semantics of the specific function being called) before consulting
// the error. The error always has type [Errno].
//
// On amd64, Call can pass and return floating-point values. To pass
// an argument x with C type "float", use
// uintptr(math.Float32bits(x)). To pass an argument with C type
// "double", use uintptr(math.Float64bits(x)). Floating-point return
// values are returned in r2. The return value for C type "float" is
// [math.Float32frombits](uint32(r2)). For C type "double", it is
// [math.Float64frombits](uint64(r2)).
//
//go:uintptrescapes
func (p *Proc) Call(a ...uintptr) (uintptr, uintptr, error)

// A LazyDLL implements access to a single [DLL].
// It will delay the load of the DLL until the first
// call to its [LazyDLL.Handle] method or to one of its
// [LazyProc]'s Addr method.
//
// LazyDLL is subject to the same DLL preloading attacks as documented
// on [LoadDLL].
//
// Use LazyDLL in golang.org/x/sys/windows for a secure way to
// load system DLLs.
type LazyDLL struct {
	mu   sync.Mutex
	dll  *DLL
	Name string
}

// Load loads DLL file d.Name into memory. It returns an error if fails.
// Load will not try to load DLL, if it is already loaded into memory.
func (d *LazyDLL) Load() error

// Handle returns d's module handle.
func (d *LazyDLL) Handle() uintptr

// NewProc returns a [LazyProc] for accessing the named procedure in the [DLL] d.
func (d *LazyDLL) NewProc(name string) *LazyProc

// NewLazyDLL creates new [LazyDLL] associated with [DLL] file.
func NewLazyDLL(name string) *LazyDLL

// A LazyProc implements access to a procedure inside a [LazyDLL].
// It delays the lookup until the [LazyProc.Addr], [LazyProc.Call], or [LazyProc.Find] method is called.
type LazyProc struct {
	mu   sync.Mutex
	Name string
	l    *LazyDLL
	proc *Proc
}

// Find searches [DLL] for procedure named p.Name. It returns
// an error if search fails. Find will not search procedure,
// if it is already found and loaded into memory.
func (p *LazyProc) Find() error

// Addr returns the address of the procedure represented by p.
// The return value can be passed to Syscall to run the procedure.
func (p *LazyProc) Addr() uintptr

// Call executes procedure p with arguments a. See the documentation of
// Proc.Call for more information.
//
//go:uintptrescapes
func (p *LazyProc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error)
