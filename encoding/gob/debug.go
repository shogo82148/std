// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Delete the next line to include in the gob package.
//
//go:build ignore

package gob

import (
	"github.com/shogo82148/std/io"
)

// Debug prints a human-readable representation of the gob data read from r.
// It is a no-op unless debugging was enabled when the package was built.
func Debug(r io.Reader)
