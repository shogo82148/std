// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Tests for transport.go.
//
// More tests are in clientserver_test.go (for things testing both client & server for both
// HTTP/1 and HTTP/2). This

package http_test

import (
	. "net/http"
)

// hostPortHandler writes back the client's "host:port".

// testCloseConn is a net.Conn tracked by a testConnSet.

// testConnSet tracks a set of TCP connections and whether they've
// been closed.

// byteFromChanReader is an io.Reader that reads a single byte at a
// time from the channel.  When the channel is closed, the reader
// returns io.EOF.

// logWritesConn is a net.Conn that logs each Write call to writes
// and then proxies to w.
// It proxies Read calls to a reader it receives from rch.

// rgz is a gzip quine that uncompresses to itself.
