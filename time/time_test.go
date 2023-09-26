// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time_test

import (
	. "time"
)

// parsedTime is the struct representing a parsed time value.

type TimeTest struct {
	seconds int64
	golden  parsedTime
}

// The time routines provide no way to get absolute time
// (seconds since zero), but we need it to compute the right
// answer for bizarre roundings like "to the nearest 3 ns".
// Compute as t - year1 = (t - 1970) + (1970 - 2001) + (2001 - 1).
// t - 1970 is returned by Unix and Nanosecond.
// 1970 - 2001 is -(31*365+8)*86400 = -978307200 seconds.
// 2001 - 1 is 2000*365.2425*86400 = 63113904000 seconds.

type TimeFormatTest struct {
	time           Time
	formattedValue string
}

type FormatTest struct {
	name   string
	format string
	result string
}

type ParseTest struct {
	name       string
	format     string
	value      string
	hasTZ      bool
	hasWD      bool
	yearSign   int
	fracDigits int
}

type ParseErrorTest struct {
	format string
	value  string
	expect string
}

type ISOWeekTest struct {
	year       int
	month, day int
	yex        int
	wex        int
}

type YearDayTest struct {
	year, month, day int
	yday             int
}

// Test YearDay in several different scenarios
// and corner cases

// Check to see if YearDay is location sensitive

// Several ways of getting from
// Fri Nov 18 7:56:35 PST 2011
// to
// Thu Mar 19 7:56:35 PST 2016
