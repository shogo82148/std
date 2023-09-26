// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aes

// gcmAble is implemented by cipher.Blocks that can provide an optimized
// implementation of GCM through the AEAD interface.
// See crypto/cipher/gcm.go.

// cbcEncAble is implemented by cipher.Blocks that can provide an optimized
// implementation of CBC encryption through the cipher.BlockMode interface.
// See crypto/cipher/cbc.go.

// cbcDecAble is implemented by cipher.Blocks that can provide an optimized
// implementation of CBC decryption through the cipher.BlockMode interface.
// See crypto/cipher/cbc.go.

// ctrAble is implemented by cipher.Blocks that can provide an optimized
// implementation of CTR through the cipher.Stream interface.
// See crypto/cipher/ctr.go.
