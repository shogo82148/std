// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objabi

// PkgSpecial indicates special build properties of a given runtime-related
// package.
type PkgSpecial struct {
	// Runtime indicates that this package is "runtime" or imported by
	// "runtime". This has several effects (which maybe should be split out):
	//
	// - Implicit allocation is disallowed.
	//
	// - Various runtime pragmas are enabled.
	//
	// - Optimizations are always enabled.
	//
	// This should be set for runtime and all packages it imports, and may be
	// set for additional packages.
	Runtime bool

	// NoInstrument indicates this package should not receive sanitizer
	// instrumentation. In many of these, instrumentation could cause infinite
	// recursion. This is all runtime packages, plus those that support the
	// sanitizers.
	NoInstrument bool

	// NoRaceFunc indicates functions in this package should not get
	// racefuncenter/racefuncexit instrumentation Memory accesses in these
	// packages are either uninteresting or will cause false positives.
	NoRaceFunc bool

	// AllowAsmABI indicates that assembly in this package is allowed to use ABI
	// selectors in symbol names. Generally this is needed for packages that
	// interact closely with the runtime package or have performance-critical
	// assembly.
	AllowAsmABI bool
}

// LookupPkgSpecial returns special build properties for the given package path.
func LookupPkgSpecial(pkgPath string) PkgSpecial
