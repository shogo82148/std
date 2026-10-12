// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package specdoc validates exported SIMD declarations against package
// simd/internal/spec and manages the composition and injection of spec-derived
// documentation into Go source files.
//
// Generators and tools pass Go source buffers through [Fill] before writing
// them to disk. Comparing generated or hand-written declarations against spec
// ensures that the API implementation and the specification agree on types and
// parameter names, and allows shared documentation to be maintained in a single
// place.
//
// # Doc-Ownership Convention
//
// Doc comments follow a three-part structural convention:
//
// 1. Spec-owned paragraphs: Describe the behavioral semantics of the operation.
// These are managed by specdoc and replaced with documentation from spec when a
// matching spec entry exists.
//
// 2. Implementation notes: Describe target-specific implementation details,
// such as assembly instructions, required CPU features, emulation notes,
// performance caveats, or deprecations. An implementation note is any paragraph
// whose first line begins with one of a fixed set of prefixes: "Asm:", "CPU
// Feature:", "Emulated", "Deprecated:", "Performance:", or "Note:".
//
// 3. Directive comments: Tool and compiler directives, such as //go:noescape.
//
// These sections must appear strictly in this order: spec-owned paragraphs
// strictly before implementation notes, which must appear strictly before
// directive comments.
//
// # Rewriting and Idempotence
//
// When rewriting a declaration, [Fill] replaces existing spec-owned paragraphs
// with the corresponding spec documentation (unless [Options.NoFillDoc] is set)
// and fills unnamed parameters and results with canonical names from spec
// (unless [Options.NoFillNames] is set), while preserving existing
// implementation notes and directives unchanged. Because spec documentation is
// guaranteed never to begin with an implementation note prefix, the rewrite
// operation is idempotent: running [Fill] repeatedly on rewritten source
// produces byte-identical output.
package specdoc

import (
	"github.com/shogo82148/std/simd/archsimd/_gen/specgen"
)

// Options configures the behavior of [Fill].
type Options struct {
	// RejectUnknown controls whether an exported declaration with no matching
	// spec entry is reported as an error.
	//
	// When false, un-specced declarations are tolerated and omitted from
	// signature and doc checks, and passed through the rewrite verbatim. When
	// true, any exported declaration lacking a spec definition causes
	// generation to fail, preventing new API operations from escaping spec
	// coverage.
	RejectUnknown bool

	// AllowDocRewrite controls whether an exported declaration containing
	// existing spec-owned doc comments is tolerated.
	//
	// When false (default), any spec-owned doc comment found in the input is
	// reported as an error (recorded in [Report.UnexpectedDocs]). This is used
	// by generators to ensure they do not emit doc comments that would be
	// overwritten by spec. Implementation notes and directives are still
	// permitted.
	//
	// When true, existing spec-owned doc comments are permitted and rewritten.
	// This is used when processing hand-written source files.
	AllowDocRewrite bool

	// AllowNameMismatches controls whether parameter and result name mismatches
	// between an AST declaration and spec are tolerated.
	//
	// When false (default), name mismatches are reported as errors.
	// When true, name mismatches are omitted from the report.
	AllowNameMismatches bool

	// NoFillDoc disables replacing doc comments with spec-derived comments.
	//
	// When false, Fill updates spec-owned doc paragraphs.
	NoFillDoc bool

	// NoFillNames disables filling unnamed parameters and results with names
	// from spec.
	//
	// When false, Fill fills unnamed parameters and results in exported
	// declarations from the matching spec function.
	NoFillNames bool

	// Filename optionally provides the name of the file being processed,
	// used for positions in report diagnostics.
	Filename string
}

// Filler returns a post-processing hook that runs [Fill] on generated Go source files.
// The returned function satisfies [gentools.PostProcessor].
func Filler(idx *specgen.Index, opts Options) func(relPath string, isGo bool, content []byte) ([]byte, error)

// Fill parses Go source code, extracts exported declarations, verifies their
// signatures and doc comments against spec, and returns the modified source.
// If any verification checks fail, Fill returns a [*Report] error containing
// all findings.
func Fill(src []byte, spec *specgen.Index, opts Options) ([]byte, error)
