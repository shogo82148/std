// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gover implements support for Go toolchain versions like 1.21.0 and 1.21rc1.
// (For historical reasons, Go does not use semver for its toolchains.)
// This package provides the same basic analysis that golang.org/x/mod/semver does for semver.
//
// The go/version package should be imported instead of this one when possible.
// Note that this package works on "1.21" while go/version works on "go1.21".
package gover

// A Version is a parsed Go version: major[.Minor[.Patch]][kind[pre]]
// The numbers are the original decimal strings to avoid integer overflows
// and since there is very little actual math. (Probably overflow doesn't matter in practice,
// but at the time this code was written, there was an existing test that used
// go1.99999999999, which does not fit in an int on 32-bit platforms.
// The "big decimal" representation avoids the problem entirely.)
type Version struct {
	Major string
	Minor string
	Patch string
	Kind  string
	Pre   string
}

// Compare returns -1, 0, or +1 depending on whether
// x < y, x == y, or x > y, interpreted as toolchain versions.
// The versions x and y must not begin with a "go" prefix: just "1.21" not "go1.21".
// Malformed versions compare less than well-formed versions and equal to each other.
// The language version "1.21" compares less than the release candidate and eventual releases "1.21rc1" and "1.21.0".
func Compare(x, y string) int

// Max returns the maximum of x and y interpreted as toolchain versions,
// compared using Compare.
// If x and y compare equal, Max returns x.
func Max(x, y string) string

// IsLang reports whether v denotes the overall Go language version
// and not a specific release. Starting with the Go 1.21 release, "1.x" denotes
// the overall language version; the first release is "1.x.0".
// The distinction is important because the relative ordering is
//
//	1.21 < 1.21rc1 < 1.21.0
//
// meaning that Go 1.21rc1 and Go 1.21.0 will both handle go.mod files that
// say "go 1.21", but Go 1.21rc1 will not handle files that say "go 1.21.0".
func IsLang(x string) bool

// Lang returns the Go language version. For example, Lang("1.2.3") == "1.2".
func Lang(x string) string

// IsValid reports whether the version x is valid.
func IsValid(x string) bool

// Parse parses the Go version string x into a version.
// It returns the zero version if x is malformed.
func Parse(x string) Version

// CmpInt returns cmp.Compare(x, y) interpreting x and y as decimal numbers.
// (Copied from golang.org/x/mod/semver's compareInt.)
func CmpInt(x, y string) int

// DecInt returns the decimal string decremented by 1, or the empty string
// if the decimal is all zeroes.
// (Copied from golang.org/x/mod/module's decDecimal.)
func DecInt(decimal string) string
