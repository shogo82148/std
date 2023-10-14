// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bbig

import (
	"github.com/shogo82148/std/crypto/internal/boring"
	"github.com/shogo82148/std/math/big"
)

func Enc(b *big.Int) boring.BigInt

func Dec(b boring.BigInt) *big.Int
