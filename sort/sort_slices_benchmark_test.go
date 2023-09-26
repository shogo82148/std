// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort_test

import (
	. "sort"
	stringspkg "strings"
)

const N = 100_000

// These benchmarks compare sorting a slice of structs with sort.Sort vs.
// slices.SortFunc.
