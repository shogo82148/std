// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nistec

// P256OrdInverse sets k to the inverse of k modulo ord(G). k is encoded as four
// uint64 limbs in little-endian order, and must be reduced. If k is zero, the
// result is zero.
func P256OrdInverse(k *[4]uint64)
