// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements Sizes.

package types2

// Sizes defines the sizing functions for package unsafe.
type Sizes interface {
	Alignof(T Type) int64

	Offsetsof(fields []*Var) []int64

	Sizeof(T Type) int64
}

// StdSizes is a convenience type for creating commonly used Sizes.
// It makes the following simplifying assumptions:
//
//   - The size of explicitly sized basic types (int16, etc.) is the
//     specified size.
//   - The size of strings and interfaces is 2*WordSize.
//   - The size of slices is 3*WordSize.
//   - The size of an array of n elements corresponds to the size of
//     a struct of n consecutive fields of the array's element type.
//   - The size of a struct is the offset of the last field plus that
//     field's size. As with all element types, if the struct is used
//     in an array its size must first be aligned to a multiple of the
//     struct's alignment.
//   - All other types have size WordSize.
//   - Arrays and structs are aligned per spec definition; all other
//     types are naturally aligned with a maximum alignment MaxAlign.
//
// *StdSizes implements Sizes.
type StdSizes struct {
	WordSize int64
	MaxAlign int64
}

func (s *StdSizes) Alignof(T Type) (result int64)

func IsSyncAtomicAlign64(T Type) bool

func (s *StdSizes) Offsetsof(fields []*Var) []int64

func (s *StdSizes) Sizeof(T Type) int64

// SizesFor returns the Sizes used by a compiler for an architecture.
// The result is nil if a compiler/architecture pair is not known.
//
// Supported architectures for compiler "gc":
// "386", "amd64", "amd64p32", "arm", "arm64", "loong64", "mips", "mipsle",
// "mips64", "mips64le", "ppc64", "ppc64le", "riscv64", "s390x", "sparc64", "wasm".
func SizesFor(compiler, arch string) Sizes
