// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto && linux && (amd64 || arm64) && !android && !msan

package boring

import "github.com/shogo82148/std/C"

type PublicKeyECDH struct {
	curve string
	key   *C.GO_EC_POINT
	group *C.GO_EC_GROUP
	bytes []byte
}

type PrivateKeyECDH struct {
	curve string
	key   *C.GO_EC_KEY
}

func NewPublicKeyECDH(curve string, bytes []byte) (*PublicKeyECDH, error)

func (k *PublicKeyECDH) Bytes() []byte

func NewPrivateKeyECDH(curve string, bytes []byte) (*PrivateKeyECDH, error)

func (k *PrivateKeyECDH) PublicKey() (*PublicKeyECDH, error)

func ECDH(priv *PrivateKeyECDH, pub *PublicKeyECDH) ([]byte, error)

func GenerateKeyECDH(curve string) (*PrivateKeyECDH, []byte, error)
