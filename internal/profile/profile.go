// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package profile provides a representation of
// github.com/google/pprof/proto/profile.proto and
// methods to encode/decode/merge profiles in this format.
package profile

import (
	"github.com/shogo82148/std/io"
)

// Profile is an in-memory representation of profile.proto.
type Profile struct {
	SampleType        []*ValueType
	DefaultSampleType string
	Sample            []*Sample
	Mapping           []*Mapping
	Location          []*Location
	Function          []*Function
	Comments          []string

	DropFrames string
	KeepFrames string

	TimeNanos     int64
	DurationNanos int64
	PeriodType    *ValueType
	Period        int64

	commentX           []int64
	dropFramesX        int64
	keepFramesX        int64
	stringTable        []string
	defaultSampleTypeX int64
}

// ValueType corresponds to Profile.ValueType
type ValueType struct {
	Type string
	Unit string

	typeX int64
	unitX int64
}

// Sample corresponds to Profile.Sample
type Sample struct {
	Location []*Location
	Value    []int64
	Label    map[string][]string
	NumLabel map[string][]int64
	NumUnit  map[string][]string

	locationIDX []uint64
	labelX      []Label
}

// Label corresponds to Profile.Label
type Label struct {
	keyX int64
	// Exactly one of the two following values must be set
	strX int64
	numX int64
}

// Mapping corresponds to Profile.Mapping
type Mapping struct {
	ID              uint64
	Start           uint64
	Limit           uint64
	Offset          uint64
	File            string
	BuildID         string
	HasFunctions    bool
	HasFilenames    bool
	HasLineNumbers  bool
	HasInlineFrames bool

	fileX    int64
	buildIDX int64
}

// Location corresponds to Profile.Location
type Location struct {
	ID       uint64
	Mapping  *Mapping
	Address  uint64
	Line     []Line
	IsFolded bool

	mappingIDX uint64
}

// Line corresponds to Profile.Line
type Line struct {
	Function *Function
	Line     int64

	functionIDX uint64
}

// Function corresponds to Profile.Function
type Function struct {
	ID         uint64
	Name       string
	SystemName string
	Filename   string
	StartLine  int64

	nameX       int64
	systemNameX int64
	filenameX   int64
}

// Parse parses a profile and checks for its validity. The input
// may be a gzip-compressed encoded protobuf or one of many legacy
// profile formats which may be unsupported in the future.
func Parse(r io.Reader) (*Profile, error)

// Write writes the profile as a gzip-compressed marshaled protobuf.
func (p *Profile) Write(w io.Writer) error

// CheckValid tests whether the profile is valid. Checks include, but are
// not limited to:
//   - len(Profile.Sample[n].value) == len(Profile.value_unit)
//   - Sample.id has a corresponding Profile.Location
func (p *Profile) CheckValid() error

// Aggregate merges the locations in the profile into equivalence
// classes preserving the request attributes. It also updates the
// samples to point to the merged locations.
func (p *Profile) Aggregate(inlineFrame, function, filename, linenumber, address bool) error

// Print dumps a text representation of a profile. Intended mainly
// for debugging purposes.
func (p *Profile) String() string

// Merge adds profile p adjusted by ratio r into profile p. Profiles
// must be compatible (same Type and SampleType).
// TODO(rsilvera): consider normalizing the profiles based on the
// total samples collected.
func (p *Profile) Merge(pb *Profile, r float64) error

// Compatible determines if two profiles can be compared/merged.
// returns nil if the profiles are compatible; otherwise an error with
// details on the incompatibility.
func (p *Profile) Compatible(pb *Profile) error

// HasFunctions determines if all locations in this profile have
// symbolized function information.
func (p *Profile) HasFunctions() bool

// HasFileLines determines if all locations in this profile have
// symbolized file and line number information.
func (p *Profile) HasFileLines() bool

// Copy makes a fully independent copy of a profile.
func (p *Profile) Copy() *Profile

// Demangler maps symbol names to a human-readable form. This may
// include C++ demangling and additional simplification. Names that
// are not demangled may be missing from the resulting map.
type Demangler func(name []string) (map[string]string, error)

// Demangle attempts to demangle and optionally simplify any function
// names referenced in the profile. It works on a best-effort basis:
// it will silently preserve the original names in case of any errors.
func (p *Profile) Demangle(d Demangler) error

// Empty reports whether the profile contains no samples.
func (p *Profile) Empty() bool

// Scale multiplies all sample values in a profile by a constant.
func (p *Profile) Scale(ratio float64)

// ScaleN multiplies each sample values in a sample by a different amount.
func (p *Profile) ScaleN(ratios []float64) error
