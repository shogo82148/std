// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/io"
)

// GenerateKey generates a new RSA key pair of the given bit size.
// bits must be at least 128.
func GenerateKey(rand io.Reader, bits int) (*PrivateKey, error)
