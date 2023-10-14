// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mvs

import (
	"golang.org/x/mod/module"
)

// BuildListError decorates an error that occurred gathering requirements
// while constructing a build list. BuildListError prints the chain
// of requirements to the module where the error occurred.
type BuildListError struct {
	Err   error
	stack []buildListErrorElem
}

// NewBuildListError returns a new BuildListError wrapping an error that
// occurred at a module found along the given path of requirements and/or
// upgrades, which must be non-empty.
//
// The isVersionChange function reports whether a path step is due to an
// explicit upgrade or downgrade (as opposed to an existing requirement in a
// go.mod file). A nil isVersionChange function indicates that none of the path
// steps are due to explicit version changes.
func NewBuildListError(err error, path []module.Version, isVersionChange func(from, to module.Version) bool) *BuildListError

// Module returns the module where the error occurred. If the module stack
// is empty, this returns a zero value.
func (e *BuildListError) Module() module.Version

func (e *BuildListError) Error() string

func (e *BuildListError) Unwrap() error
