// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Indexed package import.
// See iexport.go for the export data format.

package typecheck

import (
	"github.com/shogo82148/std/cmd/compile/internal/base"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

// HaveInlineBody reports whether we have fn's inline body available
// for inlining.
//
// It's a function literal so that it can be overridden for
// GOEXPERIMENT=unified.
var HaveInlineBody = func(fn *ir.Func) bool {
	base.Fatalf("HaveInlineBody not overridden")
	panic("unreachable")
}

func SetBaseTypeIndex(t *types.Type, i, pi int64)

func BaseTypeIndex(t *types.Type) int64
