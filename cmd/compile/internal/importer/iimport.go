// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Indexed package import.
// See cmd/compile/internal/typecheck/iexport.go for the export data format.

package importer

import (
	"github.com/shogo82148/std/cmd/compile/internal/types2"
)

// ImportData imports a package from the serialized package data
// and returns the number of bytes consumed and a reference to the package.
// If the export data version is not recognized or the format is otherwise
// compromised, an error is returned.
func ImportData(imports map[string]*types2.Package, data, path string) (pkg *types2.Package, err error)
