// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptotest

import (
	"github.com/shogo82148/std/crypto/cipher"
	"github.com/shogo82148/std/testing"
)

// MakeAEAD returns a cipher.AEAD instance.
//
// Multiple calls to MakeAEAD must return equivalent instances, so for example
// the key must be fixed.
type MakeAEAD func() (cipher.AEAD, error)

// TestAEAD performs a set of tests on cipher.AEAD implementations, checking
// the documented requirements of NonceSize, Overhead, Seal and Open.
func TestAEAD(t *testing.T, mAEAD MakeAEAD)
