// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

import "github.com/shogo82148/std/errors"

// ErrRange indicates that a value is out of range for the target type.
var ErrRange = errors.New("value out of range")

// ErrSyntax indicates that a value does not have the right syntax for the target type.
var ErrSyntax = errors.New("invalid syntax")

// A NumError records a failed conversion.
type NumError struct {
	Func string
	Num  string
	Err  error
}

func (e *NumError) Error() string

// IntSize is the size in bits of an int or uint value.
const IntSize = intSize

// ParseUint is like ParseInt but for unsigned numbers.
func ParseUint(s string, base int, bitSize int) (n uint64, err error)

// ParseInt interprets a string s in the given base (2 to 36) and
// returns the corresponding value i.  If base == 0, the base is
// implied by the string's prefix: base 16 for "0x", base 8 for
// "0", and base 10 otherwise.
//
// The bitSize argument specifies the integer type
// that the result must fit into.  Bit sizes 0, 8, 16, 32, and 64
// correspond to int, int8, int16, int32, and int64.
//
// The errors that ParseInt returns have concrete type *NumError
// and include err.Num = s.  If s is empty or contains invalid
// digits, err.Err = ErrSyntax; if the value corresponding
// to s cannot be represented by a signed integer of the
// given size, err.Err = ErrRange.
func ParseInt(s string, base int, bitSize int) (i int64, err error)

// Atoi is shorthand for ParseInt(s, 10, 0).
func Atoi(s string) (i int, err error)
