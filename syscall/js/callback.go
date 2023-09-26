// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm
// +build js,wasm

package js

// Callback is a Go function that got wrapped for use as a JavaScript callback.
type Callback struct {
	Value
	id    uint32
}

// NewCallback returns a wrapped callback function.
//
// Invoking the callback in JavaScript will queue the Go function fn for execution.
// This execution happens asynchronously on a special goroutine that handles all callbacks and preserves
// the order in which the callbacks got called.
// As a consequence, if one callback blocks this goroutine, other callbacks will not be processed.
// A blocking callback should therefore explicitly start a new goroutine.
//
// Callback.Release must be called to free up resources when the callback will not be used any more.
func NewCallback(fn func(args []Value)) Callback

type EventCallbackFlag int

const (
	// PreventDefault can be used with NewEventCallback to call event.preventDefault synchronously.
	PreventDefault EventCallbackFlag = 1 << iota
	// StopPropagation can be used with NewEventCallback to call event.stopPropagation synchronously.
	StopPropagation
	// StopImmediatePropagation can be used with NewEventCallback to call event.stopImmediatePropagation synchronously.
	StopImmediatePropagation
)

// NewEventCallback returns a wrapped callback function, just like NewCallback, but the callback expects to have
// exactly one argument, the event. Depending on flags, it will synchronously call event.preventDefault,
// event.stopPropagation and/or event.stopImmediatePropagation before queuing the Go function fn for execution.
func NewEventCallback(flags EventCallbackFlag, fn func(event Value)) Callback

// Release frees up resources allocated for the callback.
// The callback must not be invoked after calling Release.
func (c Callback) Release()
