// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

// ServerErrorは、RPC接続のリモート側から返されたエラーを表します。
type ServerError string

func (e ServerError) Error() string

var ErrShutdown = errors.New("connection is shut down")

// CallはアクティブなRPCを表します。
type Call struct {
	ServiceMethod string
	Args          any
	Reply         any
	Error         error
	Done          chan *Call
}

// ClientはRPCクライアントを表します。
// 単一のクライアントに関連付けられている複数の保留中の呼び出しがあり、
// クライアントは同時に複数のゴルーチンによって使用される場合があります。
type Client struct {
	codec ClientCodec

	reqMutex sync.Mutex
	request  Request

	mutex    sync.Mutex
	seq      uint64
	pending  map[uint64]*Call
	closing  bool
	shutdown bool
}

<<<<<<< HEAD
// ClientCodecは、RPCセッションのクライアント側において、RPCリクエストの書き込みとRPCレスポンスの読み取りを実装します。
// クライアントはWriteRequestを呼び出して接続にリクエストを書き込み、
// ReadResponseHeaderとReadResponseBodyをペアで呼び出してレスポンスを読み込みます。
// 接続が終了したら、クライアントはCloseを呼び出します。
// ReadResponseBodyは、nilの引数で呼び出されることがあり、レスポンスの本文を読み取り、その後破棄するように強制することができます。
// 同時アクセスに関する情報については、NewClientのコメントを参照してください。
=======
// A ClientCodec implements writing of RPC requests and
// reading of RPC responses for the client side of an RPC session.
// The client calls [ClientCodec.WriteRequest] to write a request to the connection
// and calls [ClientCodec.ReadResponseHeader] and [ClientCodec.ReadResponseBody] in pairs
// to read responses. The client calls [ClientCodec.Close] when finished with the
// connection. ReadResponseBody may be called with a nil
// argument to force the body of the response to be read and then
// discarded.
// See [NewClient]'s comment for information about concurrent access.
>>>>>>> upstream/release-branch.go1.22
type ClientCodec interface {
	WriteRequest(*Request, any) error
	ReadResponseHeader(*Response) error
	ReadResponseBody(any) error

	Close() error
}

<<<<<<< HEAD
// NewClientは、接続先のサービスセットに対するリクエストを処理するための新しいクライアントを返します。
// 接続の書き込み側にはバッファが追加されるため、ヘッダとペイロードがまとめて送信されます。
=======
// NewClient returns a new [Client] to handle requests to the
// set of services at the other end of the connection.
// It adds a buffer to the write side of the connection so
// the header and payload are sent as a unit.
>>>>>>> upstream/release-branch.go1.22
//
// 接続の読み込み側と書き込み側はそれぞれ独立してシリアライズされるため、相互ロックは必要ありません。ただし、各半分は同時にアクセスされる可能性があるため、connの実装は同時読み取りや同時書き込みに対して保護する必要があります。
func NewClient(conn io.ReadWriteCloser) *Client

<<<<<<< HEAD
// NewClientWithCodecは、指定されたコーデックを使用してリクエストをエンコードし、レスポンスをデコードするNewClientと同様です。
=======
// NewClientWithCodec is like [NewClient] but uses the specified
// codec to encode requests and decode responses.
>>>>>>> upstream/release-branch.go1.22
func NewClientWithCodec(codec ClientCodec) *Client

// DialHTTPは、デフォルトのHTTP RPCパスで待ち受けている、指定されたネットワークアドレスのHTTP RPCサーバーに接続します。
func DialHTTP(network, address string) (*Client, error)

// DialHTTPPathは指定したネットワークアドレスとパスでHTTP RPCサーバに接続します。
func DialHTTPPath(network, address, path string) (*Client, error)

// Dialは指定されたネットワークアドレスのRPCサーバに接続します。
func Dial(network, address string) (*Client, error)

<<<<<<< HEAD
// Closeは基礎となるコーデックのCloseメソッドを呼び出します。接続がすでに
// シャットダウン中の場合、ErrShutdownが返されます。
=======
// Close calls the underlying codec's Close method. If the connection is already
// shutting down, [ErrShutdown] is returned.
>>>>>>> upstream/release-branch.go1.22
func (client *Client) Close() error

// Go invokes the function asynchronously. It returns the [Call] structure representing
// the invocation. The done channel will signal when the call is complete by returning
// the same Call object. If done is nil, Go will allocate a new channel.
// If non-nil, done must be buffered or Go will deliberately crash.
func (client *Client) Go(serviceMethod string, args any, reply any, done chan *Call) *Call

// Callは指定された関数を呼び出し、その完了を待ち、エラー状態を返します。
func (client *Client) Call(serviceMethod string, args any, reply any) error
