// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package search

// A Match represents the result of matching a single package pattern.
type Match struct {
	pattern string
	Dirs    []string
	Pkgs    []string
	Errs    []error
}

// NewMatch returns a Match describing the given pattern,
// without resolving its packages or errors.
func NewMatch(pattern string) *Match

// Pattern returns the pattern to be matched.
func (m *Match) Pattern() string

// AddError appends a MatchError wrapping err to m.Errs.
func (m *Match) AddError(err error)

// IsLiteral reports whether the pattern is free of wildcards and meta-patterns.
//
// A literal pattern must match at most one package.
func (m *Match) IsLiteral() bool

// IsLocal reports whether the pattern must be resolved from a specific root or
// directory, such as a filesystem path or a single module.
func (m *Match) IsLocal() bool

// IsMeta reports whether the pattern is a “meta-package” keyword that represents
// multiple packages, such as "std", "cmd", "tool", or "all".
func (m *Match) IsMeta() bool

// IsMetaPackage checks if name is a reserved package name that expands to multiple packages.
func IsMetaPackage(name string) bool

// A MatchError indicates an error that occurred while attempting to match a
// pattern.
type MatchError struct {
	Match *Match
	Err   error
}

func (e *MatchError) Error() string

func (e *MatchError) Unwrap() error

// MatchPackages sets m.Pkgs to a non-nil slice containing all the packages that
// can be found under the $GOPATH directories and $GOROOT that match the
// pattern. The pattern must be either "all" (all packages), "std" (standard
// packages), "cmd" (standard commands), or a path including "...".
//
// If any errors may have caused the set of packages to be incomplete,
// MatchPackages appends those errors to m.Errs.
func (m *Match) MatchPackages()

// MatchDirs sets m.Dirs to a non-nil slice containing all directories that
// potentially match a local pattern. The pattern must begin with an absolute
// path, or "./", or "../". On Windows, the pattern may use slash or backslash
// separators or a mix of both.
//
// If any errors may have caused the set of directories to be incomplete,
// MatchDirs appends those errors to m.Errs.
func (m *Match) MatchDirs(modRoots []string)

// WarnUnmatched warns about patterns that didn't match any packages.
func WarnUnmatched(matches []*Match)

// ImportPaths returns the matching paths to use for the given command line.
// It calls ImportPathsQuiet and then WarnUnmatched.
func ImportPaths(patterns, modRoots []string) []*Match

// ImportPathsQuiet is like ImportPaths but does not warn about patterns with no matches.
func ImportPathsQuiet(patterns, modRoots []string) []*Match

// CleanPatterns returns the patterns to use for the given command line. It
// canonicalizes the patterns but does not evaluate any matches. For patterns
// that are not local or absolute paths, it preserves text after '@' to avoid
// modifying version queries.
func CleanPatterns(patterns []string) []string

// IsStandardImportPath reports whether $GOROOT/src/path should be considered
// part of the standard distribution. For historical reasons we allow people to add
// their own code to $GOROOT instead of using $GOPATH, but we assume that
// code will start with a domain name (dot in the first element).
//
// Note that this function is meant to evaluate whether a directory found in GOROOT
// should be treated as part of the standard library. It should not be used to decide
// that a directory found in GOPATH should be rejected: directories in GOPATH
// need not have dots in the first element, and they just take their chances
// with future collisions in the standard library.
func IsStandardImportPath(path string) bool

// IsRelativePath reports whether pattern should be interpreted as a directory
// path relative to the current directory, as opposed to a pattern matching
// import paths.
func IsRelativePath(pattern string) bool

// InDir checks whether path is in the file tree rooted at dir.
// If so, InDir returns an equivalent path relative to dir.
// If not, InDir returns an empty string.
// InDir makes some effort to succeed even in the presence of symbolic links.
func InDir(path, dir string) string
