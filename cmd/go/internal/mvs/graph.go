// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mvs

import (
	"github.com/shogo82148/std/golang.org/x/mod/module"
)

// Graph implements an incremental version of the MVS algorithm, with the
// requirements pushed by the caller instead of pulled by the MVS traversal.
type Graph struct {
	cmp   func(p, v1, v2 string) int
	roots []module.Version

	required map[module.Version][]module.Version

	isRoot   map[module.Version]bool
	selected map[string]string
}

// NewGraph returns an incremental MVS graph containing only a set of root
// dependencies and using the given max function for version strings.
//
// The caller must ensure that the root slice is not modified while the Graph
// may be in use.
func NewGraph(cmp func(p, v1, v2 string) int, roots []module.Version) *Graph

// Require adds the information that module m requires all modules in reqs.
// The reqs slice must not be modified after it is passed to Require.
//
// m must be reachable by some existing chain of requirements from g's target,
// and Require must not have been called for it already.
//
// If any of the modules in reqs has the same path as g's target,
// the target must have higher precedence than the version in req.
func (g *Graph) Require(m module.Version, reqs []module.Version)

// RequiredBy returns the slice of requirements passed to Require for m, if any,
// with its capacity reduced to its length.
// If Require has not been called for m, RequiredBy(m) returns ok=false.
//
// The caller must not modify the returned slice, but may safely append to it
// and may rely on it not to be modified.
func (g *Graph) RequiredBy(m module.Version) (reqs []module.Version, ok bool)

// Selected returns the selected version of the given module path.
//
// If no version is selected, Selected returns version "none".
func (g *Graph) Selected(path string) (version string)

// BuildList returns the selected versions of all modules present in the Graph,
// beginning with the selected versions of each module path in the roots of g.
//
// The order of the remaining elements in the list is deterministic
// but arbitrary.
func (g *Graph) BuildList() []module.Version

// WalkBreadthFirst invokes f once, in breadth-first order, for each module
// version other than "none" that appears in the graph, regardless of whether
// that version is selected.
func (g *Graph) WalkBreadthFirst(f func(m module.Version))

// FindPath reports a shortest requirement path starting at one of the roots of
// the graph and ending at a module version m for which f(m) returns true, or
// nil if no such path exists.
func (g *Graph) FindPath(f func(module.Version) bool) []module.Version
