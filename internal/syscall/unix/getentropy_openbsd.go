// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build openbsd && !mips64

package unix

// GetEntropy calls the OpenBSD getentropy system call.
func GetEntropy(p []byte) error
