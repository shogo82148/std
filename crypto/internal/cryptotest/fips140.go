// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptotest

import (
	"github.com/shogo82148/std/testing"
)

func MustSupportFIPS140(t *testing.T)

// MustMinimumFIPS140ModuleVersion は、コンパイル時に使用された FIPS 140-3 モジュールの
// マイナーバージョンが min（例: "v1.26.0"）より低い場合にテストをスキップします。
func MustMinimumFIPS140ModuleVersion(tb testing.TB, min string)
