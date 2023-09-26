// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/io/fs"
)

// fileTransport implements RoundTripper for the 'file' protocol.

// NewFileTransport returns a new RoundTripper, serving the provided
// FileSystem. The returned RoundTripper ignores the URL host in its
// incoming requests, as well as most other properties of the
// request.
//
// The typical use case for NewFileTransport is to register the "file"
// protocol with a Transport, as in:
//
//	t := &http.Transport{}
//	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
//	c := &http.Client{Transport: t}
//	res, err := c.Get("file:///etc/passwd")
//	...
func NewFileTransport(fs FileSystem) RoundTripper

// NewFileTransportFS returns a new RoundTripper, serving the provided
// file system fsys. The returned RoundTripper ignores the URL host in its
// incoming requests, as well as most other properties of the
// request.
//
// The typical use case for NewFileTransportFS is to register the "file"
// protocol with a Transport, as in:
//
//	fsys := os.DirFS("/")
//	t := &http.Transport{}
//	t.RegisterProtocol("file", http.NewFileTransportFS(fsys))
//	c := &http.Client{Transport: t}
//	res, err := c.Get("file:///etc/passwd")
//	...
func NewFileTransportFS(fsys fs.FS) RoundTripper

// populateResponse is a ResponseWriter that populates the *Response
// in res, and writes its body to a pipe connected to the response
// body. Once writes begin or finish() is called, the response is sent
// on ch.
