// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/encoding/gob"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
)

// This example transmits a value that implements the custom encoding and decoding methods.
func Example_encodeDecode() {
	var network bytes.Buffer // Stand-in for the network.

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatal("encode:", err)
	}

	// Create a decoder and receive a value.
	dec := gob.NewDecoder(&network)
	var v Vector
	err = dec.Decode(&v)
	if err != nil {
		log.Fatal("decode:", err)
	}
	fmt.Println(v)

	// Output:
	// {3 4 5}
}
