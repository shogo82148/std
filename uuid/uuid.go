// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package uuid provides support for generating and manipulating UUIDs.
//
// See [RFC 9562] for details.
//
// Random components of new UUIDs are generated with a
// cryptographically secure random number generator.
//
// UUIDs may be generated using various algorithms.
// The [New] function returns a new UUID generated using
// an algorithm suitable for most purposes.
//
// [RFC 9562]: https://www.rfc-editor.org/rfc/rfc9562.html
package uuid

// A UUID is a Universally Unique Identifier as specified in RFC 9562.
//
// UUIDs are comparable, such as with the == operator.
type UUID [16]byte

// Parse returns the UUID represented by s.
//
// It accepts strings in the following forms:
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6
//	{f81d4fae-7dec-11d0-a765-00a0c91e6bf6}
//	urn:uuid:f81d4fae-7dec-11d0-a765-00a0c91e6bf6
//	f81d4fae7dec11d0a76500a0c91e6bf6
//
// Alphabetic characters in the hexadecimal digits may be any case.
func Parse(s string) (UUID, error)

// MustParse returns the UUID represented by s.
//
// It panics if s is not a valid string representation of a UUID as defined by [Parse].
func MustParse(s string) UUID

// New returns a new UUID.
//
// Programs which do not have a need for a specific UUID generation algorithm should use New.
// At this time, New is equivalent to [NewV4].
func New() UUID

// Nil returns the Nil UUID 00000000-0000-0000-0000-000000000000.
//
// The Nil UUID is defined in [Section 5.9 of RFC 9562].
// Note that this is not the same as the Go value nil.
//
// [Section 5.9 of RFC 9562]: https://www.rfc-editor.org/rfc/rfc9562#section-5.9
func Nil() UUID

// Max returns the Max UUID ffffffff-ffff-ffff-ffff-ffffffffffff.
//
// The Max UUID is defined in [Section 5.10 of RFC 9562].
//
// [Section 5.10 of RFC 9562]: https://www.rfc-editor.org/rfc/rfc9562#section-5.10
func Max() UUID

// String returns the string representation of u.
//
// It uses the lowercase hex-and-dash representation defined in RFC 9562.
func (u UUID) String() string

// MarshalText implements the [encoding.TextMarshaler] interface.
// The encoding is the same as returned by [UUID.String]
func (u UUID) MarshalText() ([]byte, error)

// AppendText implements the [encoding.TextAppender] interface.
// The encoding is the same as returned by [UUID.String]
func (u UUID) AppendText(b []byte) ([]byte, error)

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
// The UUID is expected in a form accepted by [Parse].
func (u *UUID) UnmarshalText(b []byte) error

// Compare compares the UUID u with v.
// If u is before v, it returns -1.
// If u is after v, it returns +1.
// If they are the same, it returns 0.
//
// Compare uses the big-endian byte order defined in
// [Section 6.11 of RFC 9562] for sorting.
//
// [Section 6.11 of RFC 9562]: https://www.rfc-editor.org/rfc/rfc9562#section-6.11
func (u UUID) Compare(v UUID) int

// NewV4 returns a new version 4 UUID.
//
// Version 4 UUIDs contain 122 bits of random data.
func NewV4() UUID

// NewV7 returns a new version 7 UUID.
//
// Version 7 UUIDs contain a timestamp in the most significant 48 bits,
// and at least 62 bits of random data.
//
// NewV7 always returns UUIDs which sort in increasing order,
// except when the system clock moves backwards.
func NewV7() UUID
