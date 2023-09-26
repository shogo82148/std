// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bytes implements functions for the manipulation of byte slices.
// It is analogous to the facilities of the strings package.
package bytes

import (
	"github.com/shogo82148/std/unicode"
)

// Compare returns an integer comparing two byte slices lexicographically.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
// A nil argument is equivalent to an empty slice.
func Compare(a, b []byte) int

// Count counts the number of non-overlapping instances of sep in s.
func Count(s, sep []byte) int

// Contains returns whether subslice is within b.
func Contains(b, subslice []byte) bool

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
func Index(s, sep []byte) int

// LastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.
func LastIndex(s, sep []byte) int

// IndexRune interprets s as a sequence of UTF-8-encoded Unicode code points.
// It returns the byte index of the first occurrence in s of the given rune.
// It returns -1 if rune is not present in s.
func IndexRune(s []byte, r rune) int

// IndexAny interprets s as a sequence of UTF-8-encoded Unicode code points.
// It returns the byte index of the first occurrence in s of any of the Unicode
// code points in chars.  It returns -1 if chars is empty or if there is no code
// point in common.
func IndexAny(s []byte, chars string) int

// LastIndexAny interprets s as a sequence of UTF-8-encoded Unicode code
// points.  It returns the byte index of the last occurrence in s of any of
// the Unicode code points in chars.  It returns -1 if chars is empty or if
// there is no code point in common.
func LastIndexAny(s []byte, chars string) int

// SplitN slices s into subslices separated by sep and returns a slice of
// the subslices between those separators.
// If sep is empty, SplitN splits after each UTF-8 sequence.
// The count determines the number of subslices to return:
//
//	n > 0: at most n subslices; the last subslice will be the unsplit remainder.
//	n == 0: the result is nil (zero subslices)
//	n < 0: all subslices
func SplitN(s, sep []byte, n int) [][]byte

// SplitAfterN slices s into subslices after each instance of sep and
// returns a slice of those subslices.
// If sep is empty, SplitAfterN splits after each UTF-8 sequence.
// The count determines the number of subslices to return:
//
//	n > 0: at most n subslices; the last subslice will be the unsplit remainder.
//	n == 0: the result is nil (zero subslices)
//	n < 0: all subslices
func SplitAfterN(s, sep []byte, n int) [][]byte

// Split slices s into all subslices separated by sep and returns a slice of
// the subslices between those separators.
// If sep is empty, Split splits after each UTF-8 sequence.
// It is equivalent to SplitN with a count of -1.
func Split(s, sep []byte) [][]byte

// SplitAfter slices s into all subslices after each instance of sep and
// returns a slice of those subslices.
// If sep is empty, SplitAfter splits after each UTF-8 sequence.
// It is equivalent to SplitAfterN with a count of -1.
func SplitAfter(s, sep []byte) [][]byte

// Fields splits the slice s around each instance of one or more consecutive white space
// characters, returning a slice of subslices of s or an empty list if s contains only white space.
func Fields(s []byte) [][]byte

// FieldsFunc interprets s as a sequence of UTF-8-encoded Unicode code points.
// It splits the slice s at each run of code points c satisfying f(c) and
// returns a slice of subslices of s.  If no code points in s satisfy f(c), an
// empty slice is returned.
func FieldsFunc(s []byte, f func(rune) bool) [][]byte

// Join concatenates the elements of s to create a new byte slice. The separator
// sep is placed between elements in the resulting slice.
func Join(s [][]byte, sep []byte) []byte

// HasPrefix tests whether the byte slice s begins with prefix.
func HasPrefix(s, prefix []byte) bool

// HasSuffix tests whether the byte slice s ends with suffix.
func HasSuffix(s, suffix []byte) bool

// Map returns a copy of the byte slice s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.  The characters in s and the
// output are interpreted as UTF-8-encoded Unicode code points.
func Map(mapping func(r rune) rune, s []byte) []byte

// Repeat returns a new byte slice consisting of count copies of b.
func Repeat(b []byte, count int) []byte

// ToUpper returns a copy of the byte slice s with all Unicode letters mapped to their upper case.
func ToUpper(s []byte) []byte

// ToLower returns a copy of the byte slice s with all Unicode letters mapped to their lower case.
func ToLower(s []byte) []byte

// ToTitle returns a copy of the byte slice s with all Unicode letters mapped to their title case.
func ToTitle(s []byte) []byte

// ToUpperSpecial returns a copy of the byte slice s with all Unicode letters mapped to their
// upper case, giving priority to the special casing rules.
func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte

// ToLowerSpecial returns a copy of the byte slice s with all Unicode letters mapped to their
// lower case, giving priority to the special casing rules.
func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte

// ToTitleSpecial returns a copy of the byte slice s with all Unicode letters mapped to their
// title case, giving priority to the special casing rules.
func ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte

// Title returns a copy of s with all Unicode letters that begin words
// mapped to their title case.
//
// BUG: The rule Title uses for word boundaries does not handle Unicode punctuation properly.
func Title(s []byte) []byte

// TrimLeftFunc returns a subslice of s by slicing off all leading UTF-8-encoded
// Unicode code points c that satisfy f(c).
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte

// TrimRightFunc returns a subslice of s by slicing off all trailing UTF-8
// encoded Unicode code points c that satisfy f(c).
func TrimRightFunc(s []byte, f func(r rune) bool) []byte

// TrimFunc returns a subslice of s by slicing off all leading and trailing
// UTF-8-encoded Unicode code points c that satisfy f(c).
func TrimFunc(s []byte, f func(r rune) bool) []byte

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix(s, prefix []byte) []byte

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s, suffix []byte) []byte

// IndexFunc interprets s as a sequence of UTF-8-encoded Unicode code points.
// It returns the byte index in s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func IndexFunc(s []byte, f func(r rune) bool) int

// LastIndexFunc interprets s as a sequence of UTF-8-encoded Unicode code points.
// It returns the byte index in s of the last Unicode
// code point satisfying f(c), or -1 if none do.
func LastIndexFunc(s []byte, f func(r rune) bool) int

// Trim returns a subslice of s by slicing off all leading and
// trailing UTF-8-encoded Unicode code points contained in cutset.
func Trim(s []byte, cutset string) []byte

// TrimLeft returns a subslice of s by slicing off all leading
// UTF-8-encoded Unicode code points contained in cutset.
func TrimLeft(s []byte, cutset string) []byte

// TrimRight returns a subslice of s by slicing off all trailing
// UTF-8-encoded Unicode code points that are contained in cutset.
func TrimRight(s []byte, cutset string) []byte

// TrimSpace returns a subslice of s by slicing off all leading and
// trailing white space, as defined by Unicode.
func TrimSpace(s []byte) []byte

// Runes returns a slice of runes (Unicode code points) equivalent to s.
func Runes(s []byte) []rune

// Replace returns a copy of the slice s with the first n
// non-overlapping instances of old replaced by new.
// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new []byte, n int) []byte

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding.
func EqualFold(s, t []byte) bool
