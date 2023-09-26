// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// The vet/all command runs go vet on the standard library and commands.
// It compares the output against a set of whitelists
// maintained in the whitelist directory.
package main

// ignorePathPrefixes are file path prefixes that should be ignored wholesale.

// nbits maps from architecture names to the number of bits in a pointer.
// TODO: figure out a clean way to avoid get this info rather than listing it here yet again.

// archAsmX maps architectures to the suffix usually used for their assembly files,
// if different than the arch name itself.
