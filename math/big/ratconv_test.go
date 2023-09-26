// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

type StringTest struct {
	in, out string
	ok      bool
}

// These are not supported by fmt.Fscanf.

// Test inputs to Rat.SetString. The prefix "long:" causes the test
// to be skipped in --test.short mode.  (The threshold is about 500us.)
