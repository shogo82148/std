// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rand implements pseudo-random number generators.
package rand

// A Source represents a source of uniformly-distributed
// pseudo-random int64 values in the range [0, 1<<63).
type Source interface {
	Int63() int64
	Seed(seed int64)
}

// NewSource returns a new pseudo-random Source seeded with the given value.
func NewSource(seed int64) Source

// A Rand is a source of random numbers.
type Rand struct {
	src Source
}

// New returns a new Rand that uses random values from src
// to generate other random values.
func New(src Source) *Rand

// Seed uses the provided seed value to initialize the generator to a deterministic state.
func (r *Rand) Seed(seed int64)

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand) Int63() int64

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r *Rand) Int31() int32

// Int returns a non-negative pseudo-random int.
func (r *Rand) Int() int

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int63n(n int64) int64

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int31n(n int32) int32

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Intn(n int) int

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float64() float64

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float32() float32

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func (r *Rand) Perm(n int) []int

// Seed uses the provided seed value to initialize the generator to a
// deterministic state. If Seed is not called, the generator behaves as
// if seeded by Seed(1).
func Seed(seed int64)

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func Int63() int64

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func Uint32() uint32

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func Int31() int32

// Int returns a non-negative pseudo-random int.
func Int() int

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Int63n(n int64) int64

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Int31n(n int32) int32

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Intn(n int) int

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func Float64() float64

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func Float32() float32

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func Perm(n int) []int

// NormFloat64 returns a normally distributed float64 in the range
// [-math.MaxFloat64, +math.MaxFloat64] with
// standard normal distribution (mean = 0, stddev = 1).
// To produce a different normal distribution, callers can
// adjust the output using:
//
//	sample = NormFloat64() * desiredStdDev + desiredMean
func NormFloat64() float64

// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1).
// To produce a distribution with a different rate parameter,
// callers can adjust the output using:
//
//	sample = ExpFloat64() / desiredRateParameter
func ExpFloat64() float64
