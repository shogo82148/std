// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

// These are predefined layouts for use in Time.Format and time.Parse.
// The reference time used in these layouts is the specific time stamp:
//
//	01/02 03:04:05PM '06 -0700
//
// (January 2, 15:04:05, 2006, in time zone seven hours west of GMT).
// That value is recorded as the constant named Layout, listed below. As a Unix
// time, this is 1136239445. Since MST is GMT-0700, the reference would be
// printed by the Unix date command as:
//
//	Mon Jan 2 15:04:05 MST 2006
//
// It is a regrettable historic error that the date uses the American convention
// of putting the numerical month before the day.
//
// The example for Time.Format demonstrates the working of the layout string
// in detail and is a good reference.
//
// Note that the RFC822, RFC850, and RFC1123 formats should be applied
// only to local times. Applying them to UTC times will use "UTC" as the
// time zone abbreviation, while strictly speaking those RFCs require the
// use of "GMT" in that case.
// In general RFC1123Z should be used instead of RFC1123 for servers
// that insist on that format, and RFC3339 should be preferred for new protocols.
// RFC3339, RFC822, RFC822Z, RFC1123, and RFC1123Z are useful for formatting;
// when used with time.Parse they do not accept all the time formats
// permitted by the RFCs and they do accept time formats not formally defined.
// The RFC3339Nano format removes trailing zeros from the seconds field
// and thus may not sort correctly once formatted.
//
// Most programs can use one of the defined constants as the layout passed to
// Format or Parse. The rest of this comment can be ignored unless you are
// creating a custom layout string.
//
// To define your own format, write down what the reference time would look like
// formatted your way; see the values of constants like ANSIC, StampMicro or
// Kitchen for examples. The model is to demonstrate what the reference time
// looks like so that the Format and Parse methods can apply the same
// transformation to a general time value.
//
// Here is a summary of the components of a layout string. Each element shows by
// example the formatting of an element of the reference time. Only these values
// are recognized. Text in the layout string that is not recognized as part of
// the reference time is echoed verbatim during Format and expected to appear
// verbatim in the input to Parse.
//
//	Year: "2006" "06"
//	Month: "Jan" "January"
//	Textual day of the week: "Mon" "Monday"
//	Numeric day of the month: "2" "_2" "02"
//	Numeric day of the year: "__2" "002"
//	Hour: "15" "3" "03" (PM or AM)
//	Minute: "4" "04"
//	Second: "5" "05"
//	AM/PM mark: "PM"
//
// Numeric time zone offsets format as follows:
//
//	"-0700"  ±hhmm
//	"-07:00" ±hh:mm
//	"-07"    ±hh
//
// Replacing the sign in the format with a Z triggers
// the ISO 8601 behavior of printing Z instead of an
// offset for the UTC zone. Thus:
//
//	"Z0700"  Z or ±hhmm
//	"Z07:00" Z or ±hh:mm
//	"Z07"    Z or ±hh
//
// Within the format string, the underscores in "_2" and "__2" represent spaces
// that may be replaced by digits if the following number has multiple digits,
// for compatibility with fixed-width Unix time formats. A leading zero represents
// a zero-padded value.
//
// The formats __2 and 002 are space-padded and zero-padded
// three-character day of year; there is no unpadded day of year format.
//
// A comma or decimal point followed by one or more zeros represents
// a fractional second, printed to the given number of decimal places.
// A comma or decimal point followed by one or more nines represents
// a fractional second, printed to the given number of decimal places, with
// trailing zeros removed.
// For example "15:04:05,000" or "15:04:05.000" formats or parses with
// millisecond precision.
//
// Some valid layouts are invalid time values for time.Parse, due to formats
// such as _ for space padding and Z for zone information.
const (
	Layout      = "01/02 03:04:05PM '06 -0700"
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

const (
	_ = iota
)

// std0x records the std values for "01", "02", ..., "06".

// Never printed, just needs to be non-nil for return by atoi.

// String returns the time formatted using the format string
//
//	"2006-01-02 15:04:05.999999999 -0700 MST"
//
// If the time has a monotonic clock reading, the returned string
// includes a final field "m=±<value>", where value is the monotonic
// clock reading formatted as a decimal number of seconds.
//
// The returned string is meant for debugging; for a stable serialized
// representation, use t.MarshalText, t.MarshalBinary, or t.Format
// with an explicit format string.
func (t Time) String() string

// GoString implements fmt.GoStringer and formats t to be printed in Go source
// code.
func (t Time) GoString() string

// Format returns a textual representation of the time value formatted according
// to the layout defined by the argument. See the documentation for the
// constant called Layout to see how to represent the layout format.
//
// The executable example for Time.Format demonstrates the working
// of the layout string in detail and is a good reference.
func (t Time) Format(layout string) string

// AppendFormat is like Format but appends the textual
// representation to b and returns the extended buffer.
func (t Time) AppendFormat(b []byte, layout string) []byte

// ParseError describes a problem parsing a time string.
type ParseError struct {
	Layout     string
	Value      string
	LayoutElem string
	ValueElem  string
	Message    string
}

// These are borrowed from unicode/utf8 and strconv and replicate behavior in
// that package, since we can't take a dependency on either.

// Error returns the string representation of a ParseError.
func (e *ParseError) Error() string

// Parse parses a formatted string and returns the time value it represents.
// See the documentation for the constant called Layout to see how to
// represent the format. The second argument must be parseable using
// the format string (layout) provided as the first argument.
//
// The example for Time.Format demonstrates the working of the layout string
// in detail and is a good reference.
//
// When parsing (only), the input may contain a fractional second
// field immediately after the seconds field, even if the layout does not
// signify its presence. In that case either a comma or a decimal point
// followed by a maximal series of digits is parsed as a fractional second.
// Fractional seconds are truncated to nanosecond precision.
//
// Elements omitted from the layout are assumed to be zero or, when
// zero is impossible, one, so parsing "3:04pm" returns the time
// corresponding to Jan 1, year 0, 15:04:00 UTC (note that because the year is
// 0, this time is before the zero Time).
// Years must be in the range 0000..9999. The day of the week is checked
// for syntax but it is otherwise ignored.
//
// For layouts specifying the two-digit year 06, a value NN >= 69 will be treated
// as 19NN and a value NN < 69 will be treated as 20NN.
//
// The remainder of this comment describes the handling of time zones.
//
// In the absence of a time zone indicator, Parse returns a time in UTC.
//
// When parsing a time with a zone offset like -0700, if the offset corresponds
// to a time zone used by the current location (Local), then Parse uses that
// location and zone in the returned time. Otherwise it records the time as
// being in a fabricated location with time fixed at the given zone offset.
//
// When parsing a time with a zone abbreviation like MST, if the zone abbreviation
// has a defined offset in the current location, then that offset is used.
// The zone abbreviation "UTC" is recognized as UTC regardless of location.
// If the zone abbreviation is unknown, Parse records the time as being
// in a fabricated location with the given zone abbreviation and a zero offset.
// This choice means that such a time can be parsed and reformatted with the
// same layout losslessly, but the exact instant used in the representation will
// differ by the actual zone offset. To avoid such problems, prefer time layouts
// that use a numeric zone offset, or use ParseInLocation.
func Parse(layout, value string) (Time, error)

// ParseInLocation is like Parse but differs in two important ways.
// First, in the absence of time zone information, Parse interprets a time as UTC;
// ParseInLocation interprets the time as in the given location.
// Second, when given a zone offset or abbreviation, Parse tries to match it
// against the Local location; ParseInLocation uses the given location.
func ParseInLocation(layout, value string, loc *Location) (Time, error)

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func ParseDuration(s string) (Duration, error)
