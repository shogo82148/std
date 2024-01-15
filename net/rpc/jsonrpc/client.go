// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package jsonrpc はRPCパッケージのためのJSON-RPC 1.0のClientCodecとServerCodecを実装します。
// JSON-RPC 2.0のサポートについては、 https://godoc.org/?q=json-rpc+2.0 を参照してください。
package jsonrpc

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/rpc"
)

<<<<<<< HEAD
// NewClientCodecは、conn上でJSON-RPCを使用して新しいrpc.ClientCodecを返します。
func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec

// NewClientは、接続先の一連のサービスへのリクエストを処理する新しいrpc.Clientを返します。
=======
// NewClientCodec returns a new [rpc.ClientCodec] using JSON-RPC on conn.
func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec

// NewClient returns a new [rpc.Client] to handle requests to the
// set of services at the other end of the connection.
>>>>>>> upstream/master
func NewClient(conn io.ReadWriteCloser) *rpc.Client

// Dialは指定されたネットワークアドレスのJSON-RPCサーバに接続します。
func Dial(network, address string) (*rpc.Client, error)
