// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto && linux && (amd64 || arm64) && !android && !cmd_go_bootstrap && !msan

package boring

// Unreachable marks code that should be unreachable
// when BoringCrypto is in use. It panics.
func Unreachable()

// UnreachableExceptTests marks code that should be unreachable
// when BoringCrypto is in use. It panics.
func UnreachableExceptTests()