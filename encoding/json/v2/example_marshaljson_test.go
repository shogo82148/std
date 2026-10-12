// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"

	"github.com/shogo82148/std/encoding/json/jsontext"
)

// Custom types may define custom marshal behavior with [MarshalerTo].
func ExampleMarshalerTo() {
	set := IntSet{
		1: {},
		2: {},
		3: {},
	}

	b, err := json.Marshal(&set, json.Deterministic(true))
	if err != nil {
		log.Fatal(err)
	}

	// Indent output for readability.
	v := jsontext.Value(b)
	v.Indent()
	fmt.Println(string(v))

	// Output:
	// [
	// 	1,
	// 	2,
	// 	3
	// ]
}

// Custom types may define custom unmarshal behavior with [UnmarshalerFrom].
func ExampleUnmarshalerFrom() {
	s := "[1,2,3]"

	var set IntSet
	err := json.Unmarshal([]byte(s), &set)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(set)

	// Output:
	// map[1:{} 2:{} 3:{}]
}
