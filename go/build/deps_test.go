// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file exercises the import parser but also checks that
// some low-level packages do not have new dependencies added.

package build

// pkgDeps defines the expected dependencies between packages in
// the Go source tree.  It is a statement of policy.
// Changes should not be made to this map without prior discussion.
//
// The map contains two kinds of entries:
// 1) Lower-case keys are standard import paths and list the
// allowed imports in that package.
// 2) Upper-case keys define aliases for package sets, which can then
// be used as dependencies by other rules.
//
// DO NOT CHANGE THIS DATA TO FIX BUILDS.
//

// allowedErrors are the operating systems and packages known to contain errors
// (currently just "no Go source files")
