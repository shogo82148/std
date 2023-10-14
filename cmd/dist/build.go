// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main
<<<<<<< HEAD
=======

// IsRuntimePackagePathは 'pkgpath' を調べ、それが「ランタイム関連」のパッケージコレクションに属している場合はTRUEを返します。これには、"runtime"自体、"reflect"、"syscall"、および"runtime/internal/*"パッケージが含まれます。
// cmd/internal/objabi/path.goとの同期を維持してください。:IsRuntimePackagePath
func IsRuntimePackagePath(pkgpath string) bool
>>>>>>> release-branch.go1.21
