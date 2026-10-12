// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/sync"
)

// ClientConnはHTTPサーバーへのクライアント接続です。
//
// [Transport] とは異なり、ClientConnは単一の接続を表します。
// ほとんどのユーザーは、クライアント接続を直接作成するのではなくTransportを使用するべきです。
type ClientConn struct {
	cc genericClientConn

	stateHookMu      sync.Mutex
	userStateHook    func(*ClientConn)
	stateHookRunning bool
	lastAvailable    int
	lastInFlight     int
	lastClosed       bool
}

// NewClientConnは、指定されたアドレスへの新しいクライアント接続を作成します。
//
// schemeが"http"の場合、接続は暗号化されません。
// schemeが"https"の場合、接続はTLSを使用します。
//
// 新しい接続で使用されるプロトコルは、scheme、Transport.Protocols設定フィールド、
// およびサーバーがサポートするプロトコルによって決定されます。
// 詳細はTransport.Protocolsを参照してください。
//
// Transport.Proxyが設定されていて、指定されたアドレスに送信されるリクエストが
// プロキシーを使用するべきであることを示している場合、新しい接続はそのプロキシーを使用します。
//
// NewClientConnは、Transportが指定されたホストへのキャッシュされた接続を
// すでに持っている場合でも、常に新しい接続を作成します。
//
// 新しい接続はTransportの接続キャッシュには追加されず、
// [Transport.RoundTrip] によって使用されることはありません。
// MaxIdleConnsおよびMaxConnsPerHostの制限にはカウントされません。
//
// 新しい接続を閉じるのは呼び出し元の責任です。
func (t *Transport) NewClientConn(ctx context.Context, scheme, address string) (*ClientConn, error)

// Closeは接続を閉じます。
// 未完了のRoundTrip呼び出しは中断されます。
func (cc *ClientConn) Close() error

// Errは致命的な接続エラーを報告します。
// 接続が使用可能な場合はnilを返します。
// nil以外を返した場合、その接続はもう使用できません。
func (cc *ClientConn) Err() error

// RoundTripは [RoundTripper] インターフェースを実装します。
//
// リクエストは、要求されたURLやプロキシー設定に関係なく、
// クライアント接続上で送信されます。
//
// 接続が同時実行数の上限に達している場合、RoundTripは
// リクエストを送信する前に接続が利用可能になるまで待機します。
func (cc *ClientConn) RoundTrip(req *Request) (*Response, error)

// Availableは、ブロックせずに接続に送信できるリクエストの数を報告します。
// 接続が閉じられている場合は0を返します。
func (cc *ClientConn) Available() int

// InFlightは、予約されたリクエストを含む、処理中のリクエストの数を報告します。
// 接続が閉じられている場合は0を返します。
func (cc *ClientConn) InFlight() int

// Reserveは接続上の同時実行スロットを予約します。
// Reserveがnilを返した場合、既存のリクエストの完了を待たずに
// RoundTripをもう1回呼び出すことができます。
//
// 予約された同時実行スロットは処理中のリクエストとしてカウントされます。
// RoundTripの呼び出しが成功すると、Availableのカウントが減少し、
// InFlightのカウントが増加します。
//
// Reserveの呼び出しが成功するたびに、その後にRoundTripまたはReleaseを
// ちょうど1回呼び出して、予約を消費または解放するべきです。
//
// 接続が閉じられているか同時実行数の上限に達している場合、
// Reserveはエラーを返します。
func (cc *ClientConn) Reserve() error

// Releaseは、Reserveによって予約された未使用の同時実行スロットを解放します。
// 予約された同時実行スロットがない場合は何もしません。
func (cc *ClientConn) Release()

// SetStateHookは、接続の状態が変化したときにfが呼び出されるようにします。
// fは同時に最大1回まで呼び出されます。
// 接続の作成後に接続の状態が変化している場合、
// fは別のゴルーチンで即座に呼び出されます。
// fはRoundTripまたはResponse.Body.Closeから同期的に呼び出される場合があります。
//
// SetStateHookが複数回呼び出された場合、新しいフックが古いフックを置き換えます。
// fがnilの場合、SetStateHookが戻った後にfが呼び出されることはありません。
//
// fは、Availableが増加したとき（接続でさらにリクエストを送信できる）、
// InFlightが減少したとき（既存のリクエストが完了した）、またはErrがnil以外を
// 返し始めたとき（接続がもう使用できない）に呼び出されます。
func (cc *ClientConn) SetStateHook(f func(*ClientConn))
