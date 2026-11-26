// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux || darwin

package cryptotest

import (
	"github.com/shogo82148/std/testing"
)

// BoundarySlices は、指定されたサイズのスライスのペアを割り当てます。
// 一方はページの先頭に、もう一方はページの末尾に配置されます。
// スライスの境界より前または後ろへのアクセスは、フォルトを引き起こします。
func BoundarySlices(t *testing.T, size int) (start, end []byte)
