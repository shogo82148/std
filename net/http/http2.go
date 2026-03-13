// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !nethttpomithttp2

package http

import (
	"github.com/shogo82148/std/io"
)

// Optional http.ResponseWriter interfaces implemented.
var (
	_ CloseNotifier   = http2ResponseWriter{}
	_ Flusher         = http2ResponseWriter{}
	_ io.StringWriter = http2ResponseWriter{}
)
