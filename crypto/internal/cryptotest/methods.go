// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptotest

import (
	"github.com/shogo82148/std/testing"
)

// NoExtraMethods checks that the concrete type of *ms has no exported methods
// beyond the methods of the interface type of *ms, and any others specified in
// the allowed list.
//
// These methods are accessible through interface upgrades, so they end up part
// of the API even if undocumented per Hyrum's Law.
//
// ms must be a pointer to a non-nil interface.
func NoExtraMethods(t *testing.T, ms any, allowed ...string)
