// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http_test

import (
	. "net/http"
)

var ServeFileRangeTests = []struct {
	r      string
	code   int
	ranges []wantRange
}{
	{r: "", code: StatusOK},
	{r: "bytes=0-4", code: StatusPartialContent, ranges: []wantRange{{0, 5}}},
	{r: "bytes=2-", code: StatusPartialContent, ranges: []wantRange{{2, testFileLen}}},
	{r: "bytes=-5", code: StatusPartialContent, ranges: []wantRange{{testFileLen - 5, testFileLen}}},
	{r: "bytes=3-7", code: StatusPartialContent, ranges: []wantRange{{3, 8}}},
	{r: "bytes=0-0,-2", code: StatusPartialContent, ranges: []wantRange{{0, 1}, {testFileLen - 2, testFileLen}}},
	{r: "bytes=0-1,5-8", code: StatusPartialContent, ranges: []wantRange{{0, 2}, {5, 9}}},
	{r: "bytes=0-1,5-", code: StatusPartialContent, ranges: []wantRange{{0, 2}, {5, testFileLen}}},
	{r: "bytes=5-1000", code: StatusPartialContent, ranges: []wantRange{{5, testFileLen}}},
	{r: "bytes=0-,1-,2-,3-,4-", code: StatusOK},
	{r: "bytes=0-9", code: StatusPartialContent, ranges: []wantRange{{0, testFileLen - 1}}},
	{r: "bytes=0-10", code: StatusPartialContent, ranges: []wantRange{{0, testFileLen}}},
	{r: "bytes=0-11", code: StatusPartialContent, ranges: []wantRange{{0, testFileLen}}},
	{r: "bytes=10-11", code: StatusPartialContent, ranges: []wantRange{{testFileLen - 1, testFileLen}}},
	{r: "bytes=10-", code: StatusPartialContent, ranges: []wantRange{{testFileLen - 1, testFileLen}}},
	{r: "bytes=11-", code: StatusRequestedRangeNotSatisfiable},
	{r: "bytes=11-12", code: StatusRequestedRangeNotSatisfiable},
	{r: "bytes=12-12", code: StatusRequestedRangeNotSatisfiable},
	{r: "bytes=11-100", code: StatusRequestedRangeNotSatisfiable},
	{r: "bytes=12-100", code: StatusRequestedRangeNotSatisfiable},
	{r: "bytes=100-", code: StatusRequestedRangeNotSatisfiable},
	{r: "bytes=100-1000", code: StatusRequestedRangeNotSatisfiable},
}
