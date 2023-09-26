// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time_test

import (
	. "time"
)

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

type ParseTimeZoneTest struct {
	value  string
	length int
	ok     bool
}

type ParseErrorTest struct {
	format string
	value  string
	expect string
}

type SecondsTimeZoneOffsetTest struct {
	format         string
	value          string
	expectedoffset int
}
