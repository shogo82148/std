// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package time provides functionality for measuring and displaying time.
//
// The calendrical calculations always assume a Gregorian calendar, with
// no leap seconds.
//
// # Monotonic Clocks
//
// Operating systems provide both a “wall clock,” which is subject to
// changes for clock synchronization, and a “monotonic clock,” which is
// not. The general rule is that the wall clock is for telling time and
// the monotonic clock is for measuring time. Rather than split the API,
// in this package the Time returned by time.Now contains both a wall
// clock reading and a monotonic clock reading; later time-telling
// operations use the wall clock reading, but later time-measuring
// operations, specifically comparisons and subtractions, use the
// monotonic clock reading.
//
// For example, this code always computes a positive elapsed time of
// approximately 20 milliseconds, even if the wall clock is changed during
// the operation being timed:
//
//	start := time.Now()
//	... operation that takes 20 milliseconds ...
//	t := time.Now()
//	elapsed := t.Sub(start)
//
// Other idioms, such as time.Since(start), time.Until(deadline), and
// time.Now().Before(deadline), are similarly robust against wall clock
// resets.
//
// The rest of this section gives the precise details of how operations
// use monotonic clocks, but understanding those details is not required
// to use this package.
//
// The Time returned by time.Now contains a monotonic clock reading.
// If Time t has a monotonic clock reading, t.Add adds the same duration to
// both the wall clock and monotonic clock readings to compute the result.
// Because t.AddDate(y, m, d), t.Round(d), and t.Truncate(d) are wall time
// computations, they always strip any monotonic clock reading from their results.
// Because t.In, t.Local, and t.UTC are used for their effect on the interpretation
// of the wall time, they also strip any monotonic clock reading from their results.
// The canonical way to strip a monotonic clock reading is to use t = t.Round(0).
//
// If Times t and u both contain monotonic clock readings, the operations
// t.After(u), t.Before(u), t.Equal(u), t.Compare(u), and t.Sub(u) are carried out
// using the monotonic clock readings alone, ignoring the wall clock
// readings. If either t or u contains no monotonic clock reading, these
// operations fall back to using the wall clock readings.
//
// On some systems the monotonic clock will stop if the computer goes to sleep.
// On such a system, t.Sub(u) may not accurately reflect the actual
// time that passed between t and u.
//
// Because the monotonic clock reading has no meaning outside
// the current process, the serialized forms generated by t.GobEncode,
// t.MarshalBinary, t.MarshalJSON, and t.MarshalText omit the monotonic
// clock reading, and t.Format provides no format for it. Similarly, the
// constructors time.Date, time.Parse, time.ParseInLocation, and time.Unix,
// as well as the unmarshalers t.GobDecode, t.UnmarshalBinary.
// t.UnmarshalJSON, and t.UnmarshalText always create times with
// no monotonic clock reading.
//
// The monotonic clock reading exists only in Time values. It is not
// a part of Duration values or the Unix times returned by t.Unix and
// friends.
//
// Note that the Go == operator compares not just the time instant but
// also the Location and the monotonic clock reading. See the
// documentation for the Time type for a discussion of equality
// testing for Time values.
//
// For debugging, the result of t.String does include the monotonic
// clock reading if present. If t != u because of different monotonic clock readings,
// that difference will be visible when printing t.String() and u.String().
//
// # Timer Resolution
//
// Timer resolution varies depending on the Go runtime, the operating system
// and the underlying hardware.
// On Unix, the resolution is approximately 1ms.
// On Windows, the default resolution is approximately 16ms, but
// a higher resolution may be requested using [golang.org/x/sys/windows.TimeBeginPeriod].
package time

import (
	_ "github.com/shogo82148/std/unsafe"
)

