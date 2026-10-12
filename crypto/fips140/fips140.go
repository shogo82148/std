// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// fips140パッケージは、FIPS 140-3 Go Cryptographic Moduleおよび
// FIPS 140-3モードに関する情報を提供します。
//
// 詳細については、[FIPS 140-3 documentation] を参照してください。
//
// [FIPS 140-3 documentation]: https://go.dev/doc/security/fips140
package fips140

// Enabled reports whether the cryptography libraries are operating in FIPS
// 140-3 mode.
//
// It can be controlled at runtime using the GODEBUG setting "fips140". If set
// to "on", FIPS 140-3 mode is enabled. If set to "only", non-approved
// cryptography functions will additionally return errors or panic.
//
// This can't be changed after the program has started.
func Enabled() bool
