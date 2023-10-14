// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements FindExportData.

package importer

import (
	"github.com/shogo82148/std/bufio"
)

// FindExportData positions the reader r at the beginning of the
// export data section of an underlying GC-created object/archive
// file by reading from it. The reader must be positioned at the
// start of the file before calling this function. The hdr result
// is the string before the export data, either "$$" or "$$B".
//
// If size is non-negative, it's the number of bytes of export data
// still available to read from r.
func FindExportData(r *bufio.Reader) (hdr string, size int, err error)
