// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copied from Go distribution src/go/build/build.go, syslist.go.
// That package does not export the ability to process raw file data,
// although we could fake it with an appropriate build.Context
// and a lot of unwrapping.
// More importantly, that package does not implement the tags["*"]
// special case, in which both tag and !tag are considered to be true
// for essentially all tags (except "ignore").
//
// If we added this API to go/build directly, we wouldn't need this
// file anymore, but this API is not terribly general-purpose and we
// don't really want to commit to any public form of it, nor do we
// want to move the core parts of go/build into a top-level internal package.
// These details change very infrequently, so the copy is fine.

package imports

import (
	"github.com/shogo82148/std/go/build/constraint"
)

// ShouldBuild reports whether it is okay to use this file,
// The rule is that in the file's leading run of // comments
// and blank lines, which must be followed by a blank line
// (to avoid including a Go package clause doc comment),
// lines beginning with '// +build' are taken as build directives.
//
// The file is accepted only if each such line lists something
// matching the file. For example:
//
//	// +build windows linux
//
// marks the file as applicable only on Windows and Linux.
//
// If tags["*"] is true, then ShouldBuild will consider every
// build tag except "ignore" to be both true and false for
// the purpose of satisfying build tags, in order to estimate
// (conservatively) whether a file could ever possibly be used
// in any build.
func ShouldBuild(content []byte, tags map[string]bool) bool

// Eval is like
//
//	x.Eval(func(tag string) bool { return matchTag(tag, tags) })
//
// except that it implements the special case for tags["*"] meaning
// all tags are both true and false at the same time.
func Eval(x constraint.Expr, tags map[string]bool, prefer bool) bool

// MatchFile returns false if the name contains a $GOOS or $GOARCH
// suffix which does not match the current system.
// The recognized name formats are:
//
//	name_$(GOOS).*
//	name_$(GOARCH).*
//	name_$(GOOS)_$(GOARCH).*
//	name_$(GOOS)_test.*
//	name_$(GOARCH)_test.*
//	name_$(GOOS)_$(GOARCH)_test.*
//
// Exceptions:
//
//	if GOOS=android, then files with GOOS=linux are also matched.
//	if GOOS=illumos, then files with GOOS=solaris are also matched.
//	if GOOS=ios, then files with GOOS=darwin are also matched.
//
// If tags["*"] is true, then MatchFile will consider all possible
// GOOS and GOARCH to be available and will consequently
// always return true.
func MatchFile(name string, tags map[string]bool) bool
