// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modload

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"

	"github.com/shogo82148/std/cmd/go/internal/mvs"
	"github.com/shogo82148/std/cmd/go/internal/par"

	"github.com/shogo82148/std/golang.org/x/mod/module"
)

// A Requirements represents a logically-immutable set of root module requirements.
type Requirements struct {
	// pruning is the pruning at which the requirement graph is computed.
	//
	// If unpruned, the graph includes all transitive requirements regardless
	// of whether the requiring module supports pruning.
	//
	// If pruned, the graph includes only the root modules, the explicit
	// requirements of those root modules, and the transitive requirements of only
	// the root modules that do not support pruning.
	//
	// If workspace, the graph includes only the workspace modules, the explicit
	// requirements of the workspace modules, and the transitive requirements of
	// the workspace modules that do not support pruning.
	pruning modPruning

	// rootModules is the set of root modules of the graph, sorted and capped to
	// length. It may contain duplicates, and may contain multiple versions for a
	// given module path. The root modules of the graph are the set of main
	// modules in workspace mode, and the main module's direct requirements
	// outside workspace mode.
	//
	// The roots are always expected to contain an entry for the "go" module,
	// indicating the Go language version in use.
	rootModules    []module.Version
	maxRootVersion map[string]string

	// direct is the set of module paths for which we believe the module provides
	// a package directly imported by a package or test in the main module.
	//
	// The "direct" map controls which modules are annotated with "// indirect"
	// comments in the go.mod file, and may impact which modules are listed as
	// explicit roots (vs. indirect-only dependencies). However, it should not
	// have a semantic effect on the build list overall.
	//
	// The initial direct map is populated from the existing "// indirect"
	// comments (or lack thereof) in the go.mod file. It is updated by the
	// package loader: dependencies may be promoted to direct if new
	// direct imports are observed, and may be demoted to indirect during
	// 'go mod tidy' or 'go mod vendor'.
	//
	// The direct map is keyed by module paths, not module versions. When a
	// module's selected version changes, we assume that it remains direct if the
	// previous version was a direct dependency. That assumption might not hold in
	// rare cases (such as if a dependency splits out a nested module, or merges a
	// nested module back into a parent module).
	direct map[string]bool

	graphOnce sync.Once
	graph     atomic.Pointer[cachedGraph]
}

// String returns a string describing the Requirements for debugging.
func (rs *Requirements) String() string

// GoVersion returns the Go language version for the Requirements.
func (rs *Requirements) GoVersion() string

// Graph returns the graph of module requirements loaded from the current
// root modules (as reported by RootModules).
//
// Graph always makes a best effort to load the requirement graph despite any
// errors, and always returns a non-nil *ModuleGraph.
//
// If the requirements of any relevant module fail to load, Graph also
// returns a non-nil error of type *mvs.BuildListError.
func (rs *Requirements) Graph(ctx context.Context) (*ModuleGraph, error)

// IsDirect returns whether the given module provides a package directly
// imported by a package or test in the main module.
func (rs *Requirements) IsDirect(path string) bool

// A ModuleGraph represents the complete graph of module dependencies
// of a main module.
//
// If the main module supports module graph pruning, the graph does not include
// transitive dependencies of non-root (implicit) dependencies.
type ModuleGraph struct {
	g         *mvs.Graph
	loadCache par.ErrCache[module.Version, *modFileSummary]

	buildListOnce sync.Once
	buildList     []module.Version
}

// RequiredBy returns the dependencies required by module m in the graph,
// or ok=false if module m's dependencies are pruned out.
//
// The caller must not modify the returned slice, but may safely append to it
// and may rely on it not to be modified.
func (mg *ModuleGraph) RequiredBy(m module.Version) (reqs []module.Version, ok bool)

// Selected returns the selected version of the module with the given path.
//
// If no version is selected, Selected returns version "none".
func (mg *ModuleGraph) Selected(path string) (version string)

// WalkBreadthFirst invokes f once, in breadth-first order, for each module
// version other than "none" that appears in the graph, regardless of whether
// that version is selected.
func (mg *ModuleGraph) WalkBreadthFirst(f func(m module.Version))

// BuildList returns the selected versions of all modules present in the graph,
// beginning with the main modules.
//
// The order of the remaining elements in the list is deterministic
// but arbitrary.
//
// The caller must not modify the returned list, but may safely append to it
// and may rely on it not to be modified.
func (mg *ModuleGraph) BuildList() []module.Version

// LoadModGraph loads and returns the graph of module dependencies of the main module,
// without loading any packages.
//
// If the goVersion string is non-empty, the returned graph is the graph
// as interpreted by the given Go version (instead of the version indicated
// in the go.mod file).
//
// Modules are loaded automatically (and lazily) in LoadPackages:
// LoadModGraph need only be called if LoadPackages is not,
// typically in commands that care about modules but no particular package.
func LoadModGraph(ctx context.Context, goVersion string) (*ModuleGraph, error)

// EditBuildList edits the global build list by first adding every module in add
// to the existing build list, then adjusting versions (and adding or removing
// requirements as needed) until every module in mustSelect is selected at the
// given version.
//
// (Note that the newly-added modules might not be selected in the resulting
// build list: they could be lower than existing requirements or conflict with
// versions in mustSelect.)
//
// If the versions listed in mustSelect are mutually incompatible (due to one of
// the listed modules requiring a higher version of another), EditBuildList
// returns a *ConstraintError and leaves the build list in its previous state.
//
// On success, EditBuildList reports whether the selected version of any module
// in the build list may have been changed (possibly to or from "none") as a
// result.
func EditBuildList(ctx context.Context, add, mustSelect []module.Version) (changed bool, err error)

// OverrideRoots edits the global requirement roots by replacing the specific module versions.
func OverrideRoots(ctx context.Context, replace []module.Version)

// A ConstraintError describes inconsistent constraints in EditBuildList
type ConstraintError struct {
	// Conflict lists the source of the conflict for each version in mustSelect
	// that could not be selected due to the requirements of some other version in
	// mustSelect.
	Conflicts []Conflict
}

func (e *ConstraintError) Error() string

// A Conflict is a path of requirements starting at a root or proposed root in
// the requirement graph, explaining why that root either causes a module passed
// in the mustSelect list to EditBuildList to be unattainable, or introduces an
// unresolvable error in loading the requirement graph.
type Conflict struct {
	// Path is a path of requirements starting at some module version passed in
	// the mustSelect argument and ending at a module whose requirements make that
	// version unacceptable. (Path always has len ≥ 1.)
	Path []module.Version

	// If Err is nil, Constraint is a module version passed in the mustSelect
	// argument that has the same module path as, and a lower version than,
	// the last element of the Path slice.
	Constraint module.Version

	// If Constraint is unset, Err is an error encountered when loading the
	// requirements of the last element in Path.
	Err error
}

// UnwrapModuleError returns c.Err, but unwraps it if it is a module.ModuleError
// with a version and path matching the last entry in the Path slice.
func (c Conflict) UnwrapModuleError() error

// Summary returns a string that describes only the first and last modules in
// the conflict path.
func (c Conflict) Summary() string

// String returns a string that describes the full conflict path.
func (c Conflict) String() string
