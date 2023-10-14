// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bio

import (
	"github.com/shogo82148/std/io"
)

// MustClose closes Closer c and calls log.Fatal if it returns a non-nil error.
func MustClose(c io.Closer)

// MustWriter returns a Writer that wraps the provided Writer,
// except that it calls log.Fatal instead of returning a non-nil error.
func MustWriter(w io.Writer) io.Writer
