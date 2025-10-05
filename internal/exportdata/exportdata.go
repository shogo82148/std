// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package exportdata implements common utilities for finding
// and reading gc-generated object files.
package exportdata

import (
	"github.com/shogo82148/std/bufio"
)

// ReadUnified reads the contents of the unified export data from a reader r
// that contains the contents of a GC-created archive file.
//
// On success, the reader will be positioned after the end-of-section marker "\n$$\n".
//
// Supported GC-created archive files have 4 layers of nesting:
//   - An archive file containing a package definition file.
//   - The package definition file contains headers followed by a data section.
//     Headers are lines (≤ 4kb) that do not start with "$$".
//   - The data section starts with "$$B\n" followed by export data followed
//     by an end of section marker "\n$$\n". (The section start "$$\n" is no
//     longer supported.)
//   - The export data starts with a format byte ('u') followed by the <data> in
//     the given format. (See ReadExportDataHeader for older formats.)
//
// Putting this together, the bytes in a GC-created archive files are expected
// to look like the following.
// See cmd/internal/archive for more details on ar file headers.
//
// | <!arch>\n             | ar file signature
// | __.PKGDEF...size...\n | ar header for __.PKGDEF including size.
// | go object <...>\n     | objabi header
// | <optional headers>\n  | other headers such as build id
// | $$B\n                 | binary format marker
// | u<data>\n             | unified export <data>
// | $$\n                  | end-of-section marker
// | [optional padding]    | padding byte (0x0A) if size is odd
// | [ar file header]      | other ar files
// | [ar file data]        |
func ReadUnified(r *bufio.Reader) (data []byte, err error)

// FindPackageDefinition positions the reader r at the beginning of a package
// definition file ("__.PKGDEF") within a GC-created archive by reading
// from it, and returns the size of the package definition file in the archive.
//
// The reader must be positioned at the start of the archive file before calling
// this function, and "__.PKGDEF" is assumed to be the first file in the archive.
//
// See cmd/internal/archive for details on the archive format.
func FindPackageDefinition(r *bufio.Reader) (size int, err error)

// ReadObjectHeaders reads object headers from the reader. Object headers are
// lines that do not start with an end-of-section marker "$$". The first header
// is the objabi header. On success, the reader will be positioned at the beginning
// of the end-of-section marker.
//
// It returns an error if any header does not fit in r.Size() bytes.
func ReadObjectHeaders(r *bufio.Reader) (objapi string, headers []string, err error)

// ReadExportDataHeader reads the export data header and format from r.
// It returns the number of bytes read, or an error if the format is no longer
// supported or it failed to read.
//
// The only currently supported format is binary export data in the
// unified export format.
func ReadExportDataHeader(r *bufio.Reader) (n int, err error)

// FindPkg returns the filename and unique package id for an import
// path based on package information provided by build.Import (using
// the build.Default build.Context). A relative srcDir is interpreted
// relative to the current working directory.
func FindPkg(path, srcDir string) (filename, id string, err error)
