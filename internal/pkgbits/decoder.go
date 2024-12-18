// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgbits

import (
	"github.com/shogo82148/std/go/constant"
	"github.com/shogo82148/std/strings"
)

// A PkgDecoder provides methods for decoding a package's Unified IR
// export data.
type PkgDecoder struct {
	// version is the file format version.
	version Version

	// sync indicates whether the file uses sync markers.
	sync bool

	// pkgPath is the package path for the package to be decoded.
	//
	// TODO(mdempsky): Remove; unneeded since CL 391014.
	pkgPath string

	// elemData is the full data payload of the encoded package.
	// Elements are densely and contiguously packed together.
	//
	// The last 8 bytes of elemData are the package fingerprint.
	elemData string

	// elemEnds stores the byte-offset end positions of element
	// bitstreams within elemData.
	//
	// For example, element I's bitstream data starts at elemEnds[I-1]
	// (or 0, if I==0) and ends at elemEnds[I].
	//
	// Note: elemEnds is indexed by absolute indices, not
	// section-relative indices.
	elemEnds []uint32

	// elemEndsEnds stores the index-offset end positions of relocation
	// sections within elemEnds.
	//
	// For example, section K's end positions start at elemEndsEnds[K-1]
	// (or 0, if K==0) and end at elemEndsEnds[K].
	elemEndsEnds [numRelocs]uint32

	scratchRelocEnt []RelocEnt
}

// PkgPath returns the package path for the package
//
// TODO(mdempsky): Remove; unneeded since CL 391014.
func (pr *PkgDecoder) PkgPath() string

// SyncMarkers reports whether pr uses sync markers.
func (pr *PkgDecoder) SyncMarkers() bool

// NewPkgDecoder returns a PkgDecoder initialized to read the Unified
// IR export data from input. pkgPath is the package path for the
// compilation unit that produced the export data.
func NewPkgDecoder(pkgPath, input string) PkgDecoder

// NumElems returns the number of elements in section k.
func (pr *PkgDecoder) NumElems(k RelocKind) int

// TotalElems returns the total number of elements across all sections.
func (pr *PkgDecoder) TotalElems() int

// Fingerprint returns the package fingerprint.
func (pr *PkgDecoder) Fingerprint() [8]byte

// AbsIdx returns the absolute index for the given (section, index)
// pair.
func (pr *PkgDecoder) AbsIdx(k RelocKind, idx Index) int

// DataIdx returns the raw element bitstream for the given (section,
// index) pair.
func (pr *PkgDecoder) DataIdx(k RelocKind, idx Index) string

// StringIdx returns the string value for the given string index.
func (pr *PkgDecoder) StringIdx(idx Index) string

// NewDecoder returns a Decoder for the given (section, index) pair,
// and decodes the given SyncMarker from the element bitstream.
func (pr *PkgDecoder) NewDecoder(k RelocKind, idx Index, marker SyncMarker) Decoder

// TempDecoder returns a Decoder for the given (section, index) pair,
// and decodes the given SyncMarker from the element bitstream.
// If possible the Decoder should be RetireDecoder'd when it is no longer
// needed, this will avoid heap allocations.
func (pr *PkgDecoder) TempDecoder(k RelocKind, idx Index, marker SyncMarker) Decoder

func (pr *PkgDecoder) RetireDecoder(d *Decoder)

// NewDecoderRaw returns a Decoder for the given (section, index) pair.
//
// Most callers should use NewDecoder instead.
func (pr *PkgDecoder) NewDecoderRaw(k RelocKind, idx Index) Decoder

func (pr *PkgDecoder) TempDecoderRaw(k RelocKind, idx Index) Decoder

// A Decoder provides methods for decoding an individual element's
// bitstream data.
type Decoder struct {
	common *PkgDecoder

	Relocs []RelocEnt
	Data   strings.Reader

	k   RelocKind
	Idx Index
}

// Sync decodes a sync marker from the element bitstream and asserts
// that it matches the expected marker.
//
// If EnableSync is false, then Sync is a no-op.
func (r *Decoder) Sync(mWant SyncMarker)

// Bool decodes and returns a bool value from the element bitstream.
func (r *Decoder) Bool() bool

// Int64 decodes and returns an int64 value from the element bitstream.
func (r *Decoder) Int64() int64

// Uint64 decodes and returns a uint64 value from the element bitstream.
func (r *Decoder) Uint64() uint64

// Len decodes and returns a non-negative int value from the element bitstream.
func (r *Decoder) Len() int

// Int decodes and returns an int value from the element bitstream.
func (r *Decoder) Int() int

// Uint decodes and returns a uint value from the element bitstream.
func (r *Decoder) Uint() uint

// Code decodes a Code value from the element bitstream and returns
// its ordinal value. It's the caller's responsibility to convert the
// result to an appropriate Code type.
//
// TODO(mdempsky): Ideally this method would have signature "Code[T
// Code] T" instead, but we don't allow generic methods and the
// compiler can't depend on generics yet anyway.
func (r *Decoder) Code(mark SyncMarker) int

// Reloc decodes a relocation of expected section k from the element
// bitstream and returns an index to the referenced element.
func (r *Decoder) Reloc(k RelocKind) Index

// String decodes and returns a string value from the element
// bitstream.
func (r *Decoder) String() string

// Strings decodes and returns a variable-length slice of strings from
// the element bitstream.
func (r *Decoder) Strings() []string

// Value decodes and returns a constant.Value from the element
// bitstream.
func (r *Decoder) Value() constant.Value

// PeekPkgPath returns the package path for the specified package
// index.
func (pr *PkgDecoder) PeekPkgPath(idx Index) string

// PeekObj returns the package path, object name, and CodeObj for the
// specified object index.
func (pr *PkgDecoder) PeekObj(idx Index) (string, string, CodeObj)

// Version reports the version of the bitstream.
func (w *Decoder) Version() Version
