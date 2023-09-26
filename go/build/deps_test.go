// Copyright 2022 The Go Authors. All rights reserved.
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
// "a < b" means package b can import package a.
//
// See `go doc internal/dag' for the full syntax.
//
// All-caps names are pseudo-names for specific points
// in the dependency lattice.
