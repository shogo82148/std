// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package entropy implements a CPU jitter-based SP 800-90B entropy source.
package entropy

// Version returns the version of the entropy source.
//
// This is independent of the FIPS 140-3 module version, in order to reuse the
// ESV certificate across module versions.
func Version() string

// ScratchBuffer is a large buffer that will be written to using atomics, to
// generate noise from memory access timings. Its contents do not matter.
type ScratchBuffer [1 << 25]byte

// Seed returns a 384-bit seed with full entropy.
//
// memory is passed in to allow changing the allocation strategy without
// modifying the frozen and certified entropy source in this package.
//
// Seed returns an error if the entropy source startup health tests fail, which
// has a non-negligible chance of happening.
func Seed(memory *ScratchBuffer) ([48]byte, error)

// Samples starts a new entropy source, collects the requested number of
// samples, conducts startup health tests, and returns the samples or an error
// if the health tests fail.
//
// The health tests have a non-negligible chance of failing.
func Samples(samples []uint8, memory *ScratchBuffer) error

// RepetitionCountTest implements the repetition count test from SP 800-90B
// Section 4.4.1. It returns an error if any symbol is repeated C = 41 or more
// times in a row.
//
// This C value is calculated from a target failure probability α = 2⁻²⁰ and a
// claimed min-entropy per symbol h = 0.5 bits, using the formula in SP 800-90B
// Section 4.4.1.
//
//	sage: α = 2^-20
//	sage: H = 0.5
//	sage: 1 + ceil(-log(α, 2) / H)
//	41
func RepetitionCountTest(samples []uint8) error

// AdaptiveProportionTest implements the adaptive proportion test from SP 800-90B
// Section 4.4.2. It returns an error if any symbol appears C = 410 or more
// times in the last W = 512 samples.
//
// This C value is calculated from a target failure probability α = 2⁻²⁰, a
// window size W = 512, and a claimed min-entropy per symbol h = 0.5 bits, using
// the formula in SP 800-90B Section 4.4.2, equivalent to the Microsoft Excel
// formula 1+CRITBINOM(W, power(2,(−H)),1−α).
//
//	sage: from scipy.stats import binom
//	sage: α = 2^-20
//	sage: H = 0.5
//	sage: W = 512
//	sage: C = 1 + binom.ppf(1 - α, W, 2**(-H))
//	sage: ceil(C)
//	410
func AdaptiveProportionTest(samples []uint8) error
