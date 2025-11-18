// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mlkemtest provides testing functions for the ML-KEM algorithm.
package mlkemtest

import (
	"github.com/shogo82148/std/crypto/mlkem"
)

// Encapsulate768 implements derandomized ML-KEM-768 encapsulation
// (ML-KEM.Encaps_internal from FIPS 203) using the provided encapsulation key
// ek and 32 bytes of randomness.
//
// It must only be used for known-answer tests.
func Encapsulate768(ek *mlkem.EncapsulationKey768, random []byte) (sharedKey, ciphertext []byte, err error)

// Encapsulate1024 implements derandomized ML-KEM-1024 encapsulation
// (ML-KEM.Encaps_internal from FIPS 203) using the provided encapsulation key
// ek and 32 bytes of randomness.
//
// It must only be used for known-answer tests.
func Encapsulate1024(ek *mlkem.EncapsulationKey1024, random []byte) (sharedKey, ciphertext []byte, err error)
