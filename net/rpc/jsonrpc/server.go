// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonrpc

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/rpc"
)

<<<<<<< HEAD
// NewServerCodecはconn上でJSON-RPCを使用する新しいrpc.ServerCodecを返します。
=======
// NewServerCodec returns a new [rpc.ServerCodec] using JSON-RPC on conn.
>>>>>>> upstream/release-branch.go1.22
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec

// ServeConnは単一の接続でJSON-RPCサーバーを実行します。
// ServeConnはブロックし、クライアントが切断するまで接続を処理します。
// 呼び出し元は通常、goステートメントでServeConnを呼び出します。
func ServeConn(conn io.ReadWriteCloser)
