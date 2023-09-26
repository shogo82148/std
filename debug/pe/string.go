// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pe

// StringTable is a COFF string table.
type StringTable []byte

// String extracts string from COFF string table st at offset start.
func (st StringTable) String(start uint32) (string, error)
