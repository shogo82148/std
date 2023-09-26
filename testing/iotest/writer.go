// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iotest

import "github.com/shogo82148/std/io"

// TruncateWriter returns a Writer that writes to w
// but stops silently after n bytes.
func TruncateWriter(w io.Writer, n int64) io.Writer
