// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha3_test

import (
	"github.com/shogo82148/std/encoding/hex"
	"github.com/shogo82148/std/fmt"
)

func Example_sum() {
	buf := []byte("some data to hash")
	// A hash needs to be 64 bytes long to have 256-bit collision resistance.
	h := make([]byte, 64)
	// Compute a 64-byte hash of buf and put it in h.
	ShakeSum256(h, buf)
	fmt.Printf("%x\n", h)
	// Output: 0f65fe41fc353e52c55667bb9e2b27bfcc8476f2c413e9437d272ee3194a4e3146d05ec04a25d16b8f577c19b82d16b1424c3e022e783d2b4da98de3658d363d
}

func Example_mac() {
	k := []byte("this is a secret key; you should generate a strong random key that's at least 32 bytes long")
	buf := []byte("and this is some data to authenticate")
	// A MAC with 32 bytes of output has 256-bit security strength -- if you use at least a 32-byte-long key.
	h := make([]byte, 32)
	d := NewShake256()
	// Write the key into the hash.
	d.Write(k)
	// Now write the data.
	d.Write(buf)
	// Read 32 bytes of output from the hash into h.
	d.Read(h)
	fmt.Printf("%x\n", h)
	// Output: 78de2974bd2711d5549ffd32b753ef0f5fa80a0db2556db60f0987eb8a9218ff
}

func ExampleNewCShake256() {
	out := make([]byte, 32)
	msg := []byte("The quick brown fox jumps over the lazy dog")

	// Example 1: Simple cshake
	c1 := NewCShake256([]byte("NAME"), []byte("Partition1"))
	c1.Write(msg)
	c1.Read(out)
	fmt.Println(hex.EncodeToString(out))

	// Example 2: Different customization string produces different digest
	c1 = NewCShake256([]byte("NAME"), []byte("Partition2"))
	c1.Write(msg)
	c1.Read(out)
	fmt.Println(hex.EncodeToString(out))

	// Example 3: Longer output length produces longer digest
	out = make([]byte, 64)
	c1 = NewCShake256([]byte("NAME"), []byte("Partition1"))
	c1.Write(msg)
	c1.Read(out)
	fmt.Println(hex.EncodeToString(out))

	// Example 4: Next read produces different result
	c1.Read(out)
	fmt.Println(hex.EncodeToString(out))

	// Output:
	//a90a4c6ca9af2156eba43dc8398279e6b60dcd56fb21837afe6c308fd4ceb05b
	//a8db03e71f3e4da5c4eee9d28333cdd355f51cef3c567e59be5beb4ecdbb28f0
	//a90a4c6ca9af2156eba43dc8398279e6b60dcd56fb21837afe6c308fd4ceb05b9dd98c6ee866ca7dc5a39d53e960f400bcd5a19c8a2d6ec6459f63696543a0d8
	//85e73a72228d08b46515553ca3a29d47df3047e5d84b12d6c2c63e579f4fd1105716b7838e92e981863907f434bfd4443c9e56ea09da998d2f9b47db71988109
}
