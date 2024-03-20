// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package base

import (
	"github.com/shogo82148/std/os"
)

// MapFile returns length bytes from the file starting at the
// specified offset as a string.
func MapFile(f *os.File, offset, length int64) (string, error)
