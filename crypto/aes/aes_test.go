// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aes

// Appendix A of FIPS 197: Key expansion examples
type KeyTest struct {
	key []byte
	enc []uint32
	dec []uint32
}

// Appendix B, C of FIPS 197: Cipher examples, Example vectors.
type CryptTest struct {
	key []byte
	in  []byte
	out []byte
}
