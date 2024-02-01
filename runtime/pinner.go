// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// A Pinner is a set of pinned Go objects. An object can be pinned with
// the Pin method and all pinned objects of a Pinner can be unpinned with the
// Unpin method.
type Pinner struct {
	*pinner
}

// Pin pins a Go object, preventing it from being moved or freed by the garbage
// collector until the Unpin method has been called.
//
// A pointer to a pinned
// object can be directly stored in C memory or can be contained in Go memory
// passed to C functions. If the pinned object itself contains pointers to Go
// objects, these objects must be pinned separately if they are going to be
// accessed from C code.
//
// The argument must be a pointer of any type or an unsafe.Pointer.
func (p *Pinner) Pin(pointer any)

// Unpin unpins all pinned objects of the Pinner.
func (p *Pinner) Unpin()
