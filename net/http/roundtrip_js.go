// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

package http

// RoundTrip implements the RoundTripper interface using the WHATWG Fetch API.
func (t *Transport) RoundTrip(req *Request) (*Response, error)
