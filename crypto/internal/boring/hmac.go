// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto && linux && (amd64 || arm64) && !android && !cmd_go_bootstrap && !msan

package boring

import (
	"github.com/shogo82148/std/hash"
)

// NewHMAC returns a new HMAC using BoringCrypto.
// The function h must return a hash implemented by
// BoringCrypto (for example, h could be boring.NewSHA256).
// If h is not recognized, NewHMAC returns nil.
func NewHMAC(h func() hash.Hash, key []byte) hash.Hash
