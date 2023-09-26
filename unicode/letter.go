// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package unicode provides data and functions to test some properties of
// Unicode code points.
package unicode

const (
	MaxRune         = '\U0010FFFF'
	ReplacementChar = '\uFFFD'
	MaxASCII        = '\u007F'
	MaxLatin1       = '\u00FF'
)

// RangeTable defines a set of Unicode code points by listing the ranges of
// code points within the set. The ranges are listed in two slices
// to save space: a slice of 16-bit ranges and a slice of 32-bit ranges.
// The two slices must be in sorted order and non-overlapping.
// Also, R32 should contain only values >= 0x10000 (1<<16).
type RangeTable struct {
	R16         []Range16
	R32         []Range32
	LatinOffset int
}

// Range16 represents of a range of 16-bit Unicode code points.  The range runs from Lo to Hi
// inclusive and has the specified stride.
type Range16 struct {
	Lo     uint16
	Hi     uint16
	Stride uint16
}

// Range32 represents of a range of Unicode code points and is used when one or
// more of the values will not fit in 16 bits.  The range runs from Lo to Hi
// inclusive and has the specified stride. Lo and Hi must always be >= 1<<16.
type Range32 struct {
	Lo     uint32
	Hi     uint32
	Stride uint32
}

// CaseRange represents a range of Unicode code points for simple (one
// code point to one code point) case conversion.
// The range runs from Lo to Hi inclusive, with a fixed stride of 1.  Deltas
// are the number to add to the code point to reach the code point for a
// different case for that character.  They may be negative.  If zero, it
// means the character is in the corresponding case. There is a special
// case representing sequences of alternating corresponding Upper and Lower
// pairs.  It appears with a fixed Delta of
//
//	{UpperLower, UpperLower, UpperLower}
//
// The constant UpperLower has an otherwise impossible delta value.
type CaseRange struct {
	Lo    uint32
	Hi    uint32
	Delta d
}

// SpecialCase represents language-specific case mappings such as Turkish.
// Methods of SpecialCase customize (by overriding) the standard mappings.
type SpecialCase []CaseRange

// Indices into the Delta arrays inside CaseRanges for case mapping.
const (
	UpperCase = iota
	LowerCase
	TitleCase
	MaxCase
)

// If the Delta field of a CaseRange is UpperLower, it means
// this CaseRange represents a sequence of the form (say)
// Upper Lower Upper Lower.
const (
	UpperLower = MaxRune + 1
)

// linearMax is the maximum size table for linear search for non-Latin1 rune.
// Derived by running 'go test -calibrate'.

// Is reports whether the rune is in the specified table of ranges.
func Is(rangeTab *RangeTable, r rune) bool

// IsUpper reports whether the rune is an upper case letter.
func IsUpper(r rune) bool

// IsLower reports whether the rune is a lower case letter.
func IsLower(r rune) bool

// IsTitle reports whether the rune is a title case letter.
func IsTitle(r rune) bool

// To maps the rune to the specified case: UpperCase, LowerCase, or TitleCase.
func To(_case int, r rune) rune

// ToUpper maps the rune to upper case.
func ToUpper(r rune) rune

// ToLower maps the rune to lower case.
func ToLower(r rune) rune

// ToTitle maps the rune to title case.
func ToTitle(r rune) rune

// ToUpper maps the rune to upper case giving priority to the special mapping.
func (special SpecialCase) ToUpper(r rune) rune

// ToTitle maps the rune to title case giving priority to the special mapping.
func (special SpecialCase) ToTitle(r rune) rune

// ToLower maps the rune to lower case giving priority to the special mapping.
func (special SpecialCase) ToLower(r rune) rune

// caseOrbit is defined in tables.go as []foldPair.  Right now all the
// entries fit in uint16, so use uint16.  If that changes, compilation
// will fail (the constants in the composite literal will not fit in uint16)
// and the types here can change to uint32.

// SimpleFold iterates over Unicode code points equivalent under
// the Unicode-defined simple case folding.  Among the code points
// equivalent to rune (including rune itself), SimpleFold returns the
// smallest rune > r if one exists, or else the smallest rune >= 0.
//
// For example:
//
//	SimpleFold('A') = 'a'
//	SimpleFold('a') = 'A'
//
//	SimpleFold('K') = 'k'
//	SimpleFold('k') = '\u212A' (Kelvin symbol, K)
//	SimpleFold('\u212A') = 'K'
//
//	SimpleFold('1') = '1'
func SimpleFold(r rune) rune
