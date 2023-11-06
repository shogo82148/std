// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

// NewFileTransportは、提供されたFileSystemを提供する新しいRoundTripperを返します。
// 返されたRoundTripperは、その入力リクエストのURLホストを無視し、リクエストの他のほとんどのプロパティも無視します。
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
