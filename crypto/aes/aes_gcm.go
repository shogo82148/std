// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || arm64

package aes

import (
	subtleoverlap "crypto/internal/subtle"
)

// Assert that aesCipherGCM implements the gcmAble interface.
var _ gcmAble = (*aesCipherGCM)(nil)
