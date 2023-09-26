// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Tests for transport.go

package http_test

import (
	. "net/http"
)

// hostPortHandler writes back the client's "host:port".

// rgz is a gzip quine that uncompresses to itself.
