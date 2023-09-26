// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package des

type CryptTest struct {
	key []byte
	in  []byte
	out []byte
}

// some custom tests for DES

// some custom tests for TripleDES

// NIST Special Publication 800-20, Appendix A
// Key for use with Table A.1 tests

// Table A.1 Resulting Ciphertext from the Variable Plaintext Known Answer Test

// Plaintext for use with Table A.2 tests

// Table A.2 Resulting Ciphertext from the Variable Key Known Answer Test

// Plaintext for use with Table A.3 tests

// Table A.3 Values To Be Used for the Permutation Operation Known Answer Test

// Table A.4 Values To Be Used for the Substitution Table Known Answer Test

func ExampleNewTripleDESCipher() {
	// NewTripleDESCipher can also be used when EDE2 is required by
	// duplicating the first 8 bytes of the 16-byte key.
	ede2Key := []byte("example key 1234")

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)

	_, err := NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}

	// See crypto/cipher for how to use a cipher.Block for encryption and
	// decryption.
}
