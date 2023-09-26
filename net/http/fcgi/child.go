// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fcgi

import (
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http"
)

// request holds the state for an in-progress request. As soon as it's complete,
// it's converted to an http.Request.

// response implements http.ResponseWriter.

// Serve accepts incoming FastCGI connections on the listener l, creating a new
// goroutine for each. The goroutine reads requests and then calls handler
// to reply to them.
// If l is nil, Serve accepts connections from os.Stdin.
// If handler is nil, http.DefaultServeMux is used.
func Serve(l net.Listener, handler http.Handler) error
