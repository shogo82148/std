// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package synctest は、並行コードのテストを支援します。
//
// [Test] 関数は、隔離された「バブル」の中で関数を実行します。
// バブル内で開始された goroutine は、すべてそのバブルの一部です。
//
// # Time
//
// バブル内では、[time] パッケージは疑似クロックを使います。
// 各バブルは独自のクロックを持ちます。
// 初期時刻は UTC 2000-01-01 の午前0時です。
//
// バブル内の時間は、バブル内のすべての goroutine が
// 永続的にブロックされたときにのみ進みます。
// 「永続的にブロック」の正確な定義は下記を参照してください。
//
// たとえば、このテストは2秒待たずに即座に実行されます。
//
//	func TestTime(t *testing.T) {
//		synctest.Test(t, func(t *testing.T) {
//			start := time.Now() // 常に UTC 2000-01-01 の午前0時
//			go func() {
//				time.Sleep(1 * time.Second)
//				t.Log(time.Since(start)) // 常に "1s" を記録
//			}()
//			time.Sleep(2 * time.Second) // ここが返る前に上の goroutine が実行される
//			t.Log(time.Since(start))    // 常に "2s" を記録
//		})
//	}
//
// バブルのルート goroutine が終了すると、時間は進まなくなります。
//
// # Blocking
//
// バブル内の goroutine が「永続的にブロック」されているとは、
// ブロック状態であり、かつ同じバブル内の別の goroutine によってしか
// アンブロックされない状態を指します。
// バブル外のイベントによってアンブロックされ得る goroutine は、
// 永続的にブロックされているとは見なされません。
//
// [Wait] 関数は、バブル内の他のすべての goroutine が
// 永続的にブロックされるまで待機します。
//
// 例:
//
//	func TestWait(t *testing.T) {
//		synctest.Test(t, func(t *testing.T) {
//			done := false
//			go func() {
//				done = true
//			}()
//			// Wait は、上の goroutine が終了するまでブロックする。
//			synctest.Wait()
//			t.Log(done) // 常に "true" を記録
//		})
//	}
//
// バブル内のすべての goroutine が永続的にブロックされると:
//
//   - [Wait] が呼ばれていれば、それが返ります。
//   - そうでなければ、少なくとも1つの goroutine をアンブロックする
//     次の時刻まで時間が進みます（そのような時刻が存在し、かつ
//     バブルのルート goroutine がまだ終了していない場合）。
//   - それ以外の場合はデッドロックとなり、[Test] は panic します。
//
// 次の操作は、goroutine を永続的にブロックします:
//
//   - バブル内で作成されたチャネルに対するブロッキング送受信
//   - すべての case がバブル内で作成されたチャネルである
//     ブロッキング select 文
//   - [sync.Cond.Wait]
//   - [sync.WaitGroup.Wait]（[sync.WaitGroup.Add] がバブル内で
//     呼ばれた場合）
//   - [time.Sleep]
//
// 上記リストにない操作は、永続的ブロッキングではありません。
// とくに、次の操作は goroutine をブロックし得ますが、
// バブル外で発生するイベントでアンブロックされ得るため、
// 永続的ブロッキングではありません:
//
//   - [sync.Mutex] または [sync.RWMutex] のロック取得
//   - ネットワークソケットからの読み取りなどの I/O 待ち
//   - システムコール
//
// # Isolation
//
// バブル内で作成されたチャネル、[time.Timer]、[time.Ticker] は
// そのバブルに関連付けられます。バブル外から、それらのバブル付き
// チャネル・タイマー・ティッカーを操作すると panic します。
//
// [sync.WaitGroup] は、最初の Add または Go 呼び出し時に
// バブルに関連付けられます。いったん WaitGroup がバブルに関連付け
// られた後で、そのバブル外から Add または Go を呼ぶと致命的エラーです。
// （技術的制限として、"var wg sync.WaitGroup" のように
// パッケージ変数として定義された WaitGroup はバブルに関連付けられず、
// その操作は永続的ブロッキングにならない可能性があります。
// この制限は、"var wg = new(sync.WaitGroup)" のような
// パッケージ変数に格納された *WaitGroup には適用されません。）
//
// [sync.Cond.Wait] は永続的ブロッキングです。バブル外から
// Cond.Wait でブロックしているバブル内 goroutine を起こすと
// 致命的エラーになります。
//
// [runtime.AddCleanup] と [runtime.SetFinalizer] で登録された
// クリーンアップ関数とファイナライザは、いずれのバブル外でも実行されます。
//
// # Example: Context.AfterFunc
//
// この例は、[context.AfterFunc] 関数のテスト方法を示します。
//
// AfterFunc は、context がキャンセルされた後に新しい goroutine で
// 実行する関数を登録します。
//
// このテストは、その関数が context のキャンセル前には実行されず、
// キャンセル後に実行されることを検証します。
//
//	func TestContextAfterFunc(t *testing.T) {
//		synctest.Test(t, func(t *testing.T) {
//			// キャンセル可能な context.Context を作成する。
//			ctx, cancel := context.WithCancel(t.Context())
//
//			// context.AfterFunc は、context がキャンセルされたときに
//			// 呼び出される関数を登録する。
//			afterFuncCalled := false
//			context.AfterFunc(ctx, func() {
//				afterFuncCalled = true
//			})
//
//			// context はまだキャンセルされていないので、AfterFunc は呼ばれない。
//			synctest.Wait()
//			if afterFuncCalled {
//				t.Fatalf("before context is canceled: AfterFunc called")
//			}
//
//			// context をキャンセルし、AfterFunc の実行完了を待つ。
//			// AfterFunc が実行されたことを確認する。
//			cancel()
//			synctest.Wait()
//			if !afterFuncCalled {
//				t.Fatalf("before context is canceled: AfterFunc not called")
//			}
//		})
//	}
//
// # Example: Context.WithTimeout
//
// この例は、[context.WithTimeout] 関数のテスト方法を示します。
//
// WithTimeout は、タイムアウト後にキャンセルされる context を作成します。
//
// このテストは、タイムアウト前には context がキャンセルされず、
// タイムアウト後にキャンセルされることを検証します。
//
//	func TestContextWithTimeout(t *testing.T) {
//		synctest.Test(t, func(t *testing.T) {
//			// タイムアウト後にキャンセルされる context.Context を作成する。
//			const timeout = 5 * time.Second
//			ctx, cancel := context.WithTimeout(t.Context(), timeout)
//			defer cancel()
//
//			// タイムアウトよりわずかに短く待つ。
//			time.Sleep(timeout - time.Nanosecond)
//			synctest.Wait()
//			if err := ctx.Err(); err != nil {
//				t.Fatalf("before timeout: ctx.Err() = %v, want nil\n", err)
//			}
//
//			// 残り時間だけ待ってタイムアウトに到達する。
//			time.Sleep(time.Nanosecond)
//			synctest.Wait()
//			if err := ctx.Err(); err != context.DeadlineExceeded {
//				t.Fatalf("after timeout: ctx.Err() = %v, want DeadlineExceeded\n", err)
//			}
//		})
//	}
//
// # Example: HTTP 100 Continue
//
// この例は、[http.Transport] の 100 Continue 処理のテスト方法を示します。
//
// リクエストを送る HTTP クライアントは、"Expect: 100-continue" ヘッダーを
// 含めることで、追加データを送る予定であることをサーバーに伝えられます。
// サーバーは、データ送信を要求するために 100 Continue 情報レスポンスを返すか、
// あるいはデータ不要であることを伝える別のステータスを返せます。
// たとえば大きなファイルをアップロードするクライアントは、送信前に
// サーバーが受け入れる意思があることを確認するためにこの機能を使えます。
//
// このテストは、"Expect: 100-continue" ヘッダーを付けて送信したときに、
// HTTP クライアントがサーバーから要求される前にリクエスト本文を送らないこと、
// そして 100 Continue レスポンス受信後には本文を送ることを確認します。
//
//	func TestHTTPTransport100Continue(t *testing.T) {
//		synctest.Test(t, func(*testing.T) {
//			// プロセス内の疑似ネットワーク接続を作成する。
//			// このテストではループバック接続は使えない。
//			// ネットワーク I/O でブロックした goroutine があると、
//			// synctest のバブルがアイドル状態になれないため。
//			srvConn, cliConn := net.Pipe()
//			defer cliConn.Close()
//			defer srvConn.Close()
//
//			tr := &http.Transport{
//				// 上で作成した疑似ネットワーク接続を使う。
//				DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
//					return cliConn, nil
//				},
//				// "Expect: 100-continue" 処理を有効化する。
//				ExpectContinueTimeout: 5 * time.Second,
//			}
//
//			// "Expect: 100-continue" ヘッダー付きリクエストを送る。
//			// テスト終了まで完了しないため、新しい goroutine で送る。
//			body := "request body"
//			go func() {
//				req, _ := http.NewRequest("PUT", "http://test.tld/", strings.NewReader(body))
//				req.Header.Set("Expect", "100-continue")
//				resp, err := tr.RoundTrip(req)
//				if err != nil {
//					t.Errorf("RoundTrip: unexpected error %v\n", err)
//				} else {
//					resp.Body.Close()
//				}
//			}()
//
//			// クライアントが送ったリクエストヘッダーを読む。
//			req, err := http.ReadRequest(bufio.NewReader(srvConn))
//			if err != nil {
//				t.Fatalf("ReadRequest: %v\n", err)
//			}
//
//			// クライアントから送られる本文をバッファへコピーする goroutine を開始する。
//			// バブル内の全 goroutine がブロックするまで待ち、
//			// まだ本文を読んでいないことを確認する。
//			var gotBody bytes.Buffer
//			go io.Copy(&gotBody, req.Body)
//			synctest.Wait()
//			if got, want := gotBody.String(), ""; got != want {
//				t.Fatalf("before sending 100 Continue, read body: %q, want %q\n", got, want)
//			}
//
//			// クライアントへ "100 Continue" レスポンスを書き、
//			// リクエスト本文が送られることを確認する。
//			srvConn.Write([]byte("HTTP/1.1 100 Continue\r\n\r\n"))
//			synctest.Wait()
//			if got, want := gotBody.String(), body; got != want {
//				t.Fatalf("after sending 100 Continue, read body: %q, want %q\n", got, want)
//			}
//
//			// 最後に "200 OK" レスポンスを送ってリクエストを完了する。
//			srvConn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
//
//			// テスト中に複数の goroutine を開始した。
//			// synctest.Test は、戻る前にそれらがすべて終了するのを待つ。
//		})
//	}
package synctest

