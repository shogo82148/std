// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (ppc64 || ppc64le) && !purego

package aes

// Assert that aesCipherAsm implements the cbcEncAble and cbcDecAble interfaces.
var _ cbcEncAble = (*aesCipherAsm)(nil)
var _ cbcDecAble = (*aesCipherAsm)(nil)
