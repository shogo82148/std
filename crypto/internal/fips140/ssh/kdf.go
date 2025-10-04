// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ssh implements the SSH KDF as specified in RFC 4253,
// Section 7.2 and allowed by SP 800-135 Revision 1.
package ssh

import (
	"github.com/shogo82148/std/hash"
)

type Direction struct {
	ivTag     []byte
	keyTag    []byte
	macKeyTag []byte
}

var ServerKeys, ClientKeys Direction

func Keys[Hash hash.Hash](hash func() Hash, d Direction,
	K, H, sessionID []byte,
	ivKeyLen, keyLen, macKeyLen int,
) (ivKey, key, macKey []byte)
