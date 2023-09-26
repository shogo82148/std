// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The wire protocol for HTTP's "chunked" Transfer-Encoding.

// This code is duplicated in httputil/chunked.go.
// Please make any changes in both files.

package http

import (
	"github.com/shogo82148/std/errors"
)

var ErrLineTooLong = errors.New("header line too long")

// Writing to chunkedWriter translates to writing in HTTP chunked Transfer
// Encoding wire format to the underlying Wire chunkedWriter.
