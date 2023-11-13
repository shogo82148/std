// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// Goの実行トレーサーです。
// このトレーサーは、ゴルーチンの作成/ブロック/アンブロック、システムコールの入出力/ブロック、GCに関連するイベント、ヒープサイズの変更、プロセッサの開始/停止など、様々な実行イベントをキャプチャし、コンパクトな形式でバッファに書き込みます。
// ほとんどのイベントには、正確なナノ秒単位のタイムスタンプとスタックトレースが記録されます。
// 詳細については、https://golang.org/s/go15trace を参照してください。
=======
//go:build !goexperiment.exectracer2

// Go execution tracer.
// The tracer captures a wide range of execution events like goroutine
// creation/blocking/unblocking, syscall enter/exit/block, GC-related events,
// changes of heap size, processor start/stop, etc and writes them to a buffer
// in a compact form. A precise nanosecond-precision timestamp and a stack
// trace is captured for most events.
// See https://golang.org/s/go15trace for more info.
>>>>>>> upstream/master

package runtime

// StartTraceは現在のプロセスのトレースを有効にします。
// トレース中はデータがバッファされ、ReadTraceを介して利用可能です。
// トレースが既に有効化されている場合、StartTraceはエラーを返します。
// ほとんどのクライアントはruntime/traceパッケージやtestingパッケージの-test.traceフラグを直接呼び出す代わりに使用するべきです。
func StartTrace() error

// StopTraceは、以前に有効にされていた場合にトレースを停止します。
// StopTraceは、トレースのすべての読み取りが完了するまで戻りません。
func StopTrace()

// ReadTrace はバイナリ追跡データの次のチャンクを返します。データが利用可能になるまでブロックされます。もし追跡がオフになっており、オンの間に蓄積されたデータがすべて返された場合、ReadTrace は nil を返します。呼び出し元は、再度 ReadTrace を呼び出す前に返されたデータをコピーする必要があります。
// ReadTrace は一度に1つの goroutine から呼び出す必要があります。
func ReadTrace() []byte
