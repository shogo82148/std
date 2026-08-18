// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// FitsInBitsU reports whether x fits in b bits (unsigned).
func FitsInBitsU(x uint64, b uint) bool

// InitLimit sets initial constant limit for v.  This limit is based
// only on the operation itself, not any of its input arguments. This
// method is only used in two places, once when the prove pass startup
// and the other when a new ssa value is created, both for init. (unlike
// flowLimit, below, which computes additional constraints based on
// ranges of opcode arguments).
func InitLimit(v *Value) Limit

// a Limit records known upper and lower bounds for a value.
//
// If we have min>max or umin>umax, then this Limit is
// called "unsatisfiable". When we encounter such a Limit, we
// know that any code for which that Limit applies is unreachable.
// We don't particularly care how unsatisfiable limits propagate,
// including becoming satisfiable, because any optimization
// decisions based on those limits only apply to unreachable code.
type Limit struct {
	Min, Max   int64
	Umin, Umax uint64
}

func NoLimit() Limit

// If x and y can add without overflow or underflow
// (using b bits), SafeAdd returns x+y, true.
// Otherwise, returns 0, false.
func SafeAdd(x, y int64, b uint) (int64, bool)

// same as safeAdd but for subtraction.
func SafeSub(x, y int64, b uint) (int64, bool)

// same as safeAddU but for subtraction.
func SafeSubU(x, y uint64, b uint) (uint64, bool)

func ConvertIntWithBitsize[Target uint64 | int64, Source uint64 | int64](x Source, bitsize uint) Target

func NoLimitForBitsize(bitsize uint) Limit

func (l Limit) String() string

func (l Limit) Intersect(l2 Limit) Limit

func (l Limit) SignedMinMax(minimum, maximum int64) Limit

func (l Limit) UnsignedMin(m uint64) Limit

func (l Limit) UnsignedMax(m uint64) Limit

func (l Limit) UnsignedMinMax(minimum, maximum uint64) Limit

func (l Limit) MaybeZero() bool

func (l Limit) Nonnegative() bool

func (l Limit) Unsat() bool

// UnsignedFixedLeadingBits extracts the all the most significant fixed bits from the limit.
// fixed and count are an other way to represent a limit, you can convert them to a limit as follows:
//
//	umin = fixed
//	umax = fixed | (1<<(64-count) - 1)
//
// In order to be useful for bitmanip analysis fixed and count are a coarser tool than a limit:
// 1. the varying section (umax-umin) is always one less than a power of two
// 2. that section is naturally aligned inside the 64-bit space
func (l Limit) UnsignedFixedLeadingBits() (fixed uint64, count uint)

// Add returns the limit obtained by adding a value with limit l
// to a value with limit l2. The result must fit in b bits.
func (l Limit) Add(l2 Limit, b uint) Limit

// same as add but for subtraction.
func (l Limit) Sub(l2 Limit, b uint) Limit

// same as add but for multiplication.
func (l Limit) Mul(l2 Limit, b uint) Limit

// Similar to add, but compute 1 << l if it fits without overflow in b bits.
func (l Limit) Exp2(b uint) Limit

// Similar to add, but computes the complement of the limit for bitsize b.
func (l Limit) Com(b uint) Limit

// Similar to add, but computes the negation of the limit for bitsize b.
func (l Limit) Neg(b uint) Limit

// Similar to add, but computes the TrailingZeros of the limit for bitsize b.
func (l Limit) Ctz(b uint) Limit

// Similar to add, but computes the Len of the limit for bitsize b.
func (l Limit) Bitlen(b uint) Limit

// Similar to add, but computes the PopCount of the limit for bitsize b.
func (l Limit) Popcount(b uint) Limit

func (l Limit) ConstValue() (_ int64, ok bool)
