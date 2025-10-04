// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsys

import (
	"github.com/shogo82148/std/io/fs"
)

// WalkDir is like filepath.WalkDir but over the virtual file system.
func WalkDir(root string, fn fs.WalkDirFunc) error
