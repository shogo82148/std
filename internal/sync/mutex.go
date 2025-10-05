// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sync provides basic synchronization primitives such as mutual
// exclusion locks to internal packages (including ones that depend on sync).
//
// Tests are defined in package [sync].
package sync

// A Mutex is a mutual exclusion lock.
//
// See package [sync.Mutex] documentation.
type Mutex struct {
	state int32
	sema  uint32
}

// Lock locks m.
//
// See package [sync.Mutex] documentation.
func (m *Mutex) Lock()

// TryLock tries to lock m and reports whether it succeeded.
//
// See package [sync.Mutex] documentation.
func (m *Mutex) TryLock() bool

// Unlock unlocks m.
//
// See package [sync.Mutex] documentation.
func (m *Mutex) Unlock()
