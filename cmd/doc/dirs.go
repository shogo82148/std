// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Dirs is a structure for scanning the directory tree.
// Its Next method returns the next Go source directory it finds.
// Although it can be used to scan the tree multiple times, it
// only walks the tree once, caching the data it finds.
type Dirs struct {
	scan   chan string
	paths  []string
	offset int
}

// Reset puts the scan back at the beginning.
func (d *Dirs) Reset()

// Next returns the next directory in the scan. The boolean
// is false when the scan is done.
func (d *Dirs) Next() (string, bool)
