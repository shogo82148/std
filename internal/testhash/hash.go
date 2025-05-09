// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testhash

import (
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/testing"
)

type MakeHash func() hash.Hash

// TestHash performs a set of tests on hash.Hash implementations, checking the
// documented requirements of Write, Sum, Reset, Size, and BlockSize.
func TestHash(t *testing.T, mh MakeHash)
