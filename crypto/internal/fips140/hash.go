// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fips140

import "github.com/shogo82148/std/io"

// Hash is the common interface implemented by all hash functions. It is a copy
// of [hash.Hash] from the standard library, to avoid depending on security
// definitions from outside of the module.
type Hash interface {
	io.Writer

	Sum(b []byte) []byte

	Reset()

	Size() int

	BlockSize() int
}
