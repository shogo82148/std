// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hmac implements HMAC according to [FIPS 198-1].
//
// [FIPS 198-1]: https://doi.org/10.6028/NIST.FIPS.198-1
package hmac

import (
	"github.com/shogo82148/std/crypto/internal/fips"
)

type HMAC struct {
	opad, ipad   []byte
	outer, inner fips.Hash

	// If marshaled is true, then opad and ipad do not contain a padded
	// copy of the key, but rather the marshaled state of outer/inner after
	// opad/ipad has been fed into it.
	marshaled bool

	// forHKDF and keyLen are stored to inform the service indicator decision.
	forHKDF bool
	keyLen  int
}

func (h *HMAC) Sum(in []byte) []byte

func (h *HMAC) Write(p []byte) (n int, err error)

func (h *HMAC) Size() int
func (h *HMAC) BlockSize() int

func (h *HMAC) Reset()

// New returns a new HMAC hash using the given [fips.Hash] type and key.
func New[H fips.Hash](h func() H, key []byte) *HMAC

// MarkAsUsedInHKDF records that this HMAC instance is used as part of HKDF.
func MarkAsUsedInHKDF(h *HMAC)
