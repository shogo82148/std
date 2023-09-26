// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

// An extFloat represents an extended floating-point number, with more
// precision than a float64. It does not try to save bits: the
// number represented by the structure is mant*(2^exp), with a negative
// sign if neg is true.

// Powers of ten taken from double-conversion library.
// http://code.google.com/p/double-conversion/
