// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textproto

import (
	"github.com/shogo82148/std/sync"
)

// パイプラインは、順番にリクエストとレスポンスを管理するためのものです。
//
// 接続上の複数のクライアントを管理するために、Pipeline p を使用する場合、
// それぞれのクライアントは次のように実行する必要があります：
//
//	id := p.Next()	// 番号を取得する
//
//	p.StartRequest(id)	// リクエストを送信する順番を待つ
//	«リクエストを送信する»
//	p.EndRequest(id)	// リクエストの送信が完了したことをPipelineに通知する
//
//	p.StartResponse(id)	// レスポンスを読み取る順番を待つ
//	«レスポンスを読み取る»
//	p.EndResponse(id)	// レスポンスの読み取りが完了したことをPipelineに通知する
//
// パイプラインサーバーでも、同じ呼び出しを使用して、並列で計算されたレスポンスを正しい順序で書き込むことができます。
type Pipeline struct {
	mu       sync.Mutex
	id       uint
	request  sequencer
	response sequencer
}

// Nextはリクエスト/レスポンスのペアの次のIDを返します。
func (p *Pipeline) Next() uint

// StartRequestは、指定したIDでリクエストを送信（または、サーバーの場合は受信）する時間になるまでブロックします。
func (p *Pipeline) StartRequest(id uint)

// EndRequestは、与えられたIDを持つリクエストが送信されたことをpに通知します
// (または、これがサーバーの場合は受信されました)。
func (p *Pipeline) EndRequest(id uint)

// StartResponseは、指定したidのリクエストを受信する（または、サーバーの場合は送信する）時までブロックします。
func (p *Pipeline) StartResponse(id uint)

// EndResponseは、指定されたIDのレスポンスが受信されたことをpに通知します
// （または、サーバーの場合は送信されます）。
func (p *Pipeline) EndResponse(id uint)
