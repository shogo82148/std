// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// OFB（Output Feedback）モード。

package cipher

<<<<<<< HEAD
// NewOFBは、ブロック暗号bを使用して暗号化または復号化する [Stream] を返します。
// 初期化ベクトルivの長さは、bのブロックサイズと等しくなければなりません。
=======
// NewOFB returns a [Stream] that encrypts or decrypts using the block cipher b
// in output feedback mode. The initialization vector iv's length must be equal
// to b's block size.
//
// Deprecated: OFB mode is not authenticated, which generally enables active
// attacks to manipulate and recover the plaintext. It is recommended that
// applications use [AEAD] modes instead. The standard library implementation of
// OFB is also unoptimized and not validated as part of the FIPS 140-3 module.
// If an unauthenticated [Stream] mode is required, use [NewCTR] instead.
>>>>>>> upstream/release-branch.go1.25
func NewOFB(b Block, iv []byte) Stream
