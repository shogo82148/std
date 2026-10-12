// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scripttest adapts the script engine for use in tests.
package scripttest

import (
	"github.com/shogo82148/std/testing"
)

// SetupTestGoRoot sets up a temporary GOROOT for use with script test
// execution. It copies the existing goroot bin and pkg dirs using
// symlinks (if possible) or raw copying. Return value is the path to
// the newly created testgoroot dir.
func SetupTestGoRoot(t *testing.T, tmpdir string, goroot string) string

// ReplaceGoToolInTestGoRoot replaces the go tool binary toolname with
// an alternate executable newtoolpath within a test GOROOT directory
// previously created by SetupTestGoRoot.
func ReplaceGoToolInTestGoRoot(t *testing.T, testgoroot, toolname, newtoolpath string)
