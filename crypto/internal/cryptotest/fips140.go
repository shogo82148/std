// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptotest

import (
	"github.com/shogo82148/std/testing"
)

func MustSupportFIPS140(tb testing.TB)

// MustMinimumFIPS140ModuleVersion skips the test if compiled against a lower
// minor version of the FIPS 140-3 module than min (such as "v1.26.0").
func MustMinimumFIPS140ModuleVersion(tb testing.TB, min string)

func RerunWithFIPS140Enabled(t *testing.T)

func RerunWithFIPS140Enforced(t *testing.T)
