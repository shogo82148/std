// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// Unicode table generator.
// Data read from the web.

package main

// UnicodeData.txt has form:
//
//	0037;DIGIT SEVEN;Nd;0;EN;;7;7;7;N;;;;;
//	007A;LATIN SMALL LETTER Z;Ll;0;L;;;;;N;;;005A;;005A
//
// See http://www.unicode.org/reports/tr44/ for a full explanation
// The fields:
const (
	FCodePoint = iota
	FName
	FGeneralCategory
	FCanonicalCombiningClass
	FBidiClass
	FDecompositionTypeAndMapping
	FNumericType
	FNumericDigit
	FNumericValue
	FBidiMirrored
	FUnicode1Name
	FISOComment
	FSimpleUppercaseMapping
	FSimpleLowercaseMapping
	FSimpleTitlecaseMapping
	NumField

	MaxChar = 0x10FFFF
)

// This contains only the properties we're interested in.
type Char struct {
	field     []string
	codePoint rune
	category  string
	upperCase rune
	lowerCase rune
	titleCase rune
	foldCase  rune
	caseOrbit rune
}

type Script struct {
	lo, hi uint32
	script string
}

// In UnicodeData.txt, some ranges are marked like this:
//
//	3400;<CJK Ideograph Extension A, First>;Lo;0;L;;;;;N;;;;;
//	4DB5;<CJK Ideograph Extension A, Last>;Lo;0;L;;;;;N;;;;;
//
// parseCategory returns a state variable indicating the weirdness.
type State int

const (
	SNormal State = iota
	SFirst
	SLast
	SMissing
)

type Op func(code rune) bool

const (
	CaseUpper = 1 << iota
	CaseLower
	CaseTitle
	CaseNone    = 0
	CaseMissing = -1
)
