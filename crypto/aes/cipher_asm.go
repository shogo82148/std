// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || arm64 || ppc64 || ppc64le

package aes

// aesCipherGCM implements crypto/cipher.gcmAble so that crypto/cipher.NewGCM
// will use the optimised implementation in aes_gcm.go when possible.
// Instances of this type only exist when hasGCMAsm returns true. Likewise,
// the gcmAble implementation is in aes_gcm.go.
