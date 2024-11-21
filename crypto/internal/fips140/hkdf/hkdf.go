// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hkdf

import (
	"github.com/shogo82148/std/crypto/internal/fips140"
)

func Extract[H fips140.Hash](h func() H, secret, salt []byte) []byte

func Expand[H fips140.Hash](h func() H, pseudorandomKey, info []byte, keyLen int) []byte

func Key[H fips140.Hash](h func() H, secret, salt, info []byte, keyLen int) []byte
