// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdh

import (
	"github.com/shogo82148/std/io"
)

// GenerateKeyP224 generates a random P-224 private key for ECDH.
//
// If FIPS mode is disabled, privateKey is generated from rand. If FIPS mode is
// enabled, rand is ignored and the key pair is generated using the approved
// DRBG (and the function runs considerably slower).
func GenerateKeyP224(rand io.Reader) (privateKey, publicKey []byte, err error)

// GenerateKeyP256 generates a random P-256 private key for ECDH.
//
// If FIPS mode is disabled, privateKey is generated from rand. If FIPS mode is
// enabled, rand is ignored and the key pair is generated using the approved
// DRBG (and the function runs considerably slower).
func GenerateKeyP256(rand io.Reader) (privateKey, publicKey []byte, err error)

// GenerateKeyP384 generates a random P-384 private key for ECDH.
//
// If FIPS mode is disabled, privateKey is generated from rand. If FIPS mode is
// enabled, rand is ignored and the key pair is generated using the approved
// DRBG (and the function runs considerably slower).
func GenerateKeyP384(rand io.Reader) (privateKey, publicKey []byte, err error)

// GenerateKeyP521 generates a random P-521 private key for ECDH.
//
// If FIPS mode is disabled, privateKey is generated from rand. If FIPS mode is
// enabled, rand is ignored and the key pair is generated using the approved
// DRBG (and the function runs considerably slower).
func GenerateKeyP521(rand io.Reader) (privateKey, publicKey []byte, err error)

func ImportKeyP224(privateKey []byte) (publicKey []byte, err error)

func ImportKeyP256(privateKey []byte) (publicKey []byte, err error)

func ImportKeyP384(privateKey []byte) (publicKey []byte, err error)

func ImportKeyP521(privateKey []byte) (publicKey []byte, err error)

func CheckPublicKeyP224(publicKey []byte) error

func CheckPublicKeyP256(publicKey []byte) error

func CheckPublicKeyP384(publicKey []byte) error

func CheckPublicKeyP521(publicKey []byte) error

func ECDHP224(privateKey, publicKey []byte) ([]byte, error)

func ECDHP256(privateKey, publicKey []byte) ([]byte, error)

func ECDHP384(privateKey, publicKey []byte) ([]byte, error)

func ECDHP521(privateKey, publicKey []byte) ([]byte, error)
