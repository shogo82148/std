// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pbkdf2

import (
	"github.com/shogo82148/std/crypto/internal/fips140"
)

func Key[Hash fips140.Hash](h func() Hash, password string, salt []byte, iter, keyLength int) ([]byte, error)
