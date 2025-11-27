// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/io"
)

// PKCS1v15DecryptOptions is for passing options to PKCS #1 v1.5 decryption using
// the [crypto.Decrypter] interface.
//
// Deprecated: PKCS #1 v1.5 encryption is dangerous and should not be used.
// See [draft-irtf-cfrg-rsa-guidance-05] for more information. Use
// [EncryptOAEP] and [DecryptOAEP] instead.
//
// [draft-irtf-cfrg-rsa-guidance-05]: https://www.ietf.org/archive/id/draft-irtf-cfrg-rsa-guidance-05.html#name-rationale
type PKCS1v15DecryptOptions struct {
	// SessionKeyLen is the length of the session key that is being
	// decrypted. If not zero, then a padding error during decryption will
	// cause a random plaintext of this length to be returned rather than
	// an error. These alternatives happen in constant time.
	SessionKeyLen int
}

// EncryptPKCS1v15 encrypts the given message with RSA and the padding
// scheme from PKCS #1 v1.5.  The message must be no longer than the
// length of the public modulus minus 11 bytes.
//
// The random parameter is used as a source of entropy to ensure that encrypting
// the same message twice doesn't result in the same ciphertext. Since Go 1.26,
// a secure source of random bytes is always used, and the Reader is ignored
// unless GODEBUG=cryptocustomrand=1 is set. This setting will be removed in a
// future Go release. Instead, use [testing/cryptotest.SetGlobalRandom].
//
// Deprecated: PKCS #1 v1.5 encryption is dangerous and should not be used.
// See [draft-irtf-cfrg-rsa-guidance-05] for more information. Use
// [EncryptOAEP] and [DecryptOAEP] instead.
//
// [draft-irtf-cfrg-rsa-guidance-05]: https://www.ietf.org/archive/id/draft-irtf-cfrg-rsa-guidance-05.html#name-rationale
func EncryptPKCS1v15(random io.Reader, pub *PublicKey, msg []byte) ([]byte, error)

// DecryptPKCS1v15 decrypts a plaintext using RSA and the padding scheme from
// PKCS #1 v1.5. The random parameter is legacy and ignored, and it can be nil.
//
// Deprecated: PKCS #1 v1.5 encryption is dangerous and should not be used.
// Whether this function returns an error or not discloses secret information.
// If an attacker can cause this function to run repeatedly and learn whether
// each instance returned an error then they can decrypt and forge signatures as
// if they had the private key. See [draft-irtf-cfrg-rsa-guidance-05] for more
// information. Use [EncryptOAEP] and [DecryptOAEP] instead.
//
// [draft-irtf-cfrg-rsa-guidance-05]: https://www.ietf.org/archive/id/draft-irtf-cfrg-rsa-guidance-05.html#name-rationale
func DecryptPKCS1v15(random io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)

// DecryptPKCS1v15SessionKey decrypts a session key using RSA and the padding
// scheme from PKCS #1 v1.5. The random parameter is legacy and ignored, and it
// can be nil.
//
// DecryptPKCS1v15SessionKey returns an error if the ciphertext is the wrong
// length or if the ciphertext is greater than the public modulus. Otherwise, no
// error is returned. If the padding is valid, the resulting plaintext message
// is copied into key. Otherwise, key is unchanged. These alternatives occur in
// constant time. It is intended that the user of this function generate a
// random session key beforehand and continue the protocol with the resulting
// value.
//
// Note that if the session key is too small then it may be possible for an
// attacker to brute-force it. If they can do that then they can learn whether a
// random value was used (because it'll be different for the same ciphertext)
// and thus whether the padding was correct. This also defeats the point of this
// function. Using at least a 16-byte key will protect against this attack.
//
// This method implements protections against Bleichenbacher chosen ciphertext
// attacks [0] described in RFC 3218 Section 2.3.2 [1]. While these protections
// make a Bleichenbacher attack significantly more difficult, the protections
// are only effective if the rest of the protocol which uses
// DecryptPKCS1v15SessionKey is designed with these considerations in mind. In
// particular, if any subsequent operations which use the decrypted session key
// leak any information about the key (e.g. whether it is a static or random
// key) then the mitigations are defeated. This method must be used extremely
// carefully, and typically should only be used when absolutely necessary for
// compatibility with an existing protocol (such as TLS) that is designed with
// these properties in mind.
//
//   - [0] “Chosen Ciphertext Attacks Against Protocols Based on the RSA Encryption
//     Standard PKCS #1”, Daniel Bleichenbacher, Advances in Cryptology (Crypto '98)
//   - [1] RFC 3218, Preventing the Million Message Attack on CMS,
//     https://www.rfc-editor.org/rfc/rfc3218.html
//
// Deprecated: PKCS #1 v1.5 encryption is dangerous and should not be used. The
// protections implemented by this function are limited and fragile, as
// explained above. See [draft-irtf-cfrg-rsa-guidance-05] for more information.
// Use [EncryptOAEP] and [DecryptOAEP] instead.
//
// [draft-irtf-cfrg-rsa-guidance-05]: https://www.ietf.org/archive/id/draft-irtf-cfrg-rsa-guidance-05.html#name-rationale
func DecryptPKCS1v15SessionKey(random io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) error
