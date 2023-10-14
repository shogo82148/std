// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package coverage

// CovMetaMagic holds the magic string for a meta-data file.
var CovMetaMagic = [4]byte{'\x00', '\x63', '\x76', '\x6d'}

// MetaFilePref is a prefix used when emitting meta-data files; these
// files are of the form "covmeta.<hash>", where hash is a hash
// computed from the hashes of all the package meta-data symbols in
// the program.
const MetaFilePref = "covmeta"

// MetaFileVersion contains the current (most recent) meta-data file version.
const MetaFileVersion = 1

// MetaFileHeader stores file header information for a meta-data file.
type MetaFileHeader struct {
	Magic        [4]byte
	Version      uint32
	TotalLength  uint64
	Entries      uint64
	MetaFileHash [16]byte
	StrTabOffset uint32
	StrTabLength uint32
	CMode        CounterMode
	CGranularity CounterGranularity
	_            [6]byte
}

// MetaSymbolHeader stores header information for a single
// meta-data blob, e.g. the coverage meta-data payload
// computed for a given Go package.
type MetaSymbolHeader struct {
	Length     uint32
	PkgName    uint32
	PkgPath    uint32
	ModulePath uint32
	MetaHash   [16]byte
	_          byte
	_          [3]byte
	NumFiles   uint32
	NumFuncs   uint32
}

const CovMetaHeaderSize = 16 + 4 + 4 + 4 + 4 + 4 + 4 + 4

// FuncDesc encapsulates the meta-data definitions for a single Go function.
// This version assumes that we're looking at a function before inlining;
// if we want to capture a post-inlining view of the world, the
// representations of source positions would need to be a good deal more
// complicated.
type FuncDesc struct {
	Funcname string
	Srcfile  string
	Units    []CoverableUnit
	Lit      bool
}

// CoverableUnit describes the source characteristics of a single
// program unit for which we want to gather coverage info. Coverable
// units are either "simple" or "intraline"; a "simple" coverable unit
// corresponds to a basic block (region of straight-line code with no
// jumps or control transfers). An "intraline" unit corresponds to a
// logical clause nested within some other simple unit. A simple unit
// will have a zero Parent value; for an intraline unit NxStmts will
// be zero and Parent will be set to 1 plus the index of the
// containing simple statement. Example:
//
//	L7:   q := 1
//	L8:   x := (y == 101 || launch() == false)
//	L9:   r := x * 2
//
// For the code above we would have three simple units (one for each
// line), then an intraline unit describing the "launch() == false"
// clause in line 8, with Parent pointing to the index of the line 8
// unit in the units array.
//
// Note: in the initial version of the coverage revamp, only simple
// units will be in use.
type CoverableUnit struct {
	StLine, StCol uint32
	EnLine, EnCol uint32
	NxStmts       uint32
	Parent        uint32
}

// CounterMode tracks the "flavor" of the coverage counters being
// used in a given coverage-instrumented program.
type CounterMode uint8

const (
	CtrModeInvalid  CounterMode = iota
	CtrModeSet
	CtrModeCount
	CtrModeAtomic
	CtrModeRegOnly
	CtrModeTestMain
)

func (cm CounterMode) String() string

func ParseCounterMode(mode string) CounterMode

// CounterGranularity tracks the granularity of the coverage counters being
// used in a given coverage-instrumented program.
type CounterGranularity uint8

const (
	CtrGranularityInvalid CounterGranularity = iota
	CtrGranularityPerBlock
	CtrGranularityPerFunc
)

func (cm CounterGranularity) String() string

// CovCounterMagic holds the magic string for a coverage counter-data file.
var CovCounterMagic = [4]byte{'\x00', '\x63', '\x77', '\x6d'}

// CounterFileVersion stores the most recent counter data file version.
const CounterFileVersion = 1

// CounterFileHeader stores files header information for a counter-data file.
type CounterFileHeader struct {
	Magic     [4]byte
	Version   uint32
	MetaHash  [16]byte
	CFlavor   CounterFlavor
	BigEndian bool
	_         [6]byte
}

// CounterSegmentHeader encapsulates information about a specific
// segment in a counter data file, which at the moment contains
// counters data from a single execution of a coverage-instrumented
// program. Following the segment header will be the string table and
// args table, and then (possibly) padding bytes to bring the byte
// size of the preamble up to a multiple of 4. Immediately following
// that will be the counter payloads.
//
// The "args" section of a segment is used to store annotations
// describing where the counter data came from; this section is
// basically a series of key-value pairs (can be thought of as an
// encoded 'map[string]string'). At the moment we only write os.Args()
// data to this section, using pairs of the form "argc=<integer>",
// "argv0=<os.Args[0]>", "argv1=<os.Args[1]>", and so on. In the
// future the args table may also include things like GOOS/GOARCH
// values, and/or tags indicating which tests were run to generate the
// counter data.
type CounterSegmentHeader struct {
	FcnEntries uint64
	StrTabLen  uint32
	ArgsLen    uint32
}

// CounterFileFooter appears at the tail end of a counter data file,
// and stores the number of segments it contains.
type CounterFileFooter struct {
	Magic       [4]byte
	_           [4]byte
	NumSegments uint32
	_           [4]byte
}

// CounterFilePref is the file prefix used when emitting coverage data
// output files. CounterFileTemplate describes the format of the file
// name: prefix followed by meta-file hash followed by process ID
// followed by emit UnixNanoTime.
const CounterFilePref = "covcounters"
const CounterFileTempl = "%s.%x.%d.%d"
const CounterFileRegexp = `^%s\.(\S+)\.(\d+)\.(\d+)+$`

// CounterFlavor describes how function and counters are
// stored/represented in the counter section of the file.
type CounterFlavor uint8

const (
	// "Raw" representation: all values (pkg ID, func ID, num counters,
	// and counters themselves) are stored as uint32's.
	CtrRaw CounterFlavor = iota + 1

	// "ULeb" representation: all values (pkg ID, func ID, num counters,
	// and counters themselves) are stored with ULEB128 encoding.
	CtrULeb128
)

func Round4(x int) int

const NumCtrsOffset = 0
const PkgIdOffset = 1
const FuncIdOffset = 2
const FirstCtrOffset = 3
