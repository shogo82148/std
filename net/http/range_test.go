// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

var ParseRangeTests = []struct {
	s      string
	length int64
	r      []httpRange
}{
	{"", 0, nil},
	{"", 1000, nil},
	{"foo", 0, nil},
	{"bytes=", 0, nil},
	{"bytes=7", 10, nil},
	{"bytes= 7 ", 10, nil},
	{"bytes=1-", 0, nil},
	{"bytes=5-4", 10, nil},
	{"bytes=0-2,5-4", 10, nil},
	{"bytes=2-5,4-3", 10, nil},
	{"bytes=--5,4--3", 10, nil},
	{"bytes=A-", 10, nil},
	{"bytes=A- ", 10, nil},
	{"bytes=A-Z", 10, nil},
	{"bytes= -Z", 10, nil},
	{"bytes=5-Z", 10, nil},
	{"bytes=Ran-dom, garbage", 10, nil},
	{"bytes=0x01-0x02", 10, nil},
	{"bytes=         ", 10, nil},
	{"bytes= , , ,   ", 10, nil},

	{"bytes=0-9", 10, []httpRange{{0, 10}}},
	{"bytes=0-", 10, []httpRange{{0, 10}}},
	{"bytes=5-", 10, []httpRange{{5, 5}}},
	{"bytes=0-20", 10, []httpRange{{0, 10}}},
	{"bytes=15-,0-5", 10, nil},
	{"bytes=1-2,5-", 10, []httpRange{{1, 2}, {5, 5}}},
	{"bytes=-2 , 7-", 11, []httpRange{{9, 2}, {7, 4}}},
	{"bytes=0-0 ,2-2, 7-", 11, []httpRange{{0, 1}, {2, 1}, {7, 4}}},
	{"bytes=-5", 10, []httpRange{{5, 5}}},
	{"bytes=-15", 10, []httpRange{{0, 10}}},
	{"bytes=0-499", 10000, []httpRange{{0, 500}}},
	{"bytes=500-999", 10000, []httpRange{{500, 500}}},
	{"bytes=-500", 10000, []httpRange{{9500, 500}}},
	{"bytes=9500-", 10000, []httpRange{{9500, 500}}},
	{"bytes=0-0,-1", 10000, []httpRange{{0, 1}, {9999, 1}}},
	{"bytes=500-600,601-999", 10000, []httpRange{{500, 101}, {601, 399}}},
	{"bytes=500-700,601-999", 10000, []httpRange{{500, 201}, {601, 399}}},

	{"bytes=   1 -2   ,  4- 5, 7 - 8 , ,,", 11, []httpRange{{1, 2}, {4, 2}, {7, 2}}},
}
