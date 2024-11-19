// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package impl is a registry of alternative implementations of cryptographic
// primitives, to allow selecting them for testing.
package impl

// Register records an alternative implementation of a cryptographic primitive.
// The implementation might be available or not based on CPU support. If
// available is false, the implementation is unavailable and can't be tested on
// this machine. If available is true, it can be set to false to disable the
// implementation. If all alternative implementations but one are disabled, the
// remaining one must be used (i.e. disabling one implementation must not
// implicitly disable any other). Each package has an implicit base
// implementation that is selected when all alternatives are unavailable or
// disabled. pkg must be the package name, not path (e.g. "aes" not "crypto/aes").
func Register(pkg, name string, available *bool)

// List returns the names of all alternative implementations registered for the
// given package, whether available or not. The implicit base implementation is
// not included.
func List(pkg string) []string

// Select disables all implementations for the given package except the one
// with the given name. If name is empty, the base implementation is selected.
// It returns whether the selected implementation is available.
func Select(pkg, name string) bool

func Reset(pkg string)
