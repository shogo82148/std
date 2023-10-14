// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package testpty is a simple pseudo-terminal package for Unix systems,
// implemented by calling C functions via cgo.
package testpty

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/os"
)

type PtyError struct {
	FuncName    string
	ErrorString string
	Errno       error
}

func (e *PtyError) Error() string

func (e *PtyError) Unwrap() error

var ErrNotSupported = errors.New("testpty.Open not implemented on this platform")

// Open returns a control pty and the name of the linked process tty.
//
// If Open is not implemented on this platform, it returns ErrNotSupported.
func Open() (pty *os.File, processTTY string, err error)
