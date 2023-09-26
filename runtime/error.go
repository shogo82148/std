// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// The Error interface identifies a run time error.
type Error interface {
	error

	RuntimeError()
}

// A TypeAssertionError explains a failed type assertion.
type TypeAssertionError struct {
	_interface    *_type
	concrete      *_type
	asserted      *_type
	missingMethod string
}

func (*TypeAssertionError) RuntimeError()

func (e *TypeAssertionError) Error() string

// An errorString represents a runtime error described by a single string.

// plainError represents a runtime error described a string without
// the prefix "runtime error: " after invoking errorString.Error().
// See Issue #14965.

// A boundsError represents an indexing or slicing operation gone wrong.

// boundsErrorFmts provide error text for various out-of-bounds panics.
// Note: if you change these strings, you should adjust the size of the buffer
// in boundsError.Error below as well.

// boundsNegErrorFmts are overriding formats if x is negative. In this case there's no need to report y.
