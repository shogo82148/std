// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

// AllocsPerRun returns the average number of allocations during calls to f.
//
// To compute the number of allocations, the function will first be run once as
// a warm-up.  The average number of allocations over the specified number of
// runs will then be measured and returned.
//
// AllocsPerRun sets GOMAXPROCS to 1 during its measurement and will restore
// it before returning.
func AllocsPerRun(runs int, f func()) (avg float64)
