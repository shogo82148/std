// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ppc64le || ppc64

package aes

// Assert that aesCipherGCM implements the gcmAble interface.
var _ gcmAble = (*aesCipherAsm)(nil)
