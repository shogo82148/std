// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encodecounter

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/internal/coverage"
	"github.com/shogo82148/std/internal/coverage/stringtab"
	"github.com/shogo82148/std/io"
)

type CoverageDataWriter struct {
	stab    *stringtab.Writer
	w       *bufio.Writer
	csh     coverage.CounterSegmentHeader
	tmp     []byte
	cflavor coverage.CounterFlavor
	segs    uint32
	debug   bool
}

func NewCoverageDataWriter(w io.Writer, flav coverage.CounterFlavor) *CoverageDataWriter

// CounterVisitor describes a helper object used during counter file
// writing; when writing counter data files, clients pass a
// CounterVisitor to the write/emit routines, then the expectation is
// that the VisitFuncs method will then invoke the callback "f" with
// data for each function to emit to the file.
type CounterVisitor interface {
	VisitFuncs(f CounterVisitorFn) error
}

// CounterVisitorFn describes a callback function invoked when writing
// coverage counter data.
type CounterVisitorFn func(pkid uint32, funcid uint32, counters []uint32) error

// Write writes the contents of the count-data file to the writer
// previously supplied to NewCoverageDataWriter. Returns an error
// if something went wrong somewhere with the write.
func (cfw *CoverageDataWriter) Write(metaFileHash [16]byte, args map[string]string, visitor CounterVisitor) error

// AppendSegment appends a new segment to a counter data, with a new
// args section followed by a payload of counter data clauses.
func (cfw *CoverageDataWriter) AppendSegment(args map[string]string, visitor CounterVisitor) error
