// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls12

import (
	"github.com/shogo82148/std/hash"
)

// PRF implements the TLS 1.2 pseudo-random function, as defined in RFC 5246,
// Section 5 and allowed by SP 800-135, Revision 1, Section 4.2.2.
func PRF[H hash.Hash](hash func() H, secret []byte, label string, seed []byte, keyLen int) []byte

// MasterSecret implements the TLS 1.2 extended master secret derivation, as
// defined in RFC 7627 and allowed by SP 800-135, Revision 1, Section 4.2.2.
func MasterSecret[H hash.Hash](hash func() H, preMasterSecret, transcript []byte) []byte
