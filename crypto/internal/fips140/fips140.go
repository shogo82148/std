// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fips140

import (
	"github.com/shogo82148/std/hash"
)

var Enabled bool

// Supported returns an error if FIPS 140-3 mode can't be enabled.
func Supported() error

func Name() string

// Version returns the formal version (such as "v1.0") if building against a
// frozen module with GOFIPS140. Otherwise, it returns "latest".
func Version() string

// Hash is a legacy compatibility alias for hash.Hash.
//
// It's only here because [crypto/internal/fips140/ecdsa.TestingOnlyNewDRBG]
// takes a "func() fips140.Hash" in v1.0.0, instead of being generic.
type Hash = hash.Hash
