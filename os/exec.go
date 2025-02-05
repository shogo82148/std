// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/time"
)

// ErrProcessDone indicates a [Process] has finished.
var ErrProcessDone = errors.New("os: process already finished")

// Process stores the information about a process created by [StartProcess].
type Process struct {
	Pid int

	// State contains the atomic process state.
	//
	// If handle is nil, this consists only of the processStatus fields,
	// which indicate if the process is done/released.
	//
	// In handle is not nil, the lower bits also contain a reference
	// count for the handle field.
	//
	// The Process itself initially holds 1 persistent reference. Any
	// operation that uses the handle with a system call temporarily holds
	// an additional transient reference. This prevents the handle from
	// being closed prematurely, which could result in the OS allocating a
	// different handle with the same value, leading to Process' methods
	// operating on the wrong process.
	//
	// Release and Wait both drop the Process' persistent reference, but
	// other concurrent references may delay actually closing the handle
	// because they hold a transient reference.
	//
	// Regardless, we want new method calls to immediately treat the handle
	// as unavailable after Release or Wait to avoid extending this delay.
	// This is achieved by setting either processStatus flag when the
	// Process' persistent reference is dropped. The only difference in the
	// flags is the reason the handle is unavailable, which affects the
	// errors returned by concurrent calls.
	state atomic.Uint64

	// Used only when handle is nil
	sigMu sync.RWMutex

	// handle, if not nil, is a pointer to a struct
	// that holds the OS-specific process handle.
	// This pointer is set when Process is created,
	// and never changed afterward.
	// This is a pointer to a separate memory allocation
	// so that we can use runtime.AddCleanup.
	handle *processHandle
}

// ProcAttr holds the attributes that will be applied to a new process
// started by StartProcess.
type ProcAttr struct {
	// If Dir is non-empty, the child changes into the directory before
	// creating the process.
	Dir string
	// If Env is non-nil, it gives the environment variables for the
	// new process in the form returned by Environ.
	// If it is nil, the result of Environ will be used.
	Env []string
	// Files specifies the open files inherited by the new process. The
	// first three entries correspond to standard input, standard output, and
	// standard error. An implementation may support additional entries,
	// depending on the underlying operating system. A nil entry corresponds
	// to that file being closed when the process starts.
	// On Unix systems, StartProcess will change these File values
	// to blocking mode, which means that SetDeadline will stop working
	// and calling Close will not interrupt a Read or Write.
	Files []*File

	// Operating system-specific process creation attributes.
	// Note that setting this field means that your program
	// may not execute properly or even compile on some
	// operating systems.
	Sys *syscall.SysProcAttr
}

// A Signal represents an operating system signal.
// The usual underlying implementation is operating system-dependent:
// on Unix it is syscall.Signal.
type Signal interface {
	String() string
	Signal()
}

// Getpid returns the process id of the caller.
func Getpid() int

// Getppid returns the process id of the caller's parent.
func Getppid() int

// FindProcess looks for a running process by its pid.
//
// The [Process] it returns can be used to obtain information
// about the underlying operating system process.
//
// On Unix systems, FindProcess always succeeds and returns a Process
// for the given pid, regardless of whether the process exists. To test whether
// the process actually exists, see whether p.Signal(syscall.Signal(0)) reports
// an error.
func FindProcess(pid int) (*Process, error)

// StartProcess starts a new process with the program, arguments and attributes
// specified by name, argv and attr. The argv slice will become [os.Args] in the
// new process, so it normally starts with the program name.
//
// If the calling goroutine has locked the operating system thread
// with [runtime.LockOSThread] and modified any inheritable OS-level
// thread state (for example, Linux or Plan 9 name spaces), the new
// process will inherit the caller's thread state.
//
// StartProcess is a low-level interface. The [os/exec] package provides
// higher-level interfaces.
//
// If there is an error, it will be of type [*PathError].
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)

// Release releases any resources associated with the [Process] p,
// rendering it unusable in the future.
// Release only needs to be called if [Process.Wait] is not.
func (p *Process) Release() error

// Kill causes the [Process] to exit immediately. Kill does not wait until
// the Process has actually exited. This only kills the Process itself,
// not any other processes it may have started.
func (p *Process) Kill() error

// Wait waits for the [Process] to exit, and then returns a
// ProcessState describing its status and an error, if any.
// Wait releases any resources associated with the Process.
// On most operating systems, the Process must be a child
// of the current process or an error will be returned.
func (p *Process) Wait() (*ProcessState, error)

// Signal sends a signal to the [Process].
// Sending [Interrupt] on Windows is not implemented.
func (p *Process) Signal(sig Signal) error

// UserTime returns the user CPU time of the exited process and its children.
func (p *ProcessState) UserTime() time.Duration

// SystemTime returns the system CPU time of the exited process and its children.
func (p *ProcessState) SystemTime() time.Duration

// Exited reports whether the program has exited.
// On Unix systems this reports true if the program exited due to calling exit,
// but false if the program terminated due to a signal.
func (p *ProcessState) Exited() bool

// Success reports whether the program exited successfully,
// such as with exit status 0 on Unix.
func (p *ProcessState) Success() bool

// Sys returns system-dependent exit information about
// the process. Convert it to the appropriate underlying
// type, such as [syscall.WaitStatus] on Unix, to access its contents.
func (p *ProcessState) Sys() any

// SysUsage returns system-dependent resource usage information about
// the exited process. Convert it to the appropriate underlying
// type, such as [*syscall.Rusage] on Unix, to access its contents.
// (On Unix, *syscall.Rusage matches struct rusage as defined in the
// getrusage(2) manual page.)
func (p *ProcessState) SysUsage() any
