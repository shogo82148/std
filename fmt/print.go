// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt

import (
	"github.com/shogo82148/std/io"
)

// Strings for use with buffer.WriteString.
// This is less overhead than using buffer.Write with byte arrays.

// State represents the printer state passed to custom formatters.
// It provides access to the io.Writer interface plus information about
// the flags and options for the operand's format specifier.
type State interface {
	Write(b []byte) (n int, err error)

	Width() (wid int, ok bool)

	Precision() (prec int, ok bool)

	Flag(c int) bool
}

// Formatter is implemented by any value that has a Format method.
// The implementation controls how State and rune are interpreted,
// and may call Sprint(f) or Fprint(f) etc. to generate its output.
type Formatter interface {
	Format(f State, verb rune)
}

// Stringer is implemented by any value that has a String method,
// which defines the “native” format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
type Stringer interface {
	String() string
}

// GoStringer is implemented by any value that has a GoString method,
// which defines the Go syntax for that value.
// The GoString method is used to print values passed as an operand
// to a %#v format.
type GoStringer interface {
	GoString() string
}

// FormatString returns a string representing the fully qualified formatting
// directive captured by the State, followed by the argument verb. (State does not
// itself contain the verb.) The result has a leading percent sign followed by any
// flags, the width, and the precision. Missing flags, width, and precision are
// omitted. This function allows a Formatter to reconstruct the original
// directive triggering the call to Format.
func FormatString(state State, verb rune) string

// Use simple []byte instead of bytes.Buffer to avoid large dependency.

// pp is used to store a printer's state and is reused with sync.Pool to avoid allocations.

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, format string, a ...any) (n int, err error)

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...any) (n int, err error)

// Sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(format string, a ...any) string

// Appendf formats according to a format specifier, appends the result to the byte
// slice, and returns the updated slice.
func Appendf(b []byte, format string, a ...any) []byte

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...any) (n int, err error)

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(a ...any) (n int, err error)

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func Sprint(a ...any) string

// Append formats using the default formats for its operands, appends the result to
// the byte slice, and returns the updated slice.
func Append(b []byte, a ...any) []byte

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Fprintln(w io.Writer, a ...any) (n int, err error)

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...any) (n int, err error)

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func Sprintln(a ...any) string

// Appendln formats using the default formats for its operands, appends the result
// to the byte slice, and returns the updated slice. Spaces are always added
// between operands and a newline is appended.
func Appendln(b []byte, a ...any) []byte
