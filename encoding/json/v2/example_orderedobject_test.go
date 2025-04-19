// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/reflect"

	"github.com/shogo82148/std/encoding/json/jsontext"
)

// The exact order of JSON object can be preserved through the use of a
// specialized type that implements [MarshalerTo] and [UnmarshalerFrom].
func Example_orderedObject() {
	// Round-trip marshal and unmarshal an ordered object.
	// We expect the order and duplicity of JSON object members to be preserved.
	// Specify jsontext.AllowDuplicateNames since this object contains "fizz" twice.
	want := OrderedObject[string]{
		{"fizz", "buzz"},
		{"hello", "world"},
		{"fizz", "wuzz"},
	}
	b, err := json.Marshal(&want, jsontext.AllowDuplicateNames(true))
	if err != nil {
		log.Fatal(err)
	}
	var got OrderedObject[string]
	err = json.Unmarshal(b, &got, jsontext.AllowDuplicateNames(true))
	if err != nil {
		log.Fatal(err)
	}

	// Sanity check.
	if !reflect.DeepEqual(got, want) {
		log.Fatalf("roundtrip mismatch: got %v, want %v", got, want)
	}

	// Print the serialized JSON object.
	(*jsontext.Value)(&b).Indent() // indent for readability
	fmt.Println(string(b))

	// Output:
	// {
	// 	"fizz": "buzz",
	// 	"hello": "world",
	// 	"fizz": "wuzz"
	// }
}
