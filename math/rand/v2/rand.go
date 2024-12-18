// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rand implements pseudo-random number generators suitable for tasks
// such as simulation, but it should not be used for security-sensitive work.
//
// Random numbers are generated by a [Source], usually wrapped in a [Rand].
// Both types should be used by a single goroutine at a time: sharing among
// multiple goroutines requires some kind of synchronization.
//
// Top-level functions, such as [Float64] and [Int],
// are safe for concurrent use by multiple goroutines.
//
// This package's outputs might be easily predictable regardless of how it's
// seeded. For random numbers suitable for security-sensitive work, see the
// [crypto/rand] package.
package rand

// A Source is a source of uniformly-distributed
// pseudo-random uint64 values in the range [0, 1<<64).
//
// A Source is not safe for concurrent use by multiple goroutines.
type Source interface {
	Uint64() uint64
}

// A Rand is a source of random numbers.
type Rand struct {
	src Source
}

// New returns a new Rand that uses random values from src
// to generate other random values.
func New(src Source) *Rand

// Int64 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand) Int64() int64

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (r *Rand) Uint64() uint64

// Int32 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r *Rand) Int32() int32

// Int returns a non-negative pseudo-random int.
func (r *Rand) Int() int

// Uint returns a pseudo-random uint.
func (r *Rand) Uint() uint

// Int64N returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (r *Rand) Int64N(n int64) int64

// Uint64N returns, as a uint64, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n == 0.
func (r *Rand) Uint64N(n uint64) uint64

// Int32N returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (r *Rand) Int32N(n int32) int32

// Uint32N returns, as a uint32, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n == 0.
func (r *Rand) Uint32N(n uint32) uint32

// IntN returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (r *Rand) IntN(n int) int

// UintN returns, as a uint, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n == 0.
func (r *Rand) UintN(n uint) uint

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0).
func (r *Rand) Float64() float64

// Float32 returns, as a float32, a pseudo-random number in the half-open interval [0.0,1.0).
func (r *Rand) Float32() float32

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers
// in the half-open interval [0,n).
func (r *Rand) Perm(n int) []int

// Shuffle pseudo-randomizes the order of elements.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func (r *Rand) Shuffle(n int, swap func(i, j int))

// Int64 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default Source.
func Int64() int64

// Uint32 returns a pseudo-random 32-bit value as a uint32
// from the default Source.
func Uint32() uint32

// Uint64N returns, as a uint64, a pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n == 0.
func Uint64N(n uint64) uint64

// Uint32N returns, as a uint32, a pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n == 0.
func Uint32N(n uint32) uint32

// Uint64 returns a pseudo-random 64-bit value as a uint64
// from the default Source.
func Uint64() uint64

// Int32 returns a non-negative pseudo-random 31-bit integer as an int32
// from the default Source.
func Int32() int32

// Int returns a non-negative pseudo-random int from the default Source.
func Int() int

// Uint returns a pseudo-random uint from the default Source.
func Uint() uint

// Int64N returns, as an int64, a pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n <= 0.
func Int64N(n int64) int64

// Int32N returns, as an int32, a pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n <= 0.
func Int32N(n int32) int32

// IntN returns, as an int, a pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n <= 0.
func IntN(n int) int

// UintN returns, as a uint, a pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n == 0.
func UintN(n uint) uint

// N returns a pseudo-random number in the half-open interval [0,n) from the default Source.
// The type parameter Int can be any integer type.
// It panics if n <= 0.
func N[Int intType](n Int) Int

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0)
// from the default Source.
func Float64() float64

// Float32 returns, as a float32, a pseudo-random number in the half-open interval [0.0,1.0)
// from the default Source.
func Float32() float32

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers
// in the half-open interval [0,n) from the default Source.
func Perm(n int) []int

// Shuffle pseudo-randomizes the order of elements using the default Source.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func Shuffle(n int, swap func(i, j int))

// NormFloat64 returns a normally distributed float64 in the range
// [-math.MaxFloat64, +math.MaxFloat64] with
// standard normal distribution (mean = 0, stddev = 1)
// from the default Source.
// To produce a different normal distribution, callers can
// adjust the output using:
//
//	sample = NormFloat64() * desiredStdDev + desiredMean
func NormFloat64() float64

// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1) from the default Source.
// To produce a distribution with a different rate parameter,
// callers can adjust the output using:
//
//	sample = ExpFloat64() / desiredRateParameter
func ExpFloat64() float64
