// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

// Responseは、すべてのRPCの戻り値の前に書かれるヘッダーです。内部で使用されますが、ネットワークトラフィックを分析する際など、デバッグの支援としてここで文書化されています。
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
//   - 2番目の引数がポインターであること
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

// Registerはレシーバーのメソッドを [DefaultServer] に登録します。
func Register(rcvr any) error

// RegisterNameは、レシーバーの具体的な型ではなく、与えられた名前を型として使用します。[Register] と同様の動作です。
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

// HandleHTTPはrpcPathでRPCメッセージのためのHTTPハンドラーを登録し、debugPathではデバッグハンドラーを登録します。
// 通常はGoステートメント内で [http.Serve]()を呼び出す必要があります。
func (server *Server) HandleHTTP(rpcPath, debugPath string)

// HandleHTTPはRPCメッセージのためのHTTPハンドラーを [DefaultServer] に登録し、[DefaultRPCPath] にデバッグハンドラーを登録します。
// 通常はgoステートメントで [http.Serve]()を呼び出す必要があります。
func HandleHTTP()
