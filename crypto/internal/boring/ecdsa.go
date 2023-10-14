// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto && linux && (amd64 || arm64) && !android && !msan

package boring

import "C"

type PrivateKeyECDSA struct {
	key *C.GO_EC_KEY
}

type PublicKeyECDSA struct {
	key *C.GO_EC_KEY
}

func NewPublicKeyECDSA(curve string, X, Y BigInt) (*PublicKeyECDSA, error)

func NewPrivateKeyECDSA(curve string, X, Y BigInt, D BigInt) (*PrivateKeyECDSA, error)

func SignMarshalECDSA(priv *PrivateKeyECDSA, hash []byte) ([]byte, error)

func VerifyECDSA(pub *PublicKeyECDSA, hash []byte, sig []byte) bool

func GenerateKeyECDSA(curve string) (X, Y, D BigInt, err error)
