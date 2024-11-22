// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// generrordocs creates a Markdown file for each (compiler) error code
// and its associated documentation.
// Note: this program must be run in this directory.
//   go run generrordocs.go <dir>

//go:generate go run generrordocs.go errors_markdown

package main
