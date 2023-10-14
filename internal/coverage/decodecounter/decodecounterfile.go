// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decodecounter

import (
	"github.com/shogo82148/std/internal/coverage"
	"github.com/shogo82148/std/internal/coverage/stringtab"
	"github.com/shogo82148/std/io"
)

type CounterDataReader struct {
	stab     *stringtab.Reader
	args     map[string]string
	osargs   []string
	goarch   string
	goos     string
	mr       io.ReadSeeker
	hdr      coverage.CounterFileHeader
	ftr      coverage.CounterFileFooter
	shdr     coverage.CounterSegmentHeader
	u32b     []byte
	u8b      []byte
	fcnCount uint32
	segCount uint32
	debug    bool
}

func NewCounterDataReader(fn string, rs io.ReadSeeker) (*CounterDataReader, error)

// OsArgs returns the program arguments (saved from os.Args during
// the run of the instrumented binary) read from the counter
// data file. Not all coverage data files will have os.Args values;
// for example, if a data file is produced by merging coverage
// data from two distinct runs, no os args will be available (an
// empty list is returned).
func (cdr *CounterDataReader) OsArgs() []string

// Goos returns the GOOS setting in effect for the "-cover" binary
// that produced this counter data file. The GOOS value may be
// empty in the case where the counter data file was produced
// from a merge in which more than one GOOS value was present.
func (cdr *CounterDataReader) Goos() string

// Goarch returns the GOARCH setting in effect for the "-cover" binary
// that produced this counter data file. The GOARCH value may be
// empty in the case where the counter data file was produced
// from a merge in which more than one GOARCH value was present.
func (cdr *CounterDataReader) Goarch() string

// FuncPayload encapsulates the counter data payload for a single
// function as read from a counter data file.
type FuncPayload struct {
	PkgIdx   uint32
	FuncIdx  uint32
	Counters []uint32
}

// NumSegments returns the number of execution segments in the file.
func (cdr *CounterDataReader) NumSegments() uint32

// BeginNextSegment sets up the reader to read the next segment,
// returning TRUE if we do have another segment to read, or FALSE
// if we're done with all the segments (also an error if
// something went wrong).
func (cdr *CounterDataReader) BeginNextSegment() (bool, error)

// NumFunctionsInSegment returns the number of live functions
// in the currently selected segment.
func (cdr *CounterDataReader) NumFunctionsInSegment() uint32

// NextFunc reads data for the next function in this current segment
// into "p", returning TRUE if the read was successful or FALSE
// if we've read all the functions already (also an error if
// something went wrong with the read or we hit a premature
// EOF).
func (cdr *CounterDataReader) NextFunc(p *FuncPayload) (bool, error)
