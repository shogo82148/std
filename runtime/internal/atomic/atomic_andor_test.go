//go:build 386 || amd64 || arm || arm64 || ppc64 || ppc64le || riscv64 || wasm

//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(61395): move these tests to atomic_test.go once And/Or have
// implementations for all architectures.
package atomic_test
