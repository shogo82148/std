// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Run all SIMD-related code generators.
package main

// ToolSet is a [flag.Value] that accepts a comma-separated list of tool names.
// It rejects any tool names that aren't in the map. A list like "a,c" sets only
// "a" and "c" to true and all other tools to false. Alternatively, the list
// items may each start with + or -, which enables or disables (respectively)
// only the named tools.
type ToolSet map[string]bool

func (s ToolSet) String() string

func (s ToolSet) Set(list string) error
