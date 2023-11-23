// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

// Map constants common to several packages
// runtime/runtime-gdb.py:MapTypePrinter contains its own copy
const (
	MapBucketCountBits = 3
	MapBucketCount     = 1 << MapBucketCountBits
	MapMaxKeyBytes     = 128
	MapMaxElemBytes    = 128
)

// ZeroValSize is the size in bytes of runtime.zeroVal.
const ZeroValSize = 1024
