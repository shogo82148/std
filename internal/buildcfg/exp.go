// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buildcfg

import (
	"github.com/shogo82148/std/internal/goexperiment"
)

// ExperimentFlags represents a set of GOEXPERIMENT flags relative to a baseline
// (platform-default) experiment configuration.
type ExperimentFlags struct {
	goexperiment.Flags
	baseline goexperiment.Flags
}

// Experiment contains the toolchain experiments enabled for the
// current build.
//
// (This is not necessarily the set of experiments the compiler itself
// was built with.)
//
// experimentBaseline specifies the experiment flags that are enabled by
// default in the current toolchain. This is, in effect, the "control"
// configuration and any variation from this is an experiment.
var Experiment ExperimentFlags = func() ExperimentFlags {
	flags, err := ParseGOEXPERIMENT(GOOS, GOARCH, envOr("GOEXPERIMENT", defaultGOEXPERIMENT))
	if err != nil {
		Error = err
		return ExperimentFlags{}
	}
	return *flags
}()

// DefaultGOEXPERIMENT is the embedded default GOEXPERIMENT string.
// It is not guaranteed to be canonical.
const DefaultGOEXPERIMENT = defaultGOEXPERIMENT

// FramePointerEnabled enables the use of platform conventions for
// saving frame pointers.
//
// This used to be an experiment, but now it's always enabled on
// platforms that support it.
//
// Note: must agree with runtime.framepointer_enabled.
var FramePointerEnabled = GOARCH == "amd64" || GOARCH == "arm64"

// ParseGOEXPERIMENT parses a (GOOS, GOARCH, GOEXPERIMENT)
// configuration tuple and returns the enabled and baseline experiment
// flag sets.
//
// TODO(mdempsky): Move to internal/goexperiment.
func ParseGOEXPERIMENT(goos, goarch, goexp string) (*ExperimentFlags, error)

// String returns the canonical GOEXPERIMENT string to enable this experiment
// configuration. (Experiments in the same state as in the baseline are elided.)
func (exp *ExperimentFlags) String() string

// Enabled returns a list of enabled experiments, as
// lower-cased experiment names.
func (exp *ExperimentFlags) Enabled() []string

// All returns a list of all experiment settings.
// Disabled experiments appear in the list prefixed by "no".
func (exp *ExperimentFlags) All() []string
