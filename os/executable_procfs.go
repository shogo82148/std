// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux || netbsd || dragonfly || nacl || (js && wasm)
// +build linux netbsd dragonfly nacl js,wasm

package os

// We query the executable path at init time to avoid the problem of
// readlink returns a path appended with " (deleted)" when the original
// binary gets deleted.
