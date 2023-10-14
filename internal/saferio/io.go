// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package saferio provides I/O functions that avoid allocating large
// amounts of memory unnecessarily. This is intended for packages that
// read data from an [io.Reader] where the size is part of the input
// data but the input may be corrupt, or may be provided by an
// untrustworthy attacker.
package saferio

import (
	"github.com/shogo82148/std/io"
)

// ReadData reads n bytes from the input stream, but avoids allocating
// all n bytes if n is large. This avoids crashing the program by
// allocating all n bytes in cases where n is incorrect.
//
// The error is io.EOF only if no bytes were read.
// If an io.EOF happens after reading some but not all the bytes,
// ReadData returns io.ErrUnexpectedEOF.
func ReadData(r io.Reader, n uint64) ([]byte, error)

// ReadDataAt reads n bytes from the input stream at off, but avoids
// allocating all n bytes if n is large. This avoids crashing the program
// by allocating all n bytes in cases where n is incorrect.
func ReadDataAt(r io.ReaderAt, n uint64, off int64) ([]byte, error)

// SliceCapWithSize returns the capacity to use when allocating a slice.
// After the slice is allocated with the capacity, it should be
// built using append. This will avoid allocating too much memory
// if the capacity is large and incorrect.
//
// A negative result means that the value is always too big.
func SliceCapWithSize(size, c uint64) int

// SliceCap is like SliceCapWithSize but using generics.
func SliceCap[E any](c uint64) int
