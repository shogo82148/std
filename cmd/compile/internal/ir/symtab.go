// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

import (
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

// Syms holds known symbols.
var Syms symsStruct

// Pkgs holds known packages.
var Pkgs struct {
	Go           *types.Pkg
	Itab         *types.Pkg
	Runtime      *types.Pkg
	InternalMaps *types.Pkg
	Coverage     *types.Pkg
}
