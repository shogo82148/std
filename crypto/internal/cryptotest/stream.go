// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptotest

import (
	"github.com/shogo82148/std/crypto/cipher"
	"github.com/shogo82148/std/testing"
)

// MakeStream returns a cipher.Stream instance.
//
// Multiple calls to MakeStream must return equivalent instances,
// so for example the key and/or IV must be fixed.
type MakeStream func() cipher.Stream

// TestStream performs a set of tests on cipher.Stream implementations,
// checking the documented requirements of XORKeyStream.
func TestStream(t *testing.T, ms MakeStream)

// TestStreamFromBlock creates a Stream from a cipher.Block used in a
// cipher.BlockMode. It addresses Issue 68377 by checking for a panic when the
// BlockMode uses an IV with incorrect length.
// For a valid IV, it also runs all TestStream tests on the resulting stream.
func TestStreamFromBlock(t *testing.T, block cipher.Block, blockMode func(b cipher.Block, iv []byte) cipher.Stream)
