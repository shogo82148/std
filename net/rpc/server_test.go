// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith int

type Embed struct {
	hidden
}

// CodecEmulator provides a client-like api and a ServerCodec interface.
// Can be used to test ServeRequest.
type CodecEmulator struct {
	server        *Server
	serviceMethod string
	args          *Args
	reply         *Reply
	err           error
}

type ReplyNotPointer int
type ArgNotPublic int
type ReplyNotPublic int
type NeedsPtrType int

type WriteFailCodec int
