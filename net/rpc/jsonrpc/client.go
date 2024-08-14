// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// jsonrpcパッケージはRPCパッケージのためのJSON-RPC 1.0のClientCodecとServerCodecを実装します。
// JSON-RPC 2.0のサポートについては、 https://godoc.org/?q=json-rpc+2.0 を参照してください。
package jsonrpc

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/rpc"
)

// NewClientCodecは、conn上でJSON-RPCを使用して新しい [rpc.ClientCodec] を返します。
func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec

// NewClientは、接続先の一連のサービスへのリクエストを処理する新しい [rpc.Client] を返します。
func NewClient(conn io.ReadWriteCloser) *rpc.Client

// Dialは指定されたネットワークアドレスのJSON-RPCサーバに接続します。
func Dial(network, address string) (*rpc.Client, error)
