// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gate contains an alternative condition variable.
package gate

import "github.com/shogo82148/std/context"

// A Gate is a monitor (mutex + condition variable) with one bit of state.
//
// The condition may be either set or unset.
// Lock operations may be unconditional, or wait for the condition to be set.
// Unlock operations record the new state of the condition.
type Gate struct {
	// When unlocked, exactly one of set or unset contains a value.
	// When locked, neither chan contains a value.
	set   chan struct{}
	unset chan struct{}
}

// New returns a new, unlocked gate.
func New(set bool) Gate

// Lock acquires the gate unconditionally.
// It reports whether the condition is set.
func (g *Gate) Lock() (set bool)

// WaitAndLock waits until the condition is set before acquiring the gate.
// If the context expires, WaitAndLock returns an error and does not acquire the gate.
func (g *Gate) WaitAndLock(ctx context.Context) error

// LockIfSet acquires the gate if and only if the condition is set.
func (g *Gate) LockIfSet() (acquired bool)

// Unlock sets the condition and releases the gate.
func (g *Gate) Unlock(set bool)
