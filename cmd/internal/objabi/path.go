// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objabi

// PathToPrefix converts raw string to the prefix that will be used in the
// symbol table. All control characters, space, '%' and '"', as well as
// non-7-bit clean bytes turn into %xx. The period needs escaping only in the
// last segment of the path, and it makes for happier users if we escape that as
// little as possible.
func PathToPrefix(s string) string

// IsRuntimePackagePath examines 'pkgpath' and returns TRUE if it
// belongs to the collection of "runtime-related" packages, including
// "runtime" itself, "reflect", "syscall", and the
// "runtime/internal/*" packages. The compiler and/or assembler in
// some cases need to be aware of when they are building such a
// package, for example to enable features such as ABI selectors in
// assembly sources.
//
// Keep in sync with cmd/dist/build.go:IsRuntimePackagePath.
func IsRuntimePackagePath(pkgpath string) bool
