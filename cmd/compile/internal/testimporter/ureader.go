// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package testimporter implements reading of compiler-internal unified export
// formats.
package testimporter

import (
	"github.com/shogo82148/std/cmd/compile/internal/types2"
	"github.com/shogo82148/std/internal/pkgbits"
)

func ReadPackage(ctxt *types2.Context, imports map[string]*types2.Package, input pkgbits.PkgDecoder) *types2.Package
