// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hashing algorithm inspired by
//   xxhash: https://code.google.com/p/xxhash/
// cityhash: https://code.google.com/p/cityhash/

//go:build amd64 || amd64p32 || arm64 || ppc64 || ppc64le
// +build amd64 amd64p32 arm64 ppc64 ppc64le

package runtime
