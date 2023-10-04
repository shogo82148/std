// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

<<<<<<< HEAD
// fileTransport implements RoundTripper for the 'file' protocol.

// NewFileTransport は、提供された FileSystem を提供する新しい RoundTripper を返します。
// 返された RoundTripper は、リクエストの URL ホストやその他のプロパティを無視します。
=======
// NewFileTransport returns a new RoundTripper, serving the provided
// FileSystem. The returned RoundTripper ignores the URL host in its
// incoming requests, as well as most other properties of the
// request.
>>>>>>> release-branch.go1.21
//
// NewFileTransport の典型的な使用例は、Transport に "file" プロトコルを登録することです。
// 例:
//
//	t := &http.Transport{}
//	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
//	c := &http.Client{Transport: t}
//	res, err := c.Get("file:///etc/passwd")
//	...
func NewFileTransport(fs FileSystem) RoundTripper
