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

// IsEncryptedPEMBlock returns whether the PEM block is password encrypted
// according to RFC 1423.
//
// Deprecated: Legacy PEM encryption as specified in RFC 1423 is insecure by
// design. Since it does not authenticate the ciphertext, it is vulnerable to
// padding oracle attacks that can let an attacker recover the plaintext.
func IsEncryptedPEMBlock(b *pem.Block) bool

// IncorrectPasswordError is returned when an incorrect password is detected.
var IncorrectPasswordError = errors.New("x509: decryption password incorrect")

// DecryptPEMBlock takes a PEM block encrypted according to RFC 1423 and the
// password used to encrypt it and returns a slice of decrypted DER encoded
// bytes. It inspects the DEK-Info header to determine the algorithm used for
// decryption. If no DEK-Info header is present, an error is returned. If an
// incorrect password is detected an [IncorrectPasswordError] is returned. Because
// of deficiencies in the format, it's not always possible to detect an
// incorrect password. In these cases no error will be returned but the
// decrypted DER bytes will be random noise.
//
// Deprecated: Legacy PEM encryption as specified in RFC 1423 is insecure by
// design. Since it does not authenticate the ciphertext, it is vulnerable to
// padding oracle attacks that can let an attacker recover the plaintext.
func DecryptPEMBlock(b *pem.Block, password []byte) ([]byte, error)

// EncryptPEMBlock returns a PEM block of the specified type holding the
// given DER encoded data encrypted with the specified algorithm and
// password according to RFC 1423.
//
// Deprecated: Legacy PEM encryption as specified in RFC 1423 is insecure by
// design. Since it does not authenticate the ciphertext, it is vulnerable to
// padding oracle attacks that can let an attacker recover the plaintext.
func EncryptPEMBlock(rand io.Reader, blockType string, data, password []byte, alg PEMCipher) (*pem.Block, error)
