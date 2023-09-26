// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file exercises the import parser but also checks that
// some low-level packages do not have new dependencies added.

package build

// depsRules defines the expected dependencies between packages in
// the Go source tree. It is a statement of policy.
//
// DO NOT CHANGE THIS DATA TO FIX BUILDS.
// Existing packages should not have their constraints relaxed
// without prior discussion.
// Negative assertions should almost never be removed.
//
// The general syntax of a rule is:
//
//	a, b < c, d;
//
// which means c and d come after a and b in the partial order
// (that is, c and d can import a and b),
// but doesn't provide a relative order between a vs b or c vs d.
//
// The rules can chain together, as in:
//
//	e < f, g < h;
//
// which is equivalent to
//
//	e < f, g;
//	f, g < h;
//
// Except for the special bottom element "NONE", each name
// must appear exactly once on the right-hand side of a rule.
// That rule serves as the definition of the allowed dependencies
// for that name. The definition must appear before any uses
// of the name on the left-hand side of a rule. (That is, the
// rules themselves must be ordered according to the partial
// order, for easier reading by people.)
//
// Negative assertions double-check the partial order:
//
//	i !< j
//
// means that it must NOT be the case that i < j.
// Negative assertions may appear anywhere in the rules,
// even before i and j have been defined.
//
// Comments begin with #.
//
// All-caps names are pseudo-names for specific points
// in the dependency lattice.

// A depsParser parses the depsRules syntax described above.
