// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// This program is run via "go generate" (via a directive in sort.go)
// to generate implementation variants of the underlying sorting algorithm.
// When passed the -generic flag it generates generic variants of sorting;
// otherwise it generates the non-generic variants used by the sort package.

package main

import (
	"github.com/shogo82148/std/text/template"
)

type Variant struct {
	Name string

	Path string

	Package string

	Imports string

	FuncSuffix string

	DataType string

	TypeParam string

	ExtraParam string

	ExtraArg string

	Funcs template.FuncMap
}
