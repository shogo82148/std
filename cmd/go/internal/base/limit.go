// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import (
	"github.com/shogo82148/std/internal/godebug"
)

var NetLimitGodebug = godebug.New("#cmdgonetlimit")

// NetLimit returns the limit on concurrent network operations
// configured by GODEBUG=cmdgonetlimit, if any.
//
// A limit of 0 (indicated by 0, true) means that network operations should not
// be allowed.
func NetLimit() (int, bool)

// AcquireNet acquires a semaphore token for a network operation.
func AcquireNet() (release func(), err error)
