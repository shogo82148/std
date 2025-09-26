// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdsa

import (
	"github.com/shogo82148/std/crypto/internal/fips140"
)

// TestingOnlyNewDRBG creates an SP 800-90A Rev. 1 HMAC_DRBG with a plain
// personalization string.
//
// This should only be used for ACVP testing. hmacDRBG is not intended to be
// used directly.
func TestingOnlyNewDRBG[H fips140.Hash](hash func() H, entropy, nonce []byte, s []byte) *hmacDRBG