import (
	"github.com/shogo82148/std/testing"
	"github.com/shogo82148/std/time"
)

// Test は、新しいバブル内で f を実行します。
//
// Test は、戻る前にバブル内のすべての goroutine が終了するのを待ちます。
// バブル内 goroutine がデッドロックすると、テストは失敗します。
//
// Test をバブル内から呼び出してはいけません。
//
// f に渡される [*testing.T] には、次の性質があります:
//
//   - T.Cleanup 関数はバブル内で実行され、
//     Test が返る直前に呼ばれます。
//   - T.Context は、バブルに関連付けられた Done チャネルを持つ
//     [context.Context] を返します。
//   - T.Run、T.Parallel、T.Deadline は呼び出してはいけません。
func Test(t *testing.T, f func(*testing.T))

// Wait は、現在のバブル内の goroutine のうち、
// 現在の goroutine 以外のすべてが永続的にブロックされるまで待機します。
//
// Wait をバブル外から呼び出してはいけません。
// 同じバブル内で、複数 goroutine から同時に Wait を呼び出してはいけません。
func Wait()

// Sleep blocks until the current bubble's clock has advanced
// by the duration of d and every goroutine within the current bubble,
// other than the current goroutine, is durably blocked.
//
// This is exactly equivalent to
//
//	time.Sleep(d)
//	synctest.Wait()
//
// In tests, this is often preferable to calling only [time.Sleep].
// If the test itself and another goroutine running the system under test
// sleeps for the exact same amount of time, it's unpredictable which
// of the two goroutines will run first. The test itself usually wants
// to wait for the system under test to "settle" after sleeping.
// This is what Sleep accomplishes.
func Sleep(d time.Duration)
