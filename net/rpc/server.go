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

メソッドの最初の引数は呼び出し元から提供される引数を表し、
2番目の引数は呼び出し元に返される結果パラメータを表します。
メソッドの戻り値がnilでない場合、それはクライアントが [errors.New] によって作成されたかのようにクライアントが確認する文字列として送り返されます。
エラーが返された場合、応答パラメータはクライアントに送り返されません。

サーバーは、[ServeConn] を呼び出すことによって単一の接続上のリクエストを処理することができます。また、通常はネットワークリスナーを作成し、[Accept] を呼び出すか、HTTPリスナーの場合は [HandleHTTP] と [http.Serve] を呼び出します。

サービスを使用するためには、クライアントは接続を確立し、その後、接続上で [NewClient] を呼び出します。[Dial]（[DialHTTP]）という便利な関数は、生のネットワーク接続（HTTP接続）に対して両方の手順を実行します。結果として得られる [Client] オブジェクトには、サービスとメソッドを指定するための2つのメソッド、[Call] とGoがあり、引数を含むポインタと結果パラメータを受け取るポインタを指定します。

Callメソッドは、リモート呼び出しが完了するまで待機し、
Goメソッドは非同期に呼び出しを開始し、Call構造体のDoneチャネルを使用して完了をシグナルします。

明示的なコーデックが設定されていない場合、データの転送には [encoding/gob] パッケージが使用されます。

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

// NewServerは新しい [Server] を返します。
func NewServer() *Server

// DefaultServerは [*Server] のデフォルトインスタンスです。
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

// RegisterNameは [Register] と同様ですが、レシーバーの具体的な型の代わりに提供された名前を型に使用します。
func (server *Server) RegisterName(name string, rcvr any) error

// ServeConnは、単一の接続上でサーバーを実行します。
// ServeConnは、クライアントが切断するまで接続を提供するためにブロックします。
// 呼び出し元は通常、goステートメントでServeConnを呼び出します。
// ServeConnは、接続上でgobワイヤーフォーマット（パッケージgobを参照）を使用します。
// 代替のコーデックを使用するには、[ServeCodec] を使用します。
// 並行アクセスに関する情報については、[NewClient] のコメントを参照してください。
func (server *Server) ServeConn(conn io.ReadWriteCloser)

// ServeCodecは [ServeConn] と同様ですが、指定されたコーデックを使用して
// リクエストをデコードし、レスポンスをエンコードします。
func (server *Server) ServeCodec(codec ServerCodec)

// ServeRequestは [ServeCodec] と似ていますが、単一のリクエストを同期的に提供します。
// 完了時にコーデックを閉じることはありません。
func (server *Server) ServeRequest(codec ServerCodec) error

// Acceptはリスナー上で接続を受け入れ、各受信接続のリクエストを処理します。
// Acceptはリスナーがnon-nilのエラーを返すまでブロックされます。通常、呼び出し元はgoステートメントでAcceptを呼び出します。
func (server *Server) Accept(lis net.Listener)

// Registerはレシーバのメソッドを [DefaultServer] に登録します。
func Register(rcvr any) error

// RegisterNameは、レシーバの具体的な型ではなく、与えられた名前を型として使用します。[Register] と同様の動作です。
func RegisterName(name string, rcvr any) error

// ServerCodecはRPCセッションのサーバー側でのRPCリクエストの読み取りとRPCレスポンスの書き込みを実装します。
// サーバーは [ServerCodec.ReadRequestHeader] と [ServerCodec.ReadRequestBody] をペアで呼び出して接続からリクエストを読み取り、[ServerCodec.WriteResponse] を呼び出してレスポンスを書き込みます。
// サーバーは接続が終了したら [ServerCodec.Close] を呼び出します。ReadRequestBodyはnilの引数で呼び出されることがあり、リクエストの本文を読み取って破棄するためのものです。
// 同時アクセスに関する情報については、[NewClient] のコメントを参照してください。
type ServerCodec interface {
	ReadRequestHeader(*Request) error
	ReadRequestBody(any) error
	WriteResponse(*Response, any) error

	Close() error
}

// ServeConnは [DefaultServer] を単一の接続上で実行します。
// ServeConnは、クライアントが切断するまで接続を処理するまでブロックします。
// 通常、呼び出し元はgo文でServeConnを呼び出します。
// ServeConnは、接続上でgobワイヤーフォーマット（パッケージgobを参照）を使用します。
// 別のコーデックを使用するには、[ServeCodec] を使用してください。
// 同時アクセスに関する情報については、[NewClient] のコメントを参照してください。
func ServeConn(conn io.ReadWriteCloser)

// ServeCodecは [ServeConn] と似ていますが、指定されたコーデックを使用して
// リクエストをデコードし、レスポンスをエンコードします。
func ServeCodec(codec ServerCodec)

// ServeRequest は [ServeCodec] に似ていますが、単一のリクエストを同期的に処理します。
// 処理が完了してもコーデックを閉じません。
func ServeRequest(codec ServerCodec) error

// Acceptはリスナー上で接続を受け付け、各受信された接続に対して [DefaultServer] にリクエストを処理します。
// Acceptはブロックします。通常、呼び出し元はgo文でそれを呼び出します。
func Accept(lis net.Listener)

// ServeHTTPはRPCリクエストに答えるための [http.Handler] を実装します。
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)

// HandleHTTPはrpcPathでRPCメッセージのためのHTTPハンドラを登録し、debugPathではデバッグハンドラを登録します。
// 通常はGoステートメント内で [http.Serve]()を呼び出す必要があります。
func (server *Server) HandleHTTP(rpcPath, debugPath string)

// HandleHTTPはRPCメッセージのためのHTTPハンドラを [DefaultServer] に登録し、[DefaultRPCPath] にデバッグハンドラを登録します。
// 通常はgoステートメントで [http.Serve]()を呼び出す必要があります。
func HandleHTTP()
