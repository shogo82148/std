// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sgutil

// CompareNatural performs a "natural sort" comparison of two strings.
// It compares non-digit sections lexicographically and digit sections
// numerically.  In the case of string-unequal "equal" strings like
// "a01b" and "a1b", strings.Compare breaks the tie.
//
// It returns:
//
//	-1 if s1 < s2
//	 0 if s1 == s2
//	+1 if s1 > s2
func CompareNatural(s1, s2 string) int
