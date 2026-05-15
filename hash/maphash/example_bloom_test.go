// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maphash_test

import (
	"github.com/shogo82148/std/fmt"
)

func Example_bloomFilter() {
	// Create a Bloom filter optimized for 2 elements with
	// a one-in-a-billion false-positive rate.
	//
	// (This low rate demands a lot of space: 88 bits and
	// 30 hash functions. More typical rates are 1-5%;
	// at 5%, we need only 16 bits and 4 hash functions.)
	f := NewComparableBloomFilter[string](2, 1e-9)

	// Insert two elements.
	f.Insert("apple")
	f.Insert("banana")

	// Check whether elements are present.
	//
	// "cherry" was not inserted, but Contains is probabilistic, so
	// this test will spuriously report Contains("cherry") = true
	// about once every billion runs.
	for _, fruit := range []string{"apple", "banana", "cherry"} {
		fmt.Printf("Contains(%q) = %v\n", fruit, f.Contains(fruit))
	}

	// Output:
	//
	// Contains("apple") = true
	// Contains("banana") = true
	// Contains("cherry") = false
}
