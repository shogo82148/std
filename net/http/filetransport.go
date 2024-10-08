// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/io/fs"
)

// NewFileTransportは、提供された [FileSystem] を提供する新しい [RoundTripper] を返します。
// 返されたRoundTripperは、その入力リクエストのURLホストを無視します。また、リクエストの
// 他のほとんどのプロパティも無視します。
//
// NewFileTransport の典型的な使用例は、[Transport] に "file" プロトコルを登録することです。
// 例:
//
//	t := &http.Transport{}
//	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
//	c := &http.Client{Transport: t}
//	res, err := c.Get("file:///etc/passwd")
//	...
func NewFileTransport(fs FileSystem) RoundTripper

// NewFileTransportFS returns a new [RoundTripper], serving the provided
// file system fsys. The returned RoundTripper ignores the URL host in its
// incoming requests, as well as most other properties of the
// request. The files provided by fsys must implement [io.Seeker].
//
// The typical use case for NewFileTransportFS is to register the "file"
// protocol with a [Transport], as in:
//
//	fsys := os.DirFS("/")
//	t := &http.Transport{}
//	t.RegisterProtocol("file", http.NewFileTransportFS(fsys))
//	c := &http.Client{Transport: t}
//	res, err := c.Get("file:///etc/passwd")
//	...
func NewFileTransportFS(fsys fs.FS) RoundTripper
