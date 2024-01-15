// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

package http

<<<<<<< HEAD
// RoundTripは、WHATWG Fetch APIを使用してRoundTripperインターフェースを実装します。
=======
// RoundTrip implements the [RoundTripper] interface using the WHATWG Fetch API.
>>>>>>> upstream/master
func (t *Transport) RoundTrip(req *Request) (*Response, error)
