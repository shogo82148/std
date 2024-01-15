// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package rpcは、オブジェクトのエクスポートされたメソッドに、ネットワークやその他のI/O接続を通じてアクセスする機能を提供します。サーバーはオブジェクトを登録し、オブジェクトのタイプ名に基づいてサービスとして表示されるようにします。登録後、オブジェクトのエクスポートされたメソッドはリモートからアクセス可能になります。サーバーは、異なるタイプの複数のオブジェクト（サービス）を登録することができますが、同じタイプの複数のオブジェクトを登録することはエラーです。

以下の条件を満たすメソッドのみがリモートアクセス可能になります。それ以外のメソッドは無視されます：

  - メソッドの型がエクスポートされていること。
  - メソッドがエクスポートされていること。
  - メソッドが2つの引数を持ち、両方の引数がエクスポートされている（または組み込み）型であること。
  - メソッドの2番目の引数がポインタであること。
  - メソッドが戻り値としてerror型を持つこと。

要するに、メソッドは次のようなスキーマである必要があります。

	func (t *T) MethodName(argType T1, replyType *T2) error

ここで、T1とT2はencoding/gobでマーシャリングできる型です。
これらの要件は、異なるコーデックが使用されている場合でも適用されます。
（将来的には、カスタムコーデックに対してこれらの要件は緩和されるかもしれません。）

<<<<<<< HEAD
メソッドの最初の引数は呼び出し元から提供される引数を表し、
2番目の引数は呼び出し元に返される結果パラメータを表します。
メソッドの戻り値がnilでない場合、それはクライアントがerrors.Newによって作成されたかのようにクライアントが確認する文字列として送り返されます。
エラーが返された場合、応答パラメータはクライアントに送り返されません。

サーバーは、ServeConnを呼び出すことによって単一の接続上のリクエストを処理することができます。また、通常はネットワークリスナーを作成し、Acceptを呼び出すか、HTTPリスナーの場合はHandleHTTPとhttp.Serveを呼び出します。

サービスを使用するためには、クライアントは接続を確立し、その後、接続上でNewClientを呼び出します。ダイヤル（DialHTTP）という便利な関数は、生のネットワーク接続（HTTP接続）に対して両方の手順を実行します。結果として得られるClientオブジェクトには、サービスとメソッドを指定するための2つのメソッド、CallとGoがあり、引数を含むポインタと結果パラメータを受け取るポインタを指定します。
=======
The method's first argument represents the arguments provided by the caller; the
second argument represents the result parameters to be returned to the caller.
The method's return value, if non-nil, is passed back as a string that the client
sees as if created by [errors.New].  If an error is returned, the reply parameter
will not be sent back to the client.

The server may handle requests on a single connection by calling [ServeConn].  More
typically it will create a network listener and call [Accept] or, for an HTTP
listener, [HandleHTTP] and [http.Serve].

A client wishing to use the service establishes a connection and then invokes
[NewClient] on the connection.  The convenience function [Dial] ([DialHTTP]) performs
both steps for a raw network connection (an HTTP connection).  The resulting
[Client] object has two methods, [Call] and Go, that specify the service and method to
call, a pointer containing the arguments, and a pointer to receive the result
parameters.
>>>>>>> upstream/master

Callメソッドは、リモート呼び出しが完了するまで待機し、
Goメソッドは非同期に呼び出しを開始し、Call構造体のDoneチャネルを使用して完了をシグナルします。

<<<<<<< HEAD
明示的なコーデックが設定されていない場合、データの転送にはencoding/gobパッケージが使用されます。
=======
Unless an explicit codec is set up, package [encoding/gob] is used to
transport the data.
>>>>>>> upstream/master

以下にシンプルな例を示します。サーバーはArithタイプのオブジェクトをエクスポートしたい場合です。

	package server

	import "errors"

	type Args struct {
		A, B int
	}

	type Quotient struct {
		Quo, Rem int
	}

	type Arith int

	func (t *Arith) Multiply(args *Args, reply *int) error {
		*reply = args.A * args.B
		return nil
	}

	func (t *Arith) Divide(args *Args, quo *Quotient) error {
		if args.B == 0 {
			return errors.New("divide by zero")
		}
		quo.Quo = args.A / args.B
		quo.Rem = args.A % args.B
		return nil
	}

サーバーの呼び出し（HTTPサービスの場合）：

	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)

この時点で、クライアントは"Arith"というサービスとそのメソッド"Arith.Multiply"、"Arith.Divide"を見ることができます。呼び出すためには、クライアントはまずサーバーにダイヤルします。

	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

そして、リモート呼び出しを行うことができます。

	// 同期呼び出し
	args := &server.Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

または

	// 非同期呼び出し
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done	// divCallと等しい
	// エラーをチェックし、出力などを行います。

