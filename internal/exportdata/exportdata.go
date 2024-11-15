// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package exportdata implements common utilities for finding
// and reading gc-generated object files.
package exportdata

import (
	"github.com/shogo82148/std/bufio"
)

// FindExportData positions the reader r at the beginning of the
// export data section of an underlying GC-created object/archive
// file by reading from it. The reader must be positioned at the
// start of the file before calling this function. The hdr result
// is the string before the export data, either "$$" or "$$B".
func FindExportData(r *bufio.Reader) (hdr string, size int, err error)

// FindPkg returns the filename and unique package id for an import
// path based on package information provided by build.Import (using
// the build.Default build.Context). A relative srcDir is interpreted
// relative to the current working directory.
func FindPkg(path, srcDir string) (filename, id string, err error)
