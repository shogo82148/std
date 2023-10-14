// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package coverage

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/internal/coverage"
)

// Names records state information collected in the first fixup
// phase so that it can be passed to the second fixup phase.
type Names struct {
	MetaVar     *ir.Name
	PkgIdVar    *ir.Name
	InitFn      *ir.Func
	CounterMode coverage.CounterMode
	CounterGran coverage.CounterGranularity
}

// FixupVars is the first of two entry points for coverage compiler
// fixup. It collects and returns the package ID and meta-data
// variables being used for this "-cover" build, along with the
// coverage counter mode and granularity. It also reclassifies selected
// variables (for example, tagging coverage counter variables with
// flags so that they can be handled properly downstream).
func FixupVars() Names

// FixupInit is the second main entry point for coverage compiler
// fixup. It adds calls to the pkg init function as appropriate to
// register coverage-related variables with the runtime.
func FixupInit(cnames Names)
