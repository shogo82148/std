// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Identify mismatches between assembly files and Go func declarations.

package main

// 'kind' is a kind of assembly variable.
// The kinds 1, 2, 4, 8 stand for values of that size.

// These special kinds are not valid sizes.

// An asmArch describes assembly parameters for an architecture

// An asmFunc describes the expected variables for a function on a given architecture.

// An asmVar describes a single assembly variable.

// Common architecture word sizes and alignments.

// A component is an assembly-addressable component of a composite type,
// or a composite type itself.
