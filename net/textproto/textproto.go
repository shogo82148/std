// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージtextprotoは、HTTP、NNTP、およびSMTPのスタイルでテキストベースのリクエスト/レスポンスプロトコルの汎用サポートを実装します。
//
// このパッケージでは以下を提供します:
//
// Errorは、サーバーからの数値エラーレスポンスを表します。
//
// Pipelineは、クライアントでパイプライン化されたリクエストとレスポンスを管理するためのものです。
//
// Readerは、数値応答コードライン、キー: 値のヘッダ、先行スペースで折り返された行、独自の行にドットで終わる全文テキストブロックを読み取るためのものです。
//
// Writerは、ドットエンコードされたテキストブロックを書き込むためのものです。
//
// Connは、単一のネットワーク接続で使用するための、Reader、Writer、およびPipelineの便利なパッケージングです。
package textproto

import (
	"github.com/shogo82148/std/io"
)

// Errorは、サーバーからの数値エラーレスポンスを表します。
type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() string

// ProtocolErrorは、無効なレスポンスや切断された接続など、プロトコル違反を示すものです。
type ProtocolError string

func (p ProtocolError) Error() string

// Connはテキストネットワークプロトコルの接続を表します。
// それは、I/Oを管理するためのReaderとWriter、および接続上で並行リクエストをシーケンスするためのパイプラインで構成されています。
// これらの埋め込まれた型は、それらの型のドキュメントで詳細なメソッドを持っています。
type Conn struct {
	Reader
	Writer
	Pipeline
	conn io.ReadWriteCloser
}

// NewConnはI/Oにconnを使用して新しいConnを返します。
func NewConn(conn io.ReadWriteCloser) *Conn

// Close は接続を閉じます。
func (c *Conn) Close() error

// Dialは、net.Dialを使って指定されたネットワークの指定されたアドレスに接続し、接続のための新しいConnを返します。
func Dial(network, addr string) (*Conn, error)

// Cmdはパイプラインの順番を待ってからコマンドを送る便利なメソッドです。コマンドのテキストは、formatとargsを使用してフォーマットし、\r\nを追加した結果です。CmdはコマンドのIDを返し、StartResponseとEndResponseで使用します。
// 例えば、クライアントはHELPコマンドを実行し、ドットボディを返すかもしれません：
// id, err := c.Cmd("HELP")
//
//	if err != nil {
//	    return nil, err
//	}
//
// c.StartResponse(id)
// defer c.EndResponse(id)
//
//	if _, _, err = c.ReadCodeLine(110); err != nil {
//	    return nil, err
//	}
//
// text, err := c.ReadDotBytes()
//
//	if err != nil {
//	    return nil, err
//	}
//
// return c.ReadCodeLine(250)
func (c *Conn) Cmd(format string, args ...any) (id uint, err error)

// TrimStringは、先頭と末尾のASCIIスペースを除いたsを返します。
func TrimString(s string) string

// TrimBytesは、先頭と末尾のASCIIスペースを除いたbを返します。
func TrimBytes(b []byte) []byte
