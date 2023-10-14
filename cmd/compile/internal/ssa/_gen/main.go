// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The gen command generates Go code (in the parent directory) for all
// the architecture-specific opcodes, blocks, and rewrites.
package main

type ArchsByName []arch

func (x ArchsByName) Len() int
func (x ArchsByName) Swap(i, j int)
func (x ArchsByName) Less(i, j int) bool
