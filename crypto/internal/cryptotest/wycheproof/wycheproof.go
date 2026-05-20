// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package wycheproof provides helper utilities for writing tests that
// rely on Wycheproof test vector schemas and JSON vector data.
// See https://github.com/C2SP/wycheproof for more information.
package wycheproof

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/testing"
)

// LoadVectorFile unmarshals Wycheproof JSON test vector file by name.
//
// Typically, the value argument will be a pointer to a Wycheproof schema
// type representing the in-memory structure of the JSON data.
//
// Panics if there is an error reading the Wycheproof JSON vector data file,
// or if it can't be unmarshalled into the provided value.
func LoadVectorFile(t *testing.T, filename string, value any)

// ShouldPass returns true if a test should pass informed by expected result
// and flags.
//
// flagsShouldPass is a map used to determine if an "acceptable" result test
// case should pass based on test's flags.
// Every possible flag value that's associated with an "acceptable" result
// should be explicitly specified, otherwise ShouldPass will panic.
func ShouldPass(t *testing.T, result Result, flags []string, flagsShouldPass map[string]bool) bool

// ParseHash maps from a Wycheproof hash name to a crypto.Hash implementation
// It panics if the provided hash name is unknown.
func ParseHash(h string) crypto.Hash

// TestName returns a t.Run subtest name for a Wycheproof test vector.
func TestName(file string, tv any) string

// MustDecodeHex is a helper that decodes the provided string or panics.
//
// Many Wycheproof vector values are hex encoded strings and in a test context
// we don't intend to handle decoding errors gracefully.
func MustDecodeHex(h string) []byte

// MustPanic calls fn and fails the test if fn does not panic.
//
// This is useful for testing that invalid inputs (like incorrect nonce sizes
// for AEAD ciphers) properly trigger panics rather than silently accepting
// the bad input.
func MustPanic(t *testing.T, name string, fn func())
