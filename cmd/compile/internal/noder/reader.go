// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package noder

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

// Shapify returns the shape type for targ.
//
// If basic is true, then the type argument is used to instantiate a
// type parameter whose constraint is a basic interface.
func Shapify(targ *types.Type, basic bool) *types.Type

// MakeWrappers constructs all wrapper methods needed for the target
// compilation unit.
func MakeWrappers(target *ir.Package)
