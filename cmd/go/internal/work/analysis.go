// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package work

import (
	"github.com/shogo82148/std/cmd/go/internal/load"
	"github.com/shogo82148/std/cmd/go/internal/modload"
)

// AnalysisVetAction returns an action that reports the diagnostics of
// the analyzers enabled for package p by the specified configuration,
// which must be non-nil. The action depends on the analysis of p,
// which in turn depends on the analysis of each of p's dependencies,
// which computes the facts needed by p's analyzers.
// The tool must be supplied later by SetAnalysisTool.
//
// Reporting is a separate action from analysis because the analysis
// action of a root package may be shared with the importers of that
// package (e.g. a package without test files is itself the root, not
// a test variant of it), and its diagnostics must not cause the
// analysis of its importers to fail.
func (b *Builder) AnalysisVetAction(s *modload.Loader, p *load.Package, config *load.AnalysisConfig) *Action

// SetAnalysisTool records the link action of the analysis tool. It
// creates the analysismeta action, which depends on the link action,
// and makes every analysis action depend on it.
func (b *Builder) SetAnalysisTool(link *Action)

// HasAnalysisActions reports whether any analysis actions were created.
func (b *Builder) HasAnalysisActions() bool
