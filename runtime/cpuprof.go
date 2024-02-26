// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// CPU プロファイリング。
//
// プロファイリングクロックのタイムステップのシグナルハンドラは、最新のトレースのログに新しいスタックトレースを追加します。ログは、ユーザーゴールーチンによって読み込まれ、フォーマットされたプロファイルデータに変換されます。読み取り側がログに追いついていない場合、これらの書き込みは失われたレコードのカウントとして記録されます。実際のプロファイルバッファはprofbuf.goにあります。

package runtime

// SetCPUProfileRateはCPUプロファイリングのレートをhzサンプル/秒に設定します。
// hz <= 0の場合、プロファイリングはオフになります。
// プロファイラがオンの場合、レートを変更する前にオフにする必要があります。
//
<<<<<<< HEAD
// ほとんどのクライアントは、[runtime/pprof] パッケージまたは
// [testing] パッケージの-test.cpuprofileフラグを直接呼び出す代わりに使用するべきです。
=======
// Most clients should use the [runtime/pprof] package or
// the [testing] package's -test.cpuprofile flag instead of calling
// SetCPUProfileRate directly.
>>>>>>> upstream/master
func SetCPUProfileRate(hz int)

// CPUProfileはパニックします。
// 以前はランタイムによって生成されたpprof形式のプロファイルの
// チャンクへの直接的なアクセスを提供していました。
// その形式を生成する方法が変更されたため、
// この機能は削除されました。
//
<<<<<<< HEAD
// Deprecated: [runtime/pprof] パッケージ、
// または [net/http/pprof] パッケージのハンドラ、
// または [testing] パッケージの-test.cpuprofileフラグを代わりに使用してください。
=======
// Deprecated: Use the [runtime/pprof] package,
// or the handlers in the [net/http/pprof] package,
// or the [testing] package's -test.cpuprofile flag instead.
>>>>>>> upstream/master
func CPUProfile() []byte
