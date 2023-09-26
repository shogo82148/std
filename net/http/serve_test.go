// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// End-to-end serving tests

package http_test

import (
	. "net/http"
)

// trackLastConnListener tracks the last net.Conn that was accepted.

// testHandlerBodyConsumer represents a function injected into a test handler to
// vary work done on a request Body.

// slowTestConn is a net.Conn that provides a means to simulate parts of a
// request being received piecemeal. Deadlines can be set and enforced in both
// Read and Write.

// repeatReader reads content count times, then EOFs.

// A Response that's just no bigger than 2KB, the buffer-before-chunking threshold.

// Listener for TestServerListenNotComparableListener.

// countCloseListener is a Listener wrapper that counts the number of Close calls.
