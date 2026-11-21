// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// BindingEntry holds the parsed declarations and linker flags collected
// from the //cgo: pragmas in a package.
type BindingEntry struct {
	Package string
	LdFlags []string
	Funcs   []BindingFunc
}
