// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sgutil

// NonSpecAPI reports whether the exported declaration recv.name (or name, for
// a package-level declaration) in package pkg is one of the few APIs that
// spec doesn't describe but that may still be generally available. pkg is the
// package name, "simd" or "archsimd". Any other generally available API must
// come from spec.
//
// Each of these APIs may be either generally available or experimental. Since
// spec doesn't describe them, the code that defines each one decides which,
// so the phase check allows them in files of either phase. For example,
// simdgen puts a CPU feature query in cpu.go, which is generally available, if
// some generally available operation needs that feature, and in cpu_exp.go,
// which is experimental, otherwise.
func NonSpecAPI(pkg, recv, name string) bool
