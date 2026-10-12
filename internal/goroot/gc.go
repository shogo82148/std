// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gc

package goroot

import (
	"github.com/shogo82148/std/os"
)

// IsStandardPackageは、gorootとcompilerが与えられたとき、pathが
// 標準パッケージかどうかを報告します。readDirはOSのファイルシステムパスを受け取ります。
func IsStandardPackage(readDir func(string) ([]os.DirEntry, error), goroot, compiler, path string) bool
