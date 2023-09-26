// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort_test

import (
	. "sort"
	stringspkg "strings"
)

// This is based on the "antiquicksort" implementation by M. Douglas McIlroy.
// See http://www.cs.dartmouth.edu/~doug/mdmspe.pdf for more info.