サーバーの実装では、クライアントのためのシンプルで型セーフなラッパーを提供することがよくあります。

net/rpcパッケージは凍結されており、新しい機能は受け付けていません。
*/
package rpc

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/sync"
)

const (
	// HandleHTTPで使用されるデフォルト値
	DefaultRPCPath   = "/_goRPC_"
	DefaultDebugPath = "/debug/rpc"
)

// RequestはRPC呼び出しの前に書かれるヘッダーです。内部で使用されますが、ネットワークトラフィックを分析する際などデバッグの支援のためにここで記述されています。
type Request struct {
	ServiceMethod string
	Seq           uint64
	next          *Request
}

// Responseは、すべてのRPCの戻り値の前に書かれるヘッダです。内部で使用されますが、ネットワークトラフィックを分析する際など、デバッグの支援としてここで文書化されています。
type Response struct {
	ServiceMethod string
	Seq           uint64
	Error         string
	next          *Response
}

// ServerはRPCサーバーを表します。
type Server struct {
	serviceMap sync.Map
	reqLock    sync.Mutex
	freeReq    *Request
	respLock   sync.Mutex
	freeResp   *Response
}

<<<<<<< HEAD
// NewServerは新しいServerを返します。
func NewServer() *Server

// DefaultServerは*Serverのデフォルトインスタンスです。
=======
// NewServer returns a new [Server].
func NewServer() *Server

// DefaultServer is the default instance of [*Server].
>>>>>>> upstream/master
var DefaultServer = NewServer()

// Registerは、以下の条件を満たすレシーバーのメソッドのセットをサーバーに公開します：
//   - エクスポートされた型のエクスポートされたメソッド
//   - 2つの引数、両方がエクスポートされた型
//   - 2番目の引数がポインタであること
//   - エラー型の1つの戻り値
//
// レシーバーがエクスポートされた型でないか、適切なメソッドがない場合は、エラーを返します。また、エラーをパッケージlogを使用してログに記録します。
// クライアントは "Type.Method" の形式の文字列を使用して各メソッドにアクセスします。ここで、Typeはレシーバーの具体的な型です。
func (server *Server) Register(rcvr any) error

<<<<<<< HEAD
// RegisterNameは、レシーバの具体的な型の代わりに、指定された名前を使用して型を登録します。
func (server *Server) RegisterName(name string, rcvr any) error

// ServeConnは単一の接続上でサーバーを実行します。
// ServeConnはクライアントが切断するまで接続をサーブし続け、ブロックします。
// 通常、呼び出し元はgo文でServeConnを呼び出します。
// ServeConnは接続上でgobワイヤフォーマット（package gobを参照）を使用します。
// 別のコーデックを使用するには、ServeCodecを使用してください。
// 同時アクセスに関する情報については、NewClientのコメントを参照してください。
func (server *Server) ServeConn(conn io.ReadWriteCloser)

// ServeCodecは、指定されたコーデックを使用してリクエストをデコードし、レスポンスをエンコードするためにServeConnと似ています。
func (server *Server) ServeCodec(codec ServerCodec)

// ServeRequestはServeCodecと似ていますが、1つのリクエストを同期的に処理します。
// 完了時にコーデックを閉じません。
=======
// RegisterName is like [Register] but uses the provided name for the type
// instead of the receiver's concrete type.
func (server *Server) RegisterName(name string, rcvr any) error

// ServeConn runs the server on a single connection.
// ServeConn blocks, serving the connection until the client hangs up.
// The caller typically invokes ServeConn in a go statement.
// ServeConn uses the gob wire format (see package gob) on the
// connection. To use an alternate codec, use [ServeCodec].
// See [NewClient]'s comment for information about concurrent access.
func (server *Server) ServeConn(conn io.ReadWriteCloser)

// ServeCodec is like [ServeConn] but uses the specified codec to
// decode requests and encode responses.
func (server *Server) ServeCodec(codec ServerCodec)

// ServeRequest is like [ServeCodec] but synchronously serves a single request.
// It does not close the codec upon completion.
>>>>>>> upstream/master
func (server *Server) ServeRequest(codec ServerCodec) error

// Acceptはリスナー上で接続を受け入れ、各受信接続のリクエストを処理します。
// Acceptはリスナーがnon-nilのエラーを返すまでブロックされます。通常、呼び出し元はgoステートメントでAcceptを呼び出します。
func (server *Server) Accept(lis net.Listener)

<<<<<<< HEAD
// RegisterはレシーバのメソッドをDefaultServerに登録します。
func Register(rcvr any) error

// RegisterNameは、レシーバの具体的な型ではなく、与えられた名前を型として使用します。Registerと同様の動作です。
func RegisterName(name string, rcvr any) error

