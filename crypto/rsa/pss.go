// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/io"
)

const (
	// PSSSaltLengthAuto causes the salt in a PSS signature to be as large
	// as possible when signing, and to be auto-detected when verifying.
	PSSSaltLengthAuto = 0
	// PSSSaltLengthEqualsHash causes the salt length to equal the length
	// of the hash used in the signature.
	PSSSaltLengthEqualsHash = -1
)

// PSSOptions contains options for creating and verifying PSS signatures.
type PSSOptions struct {
	SaltLength int

	Hash crypto.Hash
}

// HashFunc returns pssOpts.Hash so that PSSOptions implements
// crypto.SignerOpts.
func (pssOpts *PSSOptions) HashFunc() crypto.Hash

// SignPSS calculates the signature of hashed using RSASSA-PSS [1].
// Note that hashed must be the result of hashing the input message using the
// given hash function. The opts argument may be nil, in which case sensible
// defaults are used.
func SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte, opts *PSSOptions) ([]byte, error)

// VerifyPSS verifies a PSS signature.
// hashed is the result of hashing the input message using the given hash
// function and sig is the signature. A valid signature is indicated by
// returning a nil error. The opts argument may be nil, in which case sensible
// defaults are used.
func VerifyPSS(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte, opts *PSSOptions) error
