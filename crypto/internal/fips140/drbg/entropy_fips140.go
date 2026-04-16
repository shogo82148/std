// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

// This file contains reading from from entropy sources in FIPS-140
// mode. It uses a scratch buffer in the BSS section (see below),
// which usually doesn't cost much, except on Wasm, due to the way
// the linear memory works. FIPS-140 mode is not supported on Wasm,
// so we just use a build tag to exclude it. (Could also exclude other
// platforms that does not support FIPS-140 mode, but as the BSS
// variable doesn't cost much, don't bother.)

package drbg
