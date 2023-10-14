// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package poll supports non-blocking I/O on file descriptors with polling.
// This supports I/O operations that block only a goroutine, not a thread.
// This is used by the net and os packages.
// It uses a poller built into the runtime, with support from the
// runtime scheduler.
package poll

import (
	"github.com/shogo82148/std/errors"
)

// ErrNetClosing is returned when a network descriptor is used after
// it has been closed.
var ErrNetClosing = errNetClosing{}

// ErrFileClosing is returned when a file descriptor is used after it
// has been closed.
var ErrFileClosing = errors.New("use of closed file")

// ErrNoDeadline is returned when a request is made to set a deadline
// on a file type that does not use the poller.
var ErrNoDeadline = errors.New("file type does not support deadline")

// ErrDeadlineExceeded is returned for an expired deadline.
// This is exported by the os package as os.ErrDeadlineExceeded.
var ErrDeadlineExceeded error = &DeadlineExceededError{}

// DeadlineExceededError is returned for an expired deadline.
type DeadlineExceededError struct{}

// Implement the net.Error interface.
// The string is "i/o timeout" because that is what was returned
// by earlier Go versions. Changing it may break programs that
// match on error strings.
func (e *DeadlineExceededError) Error() string
func (e *DeadlineExceededError) Timeout() bool
func (e *DeadlineExceededError) Temporary() bool

// ErrNotPollable is returned when the file or socket is not suitable
// for event notification.
var ErrNotPollable = errors.New("not pollable")

// TestHookDidWritev is a hook for testing writev.
var TestHookDidWritev = func(wrote int) {}
