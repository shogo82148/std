// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || amd64

package cpu

// DataCacheSizes returns the size of each data cache from lowest
// level in the hierarchy to highest.
//
// Unlike other parts of this package's public API, it is not safe
// to reference early in runtime initialization because it allocates.
// It's intended for testing only.
func DataCacheSizes() []uintptr
