// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gcm

import (
	"github.com/shogo82148/std/crypto/internal/fips140/aes"
)

// CMAC implements the CMAC mode from NIST SP 800-38B.
//
// It is optimized for use in Counter KDF (SP 800-108r1) and XAES-256-GCM
// (https://c2sp.org/XAES-256-GCM), rather than for exposing it to applications
// as a stand-alone MAC.
type CMAC struct {
	b  aes.Block
	k1 [aes.BlockSize]byte
	k2 [aes.BlockSize]byte
}

func NewCMAC(b *aes.Block) *CMAC

func (c *CMAC) MAC(m []byte) [aes.BlockSize]byte
