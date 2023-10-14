// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Action graph execution methods related to coverage.

package work

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/io"
)

// CovData invokes "go tool covdata" with the specified arguments
// as part of the execution of action 'a'.
func (b *Builder) CovData(a *Action, cmdargs ...any) ([]byte, error)

// BuildActionCoverMetaFile locates and returns the path of the
// meta-data file written by the "go tool cover" step as part of the
// build action for the "go test -cover" run action 'runAct'. Note
// that if the package has no functions the meta-data file will exist
// but will be empty; in this case the return is an empty string.
func BuildActionCoverMetaFile(runAct *Action) (string, error)

// WriteCoveragePercent writes out to the writer 'w' a "percent
// statements covered" for the package whose test-run action is
// 'runAct', based on the meta-data file 'mf'. This helper is used in
// cases where a user runs "go test -cover" on a package that has
// functions but no tests; in the normal case (package has tests)
// the percentage is written by the test binary when it runs.
func WriteCoveragePercent(b *Builder, runAct *Action, mf string, w io.Writer) error

// WriteCoverageProfile writes out a coverage profile fragment for the
// package whose test-run action is 'runAct'; content is written to
// the file 'outf' based on the coverage meta-data info found in
// 'mf'. This helper is used in cases where a user runs "go test
// -cover" on a package that has functions but no tests.
func WriteCoverageProfile(b *Builder, runAct *Action, mf, outf string, w io.Writer) error

// WriteCoverMetaFilesFile writes out a summary file ("meta-files
// file") as part of the action function for the "writeCoverMeta"
// pseudo action employed during "go test -coverpkg" runs where there
// are multiple tests and multiple packages covered. It builds up a
// table mapping package import path to meta-data file fragment and
// writes it out to a file where it can be read by the various test
// run actions. Note that this function has to be called A) after the
// build actions are complete for all packages being tested, and B)
// before any of the "run test" actions for those packages happen.
// This requirement is enforced by adding making this action ("a")
// dependent on all test package build actions, and making all test
// run actions dependent on this action.
func WriteCoverMetaFilesFile(b *Builder, ctx context.Context, a *Action) error
