// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mldsa

// TestingOnlyNewPrivateKeyFromSemiExpanded creates a PrivateKey from a
// semi-expanded private key encoding, for testing purposes. It rejects
// inconsistent keys.
//
// [PrivateKey.Bytes] must NOT be called on the resulting key, as it will
// produce a random value.
func TestingOnlyNewPrivateKeyFromSemiExpanded(sk []byte) (*PrivateKey, error)

func TestingOnlyPrivateKeySemiExpandedBytes(priv *PrivateKey) []byte
