// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

import (
	"github.com/shogo82148/std/errors"
)

// ErrUnexpectedSuccess indicates that a script command that was expected to
// fail (as indicated by a "!" prefix) instead completed successfully.
var ErrUnexpectedSuccess = errors.New("unexpected success")

// A CommandError describes an error resulting from attempting to execute a
// specific command.
type CommandError struct {
	File string
	Line int
	Op   string
	Args []string
	Err  error
}

func (e *CommandError) Error() string

func (e *CommandError) Unwrap() error

// A UsageError reports the valid arguments for a command.
//
// It may be returned in response to invalid arguments.
type UsageError struct {
	Name    string
	Command Cmd
}

func (e *UsageError) Error() string

// ErrUsage may be returned by a Command to indicate that it was called with
// invalid arguments; its Usage method may be called to obtain details.
var ErrUsage = errors.New("invalid usage")
