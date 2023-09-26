// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen_sort_variants.go

// Package sort provides primitives for sorting slices and user-defined collections.
package sort

// An implementation of Interface can be sorted by the routines in this package.
// The methods refer to elements of the underlying collection by integer index.
type Interface interface {
	Len() int

	Less(i, j int) bool

	Swap(i, j int)
}

// Sort sorts data in ascending order as determined by the Less method.
// It makes one call to data.Len to determine n and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
//
// Note: in many situations, the newer slices.SortFunc function is more
// ergonomic and runs faster.
func Sort(data Interface)

// xorshift paper: https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf

// lessSwap is a pair of Less and Swap function for use with the
// auto-generated func-optimized variant of sort.go in
// zfuncversion.go.

// Reverse returns the reverse order for data.
func Reverse(data Interface) Interface

// IsSorted reports whether data is sorted.
//
// Note: in many situations, the newer slices.IsSortedFunc function is more
// ergonomic and runs faster.
func IsSorted(data Interface) bool

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int

func (x IntSlice) Len() int
func (x IntSlice) Less(i, j int) bool
func (x IntSlice) Swap(i, j int)

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x IntSlice) Sort()

// Float64Slice implements Interface for a []float64, sorting in increasing order,
// with not-a-number (NaN) values ordered before other values.
type Float64Slice []float64

func (x Float64Slice) Len() int

// Less reports whether x[i] should be ordered before x[j], as required by the sort Interface.
// Note that floating-point comparison by itself is not a transitive relation: it does not
// report a consistent ordering for not-a-number (NaN) values.
// This implementation of Less places NaN values before any others, by using:
//
//	x[i] < x[j] || (math.IsNaN(x[i]) && !math.IsNaN(x[j]))
func (x Float64Slice) Less(i, j int) bool
func (x Float64Slice) Swap(i, j int)

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x Float64Slice) Sort()

// StringSlice attaches the methods of Interface to []string, sorting in increasing order.
type StringSlice []string

func (x StringSlice) Len() int
func (x StringSlice) Less(i, j int) bool
func (x StringSlice) Swap(i, j int)

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x StringSlice) Sort()

// Ints sorts a slice of ints in increasing order.
//
// Note: consider using the newer slices.Sort function, which runs faster.
func Ints(x []int)

// Float64s sorts a slice of float64s in increasing order.
// Not-a-number (NaN) values are ordered before other values.
//
// Note: consider using the newer slices.Sort function, which runs faster.
func Float64s(x []float64)

// Strings sorts a slice of strings in increasing order.
//
// Note: consider using the newer slices.Sort function, which runs faster.
func Strings(x []string)

// IntsAreSorted reports whether the slice x is sorted in increasing order.
//
// Note: consider using the newer slices.IsSorted function, which runs faster.
func IntsAreSorted(x []int) bool

// Float64sAreSorted reports whether the slice x is sorted in increasing order,
// with not-a-number (NaN) values before any other values.
//
// Note: consider using the newer slices.IsSorted function, which runs faster.
func Float64sAreSorted(x []float64) bool

// StringsAreSorted reports whether the slice x is sorted in increasing order.
//
// Note: consider using the newer slices.IsSorted function, which runs faster.
func StringsAreSorted(x []string) bool

// Stable sorts data in ascending order as determined by the Less method,
// while keeping the original order of equal elements.
//
// It makes one call to data.Len to determine n, O(n*log(n)) calls to
// data.Less and O(n*log(n)*log(n)) calls to data.Swap.
//
// Note: in many situations, the newer slices.SortStableFunc function is more
// ergonomic and runs faster.
func Stable(data Interface)
