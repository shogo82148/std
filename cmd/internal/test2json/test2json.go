// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package test2json implements conversion of test binary output to JSON.
// It is used by cmd/test2json and cmd/go.
//
// See the cmd/test2json documentation for details of the JSON encoding.
package test2json

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

// Mode controls details of the conversion.
type Mode int

const (
	Timestamp Mode = 1 << iota
)

// A Converter holds the state of a test-to-JSON conversion.
// It implements io.WriteCloser; the caller writes test output in,
// and the converter writes JSON output to w.
type Converter struct {
	w          io.Writer
	pkg        string
	mode       Mode
	start      time.Time
	testName   string
	report     []*event
	result     string
	input      lineBuffer
	output     lineBuffer
	needMarker bool
}

// NewConverter returns a "test to json" converter.
// Writes on the returned writer are written as JSON to w,
// with minimal delay.
//
// The writes to w are whole JSON events ending in \n,
// so that it is safe to run multiple tests writing to multiple converters
// writing to a single underlying output stream w.
// As long as the underlying output w can handle concurrent writes
// from multiple goroutines, the result will be a JSON stream
// describing the relative ordering of execution in all the concurrent tests.
//
// The mode flag adjusts the behavior of the converter.
// Passing ModeTime includes event timestamps and elapsed times.
//
// The pkg string, if present, specifies the import path to
// report in the JSON stream.
func NewConverter(w io.Writer, pkg string, mode Mode) *Converter

// Write writes the test input to the converter.
func (c *Converter) Write(b []byte) (int, error)

// Exited marks the test process as having exited with the given error.
func (c *Converter) Exited(err error)

// Close marks the end of the go test output.
// It flushes any pending input and then output (only partial lines at this point)
// and then emits the final overall package-level pass/fail event.
func (c *Converter) Close() error
