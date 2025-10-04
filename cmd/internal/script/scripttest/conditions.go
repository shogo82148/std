// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scripttest

import (
	"github.com/shogo82148/std/cmd/internal/script"
	"github.com/shogo82148/std/testing"
)

// AddToolChainScriptConditions accepts a [script.Cond] map and adds into it a
// set of commonly used conditions for doing toolchains testing,
// including whether the platform supports cgo, a buildmode condition,
// support for GOEXPERIMENT testing, etc. Callers must also pass in
// current GOHOSTOOS/GOHOSTARCH settings, since some of the conditions
// introduced can be influenced by them.
func AddToolChainScriptConditions(t *testing.T, conds map[string]script.Cond, goHostOS, goHostArch string)
