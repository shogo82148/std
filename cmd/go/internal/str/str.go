// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package str provides string manipulation utilities.
package str

// StringList flattens its arguments into a single []string.
// Each argument in args must have type string or []string.
func StringList(args ...any) []string

// ToFold returns a string with the property that
//
//	strings.EqualFold(s, t) iff ToFold(s) == ToFold(t)
//
// This lets us test a large set of strings for fold-equivalent
// duplicates without making a quadratic number of calls
// to EqualFold. Note that strings.ToUpper and strings.ToLower
// do not have the desired property in some corner cases.
func ToFold(s string) string

// FoldDup reports a pair of strings from the list that are
// equal according to strings.EqualFold.
// It returns "", "" if there are no such strings.
func FoldDup(list []string) (string, string)

// Uniq removes consecutive duplicate strings from ss.
func Uniq(ss *[]string)
