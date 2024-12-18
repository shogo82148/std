// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modload

import (
	"github.com/shogo82148/std/context"

	"github.com/shogo82148/std/cmd/go/internal/modfetch"

	"golang.org/x/mod/module"
)

// Query looks up a revision of a given module given a version query string.
// The module must be a complete module path.
// The version must take one of the following forms:
//
//   - the literal string "latest", denoting the latest available, allowed
//     tagged version, with non-prereleases preferred over prereleases.
//     If there are no tagged versions in the repo, latest returns the most
//     recent commit.
//
//   - the literal string "upgrade", equivalent to "latest" except that if
//     current is a newer version, current will be returned (see below).
//
//   - the literal string "patch", denoting the latest available tagged version
//     with the same major and minor number as current (see below).
//
//   - v1, denoting the latest available tagged version v1.x.x.
//
//   - v1.2, denoting the latest available tagged version v1.2.x.
//
//   - v1.2.3, a semantic version string denoting that tagged version.
//
//   - <v1.2.3, <=v1.2.3, >v1.2.3, >=v1.2.3,
//     denoting the version closest to the target and satisfying the given operator,
//     with non-prereleases preferred over prereleases.
//
//   - a repository commit identifier or tag, denoting that commit.
//
// current denotes the currently-selected version of the module; it may be
// "none" if no version is currently selected, or "" if the currently-selected
// version is unknown or should not be considered. If query is
// "upgrade" or "patch", current will be returned if it is a newer
// semantic version or a chronologically later pseudo-version than the
// version that would otherwise be chosen. This prevents accidental downgrades
// from newer pre-release or development versions.
//
// The allowed function (which may be nil) is used to filter out unsuitable
// versions (see AllowedFunc documentation for details). If the query refers to
// a specific revision (for example, "master"; see IsRevisionQuery), and the
// revision is disallowed by allowed, Query returns the error. If the query
// does not refer to a specific revision (for example, "latest"), Query
// acts as if versions disallowed by allowed do not exist.
//
// If path is the path of the main module and the query is "latest",
// Query returns Target.Version as the version.
//
// Query often returns a non-nil *RevInfo with a non-nil error,
// to provide an info.Origin that can allow the error to be cached.
func Query(ctx context.Context, path, query, current string, allowed AllowedFunc) (*modfetch.RevInfo, error)

// AllowedFunc is used by Query and other functions to filter out unsuitable
// versions, for example, those listed in exclude directives in the main
// module's go.mod file.
//
// An AllowedFunc returns an error equivalent to ErrDisallowed for an unsuitable
// version. Any other error indicates the function was unable to determine
// whether the version should be allowed, for example, the function was unable
// to fetch or parse a go.mod file containing retractions. Typically, errors
// other than ErrDisallowed may be ignored.
type AllowedFunc func(context.Context, module.Version) error

// IsRevisionQuery returns true if vers is a version query that may refer to
// a particular version or revision in a repository like "v1.0.0", "master",
// or "0123abcd". IsRevisionQuery returns false if vers is a query that
// chooses from among available versions like "latest" or ">v1.0.0".
func IsRevisionQuery(path, vers string) bool

type QueryResult struct {
	Mod      module.Version
	Rev      *modfetch.RevInfo
	Packages []string
}

// QueryPackages is like QueryPattern, but requires that the pattern match at
// least one package and omits the non-package result (if any).
func QueryPackages(ctx context.Context, pattern, query string, current func(string) string, allowed AllowedFunc) ([]QueryResult, error)

// QueryPattern looks up the module(s) containing at least one package matching
// the given pattern at the given version. The results are sorted by module path
// length in descending order. If any proxy provides a non-empty set of candidate
// modules, no further proxies are tried.
//
// For wildcard patterns, QueryPattern looks in modules with package paths up to
// the first "..." in the pattern. For the pattern "example.com/a/b.../c",
// QueryPattern would consider prefixes of "example.com/a".
//
// If any matching package is in the main module, QueryPattern considers only
// the main module and only the version "latest", without checking for other
// possible modules.
//
// QueryPattern always returns at least one QueryResult (which may be only
// modOnly) or a non-nil error.
func QueryPattern(ctx context.Context, pattern, query string, current func(string) string, allowed AllowedFunc) (pkgMods []QueryResult, modOnly *QueryResult, err error)

// A NoMatchingVersionError indicates that Query found a module at the requested
// path, but not at any versions satisfying the query string and allow-function.
//
// NOTE: NoMatchingVersionError MUST NOT implement Is(fs.ErrNotExist).
//
// If the module came from a proxy, that proxy had to return a successful status
// code for the versions it knows about, and thus did not have the opportunity
// to return a non-400 status code to suppress fallback.
type NoMatchingVersionError struct {
	query, current string
}

func (e *NoMatchingVersionError) Error() string

// A NoPatchBaseError indicates that Query was called with the query "patch"
// but with a current version of "" or "none".
type NoPatchBaseError struct {
	path string
}

func (e *NoPatchBaseError) Error() string

// A WildcardInFirstElementError indicates that a pattern passed to QueryPattern
// had a wildcard in its first path element, and therefore had no pattern-prefix
// modules to search in.
type WildcardInFirstElementError struct {
	Pattern string
	Query   string
}

func (e *WildcardInFirstElementError) Error() string

// A PackageNotInModuleError indicates that QueryPattern found a candidate
// module at the requested version, but that module did not contain any packages
// matching the requested pattern.
//
// NOTE: PackageNotInModuleError MUST NOT implement Is(fs.ErrNotExist).
//
// If the module came from a proxy, that proxy had to return a successful status
// code for the versions it knows about, and thus did not have the opportunity
// to return a non-400 status code to suppress fallback.
type PackageNotInModuleError struct {
	MainModules []module.Version
	Mod         module.Version
	Replacement module.Version
	Query       string
	Pattern     string
}

func (e *PackageNotInModuleError) Error() string

func (e *PackageNotInModuleError) ImportPath() string

var _ versionRepo = modfetch.Repo(nil)

var _ versionRepo = emptyRepo{}

var _ versionRepo = (*replacementRepo)(nil)

// A QueryMatchesMainModulesError indicates that a query requests
// a version of the main module that cannot be satisfied.
// (The main module's version cannot be changed.)
type QueryMatchesMainModulesError struct {
	MainModules []module.Version
	Pattern     string
	Query       string
}

func (e *QueryMatchesMainModulesError) Error() string

// A QueryUpgradesAllError indicates that a query requests
// an upgrade on the all pattern.
// (The main module's version cannot be changed.)
type QueryUpgradesAllError struct {
	MainModules []module.Version
	Query       string
}

func (e *QueryUpgradesAllError) Error() string

// A QueryMatchesPackagesInMainModuleError indicates that a query cannot be
// satisfied because it matches one or more packages found in the main module.
type QueryMatchesPackagesInMainModuleError struct {
	Pattern  string
	Query    string
	Packages []string
}

func (e *QueryMatchesPackagesInMainModuleError) Error() string
