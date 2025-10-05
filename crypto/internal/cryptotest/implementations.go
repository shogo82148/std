// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptotest

import (
	"github.com/shogo82148/std/testing"
)

// TestAllImplementations runs the provided test function with each available
// implementation of the package registered with crypto/internal/impl. If there
// are no alternative implementations for pkg, f is invoked directly once.
func TestAllImplementations(t *testing.T, pkg string, f func(t *testing.T))
