// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js

package http

<<<<<<< HEAD
// RoundTripは、RoundTripperインターフェースを実装します。
//
// クッキーやリダイレクトの処理などのより高度なHTTPクライアントサポートについては、
// Get、Post、およびClient型を参照してください。
=======
// RoundTrip implements the [RoundTripper] interface.
//
// For higher-level HTTP client support (such as handling of cookies
// and redirects), see [Get], [Post], and the [Client] type.
>>>>>>> upstream/release-branch.go1.22
//
// RoundTripインターフェースと同様に、RoundTripによって返されるエラータイプは未指定です。
func (t *Transport) RoundTrip(req *Request) (*Response, error)
