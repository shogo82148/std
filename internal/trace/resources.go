// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

// ThreadID is the runtime-internal M structure's ID. This is unique
// for each OS thread.
type ThreadID int64

// NoThread indicates that the relevant events don't correspond to any
// thread in particular.
const NoThread = ThreadID(-1)

// ProcID is the runtime-internal G structure's id field. This is unique
// for each P.
type ProcID int64

// NoProc indicates that the relevant events don't correspond to any
// P in particular.
const NoProc = ProcID(-1)

// GoID is the runtime-internal G structure's goid field. This is unique
// for each goroutine.
type GoID int64

// NoGoroutine indicates that the relevant events don't correspond to any
// goroutine in particular.
const NoGoroutine = GoID(-1)

// GoState represents the state of a goroutine.
//
// New GoStates may be added in the future. Users of this type must be robust
// to that possibility.
type GoState uint8

const (
	GoUndetermined GoState = iota
	GoNotExist
	GoRunnable
	GoRunning
	GoWaiting
	GoSyscall
)

// Executing returns true if the state indicates that the goroutine is executing
// and bound to its thread.
func (s GoState) Executing() bool

// String returns a human-readable representation of a GoState.
//
// The format of the returned string is for debugging purposes and is subject to change.
func (s GoState) String() string

// ProcState represents the state of a proc.
//
// New ProcStates may be added in the future. Users of this type must be robust
// to that possibility.
type ProcState uint8

const (
	ProcUndetermined ProcState = iota
	ProcNotExist
	ProcRunning
	ProcIdle
)

// Executing returns true if the state indicates that the proc is executing
// and bound to its thread.
func (s ProcState) Executing() bool

// String returns a human-readable representation of a ProcState.
//
// The format of the returned string is for debugging purposes and is subject to change.
func (s ProcState) String() string

// ResourceKind indicates a kind of resource that has a state machine.
//
// New ResourceKinds may be added in the future. Users of this type must be robust
// to that possibility.
type ResourceKind uint8

const (
	ResourceNone      ResourceKind = iota
	ResourceGoroutine
	ResourceProc
	ResourceThread
)

// String returns a human-readable representation of a ResourceKind.
//
// The format of the returned string is for debugging purposes and is subject to change.
func (r ResourceKind) String() string

// ResourceID represents a generic resource ID.
type ResourceID struct {
	// Kind is the kind of resource this ID is for.
	Kind ResourceKind
	id   int64
}

// MakeResourceID creates a general resource ID from a specific resource's ID.
func MakeResourceID[T interface{ GoID | ProcID | ThreadID }](id T) ResourceID

// Goroutine obtains a GoID from the resource ID.
//
// r.Kind must be ResourceGoroutine or this function will panic.
func (r ResourceID) Goroutine() GoID

// Proc obtains a ProcID from the resource ID.
//
// r.Kind must be ResourceProc or this function will panic.
func (r ResourceID) Proc() ProcID

// Thread obtains a ThreadID from the resource ID.
//
// r.Kind must be ResourceThread or this function will panic.
func (r ResourceID) Thread() ThreadID

// String returns a human-readable string representation of the ResourceID.
//
// This representation is subject to change and is intended primarily for debugging.
func (r ResourceID) String() string

// StateTransition provides details about a StateTransition event.
type StateTransition struct {
	// Resource is the resource this state transition is for.
	Resource ResourceID

	// Reason is a human-readable reason for the state transition.
	Reason string

	// Stack is the stack trace of the resource making the state transition.
	//
	// This is distinct from the result (Event).Stack because it pertains to
	// the transitioning resource, not any of the ones executing the event
	// this StateTransition came from.
	//
	// An example of this difference is the NotExist -> Runnable transition for
	// goroutines, which indicates goroutine creation. In this particular case,
	// a Stack here would refer to the starting stack of the new goroutine, and
	// an (Event).Stack would refer to the stack trace of whoever created the
	// goroutine.
	Stack Stack

	// The actual transition data. Stored in a neutral form so that
	// we don't need fields for every kind of resource.
	id       int64
	oldState uint8
	newState uint8
}

// Goroutine returns the state transition for a goroutine.
//
// Transitions to and from states that are Executing are special in that
// they change the future execution context. In other words, future events
// on the same thread will feature the same goroutine until it stops running.
//
// Panics if d.Resource.Kind is not ResourceGoroutine.
func (d StateTransition) Goroutine() (from, to GoState)

// Proc returns the state transition for a proc.
//
// Transitions to and from states that are Executing are special in that
// they change the future execution context. In other words, future events
// on the same thread will feature the same goroutine until it stops running.
//
// Panics if d.Resource.Kind is not ResourceProc.
func (d StateTransition) Proc() (from, to ProcState)
