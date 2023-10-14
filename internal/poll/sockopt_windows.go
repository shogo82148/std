// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

import "github.com/shogo82148/std/syscall"

// Setsockopt wraps the setsockopt network call.
func (fd *FD) Setsockopt(level, optname int32, optval *byte, optlen int32) error

// WSAIoctl wraps the WSAIoctl network call.
func (fd *FD) WSAIoctl(iocc uint32, inbuf *byte, cbif uint32, outbuf *byte, cbob uint32, cbbr *uint32, overlapped *syscall.Overlapped, completionRoutine uintptr) error
