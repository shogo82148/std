// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hmac implements the Keyed-Hash Message Authentication Code (HMAC) as
// defined in U.S. Federal Information Processing Standards Publication 198.
// An HMAC is a cryptographic hash that uses a key to sign a message.
// The receiver verifies the hash by recomputing it using the same key.
package hmac

import (
	"github.com/shogo82148/std/hash"
)

// New returns a new HMAC hash using the given hash.Hash type and key.
func New(h func() hash.Hash, key []byte) hash.Hash
