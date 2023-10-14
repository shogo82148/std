// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

// InitLSym defines f's obj.LSym and initializes it based on the
// properties of f. This includes setting the symbol flags and ABI and
// creating and initializing related DWARF symbols.
//
// InitLSym must be called exactly once per function and must be
// called for both functions with bodies and functions without bodies.
// For body-less functions, we only create the LSym; for functions
// with bodies call a helper to setup up / populate the LSym.
func InitLSym(f *Func, hasBody bool)
