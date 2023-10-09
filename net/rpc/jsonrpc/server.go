// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonrpc

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/rpc"
)

// NewServerCodecはconn上でJSON-RPCを使用する新しいrpc.ServerCodecを返します。
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec

// ServeConnは単一の接続でJSON-RPCサーバーを実行します。
// ServeConnはブロックし、クライアントが切断するまで接続を処理します。
// 呼び出し元は通常、goステートメントでServeConnを呼び出します。
func ServeConn(conn io.ReadWriteCloser)
