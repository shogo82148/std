// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hkdf

import (
	"github.com/shogo82148/std/hash"
)

func Extract[H hash.Hash](h func() H, secret, salt []byte) []byte

func Expand[H hash.Hash](h func() H, pseudorandomKey []byte, info string, keyLen int) []byte

func Key[H hash.Hash](h func() H, secret, salt []byte, info string, keyLen int) []byte
