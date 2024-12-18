// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

// Tool returns the path to the named builtin tool (for example, "vet").
// If the tool cannot be found, Tool exits the process.
func Tool(toolName string) string

// ToolPath returns the path at which we expect to find the named tool
// (for example, "vet"), and the error (if any) from statting that path.
func ToolPath(toolName string) (string, error)
