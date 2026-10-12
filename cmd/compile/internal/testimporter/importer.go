// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testimporter

import (
	"github.com/shogo82148/std/sync"

	"github.com/shogo82148/std/cmd/compile/internal/types2"
)

// Importer implements a types2 importer for use in testing by calling "go
// build". It is safe for concurrent use; sharing importers can yield better
// performance. It understands the compiler-internal unified export formats.
type Importer struct {
	dir      string
	mu       sync.Mutex
	readPkgs map[string]*types2.Package
	bldOnces map[string]*sync.Once
	bldCache map[string]*bldResult
}

// NewImporter returns a new Importer.
func NewImporter() *Importer

// Import implements types2.Importer.
func (imp *Importer) Import(path string) (*types2.Package, error)

// ImportFrom implements types2.ImportFrom.
func (imp *Importer) ImportFrom(path, srcDir string, mode types2.ImportMode) (*types2.Package, error)
