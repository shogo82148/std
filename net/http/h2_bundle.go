//go:build !nethttpomithttp2
// +build !nethttpomithttp2

// Code generated by golang.org/x/tools/cmd/bundle. DO NOT EDIT.
//   $ bundle -o=h2_bundle.go -prefix=http2 -tags=!nethttpomithttp2 golang.org/x/net/http2

// Package http2 implements the HTTP/2 protocol.
//
// This package is low-level and intended to be used directly by very
// few people. Most users will use it indirectly through the automatic
// use by the net/http package (from Go 1.6 and later).
// For use in earlier Go versions see ConfigureServer. (Transport support
// requires Go 1.6 or later)
//
// See https://http2.github.io/ for more information on HTTP/2.
//
// See https://http2.golang.org/ for a test server running this code.
//

package http

<<<<<<< HEAD
// ClientConnPool manages a pool of HTTP/2 client connections.

// clientConnPoolIdleCloser is the interface implemented by ClientConnPool
// implementations which can close their idle connections.

=======
>>>>>>> release-branch.go1.21
var (
	_ http2clientConnPoolIdleCloser = (*http2clientConnPool)(nil)
	_ http2clientConnPoolIdleCloser = http2noDialClientConnPool{}
)

// Optional http.ResponseWriter interfaces implemented.
var (
	_ CloseNotifier     = (*http2responseWriter)(nil)
	_ Flusher           = (*http2responseWriter)(nil)
	_ http2stringWriter = (*http2responseWriter)(nil)
)

var _ Pusher = (*http2responseWriter)(nil)
