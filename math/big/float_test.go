// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

// Verify that ErrNaN implements the error interface.
var _ error = ErrNaN{}

// Selected precisions with which to run various tests.

// Selected bits with which to run various tests.
// Each entry is a list of bits representing a floating-point number (see fromBits).
