// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hashing algorithm inspired by
// wyhash: https://github.com/wangyi-fudan/wyhash/blob/ceb019b530e2c1c14d70b79bfa2bc49de7d95bc1/Modern%20Non-Cryptographic%20Hash%20Function%20and%20Pseudorandom%20Number%20Generator.pdf

//go:build 386 || arm || mips || mipsle
// +build 386 arm mips mipsle

package runtime
