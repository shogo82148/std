// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Traceはトレースファイルを表示するためのツールです。

トレースファイルは以下で生成できます：
  - runtime/trace.Start
  - net/http/pprofパッケージ
  - go test -trace

使用例：
'trace.out'というトレースファイルを'go test'で生成する：

	go test -trace trace.out pkg

ウェブブラウザでトレースを表示する：

	go tool trace trace.out

トレースからpprofのようなプロファイルを生成する：

	go tool trace -pprof=TYPE trace.out > TYPE.pprof

サポートされているプロファイルのタイプ：
  - net: ネットワークブロッキングプロファイル
  - sync: 同期ブロッキングプロファイル
  - syscall: システムコールブロッキングプロファイル
  - sched: スケジューラのレイテンシプロファイル

その後、プロファイルを分析するためにpprofツールを使用できます：

	go tool pprof TYPE.pprof

注意：'go tool trace'を起動した場合に利用可能なさまざまなプロファイルは、すべてのブラウザで動作しますが、トレースビューア自体（'view trace'ページ）はChrome/Chromiumプロジェクトから提供されており、そのブラウザでのみアクティブにテストされています。
*/package main
