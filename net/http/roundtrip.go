// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js

package http

// RoundTripは、RoundTripperインターフェースを実装します。
//
// クッキーやリダイレクトの処理などのより高度なHTTPクライアントサポートについては、
// Get、Post、およびClient型を参照してください。
//
// RoundTripインターフェースと同様に、RoundTripによって返されるエラータイプは未指定です。
func (t *Transport) RoundTrip(req *Request) (*Response, error)
