// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージpprofは、pprof可視化ツールが期待する形式で実行時プロファイリングデータをHTTPサーバー経由で提供します。
//
// このパッケージは通常、そのHTTPハンドラを登録する副作用のためにのみインポートされます。
// ハンドルされるパスはすべて/debug/pprof/で始まります。
//
// pprofを使用するためには、このパッケージをプログラムにリンクしてください：
//
//	import _ "net/http/pprof"
//
// アプリケーションがすでにHTTPサーバーを実行していない場合、サーバーを起動する必要があります。
// "net/http"と"log"をインポートし、次のコードをメイン関数に追加してください：
//
//	go func() {
//	    log.Println(http.ListenAndServe("localhost:6060", nil))
//	}()
//
// デフォルトでは、このパッケージで定義されている[Cmdline]、[Profile]、[Symbol]、[Trace]プロファイル
// に加えて[runtime/pprof.Profile]にリストされているすべてのプロファイルが使用可能です（[Handler]経由）。
// DefaultServeMuxを使用していない場合は、使用しているmuxにハンドラを登録する必要があります。
//
// ＃ パラメータ
//
// パラメータはGETクエリパラメータを介して渡すことができます：
//
//   - debug=N（すべてのプロファイル）：応答形式：N = 0：バイナリ（デフォルト）、N> 0：プレーンテキスト
//   - gc=N（ヒーププロファイル）：N> 0：プロファイリング前にガベージコレクションを実行
//   - seconds=N（allocs、block、goroutine、heap、mutex、threadcreateプロファイル）：デルタプロファイルを返す
//   - seconds=N（cpu（profile）、トレースプロファイル）：指定された期間のプロファイル
//
// ＃ 使用例
//
// ヒーププロファイルを表示するためにpprofツールを使用する場合：
//
//	go tool pprof http://localhost:6060/debug/pprof/heap
//
// または30秒のCPUプロファイルを表示する場合：
//
//	go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
//
// または、プログラムで [runtime.SetBlockProfileRate] を呼び出した後にブロッキングプロファイルのゴルーチンを表示する場合：
//
//	go tool pprof http://localhost:6060/debug/pprof/block
//
// または、プログラムで [runtime.SetMutexProfileFraction] を呼び出した後にコンテンデッドなミューテックスのホルダーを表示する場合：
//
//	go tool pprof http://localhost:6060/debug/pprof/mutex
//
// このパッケージはまた、「go tool trace」コマンドに対応する実行トレースデータを提供するハンドラもエクスポートしています。
// 5秒の実行トレースを収集するには：
//
//	curl -o trace.out http://localhost:6060/debug/pprof/trace?seconds=5
//	go tool trace trace.out
//
// 使用可能なプロファイルをすべて表示するには、ブラウザでhttp://localhost:6060/debug/pprof/を開いてください。
//
// 機能の詳細については、次のURLをご覧ください：
// https://blog.golang.org/2011/06/profiling-go-programs.html。
package pprof

import (
	"github.com/shogo82148/std/net/http"
)

// Cmdlineは実行中のプログラムのコマンドラインをNULバイトで区切られた引数として返します。
// このパッケージの初期化は/debug/pprof/cmdlineとして登録されます。
func Cmdline(w http.ResponseWriter, r *http.Request)

// Profile は、pprofの形式でCPUプロファイルを応答します。
// プロファイリングは、秒数が指定されたGETパラメータで指定された期間、または指定されていない場合は30秒間続きます。
// パッケージの初期化により、/debug/pprof/profile として登録されます。
func Profile(w http.ResponseWriter, r *http.Request)

// Traceはバイナリ形式の実行トレースで応答します。
// トレースは、秒数が指定されていない場合は1秒間、指定されている場合はGETパラメータで指定された時間の間続きます。
// パッケージの初期化により、/debug/pprof/traceとして登録されます。
func Trace(w http.ResponseWriter, r *http.Request)

// Symbolはリクエストにリストされたプログラムカウンターを検索し、
// プログラムカウンターと関数名のマッピングを返す。
// パッケージの初期化では、それを/debug/pprof/symbolとして登録します。
func Symbol(w http.ResponseWriter, r *http.Request)

// Handlerは、指定したプロファイルを提供するHTTPハンドラを返します。
// 使用可能なプロファイルは[runtime/pprof.Profile]で見つけることができます。
func Handler(name string) http.Handler

// Indexはリクエストで指定されたpprof形式のプロファイルで応答します。
// たとえば、"/debug/pprof/heap"は"heap"プロファイルを提供します。
// Indexは"/debug/pprof/"へのリクエストに対して利用可能なプロファイルをリストしたHTMLページで応答します。
func Index(w http.ResponseWriter, r *http.Request)
