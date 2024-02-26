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

// ClientCodecは、RPCセッションのクライアント側において、RPCリクエストの書き込みとRPCレスポンスの読み取りを実装します。
// クライアントは [ClientCodec.WriteRequest] を呼び出して接続にリクエストを書き込み、
// [ClientCodec.ReadResponseHeader] と [ClientCodec.ReadResponseBody] をペアで呼び出してレスポンスを読み込みます。
// 接続が終了したら、クライアントは [ClientCodec.Close] を呼び出します。
// ReadResponseBodyは、nilの引数で呼び出されることがあり、レスポンスの本文を読み取り、その後破棄するように強制することができます。
// 同時アクセスに関する情報については、[NewClient] のコメントを参照してください。
type ClientCodec interface {
	WriteRequest(*Request, any) error
	ReadResponseHeader(*Response) error
	ReadResponseBody(any) error

	Close() error
}

// NewClientは、接続先のサービスセットに対するリクエストを処理するための新しい [Client] を返します。
// 接続の書き込み側にはバッファが追加されるため、ヘッダとペイロードがまとめて送信されます。
//
// 接続の読み込み側と書き込み側はそれぞれ独立してシリアライズされるため、相互ロックは必要ありません。ただし、各半分は同時にアクセスされる可能性があるため、connの実装は同時読み取りや同時書き込みに対して保護する必要があります。
func NewClient(conn io.ReadWriteCloser) *Client

// NewClientWithCodecは、指定されたコーデックを使用してリクエストをエンコードし、レスポンスをデコードする [NewClient] と同様です。
func NewClientWithCodec(codec ClientCodec) *Client

// DialHTTPは、デフォルトのHTTP RPCパスで待ち受けている、指定されたネットワークアドレスのHTTP RPCサーバーに接続します。
func DialHTTP(network, address string) (*Client, error)

// DialHTTPPathは指定したネットワークアドレスとパスでHTTP RPCサーバに接続します。
func DialHTTPPath(network, address, path string) (*Client, error)

// Dialは指定されたネットワークアドレスのRPCサーバに接続します。
func Dial(network, address string) (*Client, error)

// Closeは基礎となるコーデックのCloseメソッドを呼び出します。接続がすでに
// シャットダウン中の場合、[ErrShutdown] が返されます。
func (client *Client) Close() error

// Go invokes the function asynchronously. It returns the [Call] structure representing
// the invocation. The done channel will signal when the call is complete by returning
// the same Call object. If done is nil, Go will allocate a new channel.
// If non-nil, done must be buffered or Go will deliberately crash.
func (client *Client) Go(serviceMethod string, args any, reply any, done chan *Call) *Call

// Callは指定された関数を呼び出し、その完了を待ち、エラー状態を返します。
func (client *Client) Call(serviceMethod string, args any, reply any) error
