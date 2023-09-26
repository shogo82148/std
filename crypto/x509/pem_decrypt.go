// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"github.com/shogo82148/std/encoding/pem"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

type PEMCipher int

// Possible values for the EncryptPEMBlock encryption algorithm.
const (
	_ PEMCipher = iota
	PEMCipherDES
	PEMCipher3DES
	PEMCipherAES128
	PEMCipherAES192
	PEMCipherAES256
)

// rfc1423Algo holds a method for enciphering a PEM block.

// rfc1423Algos holds a slice of the possible ways to encrypt a PEM
// block.  The ivSize numbers were taken from the OpenSSL source.

// IsEncryptedPEMBlock returns if the PEM block is password encrypted.
func IsEncryptedPEMBlock(b *pem.Block) bool

// IncorrectPasswordError is returned when an incorrect password is detected.
var IncorrectPasswordError = errors.New("x509: decryption password incorrect")

// DecryptPEMBlock takes a password encrypted PEM block and the password used to
// encrypt it and returns a slice of decrypted DER encoded bytes. It inspects
// the DEK-Info header to determine the algorithm used for decryption. If no
// DEK-Info header is present, an error is returned. If an incorrect password
// is detected an IncorrectPasswordError is returned.
func DecryptPEMBlock(b *pem.Block, password []byte) ([]byte, error)

// EncryptPEMBlock returns a PEM block of the specified type holding the
// given DER-encoded data encrypted with the specified algorithm and
// password.
func EncryptPEMBlock(rand io.Reader, blockType string, data, password []byte, alg PEMCipher) (*pem.Block, error)
