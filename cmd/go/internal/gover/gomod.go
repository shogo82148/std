// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gover

// GoModLookup takes go.mod or go.work content,
// finds the first line in the file starting with the given key,
// and returns the value associated with that key.
//
// Lookup should only be used with non-factored verbs
// such as "go" and "toolchain", usually to find versions
// or version-like strings.
func GoModLookup(gomod []byte, key string) string