// ServerCodecはRPCセッションのサーバー側でのRPCリクエストの読み取りとRPCレスポンスの書き込みを実装します。
// サーバーはReadRequestHeaderとReadRequestBodyをペアで呼び出して接続からリクエストを読み取り、WriteResponseを呼び出してレスポンスを書き込みます。
// サーバーは接続が終了したらCloseを呼び出します。ReadRequestBodyはnilの引数で呼び出されることがあり、リクエストの本文を読み取って破棄するためのものです。
// 同時アクセスに関する情報については、NewClientのコメントを参照してください。
=======
// Register publishes the receiver's methods in the [DefaultServer].
func Register(rcvr any) error

// RegisterName is like [Register] but uses the provided name for the type
// instead of the receiver's concrete type.
func RegisterName(name string, rcvr any) error

// A ServerCodec implements reading of RPC requests and writing of
// RPC responses for the server side of an RPC session.
// The server calls [ServerCodec.ReadRequestHeader] and [ServerCodec.ReadRequestBody] in pairs
// to read requests from the connection, and it calls [ServerCodec.WriteResponse] to
// write a response back. The server calls [ServerCodec.Close] when finished with the
// connection. ReadRequestBody may be called with a nil
// argument to force the body of the request to be read and discarded.
// See [NewClient]'s comment for information about concurrent access.
>>>>>>> upstream/master
type ServerCodec interface {
	ReadRequestHeader(*Request) error
	ReadRequestBody(any) error
	WriteResponse(*Response, any) error

	Close() error
}

<<<<<<< HEAD
// ServeConnはデフォルトサーバーを単一の接続上で実行します。
// ServeConnは、クライアントが切断するまで接続を処理するまでブロックします。
// 通常、呼び出し元はgo文でServeConnを呼び出します。
// ServeConnは、接続上でgobワイヤーフォーマット（パッケージgobを参照）を使用します。
// 別のコーデックを使用するには、ServeCodecを使用してください。
// 同時アクセスに関する情報については、NewClientのコメントを参照してください。
func ServeConn(conn io.ReadWriteCloser)

// ServeCodecはServeConnと似ていますが、指定されたコーデックを使用して
// リクエストをデコードし、レスポンスをエンコードします。
func ServeCodec(codec ServerCodec)

// ServeRequest は ServeCodec に似ていますが、単一のリクエストを同期的に処理します。
// 処理が完了してもコーデックを閉じません。
func ServeRequest(codec ServerCodec) error

// Acceptはリスナー上で接続を受け付け、各受信された接続に対してDefaultServerにリクエストを処理します。
// Acceptはブロックします。通常、呼び出し元はgo文でそれを呼び出します。
func Accept(lis net.Listener)

// ServeHTTPはRPCリクエストに答えるためのhttp.Handlerを実装します。
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)

// HandleHTTPはrpcPathでRPCメッセージのためのHTTPハンドラを登録し、debugPathではデバッグハンドラを登録します。
// 通常はGoステートメント内でhttp.Serve()を呼び出す必要があります。
func (server *Server) HandleHTTP(rpcPath, debugPath string)

// HandleHTTPはRPCメッセージのためのHTTPハンドラをDefaultServerに登録し、DefaultRPCPathにデバッグハンドラを登録します。
// 通常はgoステートメントでhttp.Serve()を呼び出す必要があります。
=======
// ServeConn runs the [DefaultServer] on a single connection.
// ServeConn blocks, serving the connection until the client hangs up.
// The caller typically invokes ServeConn in a go statement.
// ServeConn uses the gob wire format (see package gob) on the
// connection. To use an alternate codec, use [ServeCodec].
// See [NewClient]'s comment for information about concurrent access.
func ServeConn(conn io.ReadWriteCloser)

// ServeCodec is like [ServeConn] but uses the specified codec to
// decode requests and encode responses.
func ServeCodec(codec ServerCodec)

// ServeRequest is like [ServeCodec] but synchronously serves a single request.
// It does not close the codec upon completion.
func ServeRequest(codec ServerCodec) error

// Accept accepts connections on the listener and serves requests
// to [DefaultServer] for each incoming connection.
// Accept blocks; the caller typically invokes it in a go statement.
func Accept(lis net.Listener)

// ServeHTTP implements an [http.Handler] that answers RPC requests.
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)

// HandleHTTP registers an HTTP handler for RPC messages on rpcPath,
// and a debugging handler on debugPath.
// It is still necessary to invoke [http.Serve](), typically in a go statement.
func (server *Server) HandleHTTP(rpcPath, debugPath string)

// HandleHTTP registers an HTTP handler for RPC messages to [DefaultServer]
// on [DefaultRPCPath] and a debugging handler on [DefaultDebugPath].
// It is still necessary to invoke [http.Serve](), typically in a go statement.
>>>>>>> upstream/master
func HandleHTTP()
