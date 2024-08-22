// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgbits

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/go/constant"
	"github.com/shogo82148/std/io"
)

// A PkgEncoder provides methods for encoding a package's Unified IR
// export data.
type PkgEncoder struct {
	// version of the bitstream.
	version Version

	// elems holds the bitstream for previously encoded elements.
	elems [numRelocs][]string

	// stringsIdx maps previously encoded strings to their index within
	// the RelocString section, to allow deduplication. That is,
	// elems[RelocString][stringsIdx[s]] == s (if present).
	stringsIdx map[string]Index

	// syncFrames is the number of frames to write at each sync
	// marker. A negative value means sync markers are omitted.
	syncFrames int
}

// SyncMarkers reports whether pw uses sync markers.
func (pw *PkgEncoder) SyncMarkers() bool

// NewPkgEncoder returns an initialized PkgEncoder.
//
// syncFrames is the number of caller frames that should be serialized
// at Sync points. Serializing additional frames results in larger
// export data files, but can help diagnosing desync errors in
// higher-level Unified IR reader/writer code. If syncFrames is
// negative, then sync markers are omitted entirely.
func NewPkgEncoder(syncFrames int) PkgEncoder

// DumpTo writes the package's encoded data to out0 and returns the
// package fingerprint.
func (pw *PkgEncoder) DumpTo(out0 io.Writer) (fingerprint [8]byte)

// StringIdx adds a string value to the strings section, if not
// already present, and returns its index.
func (pw *PkgEncoder) StringIdx(s string) Index

// NewEncoder returns an Encoder for a new element within the given
// section, and encodes the given SyncMarker as the start of the
// element bitstream.
func (pw *PkgEncoder) NewEncoder(k RelocKind, marker SyncMarker) Encoder

// NewEncoderRaw returns an Encoder for a new element within the given
// section.
//
// Most callers should use NewEncoder instead.
func (pw *PkgEncoder) NewEncoderRaw(k RelocKind) Encoder

// An Encoder provides methods for encoding an individual element's
// bitstream data.
type Encoder struct {
	p *PkgEncoder

	Relocs   []RelocEnt
	RelocMap map[RelocEnt]uint32
	Data     bytes.Buffer

	encodingRelocHeader bool

	k   RelocKind
	Idx Index
}

// Flush finalizes the element's bitstream and returns its Index.
func (w *Encoder) Flush() Index

func (w *Encoder) Sync(m SyncMarker)

// Bool encodes and writes a bool value into the element bitstream,
// and then returns the bool value.
//
// For simple, 2-alternative encodings, the idiomatic way to call Bool
// is something like:
//
//	if w.Bool(x != 0) {
//		// alternative #1
//	} else {
//		// alternative #2
//	}
//
// For multi-alternative encodings, use Code instead.
func (w *Encoder) Bool(b bool) bool

// Int64 encodes and writes an int64 value into the element bitstream.
func (w *Encoder) Int64(x int64)

// Uint64 encodes and writes a uint64 value into the element bitstream.
func (w *Encoder) Uint64(x uint64)

// Len encodes and writes a non-negative int value into the element bitstream.
func (w *Encoder) Len(x int)

// Int encodes and writes an int value into the element bitstream.
func (w *Encoder) Int(x int)

// Uint encodes and writes a uint value into the element bitstream.
func (w *Encoder) Uint(x uint)

// Reloc encodes and writes a relocation for the given (section,
// index) pair into the element bitstream.
//
// Note: Only the index is formally written into the element
// bitstream, so bitstream decoders must know from context which
// section an encoded relocation refers to.
func (w *Encoder) Reloc(r RelocKind, idx Index)

// Code encodes and writes a Code value into the element bitstream.
func (w *Encoder) Code(c Code)

// String encodes and writes a string value into the element
// bitstream.
//
// Internally, strings are deduplicated by adding them to the strings
// section (if not already present), and then writing a relocation
// into the element bitstream.
func (w *Encoder) String(s string)

// StringRef writes a reference to the given index, which must be a
// previously encoded string value.
func (w *Encoder) StringRef(idx Index)

// Strings encodes and writes a variable-length slice of strings into
// the element bitstream.
func (w *Encoder) Strings(ss []string)

// Value encodes and writes a constant.Value into the element
// bitstream.
func (w *Encoder) Value(val constant.Value)

// Version reports the version of the bitstream.
func (w *Encoder) Version() Version
