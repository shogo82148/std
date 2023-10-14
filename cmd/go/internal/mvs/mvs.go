// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mvs implements Minimal Version Selection.
// See https://research.swtch.com/vgo-mvs.
package mvs

import (
	"golang.org/x/mod/module"
)

// A Reqs is the requirement graph on which Minimal Version Selection (MVS) operates.
//
// The version strings are opaque except for the special version "none"
// (see the documentation for module.Version). In particular, MVS does not
// assume that the version strings are semantic versions; instead, the Max method
// gives access to the comparison operation.
//
// It must be safe to call methods on a Reqs from multiple goroutines simultaneously.
// Because a Reqs may read the underlying graph from the network on demand,
// the MVS algorithms parallelize the traversal to overlap network delays.
type Reqs interface {
	// Required returns the module versions explicitly required by m itself.
	// The caller must not modify the returned list.
	Required(m module.Version) ([]module.Version, error)

	// Max returns the maximum of v1 and v2 (it returns either v1 or v2)
	// in the module with path p.
	//
	// For all versions v, Max(v, "none") must be v,
	// and for the target passed as the first argument to MVS functions,
	// Max(target, v) must be target.
	//
	// Note that v1 < v2 can be written Max(v1, v2) != v1
	// and similarly v1 <= v2 can be written Max(v1, v2) == v2.
	Max(p, v1, v2 string) string
}

// An UpgradeReqs is a Reqs that can also identify available upgrades.
type UpgradeReqs interface {
	Reqs

	// Upgrade returns the upgraded version of m,
	// for use during an UpgradeAll operation.
	// If m should be kept as is, Upgrade returns m.
	// If m is not yet used in the build, then m.Version will be "none".
	// More typically, m.Version will be the version required
	// by some other module in the build.
	//
	// If no module version is available for the given path,
	// Upgrade returns a non-nil error.
	// TODO(rsc): Upgrade must be able to return errors,
	// but should "no latest version" just return m instead?
	Upgrade(m module.Version) (module.Version, error)
}

// A DowngradeReqs is a Reqs that can also identify available downgrades.
type DowngradeReqs interface {
	Reqs

	// Previous returns the version of m.Path immediately prior to m.Version,
	// or "none" if no such version is known.
	Previous(m module.Version) (module.Version, error)
}

// BuildList returns the build list for the target module.
//
// target is the root vertex of a module requirement graph. For cmd/go, this is
// typically the main module, but note that this algorithm is not intended to
// be Go-specific: module paths and versions are treated as opaque values.
//
// reqs describes the module requirement graph and provides an opaque method
// for comparing versions.
//
// BuildList traverses the graph and returns a list containing the highest
// version for each visited module. The first element of the returned list is
// target itself; reqs.Max requires target.Version to compare higher than all
// other versions, so no other version can be selected. The remaining elements
// of the list are sorted by path.
//
// See https://research.swtch.com/vgo-mvs for details.
func BuildList(targets []module.Version, reqs Reqs) ([]module.Version, error)

// Req returns the minimal requirement list for the target module,
// with the constraint that all module paths listed in base must
// appear in the returned list.
func Req(mainModule module.Version, base []string, reqs Reqs) ([]module.Version, error)

// UpgradeAll returns a build list for the target module
// in which every module is upgraded to its latest version.
func UpgradeAll(target module.Version, reqs UpgradeReqs) ([]module.Version, error)

// Upgrade returns a build list for the target module
// in which the given additional modules are upgraded.
func Upgrade(target module.Version, reqs UpgradeReqs, upgrade ...module.Version) ([]module.Version, error)

// Downgrade returns a build list for the target module
// in which the given additional modules are downgraded,
// potentially overriding the requirements of the target.
//
// The versions to be downgraded may be unreachable from reqs.Latest and
// reqs.Previous, but the methods of reqs must otherwise handle such versions
// correctly.
func Downgrade(target module.Version, reqs DowngradeReqs, downgrade ...module.Version) ([]module.Version, error)
