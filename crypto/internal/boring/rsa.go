// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto && linux && (amd64 || arm64) && !android && !cmd_go_bootstrap && !msan

package boring

import "C"
import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/hash"
)

func GenerateKeyRSA(bits int) (N, E, D, P, Q, Dp, Dq, Qinv BigInt, err error)

type PublicKeyRSA struct {
	// _key MUST NOT be accessed directly. Instead, use the withKey method.
	_key *C.GO_RSA
}

func NewPublicKeyRSA(N, E BigInt) (*PublicKeyRSA, error)

type PrivateKeyRSA struct {
	// _key MUST NOT be accessed directly. Instead, use the withKey method.
	_key *C.GO_RSA
}

func NewPrivateKeyRSA(N, E, D, P, Q, Dp, Dq, Qinv BigInt) (*PrivateKeyRSA, error)

func DecryptRSAOAEP(h, mgfHash hash.Hash, priv *PrivateKeyRSA, ciphertext, label []byte) ([]byte, error)

func EncryptRSAOAEP(h, mgfHash hash.Hash, pub *PublicKeyRSA, msg, label []byte) ([]byte, error)

func DecryptRSAPKCS1(priv *PrivateKeyRSA, ciphertext []byte) ([]byte, error)

func EncryptRSAPKCS1(pub *PublicKeyRSA, msg []byte) ([]byte, error)

func DecryptRSANoPadding(priv *PrivateKeyRSA, ciphertext []byte) ([]byte, error)

func EncryptRSANoPadding(pub *PublicKeyRSA, msg []byte) ([]byte, error)

func SignRSAPSS(priv *PrivateKeyRSA, h crypto.Hash, hashed []byte, saltLen int) ([]byte, error)

func VerifyRSAPSS(pub *PublicKeyRSA, h crypto.Hash, hashed, sig []byte, saltLen int) error

func SignRSAPKCS1v15(priv *PrivateKeyRSA, h crypto.Hash, hashed []byte) ([]byte, error)

func VerifyRSAPKCS1v15(pub *PublicKeyRSA, h crypto.Hash, hashed, sig []byte) error
