// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hashing algorithm inspired by
//   xxhash: https://code.google.com/p/xxhash/
// cityhash: https://code.google.com/p/cityhash/

//go:build amd64 || amd64p32 || arm64 || mips64 || mips64le || ppc64 || ppc64le || s390x || wasm
// +build amd64 amd64p32 arm64 mips64 mips64le ppc64 ppc64le s390x wasm

package runtime
