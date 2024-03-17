// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

package aes

// Assert that aesCipherAsm implements the cbcEncAble and cbcDecAble interfaces.
var _ cbcEncAble = (*aesCipherAsm)(nil)
var _ cbcDecAble = (*aesCipherAsm)(nil)
