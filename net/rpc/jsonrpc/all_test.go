// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonrpc

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith int

type ArithAddResp struct {
	Id     any   `json:"id"`
	Result Reply `json:"result"`
	Error  any   `json:"error"`
}

type BuiltinTypes struct{}
