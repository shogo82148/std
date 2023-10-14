// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

package poll

import (
	"github.com/shogo82148/std/time"
)

// SetDeadline sets the read and write deadlines associated with fd.
func (fd *FD) SetDeadline(t time.Time) error

// SetReadDeadline sets the read deadline associated with fd.
func (fd *FD) SetReadDeadline(t time.Time) error

// SetWriteDeadline sets the write deadline associated with fd.
func (fd *FD) SetWriteDeadline(t time.Time) error

// IsPollDescriptor reports whether fd is the descriptor being used by the poller.
// This is only used for testing.
func IsPollDescriptor(fd uintptr) bool
