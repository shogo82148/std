// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aes

// gcmCount represents a 16-byte big-endian count value.

// gcmHashKey represents the 16-byte hash key required by the GHASH algorithm.

// Assert that aesCipherAsm implements the gcmAble interface.
var _ gcmAble = (*aesCipherAsm)(nil)
