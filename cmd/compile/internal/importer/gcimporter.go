// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the FindPkg and Import functions for tests
// to use gc-generated object files.

package importer

import (
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/cmd/compile/internal/types2"
)

// FindPkg returns the filename and unique package id for an import
// path based on package information provided by build.Import (using
// the build.Default build.Context). A relative srcDir is interpreted
// relative to the current working directory.
//
// This function should only be used in tests.
func FindPkg(path, srcDir string) (filename, id string, err error)

// Import imports a gc-generated package given its import path and srcDir, adds
// the corresponding package object to the packages map, and returns the object.
// The packages map must contain all packages already imported.
//
// This function should only be used in tests.
func Import(packages map[string]*types2.Package, path, srcDir string, lookup func(path string) (io.ReadCloser, error)) (pkg *types2.Package, err error)
