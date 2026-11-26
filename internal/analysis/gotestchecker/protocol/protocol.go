// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package protocol defines the types of the protocol between the go
// command and the analysis tool that it synthesizes for go test, whose
// driver is [internal/analysis/gotestchecker]. The tool performs two
// operations: describe (tool -describe=FILE), whose result is a
// [Metadata] file, and analyze (tool FILE), whose request is an
// [AnalyzeRequest] and whose results include a pass file (see [Pass])
// for each requested analyzer.
//
// The protocol is internal to the toolchain, not stable, and may be
// changed at any time.
//
// This package is imported by the go command, so it must not depend on
// the analysis framework (or go/types).
package protocol

import (
	"github.com/shogo82148/std/time"
)

// Metadata is the result of the tool's metadata query (tool
// -describe=FILE), written by the tool as a JSON file.
type Metadata struct {
	// Vars describes each variable linked into the tool, in the
	// order of the argument of gotestchecker.Main, which is sorted
	// by Symbol (see [Metadata.Var]).
	Vars []*VarInfo

	// Analyzers describes every analyzer reachable from a valid
	// member of a valid variable, in topological order
	// (prerequisites first). Its indexes are the canonical
	// numbering of analyzers, by which the protocol identifies
	// them.
	Analyzers []*Analyzer
}

// Var returns the description of the variable with the specified
// symbol ("pkgpath.Var"), or nil if it was not linked into the tool.
func (md *Metadata) Var(symbol string) *VarInfo

// A VarInfo describes an analyzer or suite variable linked into the
// tool.
type VarInfo struct {
	Symbol string

	// Error, if non-empty, describes a problem that invalidates
	// the whole variable: a value of the wrong type, a nil analyzer
	// or suite element, or distinct members of a suite with the
	// same Name.
	Error string `json:",omitempty"`

	// IsSuite reports whether the variable is a suite
	// ([]*analysis.Analyzer) rather than an analyzer
	// (*analysis.Analyzer).
	IsSuite bool

	// Members lists the distinct analyzers of the variable: the
	// analyzer itself, or the elements of the suite, in order.
	Members []*Member `json:",omitempty"`
}

// A Member describes an analyzer that is a member of a variable.
// Analyzer is -1 if and only if Error is non-empty.
type Member struct {
	Name string

	Analyzer int
	Error    string `json:",omitempty"`
}

// An Analyzer describes an analyzer in the Requires closure of the valid
// members. Its index in Metadata.Analyzers identifies it in the
// protocol.
type Analyzer struct {
	Name     string
	Doc      string
	URL      string `json:",omitempty"`
	RunFunc  string
	Requires []int
	Facts    bool
}

// An AnalyzeRequest is a request to analyze one package (tool FILE),
// written by the go command as a JSON file. The results are the files
// that it names: types (TypeExport), type errors (TypeErrors), and a
// pass file (see [Pass]) for each analyzer in Outputs.
type AnalyzeRequest struct {
	ID           string
	ImportPath   string
	Dir          string
	GoFiles      []string
	NonGoFiles   []string
	IgnoredFiles []string
	ImportMap    map[string]string
	GoVersion    string
	GOARCH       string
	Module       *Module

	// TypeImports maps the package path of each direct import to
	// the name of a file holding its (deep) export data.
	TypeImports map[string]string

	// TypeExport names the file to which the export data of this
	// package should be written. It is empty if the go command
	// already has it (in its cache) and needs only the results of
	// analyzers (see Outputs).
	TypeExport string

	// TypeErrors names the file to which the type errors of this
	// package (a JSON list of strings, perhaps empty) should be
	// written. It is set if and only if TypeExport is: the type
	// errors are a property of the package, reported once by the go
	// command, not by each analyzer.
	TypeErrors string

	// Facts is a sparse mapping from the index of each analyzer
	// that uses facts, and is in the Requires closure of Outputs, to a
	// mapping from the package
	// path of each direct import to the name of the pass file
	// holding that analyzer's facts for that import. A missing
	// entry denotes an empty fact set.
	Facts map[int]map[string]string

	// Outputs is a sparse mapping from the index of each required
	// analyzer to the name of the file to which its [Pass] should be
	// written: its facts, if it uses facts; its diagnostics, if it
	// is in Report; or its failure. The tool runs these analyzers
	// and their prerequisites.
	Outputs map[int]string

	// Report lists the indexes of the analyzers in Outputs whose
	// diagnostics should be included in their pass files.
	Report []int
}

// A Module describes the module to which a package belongs. Its JSON
// encoding is that of a go/analysis.Module (and of the Module of
// go list -json); this copy avoids a dependency on go/analysis.
type Module struct {
	Path      string
	Version   string
	Replace   *Module
	Time      *time.Time
	Main      bool
	Indirect  bool
	Dir       string
	GoMod     string
	GoVersion string
	Error     *ModuleError
}

// A ModuleError holds errors loading a module; see [Module].
type ModuleError struct {
	Err string
}

// A Pass holds the results of an analyzer on a package (a "pass"). A
// pass file is the encoding of a Pass (see [WritePass]).
//
// A pass that failed or was skipped has no facts, and its dependent
// passes (the same analyzer on importers, if it uses facts, and the
// analyzers that require it on the same package) are skipped. Each
// failure is reported once, by the go command, at the pass that
// failed; a skipped pass is not reported, since its cause is.
type Pass struct {
	// Facts is the encoded fact set of the analyzer, which
	// includes facts about dependencies ("deep"). It is nil if the
	// analyzer uses no facts.
	Facts []byte

	// Failure records the failure of the analyzer on this package,
	// if any: an error returned by Run, a panic in Run, or a fact
	// that cannot be encoded. A failure prevents only the dependent
	// passes; it does not stop the build.
	Failure *Failure

	// Skipped reports that the analyzer did not run on this package
	// because the package has type errors, a prerequisite failed or
	// was skipped, or (for an analyzer that uses facts) it failed or
	// was skipped on a direct import.
	Skipped bool

	// Diagnostics holds the analyzer's diagnostics. It is empty
	// unless the package's diagnostics are reported.
	Diagnostics []*Diagnostic
}

// A Failure records the failure of an analyzer on a package: an
// error returned by the analyzer, a panic in the analyzer, or a fact
// that cannot be encoded.
type Failure struct {
	Package  string
	Analyzer string
	Message  string
}

func (f *Failure) String() string

// A Diagnostic is the serialized form of an analysis.Diagnostic,
// without suggested fixes.
type Diagnostic struct {
	Posn     string
	End      string
	Category string
	Message  string
	URL      string
}

// WritePass writes the pass to the named file.
func WritePass(filename string, pass *Pass) error

// ReadPass reads the named pass file.
func ReadPass(filename string) (*Pass, error)
