// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements textual dumping of arbitrary data structures
// for debugging purposes. The code is customized for Node graphs
// and may be used for an alternative view of the node structure.

package ir

import (
	"github.com/shogo82148/std/io"
)

// DumpAny is like FDumpAny but prints to stderr.
func DumpAny(root interface{}, filter string, depth int)

// FDumpAny prints the structure of a rooted data structure
// to w by depth-first traversal of the data structure.
//
// The filter parameter is a regular expression. If it is
// non-empty, only struct fields whose names match filter
// are printed.
//
// The depth parameter controls how deep traversal recurses
// before it returns (higher value means greater depth).
// If an empty field filter is given, a good depth default value
// is 4. A negative depth means no depth limit, which may be fine
// for small data structures or if there is a non-empty filter.
//
// In the output, Node structs are identified by their Op name
// rather than their type; struct fields with zero values or
// non-matching field names are omitted, and "…" means recursion
// depth has been reached or struct fields have been omitted.
func FDumpAny(w io.Writer, root interface{}, filter string, depth int)
