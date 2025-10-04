// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/time"
)

// ResponseControllerは、HTTPハンドラーがレスポンスを制御するために使用されます。
//
// [Handler.ServeHTTP] メソッドが返された後にResponseControllerを使用することはできません。
type ResponseController struct {
	rw ResponseWriter
}

// NewResponseControllerは、リクエスト用の [ResponseController] を作成します。
//
// ResponseWriterは、[Handler.ServeHTTP] メソッドに渡された元の値である必要があります。
// または、元のResponseWriterを返すUnwrapメソッドを持っている必要があります。
//
// ResponseWriterが次のいずれかのメソッドを実装している場合、ResponseControllerは
// 適切に呼び出します。
//
//	Flush()
//	FlushError() error // エラーを返す代替Flush
//	Hijack() (net.Conn、*bufio.ReadWriter、error)
//	SetReadDeadline(deadline time.Time) error
//	SetWriteDeadline(deadline time.Time) error
//	EnableFullDuplex() error
//
// ResponseWriterがメソッドをサポートしていない場合、ResponseControllerは
// [ErrNotSupported] に一致するエラーを返します。
func NewResponseController(rw ResponseWriter) *ResponseController

// Flushは、バッファリングされたデータをクライアントにフラッシュします。
func (c *ResponseController) Flush() error

<<<<<<< HEAD
// Hijackは、呼び出し元が接続を引き継ぐことを可能にします。
// 詳細については、Hijackerインターフェースを参照してください。
=======
// Hijack lets the caller take over the connection.
// See the [Hijacker] interface for details.
>>>>>>> upstream/release-branch.go1.25
func (c *ResponseController) Hijack() (net.Conn, *bufio.ReadWriter, error)

// SetReadDeadlineは、ボディを含むリクエスト全体の読み取りの期限を設定します。
// 期限が超過した後にリクエストボディから読み取りを行うと、エラーが返されます。
// ゼロ値は期限がないことを意味します。
//
// 期限が超過した後に読み取り期限を設定しても、期限は延長されません。
func (c *ResponseController) SetReadDeadline(deadline time.Time) error

// SetWriteDeadlineは、レスポンスの書き込みの期限を設定します。
// 期限が超過した後にレスポンスボディに書き込みを行うと、ブロックされず、
// データがバッファリングされている場合は成功する可能性があります。
// ゼロ値は期限がないことを意味します。
//
// 期限が超過した後に書き込み期限を設定しても、期限は延長されません。
func (c *ResponseController) SetWriteDeadline(deadline time.Time) error

// EnableFullDuplexは、リクエストハンドラが [Request.Body] からの読み取りを交互に行い、
// [ResponseWriter] への書き込みと交互に行うことを示します。
//
// HTTP/1リクエストの場合、Go HTTPサーバーはデフォルトで、レスポンスの書き込みを開始する前に
// リクエストボディの未読部分を消費し、ハンドラがリクエストから読み取りとレスポンスの書き込みを
// 同時に行うことを防止します。EnableFullDuplexを呼び出すと、この動作が無効になり、
// ハンドラがリクエストからの読み取りを続けながらレスポンスを同時に書き込むことができるようになります。
//
// HTTP/2リクエストの場合、Go HTTPサーバーは常に並行して読み取りとレスポンスを許可します。
func (c *ResponseController) EnableFullDuplex() error
