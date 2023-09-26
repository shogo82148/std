// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run genzfunc.go

// Package sort provides primitives for sorting slices and user-defined
// collections.
package sort

// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
	Len() int

	Less(i, j int) bool

	Swap(i, j int)
}

// Sort sorts data.
// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func Sort(data Interface)

// lessSwap is a pair of Less and Swap function for use with the
// auto-generated func-optimized variant of sort.go in
// zfuncversion.go.

// Slice sorts the provided slice given the provided less function.
//
// The sort is not guaranteed to be stable. For a stable sort, use
// SliceStable.
//
// The function panics if the provided interface is not a slice.
func Slice(slice interface{}, less func(i, j int) bool)

// SliceStable sorts the provided slice given the provided less
// function while keeping the original order of equal elements.
//
// The function panics if the provided interface is not a slice.
func SliceStable(slice interface{}, less func(i, j int) bool)

// SliceIsSorted tests whether a slice is sorted.
//
// The function panics if the provided interface is not a slice.
func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool

// Reverse returns the reverse order for data.
func Reverse(data Interface) Interface

// IsSorted reports whether data is sorted.
func IsSorted(data Interface) bool

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int

func (p IntSlice) Len() int
func (p IntSlice) Less(i, j int) bool
func (p IntSlice) Swap(i, j int)

// Sort is a convenience method.
func (p IntSlice) Sort()

// Float64Slice attaches the methods of Interface to []float64, sorting in increasing order
// (not-a-number values are treated as less than other values).
type Float64Slice []float64

func (p Float64Slice) Len() int
func (p Float64Slice) Less(i, j int) bool
func (p Float64Slice) Swap(i, j int)

// Sort is a convenience method.
func (p Float64Slice) Sort()

// StringSlice attaches the methods of Interface to []string, sorting in increasing order.
type StringSlice []string

func (p StringSlice) Len() int
func (p StringSlice) Less(i, j int) bool
func (p StringSlice) Swap(i, j int)

// Sort is a convenience method.
func (p StringSlice) Sort()

// Ints sorts a slice of ints in increasing order.
func Ints(a []int)

// Float64s sorts a slice of float64s in increasing order
// (not-a-number values are treated as less than other values).
func Float64s(a []float64)

// Strings sorts a slice of strings in increasing order.
func Strings(a []string)

// IntsAreSorted tests whether a slice of ints is sorted in increasing order.
func IntsAreSorted(a []int) bool

// Float64sAreSorted tests whether a slice of float64s is sorted in increasing order
// (not-a-number values are treated as less than other values).
func Float64sAreSorted(a []float64) bool

// StringsAreSorted tests whether a slice of strings is sorted in increasing order.
func StringsAreSorted(a []string) bool

// Stable sorts data while keeping the original order of equal elements.
//
// It makes one call to data.Len to determine n, O(n*log(n)) calls to
// data.Less and O(n*log(n)*log(n)) calls to data.Swap.
func Stable(data Interface)
