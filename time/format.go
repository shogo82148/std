// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

// These are predefined layouts for use in Time.Format.
// The standard time used in the layouts is:
//
//	Mon Jan 2 15:04:05 MST 2006
//
// which is Unix time 1136243045. Since MST is GMT-0700,
// the standard time can be thought of as
//
//	01/02 03:04:05PM '06 -0700
//
// To define your own format, write down what the standard time would look
// like formatted your way; see the values of constants like ANSIC,
// StampMicro or Kitchen for examples.
//
// Within the format string, an underscore _ represents a space that may be
// replaced by a digit if the following number (a day) has two digits; for
// compatibility with fixed-width Unix time formats.
//
// A decimal point followed by one or more zeros represents a fractional
// second, printed to the given number of decimal places.  A decimal point
// followed by one or more nines represents a fractional second, printed to
// the given number of decimal places, with trailing zeros removed.
// When parsing (only), the input may contain a fractional second
// field immediately after the seconds field, even if the layout does not
// signify its presence. In that case a decimal point followed by a maximal
// series of digits is parsed as a fractional second.
//
// Numeric time zone offsets format as follows:
//
//	-0700  ±hhmm
//	-07:00 ±hh:mm
//
// Replacing the sign in the format with a Z triggers
// the ISO 8601 behavior of printing Z instead of an
// offset for the UTC zone.  Thus:
//
//	Z0700  Z or ±hhmm
//	Z07:00 Z or ±hh:mm
const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700"
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)

// Never printed, just needs to be non-nil for return by atoi.

// String returns the time formatted using the format string
//
//	"2006-01-02 15:04:05.999999999 -0700 MST"
func (t Time) String() string

// Format returns a textual representation of the time value formatted
// according to layout.  The layout defines the format by showing the
// representation of the standard time,
//
//	Mon Jan 2 15:04:05 -0700 MST 2006
//
// which is then used to describe the time to be formatted. Predefined
// layouts ANSIC, UnixDate, RFC3339 and others describe standard
// representations. For more information about the formats and the
// definition of the standard time, see the documentation for ANSIC.
func (t Time) Format(layout string) string

// ParseError describes a problem parsing a time string.
type ParseError struct {
	Layout     string
	Value      string
	LayoutElem string
	ValueElem  string
	Message    string
}

// Error returns the string representation of a ParseError.
func (e *ParseError) Error() string

// Parse parses a formatted string and returns the time value it represents.
// The layout defines the format by showing the representation of the
// standard time,
//
//	Mon Jan 2 15:04:05 -0700 MST 2006
//
// which is then used to describe the string to be parsed. Predefined layouts
// ANSIC, UnixDate, RFC3339 and others describe standard representations. For
// more information about the formats and the definition of the standard
// time, see the documentation for ANSIC.
//
// Elements omitted from the value are assumed to be zero or, when
// zero is impossible, one, so parsing "3:04pm" returns the time
// corresponding to Jan 1, year 0, 15:04:00 UTC.
// Years must be in the range 0000..9999. The day of the week is checked
// for syntax but it is otherwise ignored.
func Parse(layout, value string) (Time, error)

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func ParseDuration(s string) (Duration, error)