// A Time represents an instant in time with nanosecond precision.
//
// Programs using times should typically store and pass them as values,
// not pointers. That is, time variables and struct fields should be of
// type time.Time, not *time.Time.
//
// A Time value can be used by multiple goroutines simultaneously except
// that the methods GobDecode, UnmarshalBinary, UnmarshalJSON and
// UnmarshalText are not concurrency-safe.
//
// Time instants can be compared using the Before, After, and Equal methods.
// The Sub method subtracts two instants, producing a Duration.
// The Add method adds a Time and a Duration, producing a Time.
//
// The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
// As this time is unlikely to come up in practice, the IsZero method gives
// a simple way of detecting a time that has not been initialized explicitly.
//
// Each time has an associated Location. The methods Local, UTC, and In return a
// Time with a specific Location. Changing the Location of a Time value with
// these methods does not change the actual instant it represents, only the time
// zone in which to interpret it.
//
// Representations of a Time value saved by the GobEncode, MarshalBinary,
// MarshalJSON, and MarshalText methods store the Time.Location's offset, but not
// the location name. They therefore lose information about Daylight Saving Time.
//
// In addition to the required “wall clock” reading, a Time may contain an optional
// reading of the current process's monotonic clock, to provide additional precision
// for comparison or subtraction.
// See the “Monotonic Clocks” section in the package documentation for details.
//
// Note that the Go == operator compares not just the time instant but also the
// Location and the monotonic clock reading. Therefore, Time values should not
// be used as map or database keys without first guaranteeing that the
// identical Location has been set for all values, which can be achieved
// through use of the UTC or Local method, and that the monotonic clock reading
// has been stripped by setting t = t.Round(0). In general, prefer t.Equal(u)
// to t == u, since t.Equal uses the most accurate comparison available and
// correctly handles the case when only one of its arguments has a monotonic
// clock reading.
type Time struct {
	wall uint64
	ext  int64

	loc *Location
}

// After reports whether the time instant t is after u.
func (t Time) After(u Time) bool

// Before reports whether the time instant t is before u.
func (t Time) Before(u Time) bool

// Compare compares the time instant t with u. If t is before u, it returns -1;
// if t is after u, it returns +1; if they're the same, it returns 0.
func (t Time) Compare(u Time) int

// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with
// Time values; most code should use Equal instead.
func (t Time) Equal(u Time) bool

// A Month specifies a month of the year (January = 1, ...).
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// String returns the English name of the month ("January", "February", ...).
func (m Month) String() string

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// String returns the English name of the day ("Sunday", "Monday", ...).
func (d Weekday) String() string

// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
func (t Time) IsZero() bool

// Date returns the year, month, and day in which t occurs.
func (t Time) Date() (year int, month Month, day int)

// Year returns the year in which t occurs.
func (t Time) Year() int

// Month returns the month of the year specified by t.
func (t Time) Month() Month

// Day returns the day of the month specified by t.
func (t Time) Day() int

// Weekday returns the day of the week specified by t.
func (t Time) Weekday() Weekday

// ISOWeek returns the ISO 8601 year and week number in which t occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
// week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
// of year n+1.
func (t Time) ISOWeek() (year, week int)

// Clock returns the hour, minute, and second within the day specified by t.
func (t Time) Clock() (hour, min, sec int)

// Hour returns the hour within the day specified by t, in the range [0, 23].
func (t Time) Hour() int

// Minute returns the minute offset within the hour specified by t, in the range [0, 59].
func (t Time) Minute() int

// Second returns the second offset within the minute specified by t, in the range [0, 59].
func (t Time) Second() int

// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
func (t Time) Nanosecond() int

// YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years,
// and [1,366] in leap years.
func (t Time) YearDay() int

// A Duration represents the elapsed time between two instants
// as an int64 nanosecond count. The representation limits the
// largest representable duration to approximately 290 years.
type Duration int64

// Common durations. There is no definition for units of Day or larger
// to avoid confusion across daylight savings time zone transitions.
//
// To count the number of units in a Duration, divide:
//
//	second := time.Second
//	fmt.Print(int64(second/time.Millisecond)) // prints 1000
//
// To convert an integer number of units to a Duration, multiply:
//
//	seconds := 10
//	fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

