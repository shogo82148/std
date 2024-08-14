// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptotest

import (
	"github.com/shogo82148/std/crypto/cipher"
	"github.com/shogo82148/std/testing"
)

type MakeBlock func(key []byte) (cipher.Block, error)

// TestBlock performs a set of tests on cipher.Block implementations, checking
// the documented requirements of BlockSize, Encrypt, and Decrypt.
func TestBlock(t *testing.T, keySize int, mb MakeBlock)
