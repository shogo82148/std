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

// Several ways of getting from
// Fri Nov 18 7:56:35 PST 2011
// to
// Thu Mar 19 7:56:35 PST 2016
