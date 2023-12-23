// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !(boringcrypto && linux && (amd64 || arm64) && !android && !msan && cgo)

package boring

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/cipher"
	"github.com/shogo82148/std/hash"
)

// Unreachable marks code that should be unreachable
// when BoringCrypto is in use. It is a no-op without BoringCrypto.
func Unreachable()

// UnreachableExceptTests marks code that should be unreachable
// when BoringCrypto is in use. It is a no-op without BoringCrypto.
func UnreachableExceptTests()

const RandReader = randReader(0)

func NewSHA1() hash.Hash
func NewSHA224() hash.Hash
func NewSHA256() hash.Hash
func NewSHA384() hash.Hash
func NewSHA512() hash.Hash

func SHA1([]byte) [20]byte
func SHA224([]byte) [28]byte
func SHA256([]byte) [32]byte
func SHA384([]byte) [48]byte
func SHA512([]byte) [64]byte

func NewHMAC(h func() hash.Hash, key []byte) hash.Hash

func NewAESCipher(key []byte) (cipher.Block, error)
func NewGCMTLS(cipher.Block) (cipher.AEAD, error)
func NewGCMTLS13(cipher.Block) (cipher.AEAD, error)

type PublicKeyECDSA struct{ _ int }
type PrivateKeyECDSA struct{ _ int }

func GenerateKeyECDSA(curve string) (X, Y, D BigInt, err error)

func NewPrivateKeyECDSA(curve string, X, Y, D BigInt) (*PrivateKeyECDSA, error)

func NewPublicKeyECDSA(curve string, X, Y BigInt) (*PublicKeyECDSA, error)

func SignMarshalECDSA(priv *PrivateKeyECDSA, hash []byte) ([]byte, error)

func VerifyECDSA(pub *PublicKeyECDSA, hash []byte, sig []byte) bool

type PublicKeyRSA struct{ _ int }
type PrivateKeyRSA struct{ _ int }

func DecryptRSAOAEP(h, mgfHash hash.Hash, priv *PrivateKeyRSA, ciphertext, label []byte) ([]byte, error)

func DecryptRSAPKCS1(priv *PrivateKeyRSA, ciphertext []byte) ([]byte, error)

func DecryptRSANoPadding(priv *PrivateKeyRSA, ciphertext []byte) ([]byte, error)

func EncryptRSAOAEP(h, mgfHash hash.Hash, pub *PublicKeyRSA, msg, label []byte) ([]byte, error)

func EncryptRSAPKCS1(pub *PublicKeyRSA, msg []byte) ([]byte, error)

func EncryptRSANoPadding(pub *PublicKeyRSA, msg []byte) ([]byte, error)

func GenerateKeyRSA(bits int) (N, E, D, P, Q, Dp, Dq, Qinv BigInt, err error)

func NewPrivateKeyRSA(N, E, D, P, Q, Dp, Dq, Qinv BigInt) (*PrivateKeyRSA, error)

func NewPublicKeyRSA(N, E BigInt) (*PublicKeyRSA, error)
func SignRSAPKCS1v15(priv *PrivateKeyRSA, h crypto.Hash, hashed []byte) ([]byte, error)

func SignRSAPSS(priv *PrivateKeyRSA, h crypto.Hash, hashed []byte, saltLen int) ([]byte, error)

func VerifyRSAPKCS1v15(pub *PublicKeyRSA, h crypto.Hash, hashed, sig []byte) error

func VerifyRSAPSS(pub *PublicKeyRSA, h crypto.Hash, hashed, sig []byte, saltLen int) error

type PublicKeyECDH struct{}
type PrivateKeyECDH struct{}

func ECDH(*PrivateKeyECDH, *PublicKeyECDH) ([]byte, error)
func GenerateKeyECDH(string) (*PrivateKeyECDH, []byte, error)
func NewPrivateKeyECDH(string, []byte) (*PrivateKeyECDH, error)
func NewPublicKeyECDH(string, []byte) (*PublicKeyECDH, error)
func (*PublicKeyECDH) Bytes() []byte
func (*PrivateKeyECDH) PublicKey() (*PublicKeyECDH, error)