// String returns a string representing the duration in the form "72h3m0.5s".
// Leading zero units are omitted. As a special case, durations less than one
// second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure
// that the leading digit is non-zero. The zero duration formats as 0s.
func (d Duration) String() string

// Nanoseconds returns the duration as an integer nanosecond count.
func (d Duration) Nanoseconds() int64

// Microseconds returns the duration as an integer microsecond count.
func (d Duration) Microseconds() int64

// Milliseconds returns the duration as an integer millisecond count.
func (d Duration) Milliseconds() int64

// Seconds returns the duration as a floating point number of seconds.
func (d Duration) Seconds() float64

// Minutes returns the duration as a floating point number of minutes.
func (d Duration) Minutes() float64

// Hours returns the duration as a floating point number of hours.
func (d Duration) Hours() float64

// Truncate returns the result of rounding d toward zero to a multiple of m.
// If m <= 0, Truncate returns d unchanged.
func (d Duration) Truncate(m Duration) Duration

// Round returns the result of rounding d to the nearest multiple of m.
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration,
// Round returns the maximum (or minimum) duration.
// If m <= 0, Round returns d unchanged.
func (d Duration) Round(m Duration) Duration

// Abs returns the absolute value of d.
// As a special case, math.MinInt64 is converted to math.MaxInt64.
func (d Duration) Abs() Duration

// Add returns the time t+d.
func (t Time) Add(d Duration) Time

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
// To compute t-d for a duration d, use t.Add(-d).
func (t Time) Sub(u Time) Duration

// Since returns the time elapsed since t.
// It is shorthand for time.Now().Sub(t).
func Since(t Time) Duration

// Until returns the duration until t.
// It is shorthand for t.Sub(time.Now()).
func Until(t Time) Duration

// AddDate returns the time corresponding to adding the
// given number of years, months, and days to t.
// For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
//
// Note that dates are fundamentally coupled to timezones, and calendrical
// periods like days don't have fixed durations. AddDate uses the Location of
// the Time value to determine these durations. That means that the same
// AddDate arguments can produce a different shift in absolute time depending on
// the base Time value and its Location. For example, AddDate(0, 0, 1) applied
// to 12:00 on March 27 always returns 12:00 on March 28. At some locations and
// in some years this is a 24 hour shift. In others it's a 23 hour shift due to
// daylight savings time transitions.
//
// AddDate normalizes its result in the same way that Date does,
// so, for example, adding one month to October 31 yields
// December 1, the normalized form for November 31.
func (t Time) AddDate(years int, months int, days int) Time

// daysBefore[m] counts the number of days in a non-leap year
// before month m begins. There is an entry for m=12, counting
// the number of days before January of next year (365).

// Monotonic times are reported as offsets from startNano.
// We initialize startNano to runtimeNano() - 1 so that on systems where
// monotonic time resolution is fairly low (e.g. Windows 2008
// which appears to have a default resolution of 15ms),
// we avoid ever reporting a monotonic time of 0.
// (Callers may want to use 0 as "time not set".)

// Now returns the current local time.
func Now() Time

// UTC returns t with the location set to UTC.
func (t Time) UTC() Time

// Local returns t with the location set to local time.
func (t Time) Local() Time

// In returns a copy of t representing the same time instant, but
// with the copy's location information set to loc for display
// purposes.
//
// In panics if loc is nil.
func (t Time) In(loc *Location) Time

// Location returns the time zone information associated with t.
func (t Time) Location() *Location

// Zone computes the time zone in effect at time t, returning the abbreviated
// name of the zone (such as "CET") and its offset in seconds east of UTC.
func (t Time) Zone() (name string, offset int)

