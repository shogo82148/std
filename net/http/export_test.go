// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Bridge package to expose http internals to tests in the http_test
// package.

package http

var ExportAppendTime = appendTime

var DefaultUserAgent = defaultUserAgent

var ExportServerNewConn = (*Server).newConn

var ExportCloseWriteAndWait = (*conn).closeWriteAndWait
