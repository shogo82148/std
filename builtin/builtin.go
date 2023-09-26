// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package builtin provides documentation for Go's predeclared identifiers.
The items documented here are not actually in package builtin
but their descriptions here allow godoc to present documentation
for the language's special identifiers.
*/
package builtin

// bool is the set of boolean values, true and false.

// true and false are the two untyped boolean values.

// uint8 is the set of all unsigned 8-bit integers.
// Range: 0 through 255.

// uint16 is the set of all unsigned 16-bit integers.
// Range: 0 through 65535.

// uint32 is the set of all unsigned 32-bit integers.
// Range: 0 through 4294967295.

// uint64 is the set of all unsigned 64-bit integers.
// Range: 0 through 18446744073709551615.

// int8 is the set of all signed 8-bit integers.
// Range: -128 through 127.

// int16 is the set of all signed 16-bit integers.
// Range: -32768 through 32767.

// int32 is the set of all signed 32-bit integers.
// Range: -2147483648 through 2147483647.

// int64 is the set of all signed 64-bit integers.
// Range: -9223372036854775808 through 9223372036854775807.

// float32 is the set of all IEEE-754 32-bit floating-point numbers.

// float64 is the set of all IEEE-754 64-bit floating-point numbers.

// complex64 is the set of all complex numbers with float32 real and
// imaginary parts.

// complex128 is the set of all complex numbers with float64 real and
// imaginary parts.

// string is the set of all strings of 8-bit bytes, conventionally but not
// necessarily representing UTF-8-encoded text. A string may be empty, but
// not nil. Values of string type are immutable.

// int is a signed integer type that is at least 32 bits in size. It is a
// distinct type, however, and not an alias for, say, int32.

// uint is an unsigned integer type that is at least 32 bits in size. It is a
// distinct type, however, and not an alias for, say, uint32.

// uintptr is an integer type that is large enough to hold the bit pattern of
// any pointer.

// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.

// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.

// iota is a predeclared identifier representing the untyped integer ordinal
// number of the current const specification in a (usually parenthesized)
// const declaration. It is zero-indexed.

// nil is a predeclared identifier representing the zero value for a
// pointer, channel, func, interface, map, or slice type.

// Type is here for the purposes of documentation only. It is a stand-in
// for any Go type, but represents the same type for any given function
// invocation.
type Type int

// Type1 is here for the purposes of documentation only. It is a stand-in
// for any Go type, but represents the same type for any given function
// invocation.
type Type1 int

// IntegerType is here for the purposes of documentation only. It is a stand-in
// for any integer type: int, uint, int8 etc.
type IntegerType int

// FloatType is here for the purposes of documentation only. It is a stand-in
// for either float type: float32 or float64.
type FloatType float32

// ComplexType is here for the purposes of documentation only. It is a
// stand-in for either complex type: complex64 or complex128.
type ComplexType complex64

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
