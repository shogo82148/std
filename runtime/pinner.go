// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// A Pinner is a set of Go objects each pinned to a fixed location in memory. The
// [Pin] method pins one object, while [Unpin] unpins all pinned objects. See their
// comments for more information.
type Pinner struct {
	*pinner
}

// Pin pins a Go object, preventing it from being moved or freed by the garbage
// collector until the [Pinner.Unpin] method has been called.
//
// A pointer to a pinned object can be directly stored in C memory or can be
// contained in Go memory passed to C functions. If the pinned object itself
// contains pointers to Go objects, these objects must be pinned separately if they
// are going to be accessed from C code.
//
// The argument must be a pointer of any type or an [unsafe.Pointer].
// It's safe to call Pin on non-Go pointers, in which case Pin will do nothing.
func (p *Pinner) Pin(pointer any)

// Unpin unpins all pinned objects of the [Pinner].
func (p *Pinner) Unpin()
