// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package log implements a simple logging package. It defines a type, [Logger],
// with methods for formatting output. It also has a predefined 'standard'
// Logger accessible through helper functions Print[f|ln], Fatal[f|ln], and
// Panic[f|ln], which are easier to use than creating a Logger manually.
// That logger writes to standard error and prints the date and time
// of each logged message.
// Every log message is output on a separate line: if the message being
// printed does not end in a newline, the logger will add one.
// The Fatal functions call [os.Exit](1) after writing the log message.
// The Panic functions call panic after writing the log message.
package log

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
)

// These flags define which text to prefix to each log entry generated by the [Logger].
// Bits are or'ed together to control what's printed.
// With the exception of the Lmsgprefix flag, there is no
// control over the order they appear (the order listed here)
// or the format they present (as described in the comments).
// The prefix is followed by a colon only when Llongfile or Lshortfile
// is specified.
// For example, flags Ldate | Ltime (or LstdFlags) produce,
//
//	2009/01/23 01:23:23 message
//
// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
//
//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	Lmsgprefix
	LstdFlags = Ldate | Ltime
)

// A Logger represents an active logging object that generates lines of
// output to an [io.Writer]. Each logging operation makes a single call to
// the Writer's Write method. A Logger can be used simultaneously from
// multiple goroutines; it guarantees to serialize access to the Writer.
type Logger struct {
	outMu sync.Mutex
	out   io.Writer

	prefix    atomic.Pointer[string]
	flag      atomic.Int32
	isDiscard atomic.Bool
}

// New creates a new [Logger]. The out variable sets the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line, or
// after the log header if the [Lmsgprefix] flag is provided.
// The flag argument defines the logging properties.
func New(out io.Writer, prefix string, flag int) *Logger

// SetOutput sets the output destination for the logger.
func (l *Logger) SetOutput(w io.Writer)

// Default returns the standard logger used by the package-level output functions.
func Default() *Logger

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is used to recover the PC and is
// provided for generality, although at the moment on all pre-defined
// paths it will be 2.
func (l *Logger) Output(calldepth int, s string) error

// Print calls l.Output to print to the logger.
// Arguments are handled in the manner of [fmt.Print].
func (l *Logger) Print(v ...any)

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of [fmt.Printf].
func (l *Logger) Printf(format string, v ...any)

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of [fmt.Println].
func (l *Logger) Println(v ...any)

// Fatal is equivalent to l.Print() followed by a call to [os.Exit](1).
func (l *Logger) Fatal(v ...any)

// Fatalf is equivalent to l.Printf() followed by a call to [os.Exit](1).
func (l *Logger) Fatalf(format string, v ...any)

// Fatalln is equivalent to l.Println() followed by a call to [os.Exit](1).
func (l *Logger) Fatalln(v ...any)

// Panic is equivalent to l.Print() followed by a call to panic().
func (l *Logger) Panic(v ...any)

// Panicf is equivalent to l.Printf() followed by a call to panic().
func (l *Logger) Panicf(format string, v ...any)

// Panicln is equivalent to l.Println() followed by a call to panic().
func (l *Logger) Panicln(v ...any)

// Flags returns the output flags for the logger.
// The flag bits are [Ldate], [Ltime], and so on.
func (l *Logger) Flags() int

// SetFlags sets the output flags for the logger.
// The flag bits are [Ldate], [Ltime], and so on.
func (l *Logger) SetFlags(flag int)

// Prefix returns the output prefix for the logger.
func (l *Logger) Prefix() string

// SetPrefix sets the output prefix for the logger.
func (l *Logger) SetPrefix(prefix string)

// Writer returns the output destination for the logger.
func (l *Logger) Writer() io.Writer

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer)

// Flags returns the output flags for the standard logger.
// The flag bits are [Ldate], [Ltime], and so on.
func Flags() int

// SetFlags sets the output flags for the standard logger.
// The flag bits are [Ldate], [Ltime], and so on.
func SetFlags(flag int)

// Prefix returns the output prefix for the standard logger.
func Prefix() string

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string)

// Writer returns the output destination for the standard logger.
func Writer() io.Writer

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of [fmt.Print].
func Print(v ...any)

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of [fmt.Printf].
func Printf(format string, v ...any)

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of [fmt.Println].
func Println(v ...any)

// Fatal is equivalent to [Print] followed by a call to [os.Exit](1).
func Fatal(v ...any)

// Fatalf is equivalent to [Printf] followed by a call to [os.Exit](1).
func Fatalf(format string, v ...any)

// Fatalln is equivalent to [Println] followed by a call to [os.Exit](1).
func Fatalln(v ...any)

// Panic is equivalent to [Print] followed by a call to panic().
func Panic(v ...any)

// Panicf is equivalent to [Printf] followed by a call to panic().
func Panicf(format string, v ...any)

// Panicln is equivalent to [Println] followed by a call to panic().
func Panicln(v ...any)

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is the count of the number of
// frames to skip when computing the file name and line number
// if [Llongfile] or [Lshortfile] is set; a value of 1 will print the details
// for the caller of Output.
func Output(calldepth int, s string) error
