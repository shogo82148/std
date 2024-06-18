// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fmtsort provides a general stable ordering mechanism
// for maps, on behalf of the fmt and text/template packages.
// It is not guaranteed to be efficient and works only for types
// that are valid map keys.
package fmtsort

import (
	"github.com/shogo82148/std/reflect"
)

// SortedMap is a slice of KeyValue pairs that simplifies sorting
// and iterating over map entries.
//
// Each KeyValue pair contains a map key and its corresponding value.
type SortedMap []KeyValue

// KeyValue holds a single key and value pair found in a map.
type KeyValue struct {
	Key, Value reflect.Value
}

// Sort accepts a map and returns a SortedMap that has the same keys and
// values but in a stable sorted order according to the keys, modulo issues
// raised by unorderable key values such as NaNs.
//
// The ordering rules are more general than with Go's < operator:
//
//   - when applicable, nil compares low
//   - ints, floats, and strings order by <
//   - NaN compares less than non-NaN floats
//   - bool compares false before true
//   - complex compares real, then imag
//   - pointers compare by machine address
//   - channel values compare by machine address
//   - structs compare each field in turn
//   - arrays compare each element in turn.
//     Otherwise identical arrays compare by length.
//   - interface values compare first by reflect.Type describing the concrete type
//     and then by concrete value as described in the previous rules.
func Sort(mapValue reflect.Value) SortedMap
