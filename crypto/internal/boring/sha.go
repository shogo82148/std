// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto && linux && (amd64 || arm64) && !android && !msan

package boring

import (
	"github.com/shogo82148/std/hash"
)

func SHA1(p []byte) (sum [20]byte)

func SHA224(p []byte) (sum [28]byte)

func SHA256(p []byte) (sum [32]byte)

func SHA384(p []byte) (sum [48]byte)

func SHA512(p []byte) (sum [64]byte)

// NewSHA1 returns a new SHA1 hash.
func NewSHA1() hash.Hash

// NewSHA224 returns a new SHA224 hash.
func NewSHA224() hash.Hash

// NewSHA256 returns a new SHA256 hash.
func NewSHA256() hash.Hash

// NewSHA384 returns a new SHA384 hash.
func NewSHA384() hash.Hash

// NewSHA512 returns a new SHA512 hash.
func NewSHA512() hash.Hash
