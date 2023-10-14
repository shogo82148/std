// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Executableは、現在のプロセスを開始した実行可能ファイルのパス名を返します。
// パスがまだ正しい実行可能ファイルを指しているとは限りません。
// プロセスの開始にシンボリックリンクが使用された場合、オペレーティングシステムによって結果は
// シンボリックリンクまたはそれが指していたパスになる可能性があります。
// 安定した結果が必要な場合は、path/filepath.EvalSymlinksが役立ちます。
//
// Executableは、エラーが発生しない限り、絶対パスを返します。
//
// 主な利用ケースは、実行可能ファイルに対して相対的に配置されたリソースを見つけることです。
func Executable() (string, error)
