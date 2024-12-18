// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanner

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
)

// In an [ErrorList], an error is represented by an *Error.
// The position Pos, if valid, points to the beginning of
// the offending token, and the error condition is described
// by Msg.
type Error struct {
	Pos token.Position
	Msg string
}

// Error implements the error interface.
func (e Error) Error() string

// ErrorList is a list of *Errors.
// The zero value for an ErrorList is an empty ErrorList ready to use.
type ErrorList []*Error

// Add adds an [Error] with given position and error message to an [ErrorList].
func (p *ErrorList) Add(pos token.Position, msg string)

// Reset resets an [ErrorList] to no errors.
func (p *ErrorList) Reset()

// [ErrorList] implements the sort Interface.
func (p ErrorList) Len() int
func (p ErrorList) Swap(i, j int)

func (p ErrorList) Less(i, j int) bool

// Sort sorts an [ErrorList]. *[Error] entries are sorted by position,
// other errors are sorted by error message, and before any *[Error]
// entry.
func (p ErrorList) Sort()

// RemoveMultiples sorts an [ErrorList] and removes all but the first error per line.
func (p *ErrorList) RemoveMultiples()

// An [ErrorList] implements the error interface.
func (p ErrorList) Error() string

// Err returns an error equivalent to this error list.
// If the list is empty, Err returns nil.
func (p ErrorList) Err() error

// PrintError is a utility function that prints a list of errors to w,
// one error per line, if the err parameter is an [ErrorList]. Otherwise
// it prints the err string.
func PrintError(w io.Writer, err error)
