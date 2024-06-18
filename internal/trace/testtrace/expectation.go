// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testtrace

import (
	"github.com/shogo82148/std/regexp"
)

// Expectation represents the expected result of some operation.
type Expectation struct {
	failure      bool
	errorMatcher *regexp.Regexp
}

// ExpectSuccess returns an Expectation that trivially expects success.
func ExpectSuccess() *Expectation

// Check validates whether err conforms to the expectation. Returns
// an error if it does not conform.
//
// Conformance means that if failure is true, then err must be non-nil.
// If err is non-nil, then it must match errorMatcher.
func (e *Expectation) Check(err error) error

// ParseExpectation parses the serialized form of an Expectation.
func ParseExpectation(data []byte) (*Expectation, error)
