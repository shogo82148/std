// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Bridge package to expose http internals to tests in the http_test
// package.

package http

var (
	DefaultUserAgent             = defaultUserAgent
	NewLoggingConn               = newLoggingConn
	ExportAppendTime             = appendTime
	ExportRefererForURL          = refererForURL
	ExportServerNewConn          = (*Server).newConn
	ExportCloseWriteAndWait      = (*conn).closeWriteAndWait
	ExportErrRequestCanceled     = errRequestCanceled
	ExportErrRequestCanceledConn = errRequestCanceledConn
	ExportServeFile              = serveFile
	ExportHttp2ConfigureServer   = http2ConfigureServer
)

var (
	SetEnterRoundTripHook = hookSetter(&testHookEnterRoundTrip)
	SetRoundTripRetried   = hookSetter(&testHookRoundTripRetried)
)
