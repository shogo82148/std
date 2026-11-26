// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package load

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/iter"

	"github.com/shogo82148/std/cmd/go/internal/modload"
)

// AnalysisToolDriver is the import path of the driver package linked
// into the synthesized analysis tool.
const AnalysisToolDriver = "internal/analysis/gotestchecker"

// An AnalysisVar identifies an exported analyzer or suite variable.
type AnalysisVar struct {
	PkgPath string
	Name    string
}

func (v AnalysisVar) String() string

// An AnalysisRef is a reference to an analyzer variable (pkg.Var) or
// to members of a suite variable (pkg.Var[*] or pkg.Var[A,B]), or (if
// Negative) a negation of analyzers of a given Name.
type AnalysisRef struct {
	Negative bool
	Var      AnalysisVar
	Members  []string
	Name     string
}

// Ref returns the reference as written, e.g. "pkgpath.Var[A,B]",
// or the quoted analyzer name of a negation by name.
func (a *AnalysisRef) Ref() string

// ByName reports whether the reference is a negation that names
// analyzers by Analyzer.Name rather than by reference. Such a negation
// is applied only in the dynamic stage, to every analyzer of that name;
// it is not linked into the tool.
func (a *AnalysisRef) ByName() bool

// IsSuite reports whether the reference is to a suite (pkg.Var[...]).
func (a *AnalysisRef) IsSuite() bool

// AllMembers reports whether the reference selects all members of a
// suite (pkg.Var[*]).
func (a *AnalysisRef) AllMembers() bool

// Errorf returns an error about the reference.
func (a *AnalysisRef) Errorf(format string, args ...any) error

// CompareAnalysisRefs orders references deterministically.
func CompareAnalysisRefs(x, y *AnalysisRef) int

// An AnalysisConfig records the references that apply to a root
// package.
type AnalysisConfig struct {
	// Positive holds the references that enable analyzers for the
	// package, in order of decreasing specificity.
	Positive []*AnalysisRef

	// Negative holds the negations that apply to the package,
	// which are applied once the analyzers are known.
	Negative []*AnalysisRef
}

// VetAnalysisConfig returns go test's built-in analysis configuration:
// the members of the vet suite with the specified names, or all of
// them if names is nil.
func VetAnalysisConfig(names []string) *AnalysisConfig

// AnalysisToolPackage returns a synthesized main package for an
// analysis tool and the source of its sole file. The tool contains
// the variables referenced by the specified references (those that
// enable analyzers for some root package).
//
// The package of each reference must load. A reference whose package
// does not load is reported, and dropped. (Whether the package
// declares the variable is discovered only if the tool fails to
// compile.)
//
// It also returns errors about the offending references.
//
// The package's imports are loaded without regard to the rule for
// internal packages, like the generated testmain: the tool must
// import analyzers from internal packages, and the driver itself is
// an internal package of the standard library.
func AnalysisToolPackage(ld *modload.Loader, ctx context.Context, opts PackageOpts, refs []*AnalysisRef) (pmain *Package, src []byte, errs []error)

// AnalysisToolSource returns the source of the main package of an
// analysis tool that links the specified variables.
func AnalysisToolSource(vars iter.Seq[AnalysisVar]) []byte
