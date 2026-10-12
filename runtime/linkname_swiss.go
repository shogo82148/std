// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.swissmap

package runtime

import (
	"github.com/shogo82148/std/testing"
)

func MustSupportFIPS140(tb testing.TB)

// MustMinimumFIPS140ModuleVersionは、minより低いマイナーバージョンの
// FIPS 140-3モジュール（"v1.26.0"など）に対してコンパイルされている場合、テストをスキップします。
func MustMinimumFIPS140ModuleVersion(tb testing.TB, min string)

func RerunWithFIPS140Enabled(t *testing.T)

func RerunWithFIPS140Enforced(t *testing.T)
