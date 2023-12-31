// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// traceパッケージには、Go実行トレーサーのためのトレースを生成するプログラムの機能が含まれています。
//
// # Tracing runtime activities
//
// 実行トレースは、ゴルーチンの作成/ブロック/アンブロック、システムコールの入力/出力/ブロック、
// GC関連のイベント、ヒープサイズの変更、プロセッサの開始/停止など、幅広い実行イベントをキャプチャします。
// CPUプロファイリングがアクティブな場合、実行トレーサーはこれらのサンプルも含めるように努力します。
// ほとんどのイベントに対して、正確なナノ秒精度のタイムスタンプとスタックトレースがキャプチャされます。
// 生成されたトレースは `go tool trace` を使用して解釈することができます。
//
// 標準のテストパッケージで構築されたテストとベンチマークのトレースのサポートは、
// `go test`に組み込まれています。例えば、以下のコマンドは現在のディレクトリでテストを実行し、
// トレースファイル（trace.out）を書き出します。
//
//	go test -trace=trace.out
//
// このruntime/traceパッケージは、同等のトレーシングサポートをスタンドアロンのプログラムに追加するためのAPIを提供します。
// このAPIを使用してトレーシングを有効にする方法を示す例を参照してください。
//
// また、トレースデータには標準的なHTTPインターフェースもあります。以下の行を追加すると、
// /debug/pprof/trace URLの下にハンドラがインストールされ、ライブトレースをダウンロードできます：
//
//	import _ "net/http/pprof"
//
// このインポートによってインストールされたすべてのデバッグエンドポイントについての詳細は、
// [net/http/pprof] パッケージを参照してください。
//
// # User annotation
//
// traceパッケージは、実行中の興味深いイベントをログに記録するために使用できる
// ユーザー注釈APIを提供します。
//
// ユーザー注釈には3つのタイプがあります：ログメッセージ、領域、タスク。
//
// [Log] は、メッセージのカテゴリや [Log] を呼び出したゴルーチンなどの追加情報とともに、
// 実行トレースにタイムスタンプ付きのメッセージを発行します。実行トレーサーは、
// ログカテゴリと [Log] で提供されるメッセージを使用してゴルーチンをフィルタリング
// およびグループ化するUIを提供します。
//
// リージョンは、ゴルーチンの実行中の時間間隔をログに記録するためのものです。
// 定義上、リージョンは同じゴルーチンで開始し終了します。
// リージョンは、サブインターバルを表すためにネストすることができます。
// 例えば、次のコードは、カプチーノ作成操作の連続したステップの期間をトレースするために、
// 実行トレースに4つのリージョンを記録します。
//
//	trace.WithRegion(ctx, "makeCappuccino", func() {
//
//	   // orderIDは、多くのカプチーノ注文領域レコードの中から特定の注文を識別するために使用します。
//	   trace.Log(ctx, "orderID", orderID)
//
//	   trace.WithRegion(ctx, "steamMilk", steamMilk)
//	   trace.WithRegion(ctx, "extractCoffee", extractCoffee)
//	   trace.WithRegion(ctx, "mixMilkCoffee", mixMilkCoffee)
//	})
//
// タスクは、RPCリクエスト、HTTPリクエスト、または複数のゴルーチンが協力して行う必要がある
// 興味深いローカル操作など、論理的な操作のトレースを支援する高レベルのコンポーネントです。
// タスクは複数のゴルーチンを含む可能性があるため、[context.Context] オブジェクトを介して追跡されます。
// [NewTask] は新しいタスクを作成し、それを返された [context.Context] オブジェクトに埋め込みます。
// ログメッセージとリージョンは、[Log] と [WithRegion] に渡されたContextにあるタスク（存在する場合）に添付されます。
//
// 例えば、ミルクを泡立て、コーヒーを抽出し、ミルクとコーヒーを別々のゴルーチンで混ぜることにしました。
// タスクを使用すると、トレースツールは特定のカプチーノ注文に関与するゴルーチンを識別できます。
//
//	ctx, task := trace.NewTask(ctx, "makeCappuccino")
//	trace.Log(ctx, "orderID", orderID)
//
//	milk := make(chan bool)
//	espresso := make(chan bool)
//
//	go func() {
//	        trace.WithRegion(ctx, "steamMilk", steamMilk)
//	        milk <- true
//	}()
//	go func() {
//	        trace.WithRegion(ctx, "extractCoffee", extractCoffee)
//	        espresso <- true
//	}()
//	go func() {
//	        defer task.End() // アセンブルが完了したら、注文は完了します。
//	        <-espresso
//	        <-milk
//	        trace.WithRegion(ctx, "mixMilkCoffee", mixMilkCoffee)
//	}()
//
// トレースツールは、タスクの作成とタスクの終了の間の時間を測定することでタスクの遅延を計算し、
// トレース内で見つかった各タスクタイプの遅延分布を提供します。
package trace

import (
	"github.com/shogo82148/std/io"
)

// Start は現在のプログラムのトレースを有効にします。
// トレース中は、トレースはバッファリングされ、w に書き込まれます。
// トレースが既に有効になっている場合、Start はエラーを返します。
func Start(w io.Writer) error

// Stopは現在のトレースを停止します（存在する場合）。
// トレースのすべての書き込みが完了するまで、Stopは戻りません。
func Stop()
