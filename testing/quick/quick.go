// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package quick implements utility functions to help with black box testing.
//
// The testing/quick package is frozen and is not accepting new features.
package quick

import (
	"github.com/shogo82148/std/math/rand"
	"github.com/shogo82148/std/reflect"
)

// A Generator can generate random values of its own type.
type Generator interface {
	Generate(rand *rand.Rand, size int) reflect.Value
}

// Value returns an arbitrary value of the given type.
// If the type implements the [Generator] interface, that will be used.
// Note: To create arbitrary values for structs, all the fields must be exported.
func Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool)

// A Config structure contains options for running a test.
type Config struct {
	// MaxCount sets the maximum number of iterations.
	// If zero, MaxCountScale is used.
	MaxCount int
	// MaxCountScale is a non-negative scale factor applied to the
	// default maximum.
	// A count of zero implies the default, which is usually 100
	// but can be set by the -quickchecks flag.
	MaxCountScale float64
	// Rand specifies a source of random numbers.
	// If nil, a default pseudo-random source will be used.
	Rand *rand.Rand
	// Values specifies a function to generate a slice of
	// arbitrary reflect.Values that are congruent with the
	// arguments to the function being tested.
	// If nil, the top-level Value function is used to generate them.
	Values func([]reflect.Value, *rand.Rand)
}

// A SetupError is the result of an error in the way that check is being
// used, independent of the functions being tested.
type SetupError string

func (s SetupError) Error() string

// A CheckError is the result of Check finding an error.
type CheckError struct {
	Count int
	In    []any
}

func (s *CheckError) Error() string

// A CheckEqualError is the result [CheckEqual] finding an error.
type CheckEqualError struct {
	CheckError
	Out1 []any
	Out2 []any
}

func (s *CheckEqualError) Error() string

// Check looks for an input to f, any function that returns bool,
// such that f returns false. It calls f repeatedly, with arbitrary
// values for each argument. If f returns false on a given input,
// Check returns that input as a *[CheckError].
// For example:
//
//	func TestOddMultipleOfThree(t *testing.T) {
//		f := func(x int) bool {
//			y := OddMultipleOfThree(x)
//			return y%2 == 1 && y%3 == 0
//		}
//		if err := quick.Check(f, nil); err != nil {
//			t.Error(err)
//		}
//	}
func Check(f any, config *Config) error

// CheckEqual looks for an input on which f and g return different results.
// It calls f and g repeatedly with arbitrary values for each argument.
// If f and g return different answers, CheckEqual returns a *[CheckEqualError]
// describing the input and the outputs.
func CheckEqual(f, g any, config *Config) error
