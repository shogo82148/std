// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// White-box tests for transport.go (in package http instead of http_test).

package http

// issue22091Error acts like a golang.org/x/net/http2.ErrNoCachedConn.
