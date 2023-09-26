// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elliptic

// p384Curve is a Curve implementation based on nistec.P384Point.
//
// It's a wrapper that exposes the big.Int-based Curve interface and encodes the
// legacy idiosyncrasies it requires, such as invalid and infinity point
// handling.
//
// To interact with the nistec package, points are encoded into and decoded from
// properly formatted byte slices. All big.Int use is limited to this package.
// Encoding and decoding is 1/1000th of the runtime of a scalar multiplication,
// so the overhead is acceptable.

var _ Curve = p384
