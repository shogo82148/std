// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Preprofile creates an intermediate representation of a pprof profile for use
// during PGO in the compiler. This transformation depends only on the profile
// itself and is thus wasteful to perform in every invocation of the compiler.
//
// Usage:
//
//	go tool preprofile [-v] [-o output] -i input
//
//

package main
