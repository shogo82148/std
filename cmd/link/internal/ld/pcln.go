// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"github.com/shogo82148/std/internal/abi"
)

const (
	SUBBUCKETS    = 16
	SUBBUCKETSIZE = abi.FuncTabBucketSize / SUBBUCKETS
	NOIDX         = 0x7fffffff
)
