// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maphash

import (
	"hash"
)

// Make sure a Hash implements the hash.Hash and hash.Hash64 interfaces.
var _ hash.Hash = &Hash{}
var _ hash.Hash64 = &Hash{}
