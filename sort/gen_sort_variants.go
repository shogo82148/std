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
	// Name is the variant name: should be unique among variants.
	Name string

	// Path is the file path into which the generator will emit the code for this
	// variant.
	Path string

	// Package is the package this code will be emitted into.
	Package string

	// Imports is the imports needed for this package.
	Imports string

	// FuncSuffix is appended to all function names in this variant's code. All
	// suffixes should be unique within a package.
	FuncSuffix string

	// DataType is the type of the data parameter of functions in this variant's
	// code.
	DataType string

	// TypeParam is the optional type parameter for the function.
	TypeParam string

	// ExtraParam is an extra parameter to pass to the function. Should begin with
	// ", " to separate from other params.
	ExtraParam string

	// ExtraArg is an extra argument to pass to calls between functions; typically
	// it invokes ExtraParam. Should begin with ", " to separate from other args.
	ExtraArg string

	// Funcs is a map of functions used from within the template. The following
	// functions are expected to exist:
	//
	//    Less (name, i, j):
	//      emits a comparison expression that checks if the value `name` at
	//      index `i` is smaller than at index `j`.
	//
	//    Swap (name, i, j):
	//      emits a statement that performs a data swap between elements `i` and
	//      `j` of the value `name`.
	Funcs template.FuncMap
}
