// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

package http

// RoundTripは、WHATWG Fetch APIを使用してRoundTripperインターフェースを実装します。
func (t *Transport) RoundTrip(req *Request) (*Response, error)
