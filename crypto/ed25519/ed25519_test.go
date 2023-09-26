// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ed25519

import (
	"log"
)

func Example_ed25519ctx() {
	pub, priv, err := GenerateKey(nil)
	if err != nil {
		log.Fatal(err)
	}

	msg := []byte("The quick brown fox jumps over the lazy dog")

	sig, err := priv.Sign(nil, msg, &Options{
		Context: "Example_ed25519ctx",
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := VerifyWithOptions(pub, msg, sig, &Options{
		Context: "Example_ed25519ctx",
	}); err != nil {
		log.Fatal("invalid signature")
	}
}

// From RFC 8032, Section 7.3