// ZoneBounds returns the bounds of the time zone in effect at time t.
// The zone begins at start and the next zone begins at end.
// If the zone begins at the beginning of time, start will be returned as a zero Time.
// If the zone goes on forever, end will be returned as a zero Time.
// The Location of the returned times will be the same as t.
func (t Time) ZoneBounds() (start, end Time)

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC. The result does not depend on the
// location associated with t.
// Unix-like operating systems often record time as a 32-bit
// count of seconds, but since the method here returns a 64-bit
// value it is valid for billions of years into the past or future.
func (t Time) Unix() int64

// UnixMilli returns t as a Unix time, the number of milliseconds elapsed since
// January 1, 1970 UTC. The result is undefined if the Unix time in
// milliseconds cannot be represented by an int64 (a date more than 292 million
// years before or after 1970). The result does not depend on the
// location associated with t.
func (t Time) UnixMilli() int64

// UnixMicro returns t as a Unix time, the number of microseconds elapsed since
// January 1, 1970 UTC. The result is undefined if the Unix time in
// microseconds cannot be represented by an int64 (a date before year -290307 or
// after year 294246). The result does not depend on the location associated
// with t.
func (t Time) UnixMicro() int64

// UnixNano returns t as a Unix time, the number of nanoseconds elapsed
// since January 1, 1970 UTC. The result is undefined if the Unix time
// in nanoseconds cannot be represented by an int64 (a date before the year
// 1678 or after 2262). Note that this means the result of calling UnixNano
// on the zero Time is undefined. The result does not depend on the
// location associated with t.
func (t Time) UnixNano() int64

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (t Time) MarshalBinary() ([]byte, error)

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (t *Time) UnmarshalBinary(data []byte) error

// GobEncode implements the gob.GobEncoder interface.
func (t Time) GobEncode() ([]byte, error)

// GobDecode implements the gob.GobDecoder interface.
func (t *Time) GobDecode(data []byte) error

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in the RFC 3339 format with sub-second precision.
// If the timestamp cannot be represented as valid RFC 3339
// (e.g., the year is out of range), then an error is reported.
func (t Time) MarshalJSON() ([]byte, error)

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time must be a quoted string in the RFC 3339 format.
func (t *Time) UnmarshalJSON(data []byte) error

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format with sub-second precision.
// If the timestamp cannot be represented as valid RFC 3339
// (e.g., the year is out of range), then an error is reported.
func (t Time) MarshalText() ([]byte, error)

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time must be in the RFC 3339 format.
func (t *Time) UnmarshalText(data []byte) error

// Unix returns the local Time corresponding to the given Unix time,
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// It is valid to pass nsec outside the range [0, 999999999].
// Not all sec values have a corresponding time value. One such
// value is 1<<63-1 (the largest int64 value).
func Unix(sec int64, nsec int64) Time

// UnixMilli returns the local Time corresponding to the given Unix time,
// msec milliseconds since January 1, 1970 UTC.
func UnixMilli(msec int64) Time

// UnixMicro returns the local Time corresponding to the given Unix time,
// usec microseconds since January 1, 1970 UTC.
func UnixMicro(usec int64) Time

// IsDST reports whether the time in the configured location is in Daylight Savings Time.
func (t Time) IsDST() bool

// Date returns the Time corresponding to
//
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
//
// in the appropriate zone for that time in the given location.
//
// The month, day, hour, min, sec, and nsec values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
//
// A daylight savings time transition skips or repeats times.
// For example, in the United States, March 13, 2011 2:15am never occurred,
// while November 6, 2011 1:15am occurred twice. In such cases, the
// choice of time zone, and therefore the time, is not well-defined.
// Date returns a time that is correct in one of the two zones involved
// in the transition, but it does not guarantee which.
//
// Date panics if loc is nil.
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time

// Truncate returns the result of rounding t down to a multiple of d (since the zero time).
// If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Truncate(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
func (t Time) Truncate(d Duration) Time

// Round returns the result of rounding t to the nearest multiple of d (since the zero time).
// The rounding behavior for halfway values is to round up.
// If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Round operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Round(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
func (t Time) Round(d Duration) Time
