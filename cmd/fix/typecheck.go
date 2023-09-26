// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type TypeConfig struct {
	Type map[string]*Type
	Var  map[string]string
	Func map[string]string

	External map[string]string
}

// Type describes the Fields and Methods of a type.
// If the field or method cannot be found there, it is next
// looked for in the Embed list.
type Type struct {
	Field  map[string]string
	Method map[string]string
	Embed  []string
	Def    string
}
