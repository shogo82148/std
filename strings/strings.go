// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package strings implements simple functions to manipulate UTF-8 encoded strings.
//
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
package strings

import (
	"github.com/shogo82148/std/unicode"
)

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func Count(s, substr string) int

// Contains reports whether substr is within s.
func Contains(s, substr string) bool

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s, chars string) bool

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s string, r rune) bool

// ContainsFunc reports whether any Unicode code points r within s satisfy f(r).
func ContainsFunc(s string, f func(rune) bool) bool

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s, substr string) int

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByte(s string, c byte) int

// IndexRune returns the index of the first instance of the Unicode code point
// r, or -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any
// invalid UTF-8 byte sequence.
func IndexRune(s string, r rune) int

// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
func IndexAny(s, chars string) int

// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
func LastIndexAny(s, chars string) int

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexByte(s string, c byte) int

// SplitN slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
//
// To split around the first instance of a separator, see Cut.
func SplitN(s, sep string, n int) []string

// SplitAfterN slices s into substrings after each instance of sep and
// returns a slice of those substrings.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for SplitAfter.
func SplitAfterN(s, sep string, n int) []string

// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to SplitN with a count of -1.
//
// To split around the first instance of a separator, see Cut.
func Split(s, sep string) []string

// SplitAfter slices s into all substrings after each instance of sep and
// returns a slice of those substrings.
//
// If s does not contain sep and sep is not empty, SplitAfter returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
// both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to SplitAfterN with a count of -1.
func SplitAfter(s, sep string) []string

// Fields splits the string s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
// empty slice if s contains only white space.
func Fields(s string) []string

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// string is empty, an empty slice is returned.
//
// FieldsFunc makes no guarantees about the order in which it calls f(c)
// and assumes that f always returns the same value for a given c.
func FieldsFunc(s string, f func(rune) bool) []string

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func Join(elems []string, sep string) string

// HasPrefix reports whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool

// HasSuffix reports whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func Map(mapping func(rune) rune, s string) string

// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if the result of (len(s) * count)
// overflows.
func Repeat(s string, count int) string

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string

// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower(s string) string

// ToTitle returns a copy of the string s with all Unicode letters mapped to
// their Unicode title case.
func ToTitle(s string) string

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
// upper case using the case mapping specified by c.
func ToUpperSpecial(c unicode.SpecialCase, s string) string

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
// lower case using the case mapping specified by c.
func ToLowerSpecial(c unicode.SpecialCase, s string) string

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
// Unicode title case, giving priority to the special casing rules.
func ToTitleSpecial(c unicode.SpecialCase, s string) string

// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
// replaced by the replacement string, which may be empty.
func ToValidUTF8(s, replacement string) string

// Title returns a copy of the string s with all Unicode letters that begin words
// mapped to their Unicode title case.
//
// Deprecated: The rule Title uses for word boundaries does not handle Unicode
// punctuation properly. Use golang.org/x/text/cases instead.
func Title(s string) string

// TrimLeftFunc returns a slice of the string s with all leading
// Unicode code points c satisfying f(c) removed.
func TrimLeftFunc(s string, f func(rune) bool) string

// TrimRightFunc returns a slice of the string s with all trailing
// Unicode code points c satisfying f(c) removed.
func TrimRightFunc(s string, f func(rune) bool) string

// TrimFunc returns a slice of the string s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
func TrimFunc(s string, f func(rune) bool) string

// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func IndexFunc(s string, f func(rune) bool) int

// LastIndexFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
func LastIndexFunc(s string, f func(rune) bool) int

// asciiSet is a 32-byte value, where each bit represents the presence of a
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
// starting with the least-significant bit of the lowest word to the
// most-significant bit of the highest word, map to the full range of all
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
// ensuring that any non-ASCII character will be reported as not in the set.
// This allocates a total of 32 bytes even though the upper half
// is unused to avoid bounds checks in asciiSet.contains.

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
func Trim(s, cutset string) string

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
//
// To remove a prefix, use TrimPrefix instead.
func TrimLeft(s, cutset string) string

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use TrimSuffix instead.
func TrimRight(s, cutset string) string

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpace(s string) string

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix(s, prefix string) string

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s, suffix string) string

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new string, n int) string

// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
func ReplaceAll(s, old, new string) string

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under simple Unicode case-folding, which is a more general
// form of case-insensitivity.
func EqualFold(s, t string) bool

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func Cut(s, sep string) (before, after string, found bool)

// CutPrefix returns s without the provided leading prefix string
// and reports whether it found the prefix.
// If s doesn't start with prefix, CutPrefix returns s, false.
// If prefix is the empty string, CutPrefix returns s, true.
func CutPrefix(s, prefix string) (after string, found bool)

// CutSuffix returns s without the provided ending suffix string
// and reports whether it found the suffix.
// If s doesn't end with suffix, CutSuffix returns s, false.
// If suffix is the empty string, CutSuffix returns s, true.
func CutSuffix(s, suffix string) (before string, found bool)
