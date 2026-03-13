// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Transport code's client connection pooling.

package http2

// ClientConnPool manages a pool of HTTP/2 client connections.
type ClientConnPool interface {
	GetClientConn(req *ClientRequest, addr string) (*ClientConn, error)
	MarkDead(*ClientConn)
}

var (
	_ clientConnPoolIdleCloser = (*clientConnPool)(nil)
	_ clientConnPoolIdleCloser = noDialClientConnPool{}
)
