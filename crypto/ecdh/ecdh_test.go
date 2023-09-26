// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdh_test

import (
	"crypto"
	"crypto/ecdh"
)

// Check that PublicKey and PrivateKey implement the interfaces documented in
// crypto.PublicKey and crypto.PrivateKey.
var _ interface {
	Equal(x crypto.PublicKey) bool
} = &ecdh.PublicKey{}
var _ interface {
	Public() crypto.PublicKey
	Equal(x crypto.PrivateKey) bool
} = &ecdh.PrivateKey{}
