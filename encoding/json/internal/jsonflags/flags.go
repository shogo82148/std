// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

// jsonflags implements all the optional boolean flags.
// These flags are shared across both "json", "jsontext", and "jsonopts".
package jsonflags

import "github.com/shogo82148/std/encoding/json/internal"

// Bools represents zero or more boolean flags, all set to true or false.
// The least-significant bit is the boolean value of all flags in the set.
// The remaining bits identify which particular flags.
//
// In common usage, this is OR'd with 0 or 1. For example:
//   - (AllowInvalidUTF8 | 0) means "AllowInvalidUTF8 is false"
//   - (Multiline | Indent | 1) means "Multiline and Indent are true"
type Bools uint64

func (Bools) JSONOptions(internal.NotForPublicUse)

const (
	// AllFlags is the set of all flags.
	AllFlags = AllCoderFlags | AllArshalV2Flags | AllArshalV1Flags

	// AllCoderFlags is the set of all encoder/decoder flags.
	AllCoderFlags = (maxCoderFlag - 1) - initFlag

	// AllArshalV2Flags is the set of all v2 marshal/unmarshal flags.
	AllArshalV2Flags = (maxArshalV2Flag - 1) - (maxCoderFlag - 1)

	// AllArshalV1Flags is the set of all v1 marshal/unmarshal flags.
	AllArshalV1Flags = (maxArshalV1Flag - 1) - (maxArshalV2Flag - 1)

	// NonBooleanFlags is the set of non-boolean flags,
	// where the value is some other concrete Go type.
	// The value of the flag is stored within jsonopts.Struct.
	NonBooleanFlags = 0 |
		Indent |
		IndentPrefix |
		ByteLimit |
		DepthLimit |
		Marshalers |
		Unmarshalers

	// DefaultV1Flags is the set of booleans flags that default to true under
	// v1 semantics. None of the non-boolean flags differ between v1 and v2.
	DefaultV1Flags = 0 |
		AllowDuplicateNames |
		AllowInvalidUTF8 |
		EscapeForHTML |
		EscapeForJS |
		EscapeInvalidUTF8 |
		PreserveRawStrings |
		Deterministic |
		FormatNilMapAsNull |
		FormatNilSliceAsNull |
		MatchCaseInsensitiveNames |
		CallMethodsWithLegacySemantics |
		FormatBytesWithLegacySemantics |
		FormatTimeWithLegacySemantics |
		MatchCaseSensitiveDelimiter |
		MergeWithLegacySemantics |
		OmitEmptyWithLegacyDefinition |
		ReportErrorsWithLegacySemantics |
		StringifyWithLegacySemantics |
		UnmarshalArrayFromAnyLength

	// AnyWhitespace reports whether the encoded output might have any whitespace.
	AnyWhitespace = Multiline | SpaceAfterColon | SpaceAfterComma

	// WhitespaceFlags is the set of flags related to whitespace formatting.
	// In contrast to AnyWhitespace, this includes Indent and IndentPrefix
	// as those settings take no effect if Multiline is false.
	WhitespaceFlags = AnyWhitespace | Indent | IndentPrefix

	// AnyEscape is the set of flags related to escaping in a JSON string.
	AnyEscape = EscapeForHTML | EscapeForJS | EscapeInvalidUTF8

	// CanonicalizeNumbers is the set of flags related to raw number canonicalization.
	CanonicalizeNumbers = CanonicalizeRawInts | CanonicalizeRawFloats
)

// Encoder and decoder flags.
const (
	AllowDuplicateNames
	AllowInvalidUTF8
	WithinArshalCall
	OmitTopLevelNewline
	PreserveRawStrings
	CanonicalizeRawInts
	CanonicalizeRawFloats
	ReorderRawObjects
	EscapeForHTML
	EscapeForJS
	EscapeInvalidUTF8
	Multiline
	SpaceAfterColon
	SpaceAfterComma
	Indent
	IndentPrefix
	ByteLimit
	DepthLimit
)

// Marshal and Unmarshal flags (for v2).
const (
	_ Bools = (maxCoderFlag >> 1) << iota

	StringifyNumbers
	Deterministic
	FormatNilMapAsNull
	FormatNilSliceAsNull
	OmitZeroStructFields
	MatchCaseInsensitiveNames
	DiscardUnknownMembers
	RejectUnknownMembers
	Marshalers
	Unmarshalers
)

// Marshal and Unmarshal flags (for v1).
const (
	_ Bools = (maxArshalV2Flag >> 1) << iota

	CallMethodsWithLegacySemantics
	FormatBytesWithLegacySemantics
	FormatTimeWithLegacySemantics
	MatchCaseSensitiveDelimiter
	MergeWithLegacySemantics
	OmitEmptyWithLegacyDefinition
	ReportErrorsWithLegacySemantics
	StringifyWithLegacySemantics
	StringifyBoolsAndStrings
	UnmarshalAnyWithRawNumber
	UnmarshalArrayFromAnyLength
)

// Flags is a set of boolean flags.
// If the presence bit is zero, then the value bit must also be zero.
// The least-significant bit of both fields is always zero.
//
// Unlike Bools, which can represent a set of bools that are all true or false,
// Flags represents a set of bools, each individually may be true or false.
type Flags struct{ Presence, Values uint64 }

// Join joins two sets of flags such that the latter takes precedence.
func (dst *Flags) Join(src Flags)

// Set sets both the presence and value for the provided bool (or set of bools).
func (fs *Flags) Set(f Bools)

// Get reports whether the bool (or any of the bools) is true.
// This is generally only used with a singular bool.
// The value bit of f (i.e., the LSB) is ignored.
func (fs Flags) Get(f Bools) bool

// Has reports whether the bool (or any of the bools) is set.
// The value bit of f (i.e., the LSB) is ignored.
func (fs Flags) Has(f Bools) bool

// Clear clears both the presence and value for the provided bool or bools.
// The value bit of f (i.e., the LSB) is ignored.
func (fs *Flags) Clear(f Bools)
