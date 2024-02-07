// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !goexperiment.exectracer2

// Go execution tracer.
// The tracer captures a wide range of execution events like goroutine
// creation/blocking/unblocking, syscall enter/exit/block, GC-related events,
// changes of heap size, processor start/stop, etc and writes them to a buffer
// in a compact form. A precise nanosecond-precision timestamp and a stack
// trace is captured for most events.
// See https://golang.org/s/go15trace for more info.

package runtime

<<<<<<< HEAD
// StartTraceは現在のプロセスのトレースを有効にします。
// トレース中はデータがバッファされ、ReadTraceを介して利用可能です。
// トレースが既に有効化されている場合、StartTraceはエラーを返します。
// ほとんどのクライアントはruntime/traceパッケージやtestingパッケージの-test.traceフラグを直接呼び出す代わりに使用するべきです。
=======
// StartTrace enables tracing for the current process.
// While tracing, the data will be buffered and available via [ReadTrace].
// StartTrace returns an error if tracing is already enabled.
// Most clients should use the [runtime/trace] package or the [testing] package's
// -test.trace flag instead of calling StartTrace directly.
>>>>>>> upstream/release-branch.go1.22
func StartTrace() error

// StopTraceは、以前に有効にされていた場合にトレースを停止します。
// StopTraceは、トレースのすべての読み取りが完了するまで戻りません。
func StopTrace()

// ReadTrace はバイナリ追跡データの次のチャンクを返します。データが利用可能になるまでブロックされます。もし追跡がオフになっており、オンの間に蓄積されたデータがすべて返された場合、ReadTrace は nil を返します。呼び出し元は、再度 ReadTrace を呼び出す前に返されたデータをコピーする必要があります。
// ReadTrace は一度に1つの goroutine から呼び出す必要があります。
func ReadTrace() []byte
