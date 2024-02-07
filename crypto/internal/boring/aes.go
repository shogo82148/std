// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto && linux && (amd64 || arm64) && !android && !msan

package boring

import (
	"github.com/shogo82148/std/crypto/cipher"
)

var _ extraModes = (*aesCipher)(nil)

func NewAESCipher(key []byte) (cipher.Block, error)

func NewGCMTLS(c cipher.Block) (cipher.AEAD, error)
