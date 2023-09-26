// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ppc64le
// +build ppc64le

package elliptic

// This was ported from the s390x implementation for ppc64le.
// Some hints are included here for changes that should be
// in the big endian ppc64 implementation, however more
// investigation and testing is needed for the ppc64 big
// endian version to work.

// p256MulAsm operates in a Montgomery domain with R = 2^256 mod p, where p is the
// underlying field of the curve. (See initP256 for the value.) Thus rr here is
// R×R mod p. See comment in Inverse about how this is used.
// TODO: For big endian implementation, the bytes in these slices should be in reverse order,
// as found in the s390x implementation.

// (This is one, in the Montgomery domain.)
