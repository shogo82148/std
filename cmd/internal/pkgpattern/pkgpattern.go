// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgpattern

// TreeCanMatchPattern(pattern)(name) reports whether
// name or children of name can possibly match pattern.
// Pattern is the same limited glob accepted by MatchPattern.
func TreeCanMatchPattern(pattern string) func(name string) bool

// MatchPattern(pattern)(name) reports whether
// name matches pattern. Pattern is a limited glob
// pattern in which '...' means 'any string' and there
// is no other special syntax.
// Unfortunately, there are two special cases. Quoting "go help packages":
//
// First, /... at the end of the pattern can match an empty string,
// so that net/... matches both net and packages in its subdirectories, like net/http.
// Second, any slash-separated pattern element containing a wildcard never
// participates in a match of the "vendor" element in the path of a vendored
// package, so that ./... does not match packages in subdirectories of
// ./vendor or ./mycode/vendor, but ./vendor/... and ./mycode/vendor/... do.
// Note, however, that a directory named vendor that itself contains code
// is not a vendored package: cmd/vendor would be a command named vendor,
// and the pattern cmd/... matches it.
func MatchPattern(pattern string) func(name string) bool

// MatchSimplePattern returns a function that can be used to check
// whether a given name matches a pattern, where pattern is a limited
// glob pattern in which '...' means 'any string', with no other
// special syntax. There is one special case for MatchPatternSimple:
// according to the rules in "go help packages": a /... at the end of
// the pattern can match an empty string, so that net/... matches both
// net and packages in its subdirectories, like net/http.
func MatchSimplePattern(pattern string) func(name string) bool
