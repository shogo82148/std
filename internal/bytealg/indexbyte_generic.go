// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Avoid IndexByte and IndexByteString on Plan 9 because it uses
// SSE instructions on x86 machines, and those are classified as
// floating point instructions, which are illegal in a note handler.

//go:build !386 && (!amd64 || plan9) && !s390x && !arm && !arm64 && !loong64 && !ppc64 && !ppc64le && !mips && !mipsle && !mips64 && !mips64le && !riscv64 && !wasm

package bytealg

func IndexByte(b []byte, c byte) int

func IndexByteString(s string, c byte) int
