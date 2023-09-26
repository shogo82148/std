// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

type FloatInt struct {
	x float64
	y int
}

// belowOverflow should be a pretty-full pair of buckets;
// atOverflow is 1/8 bs larger = 13/8 buckets or two buckets
// that are 13/16 full each, which is the overflow boundary.
// Adding one to that should ensure overflow to the next higher size.
